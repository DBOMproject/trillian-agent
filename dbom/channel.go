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

	"github.com/golang/protobuf/ptypes"
	tclient "github.com/google/trillian/client"
	"github.com/google/trillian/crypto/keyspb"
	"github.com/google/trillian/crypto/sigpb"
	"github.com/opentracing/opentracing-go"

	"github.com/google/trillian"
)

var channelLogger = logger.GetLogger("DBoM:Channel")

var add = (*client.Client).Add
var getByRevision = (*client.MapClient).GetByRevision
var get = (*client.MapClient).Get
var getCurrentRevision = (*client.MapClient).GetCurrentRevision

// CreateChannel creates a channel and writes it to trillian
func CreateChannel(ctx context.Context, trillAdminClient trillian.TrillianAdminClient, trillMapClient trillian.TrillianMapClient, trillMapWriteClient trillian.TrillianMapWriteClient, revision int64, channelMapID int64, channelID string, tracer opentracing.Tracer) (int64, error) {
	channelLogger.Info().Msg("[DBoM:CreateChannel] Entered")
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "DBoM:CreateChannel")
	ctr := trillian.CreateTreeRequest{Tree: &trillian.Tree{
		TreeId:             channelMapID,
		TreeState:          trillian.TreeState(trillian.TreeState_ACTIVE),
		TreeType:           trillian.TreeType(trillian.TreeType_MAP),
		HashStrategy:       trillian.HashStrategy(trillian.HashStrategy_CONIKS_SHA256),
		HashAlgorithm:      sigpb.DigitallySigned_HashAlgorithm(sigpb.DigitallySigned_SHA256),
		SignatureAlgorithm: sigpb.DigitallySigned_SignatureAlgorithm(sigpb.DigitallySigned_ECDSA),
		DisplayName:        channelID,
		Description:        channelID,
		MaxRootDuration:    ptypes.DurationProto(time.Hour),
	}}
	ctr.KeySpec = &keyspb.Specification{}
	ctr.KeySpec.Params = &keyspb.Specification_EcdsaParams{
		EcdsaParams: &keyspb.Specification_ECDSA{},
	}
	tree, err := trillAdminClient.CreateTree(ctx, &ctr)
	if err != nil {
		tracing.LogAndTraceErr(channelLogger, span, err, responses.InternalError)
		return -1, err
	}
	req := trillian.InitMapRequest{
		MapId: tree.TreeId,
	}
	_, err = trillMapClient.InitMap(ctx, &req)
	if err != nil {
		tracing.LogAndTraceErr(channelLogger, span, err, responses.InternalError)
		return -1, err
	}
	var client = client.NewClient(trillMapWriteClient, channelMapID)
	channel := models.Channel{
		ChannelID: channelID,
		MapID:     tree.TreeId,
	}

	leaves := make([]*trillian.MapLeaf, 1)
	hasher := sha256.New()
	hasher.Write([]byte(channel.ChannelID))
	index := hasher.Sum(nil)
	val, err := channel.MarshalBinary()
	if err != nil {
		tracing.LogAndTraceErr(channelLogger, span, err, responses.InternalError)
		return -1, err
	}
	leaf := &trillian.MapLeaf{
		Index:     index,
		LeafValue: val,
	}
	leaves[0] = leaf
	err = add(client, ctx, leaves, revision, tracer)
	if err != nil {
		tracing.LogAndTraceErr(channelLogger, span, err, responses.InternalError)
		return -1, err
	}

	channelLogger.Info().Msg("[DBoM:CreateChannel] Finished")
	span.Finish()
	return tree.TreeId, nil
}

// GetChannel gets a channel from trillian
func GetChannel(ctx context.Context, client *client.MapClient, channelID string, tracer opentracing.Tracer) (*models.Channel, error) {
	channelLogger.Info().Msg("[DBoM:GetChannel] Entered")
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "DBoM:GetChannel")
	hasher := sha256.New()
	hasher.Write([]byte(channelID))
	index := hasher.Sum(nil)
	indexes := [][]byte{
		index,
	}
	inclusions, mapRoot, err := get(client, ctx, indexes, tracer)
	if err != nil {
		tracing.LogAndTraceErr(channelLogger, span, err, responses.InternalError)
		return nil, err
	}
	var resString = string(inclusions[0].GetLeaf().GetLeafValue())
	//var idString = string(inclusions[0].GetLeaf().GetIndex())
	//id, err := strconv.ParseInt(idString, 10, 64)
	if len(resString) == 0 {
		tracing.LogAndTraceErr(channelLogger, span, nil, responses.ChannelNotFound)
		return nil, nil
	}
	channelLogger.Debug().Msgf("Retrieved channel %v at revision %v", channelID, mapRoot.Revision)
	var result models.Channel
	err = result.UnmarshalBinary([]byte(resString))
	if err != nil {
		tracing.LogAndTraceErr(channelLogger, span, err, responses.InternalError)
		return nil, err
	}

	channelLogger.Info().Msg("[DBoM:GetChannel] Finished")
	span.Finish()
	return &result, nil
}

// GetChannelClient gets a channel client
func GetChannelClient(ctx context.Context, trillAdminClient trillian.TrillianAdminClient, trillMapClient trillian.TrillianMapClient, channelMapID int64, tracer opentracing.Tracer) (*tclient.MapClient, error) {
	channelLogger.Info().Msg("[DBoM:GetChannelClient] Entered")
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "DBoM:GetChannelClient")
	rqst := &trillian.GetTreeRequest{
		TreeId: channelMapID,
	}

	channelTree, treeError := trillAdminClient.GetTree(ctx, rqst)
	if treeError != nil {
		tracing.LogAndTraceErr(channelLogger, span, treeError, responses.InternalError)
		return nil, treeError
	}
	channelLogger.Info().Msg("[DBoM:GetChannelClient] Finished")
	span.Finish()
	return tclient.NewMapClientFromTree(trillMapClient, channelTree)
}
