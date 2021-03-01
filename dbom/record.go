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

package dbom

import (
	"context"
	"crypto/sha256"
	"time"
	"trillian-agent/logger"
	"trillian-agent/models"
	"trillian-agent/responses"
	"trillian-agent/tracing"
	client "trillian-agent/trillian"

	"github.com/go-openapi/strfmt"
	"github.com/google/trillian"
	"github.com/opentracing/opentracing-go"
)

var recordLogger = logger.GetLogger("DBoM:Record")

// CreateRecord creates a record and writes it to trillian
func CreateRecord(ctx context.Context, client *client.Client, revision int64, prevRevision int64, channelID string, commitType string, recordDef *models.RecordDefinition, tracer opentracing.Tracer) error {
	recordLogger.Info().Msg("[DBoM:CreateRecord] Entered")
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "DBoM:CreateRecord")

	t := strfmt.DateTime(time.Now())
	audit := models.AuditDefinition{
		ChannelID:  &channelID,
		ResourceID: recordDef.RecordID,
		EventType:  &commitType,
		Payload:    recordDef,
		Timestamp:  &t,
	}
	record := models.Record{
		AuditDefinition:  audit,
		Revision:         revision,
		PreviousRevision: prevRevision,
	}

	leaves := make([]*trillian.MapLeaf, 1)
	hasher := sha256.New()
	hasher.Write([]byte(*record.ResourceID))
	index := hasher.Sum(nil)
	val, err := record.MarshalBinary()
	if err != nil {
		tracing.LogAndTraceErr(recordLogger, span, err, responses.InternalError)
		return err
	}
	leaf := &trillian.MapLeaf{
		Index:     index,
		LeafValue: val,
	}
	leaves[0] = leaf
	recordLogger.Debug().Msgf("Adding asset %v at revision %v", *record.ResourceID, revision)
	err = add(client, ctx, leaves, revision, tracer)
	if err != nil {
		tracing.LogAndTraceErr(recordLogger, span, err, responses.InternalError)
		return err
	}
	recordLogger.Debug().Msgf("Added asset %v at revision %v", *record.ResourceID, revision)

	recordLogger.Info().Msg("[DBoM:CreateRecord] Finished")
	span.Finish()
	return nil
}

// GetRecord gets a record from trillian
func GetRecord(ctx context.Context, client *client.MapClient, recordID string, revision int64, tracer opentracing.Tracer) (*models.Record, error) {
	recordLogger.Info().Msg("[DBoM:GetRecord] Entered")
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "DBoM:GetRecord")

	hasher := sha256.New()
	hasher.Write([]byte(recordID))
	index := hasher.Sum(nil)
	indexes := [][]byte{
		index,
	}
	var inclusions []*trillian.MapLeafInclusion
	var err error
	if revision > 0 {
		inclusions, _, err = getByRevision(client, ctx, indexes, revision, tracer)
	} else {
		inclusions, _, err = get(client, ctx, indexes, tracer)
	}
	if err != nil {
		tracing.LogAndTraceErr(recordLogger, span, err, responses.InternalError)
		return nil, err
	}
	var resString = string(inclusions[0].GetLeaf().GetLeafValue())
	//var idString = string(inclusions[0].GetLeaf().GetIndex())
	//id, err := strconv.ParseInt(idString, 10, 64)
	if len(resString) == 0 {
		tracing.LogAndTraceErr(recordLogger, span, nil, responses.ResourceNotFound)
		return nil, nil
	}

	var result models.Record
	err = result.UnmarshalBinary([]byte(resString))
	if err != nil {
		tracing.LogAndTraceErr(recordLogger, span, err, responses.InternalError)
		return nil, err
	}
	recordLogger.Debug().Msgf("Retrieved asset %v at revision %v", recordID, result.Revision)

	recordLogger.Info().Msg("[DBoM:GetRecord] Finished")
	span.Finish()
	return &result, nil
}
