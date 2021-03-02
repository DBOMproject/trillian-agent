/*
 * Copyright 2020 Unisys Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"strconv"
	dbom "trillian-agent/dbom"
	"trillian-agent/helpers"
	"trillian-agent/logger"
	"trillian-agent/responses"
	"trillian-agent/tracing"
	client "trillian-agent/trillian"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/opentracing/opentracing-go"

	"trillian-agent/models"
	"trillian-agent/restapi/operations"
	"trillian-agent/restapi/operations/record"

	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/google/trillian"
	"google.golang.org/grpc"
)

//go:generate swagger generate server --target ../../trillian-agent --name TrillianAgent --spec ../../../../api-specs/agent/v2/agent.json --principal interface{}

var configLogger = logger.GetLogger("Restapi:Config")
var getChannelClient = dbom.GetChannelClient
var getCurrentRevision = (*client.MapClient).GetCurrentRevision
var getChannel = dbom.GetChannel
var getRecord = dbom.GetRecord

var (
	trillianEndpoint      = helpers.GetEnv("TRILLIAN_ENDPOINT", "localhost:8091")
	channelConfigMapID, _ = strconv.ParseInt(helpers.GetEnv("CHANNEL_CONFIG_MAP_ID", "-1"), 10, 64)
)

const (
	//CREATE commit type
	CREATE = "CREATE"
	//UPDATE commit type
	UPDATE = "UPDATE"
	//ATTACH commit type
	ATTACH = "ATTACH"
	//DETACH commit type
	DETACH = "DETACH"
	//TRANSFERIN commit type
	TRANSFERIN = "TRANSFER-IN"
	//TRANSFEROUT commit type
	TRANSFEROUT = "TRANSFER-OUT"
)

func configureFlags(api *operations.TrillianAgentAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TrillianAgentAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf
	apiLogger := logger.GetLogger("Server")
	api.Logger = apiLogger.Info().Msgf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.RecordAuditRecordHandler = record.AuditRecordHandlerFunc(func(params record.AuditRecordParams) middleware.Responder {
		tracer, closer, err := tracing.SetupGlobalTracer()
		if err != nil {
			configLogger.Err(err).Msg("Unable to initialize Jaeger tracer. Falling back to the NoopTracer")
		} else {
			defer closer.Close()
		}
		configLogger.Info().Msg("[Restapi:RecordAuditRecordHandler] Entered")
		span, ctx := opentracing.StartSpanFromContextWithTracer(params.HTTPRequest.Context(), tracer, "RecordAuditRecordHandler")
		defer span.Finish()
		conn, err := grpc.Dial(trillianEndpoint, grpc.WithInsecure())
		if err != nil {
			tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
			return responses.ErrAuditInternalServerError(err)
		}
		defer conn.Close()

		if ctx == nil {
			ctx = context.Background()
		}
		trillMapClient := trillian.NewTrillianMapClient(conn)
		trillAdminClient := trillian.NewTrillianAdminClient(conn)
		channelMapClientTree, channelErr := getChannelClient(ctx, trillAdminClient, trillMapClient, channelConfigMapID, tracer)
		if channelErr != nil {
			tracing.LogAndTraceErr(apiLogger, span, channelErr, responses.InternalError)
			return responses.ErrAuditResourceNotFound()
		}

		channelMapClient := client.MapClient{MapClient: channelMapClientTree}
		channelRevision, err := getCurrentRevision(&channelMapClient, ctx, channelConfigMapID, tracer)
		channelRevision++
		if err != nil {
			tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
			return responses.ErrAuditInternalServerError(err)
		}
		channel, err := getChannel(ctx, &channelMapClient, params.ChannelID, tracer)
		if err != nil {
			tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
			return responses.ErrRetrieveRecordInternalServerError(err)
		} else if channel == nil {
			tracing.LogAndTraceErr(apiLogger, span, nil, responses.ChannelNotFound)
			return responses.ErrRetrieveChannelNotFound()
		}
		mapClientTree, err := getChannelClient(ctx, trillAdminClient, trillMapClient, channel.MapID, tracer)
		if err != nil {
			tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
			return responses.ErrAuditResourceNotFound()
		}
		mapClient := client.MapClient{MapClient: mapClientTree}
		var recordList []*models.AuditDefinition
		result, err := getRecord(ctx, &mapClient, params.RecordID, -1, tracer)
		if err != nil {
			tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
			return responses.ErrAuditInternalServerError(err)
		} else if result == nil {
			tracing.LogAndTraceErr(apiLogger, span, nil, responses.ResourceNotFound)
			return responses.ErrAuditResourceNotFound()
		}
		var auditRecord = result.AuditDefinition
		auditRecord.ID = &result.Revision
		recordList = append(recordList, &auditRecord)
		rev := result.PreviousRevision
		for rev > 0 {
			result, err := getRecord(ctx, &mapClient, params.RecordID, rev, tracer)
			if err != nil {
				tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
				return responses.ErrAuditInternalServerError(err)
			} else if result == nil {
				tracing.LogAndTraceErr(apiLogger, span, nil, responses.ResourceNotFound)
				return responses.ErrAuditResourceNotFound()
			}
			var auditRecord = result.AuditDefinition
			auditRecord.ID = &result.Revision
			recordList = append(recordList, &auditRecord)
			rev = result.PreviousRevision
		}

		var payload = models.AuditResponseDefinition{
			History: recordList,
		}
		var res = record.AuditRecordOK{Payload: &payload}
		//configLogger.Debug().Msgf("%v",res.Payload)
		configLogger.Debug().Msgf("%v", res.Payload)
		configLogger.Info().Msg("[Restapi:RecordAuditRecordHandler] Finished")
		span.Finish()
		return &res
	})

	api.RecordCommitRecordHandler = record.CommitRecordHandlerFunc(func(params record.CommitRecordParams) middleware.Responder {
		configLogger.Info().Msg("[Restapi:CommitRecordHandlerFunc] Entered")
		tracer, closer, err := tracing.SetupGlobalTracer()
		if err != nil {
			configLogger.Err(err).Msg("Unable to initialize Jaeger tracer. Falling back to the NoopTracer")
		} else {
			defer closer.Close()
		}
		span, ctx := opentracing.StartSpanFromContextWithTracer(params.HTTPRequest.Context(), tracer, "CommitRecordHandlerFunc")
		defer span.Finish()

		conn, err := grpc.Dial(trillianEndpoint, grpc.WithInsecure())
		if err != nil {
			tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
			return responses.ErrCommitInternalServerError(err)
		}
		defer conn.Close()

		trillMapWriteClient := trillian.NewTrillianMapWriteClient(conn)
		if ctx == nil {
			ctx = context.Background()
		}
		trillMapClient := trillian.NewTrillianMapClient(conn)
		trillAdminClient := trillian.NewTrillianAdminClient(conn)

		channelMapClientTree, channelErr := dbom.GetChannelClient(ctx, trillAdminClient, trillMapClient, channelConfigMapID, tracer)
		if channelErr != nil {
			tracing.LogAndTraceErr(apiLogger, span, channelErr, responses.InternalError)
			return responses.ErrCommitInternalServerError(channelErr)
		}

		var channelMapID = int64(0)
		channelMapClient := client.MapClient{MapClient: channelMapClientTree}
		channelRevision, err := channelMapClient.GetCurrentRevision(ctx, channelConfigMapID, tracer)
		channelRevision++
		if err != nil {
			tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
			return responses.ErrCommitInternalServerError(err)
		}
		channel, err := dbom.GetChannel(ctx, &channelMapClient, params.ChannelID, tracer)
		if err != nil {
			tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
			return responses.ErrCommitInternalServerError(err)
		}

		if params.CommitType == CREATE || params.CommitType == TRANSFERIN {

			revision := uint64(1)
			err := error(nil)
			if channel == nil {
				channelMapID, err = dbom.CreateChannel(ctx, trillAdminClient, trillMapClient, trillMapWriteClient, int64(channelRevision), channelConfigMapID, params.ChannelID, tracer)
				if err != nil {
					tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
					return responses.ErrCommitInternalServerError(err)
				}
			} else {
				if err != nil {
					tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
					return responses.ErrRetrieveRecordInternalServerError(err)
				} else if channel == nil {
					tracing.LogAndTraceErr(apiLogger, span, nil, responses.ChannelNotFound)
					return responses.ErrRetrieveChannelNotFound()
				}
				channelMapID = channel.MapID
				mapClientTree, err := dbom.GetChannelClient(ctx, trillAdminClient, trillMapClient, channel.MapID, tracer)
				if err != nil {
					tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
					return responses.ErrCommitInternalServerError(err)
				}
				mapClient := client.MapClient{MapClient: mapClientTree}
				revision, err = mapClient.GetCurrentRevision(ctx, channelConfigMapID, tracer)
				if err != nil {
					tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
					return responses.ErrCommitInternalServerError(err)
				}

				result, err := dbom.GetRecord(ctx, &mapClient, *params.Body.RecordID, -1, tracer)
				if err != nil {
					tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
					return responses.ErrCommitInternalServerError(err)
				} else if result != nil {
					tracing.LogAndTraceErr(apiLogger, span, nil, responses.ResourceExists)
					return responses.ErrCommitRecordConflict()
				}

				revision++
			}

			mapWriteClient := client.NewClient(trillMapWriteClient, channelMapID)

			dbom.CreateRecord(ctx, mapWriteClient, int64(revision), 0, params.ChannelID, params.CommitType, params.Body, tracer)
			if err != nil {
				tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
				return responses.ErrCommitInternalServerError(err)
			}
			var success = true
			var resDef = models.CreateRecordResponseDefinition{Success: &success}
			var res = record.CommitRecordOK{Payload: &resDef}
			configLogger.Debug().Msgf("%v", res.Payload)
			configLogger.Info().Msg("[Restapi:CommitRecordHandlerFunc] Finished")
			span.Finish()
			return &res
		} else if params.CommitType == UPDATE || params.CommitType == ATTACH || params.CommitType == DETACH || params.CommitType == TRANSFEROUT {
			if err != nil {
				tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
				return responses.ErrRetrieveRecordInternalServerError(err)
			} else if channel == nil {
				tracing.LogAndTraceErr(apiLogger, span, nil, responses.ChannelNotFound)
				return responses.ErrRetrieveChannelNotFound()
			}
			mapClientTree, err := dbom.GetChannelClient(ctx, trillAdminClient, trillMapClient, channel.MapID, tracer)
			if mapClientTree == nil {
				tracing.LogAndTraceErr(apiLogger, span, nil, responses.ChannelNotFound)
				return responses.ErrCommitChannelNotFound()
			} else if err != nil {
				tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
				return responses.ErrCommitInternalServerError(err)
			}
			mapClient := client.MapClient{MapClient: mapClientTree}
			revision, err := mapClient.GetCurrentRevision(ctx, channel.MapID, tracer)
			if err != nil {
				tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
				return responses.ErrCommitInternalServerError(err)
			}
			updateResult, err := dbom.GetRecord(ctx, &mapClient, *params.Body.RecordID, -1, tracer)
			if err != nil {
				tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
				return responses.ErrCommitInternalServerError(err)
			} else if updateResult == nil {
				tracing.LogAndTraceErr(apiLogger, span, nil, responses.ResourceNotFound)
				return responses.ErrCommitResourceNotFound()
			}
			revision++
			mapWriteClient := client.NewClient(trillMapWriteClient, channel.MapID)
			err = dbom.CreateRecord(ctx, mapWriteClient, int64(revision), updateResult.Revision, params.ChannelID, params.CommitType, params.Body, tracer)
			if err != nil {
				tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
				return responses.ErrCommitInternalServerError(err)
			}
			var success = true
			var resDef = models.CreateRecordResponseDefinition{Success: &success}
			var res = record.CommitRecordOK{Payload: &resDef}
			configLogger.Debug().Msgf("%v", res.Payload)
			configLogger.Info().Msg("[Restapi:CommitRecordHandlerFunc] Finished")
			span.Finish()
			return &res
		}
		tracing.LogAndTraceErr(apiLogger, span, nil, responses.InvalidCommitType)
		return responses.ErrCommitInvalidCommitType()
	})
	api.RecordRetrieveRecordHandler = record.RetrieveRecordHandlerFunc(func(params record.RetrieveRecordParams) middleware.Responder {
		configLogger.Info().Msg("[Restapi:RecordRetrieveRecordHandler] Entered")
		tracer, closer, err := tracing.SetupGlobalTracer()
		if err != nil {
			configLogger.Err(err).Msg("Unable to initialize Jaeger tracer. Falling back to the NoopTracer")
		} else {
			defer closer.Close()
		}
		span, ctx := opentracing.StartSpanFromContextWithTracer(params.HTTPRequest.Context(), tracer, "RecordRetrieveRecordHandler")
		defer span.Finish()
		conn, err := grpc.Dial(trillianEndpoint, grpc.WithInsecure())
		if err != nil {
			tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
			return responses.ErrRetrieveRecordInternalServerError(err)
		}
		defer conn.Close()
		if ctx == nil {
			ctx = context.Background()
		}
		trillMapClient := trillian.NewTrillianMapClient(conn)
		trillAdminClient := trillian.NewTrillianAdminClient(conn)
		channelMapClientTree, channelErr := getChannelClient(ctx, trillAdminClient, trillMapClient, channelConfigMapID, tracer)
		if channelErr != nil {
			tracing.LogAndTraceErr(apiLogger, span, channelErr, responses.InternalError)
			return responses.ErrRetrieveResourceNotFound()
		}
		channelMapClient := client.MapClient{MapClient: channelMapClientTree}
		channelRevision, err := getCurrentRevision(&channelMapClient, ctx, channelConfigMapID, tracer)
		channelRevision++
		if err != nil {
			tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
			return responses.ErrRetrieveRecordInternalServerError(err)
		}
		channel, err := getChannel(ctx, &channelMapClient, params.ChannelID, tracer)
		if err != nil {
			tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
			return responses.ErrRetrieveRecordInternalServerError(err)
		} else if channel == nil {
			tracing.LogAndTraceErr(apiLogger, span, nil, responses.ChannelNotFound)
			return responses.ErrRetrieveChannelNotFound()
		}
		mapClientTree, err := getChannelClient(ctx, trillAdminClient, trillMapClient, channel.MapID, tracer)
		if err != nil {
			tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
			return responses.ErrRetrieveResourceNotFound()
		}
		mapClient := client.MapClient{MapClient: mapClientTree}
		result, err := getRecord(ctx, &mapClient, params.RecordID, -1, tracer)
		if err != nil {
			tracing.LogAndTraceErr(apiLogger, span, err, responses.InternalError)
			return responses.ErrRetrieveRecordInternalServerError(err)
		} else if result == nil {
			tracing.LogAndTraceErr(apiLogger, span, nil, responses.ResourceNotFound)
			return responses.ErrRetrieveResourceNotFound()
		}
		rec := result.Payload.(map[string]interface{})
		var res = record.RetrieveRecordOK{Payload: rec["recordIDPayload"]}
		configLogger.Debug().Msgf("%v", res.Payload)
		configLogger.Info().Msg("[Restapi:RecordRetrieveRecordHandler] Finished")
		span.Finish()
		return &res
	})
	api.PreServerShutdown = func() {}
	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	logger.SetLogLevelFromEnv()
	return logger.SetupLoggingMiddleware(chiMiddleware.Recoverer(chiMiddleware.RealIP(chiMiddleware.RequestID(handler))))
}

/*func addLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		configLogger.Debug().Msgf("%v","received request:", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}*/
