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

package trillian

import (
	"context"
	"trillian-agent/logger"
	"trillian-agent/responses"
	"trillian-agent/tracing"

	"github.com/google/trillian"
	tclient "github.com/google/trillian/client"
	_ "github.com/google/trillian/merkle/coniks"  // Register hasher
	_ "github.com/google/trillian/merkle/rfc6962" // Register hasher
	"github.com/google/trillian/types"
	"github.com/opentracing/opentracing-go"
)

var clientLogger = logger.GetLogger("Trillian:Client")

// Client is a type that represents a Trillian Map Write Client
type Client struct {
	client trillian.TrillianMapWriteClient
	mapID  int64
}

// NewClient is a function that creates a new Client
func NewClient(client trillian.TrillianMapWriteClient, mapID int64) *Client {
	return &Client{
		client: client,
		mapID:  mapID,
	}
}

// MapClient is a type that represents a Trillian Map Client
type MapClient struct {
	*tclient.MapClient
}

// Add is a function that adds leaves to a Map
func (c *Client) Add(ctx context.Context, leaves []*trillian.MapLeaf, revision int64, tracer opentracing.Tracer) error {
	clientLogger.Info().Msg("[Client:Add] Entered")
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "Client:Add")
	rqst := &trillian.WriteMapLeavesRequest{
		MapId:          c.mapID,
		Leaves:         leaves,
		ExpectRevision: revision,
	}
	resp, err := c.client.WriteLeaves(ctx, rqst)
	if err != nil {
		tracing.LogAndTraceErr(clientLogger, span, err, responses.InternalError)
		return err
	}

	clientLogger.Debug().Msgf("[Client:Add] %+v", resp)
	clientLogger.Info().Msg("[Client:Add] Finished")
	span.Finish()
	return nil
}

// GetByRevison is a function that gets leaves for a specific revisions from a Map
func (c *MapClient) GetByRevison(ctx context.Context, indexes [][]byte, revison int64, tracer opentracing.Tracer) ([]*trillian.MapLeafInclusion, *types.MapRootV1, error) {
	clientLogger.Info().Msg("[Client:GetByRevison] Entered")
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "Client:GetByRevison")
	clientLogger.Debug().Msg("Get Map Leaves")
	rqst := &trillian.GetMapLeavesByRevisionRequest{
		MapId:    c.MapID,
		Index:    indexes,
		Revision: revison,
	}
	resp, err := c.Conn.GetLeavesByRevision(ctx, rqst)
	if err != nil {
		tracing.LogAndTraceErr(clientLogger, span, err, responses.InternalError)
		return nil, nil, err
	}
	clientLogger.Debug().Msg("Get Map Root")
	rqst2 := &trillian.GetSignedMapRootRequest{
		MapId: c.MapID,
	}
	resp2, err2 := c.Conn.GetSignedMapRoot(ctx, rqst2)
	if err2 != nil {
		tracing.LogAndTraceErr(clientLogger, span, err2, responses.InternalError)
		return nil, nil, err2
	}
	clientLogger.Debug().Msg("Verify Map Root")
	verify, err3 := c.MapVerifier.VerifySignedMapRoot(resp2.GetMapRoot())
	clientLogger.Debug().Msgf("%v", verify)
	if err3 != nil {
		tracing.LogAndTraceErr(clientLogger, span, err3, responses.InternalError)
		return nil, nil, err3
	}

	clientLogger.Debug().Msgf("[Client:GetByRevison] %+v", resp)
	clientLogger.Info().Msg("[Client:GetByRevison] Finished")
	span.Finish()
	return resp.GetMapLeafInclusion(), verify, nil
}

// Get is a function that gets leaves for the latest revision from a Map
func (c *MapClient) Get(ctx context.Context, indexes [][]byte, tracer opentracing.Tracer) ([]*trillian.MapLeafInclusion, *types.MapRootV1, error) {
	clientLogger.Info().Msg("[Client:Get] Entered")
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "Client:Get")
	clientLogger.Debug().Msg("Get Map Leaves")
	rqst := &trillian.GetMapLeavesRequest{
		MapId: c.MapID,
		Index: indexes,
	}
	resp, err := c.Conn.GetLeaves(ctx, rqst)
	if err != nil {
		tracing.LogAndTraceErr(clientLogger, span, err, responses.InternalError)
		return nil, nil, err
	}
	clientLogger.Debug().Msg("Get Map Root")
	rqst2 := &trillian.GetSignedMapRootRequest{
		MapId: c.MapID,
	}
	resp2, err2 := c.Conn.GetSignedMapRoot(ctx, rqst2)
	if err2 != nil {
		tracing.LogAndTraceErr(clientLogger, span, err2, responses.InternalError)
		return nil, nil, err2
	}
	clientLogger.Debug().Msg("Verify Map Root")
	verify, err3 := c.MapVerifier.VerifySignedMapRoot(resp2.GetMapRoot())
	clientLogger.Debug().Msgf("%v", verify)
	if err3 != nil {
		tracing.LogAndTraceErr(clientLogger, span, err3, responses.InternalError)
		return nil, nil, err3
	}

	clientLogger.Debug().Msgf("[Client:Get] %+v", resp)
	clientLogger.Info().Msg("[Client:Get] Finished")
	span.Finish()
	return resp.GetMapLeafInclusion(), verify, nil
}

// GetCurrentRevison gets for the map
func (c *MapClient) GetCurrentRevison(ctx context.Context, mapID int64, tracer opentracing.Tracer) (uint64, error) {
	clientLogger.Info().Msg("[Client:GetCurrentRevison] Entered")
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "Client:GetCurrentRevison")
	clientLogger.Debug().Msg("Get Map Root")
	rqst2 := &trillian.GetSignedMapRootRequest{
		MapId: c.MapID,
	}
	resp2, err2 := c.Conn.GetSignedMapRoot(ctx, rqst2)
	if err2 != nil {
		tracing.LogAndTraceErr(clientLogger, span, err2, responses.InternalError)
		return 0, err2
	}
	clientLogger.Debug().Msg("Verify Map Root")
	verify, err3 := c.MapVerifier.VerifySignedMapRoot(resp2.GetMapRoot())
	clientLogger.Debug().Msgf("%v", verify)
	if err3 != nil {
		tracing.LogAndTraceErr(clientLogger, span, err3, responses.InternalError)
		return 0, err3
	}

	clientLogger.Debug().Msgf("[Client:GetCurrentRevison] %+v", verify.Revision)
	clientLogger.Info().Msg("[Client:GetCurrentRevison] Finished")
	span.Finish()
	return verify.Revision, nil
}
