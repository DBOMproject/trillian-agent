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
	"errors"
	"testing"
	"trillian-agent/models"
	"trillian-agent/test"
	"trillian-agent/tracing"
	client "trillian-agent/trillian"

	tclient "github.com/google/trillian/client"
	"github.com/google/trillian/types"

	"github.com/google/trillian"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestCreate(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	add = addMock
	assert.Equal(t, true, true)

	CreateChannel(ctx, test.NewTrillianAdminMockClient(conn, false, false), test.NewTrillianMapMockClient(conn, false, false, false), nil, 1, 1, "testChannel", tracer)
}

func TestCreateError(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	add = addErrorMock
	assert.Equal(t, true, true)
	_, err := CreateChannel(ctx, test.NewTrillianAdminMockClient(conn, false, false), test.NewTrillianMapMockClient(conn, false, false, false), nil, 1, 1, "testChannel", tracer)
	assert.Error(t, err)
}

func TestCreateErrorCreateTree(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	add = addMock
	assert.Equal(t, true, true)
	_, err := CreateChannel(ctx, test.NewTrillianAdminMockClient(conn, true, false), test.NewTrillianMapMockClient(conn, false, false, false), nil, 1, 1, "testChannel", tracer)
	assert.Error(t, err)
}

func TestCreateErrorInitMap(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	add = addMock
	assert.Equal(t, true, true)
	_, err := CreateChannel(ctx, test.NewTrillianAdminMockClient(conn, false, false), test.NewTrillianMapMockClient(conn, false, false, true), nil, 1, 1, "testChannel", tracer)
	assert.Error(t, err)
}

func TestGet(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	get = getMock
	trillMapWriteClient := test.NewTrillianMapMockClient(conn, false, false, false)
	assert.Equal(t, true, true)
	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapWriteClient}
	channel, err := GetChannel(ctx, &client.MapClient{mapClientTree}, "test-channel", tracer)
	assert.NotNil(t, channel)
	assert.Equal(t, channel.ChannelID, "test-channel")
	assert.Equal(t, channel.MapID, int64(1654))
	assert.Nil(t, err)
}

func TestGetError(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	get = getErrorMock
	trillMapWriteClient := test.NewTrillianMapMockClient(conn, false, false, false)
	assert.Equal(t, true, true)
	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapWriteClient}
	_, err := GetChannel(ctx, &client.MapClient{mapClientTree}, "test-channel", tracer)
	assert.Error(t, err)
}

func TestGetNoRes(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	get = getNoResMock
	trillMapWriteClient := test.NewTrillianMapMockClient(conn, false, false, false)
	assert.Equal(t, true, true)
	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapWriteClient}
	channel, err := GetChannel(ctx, &client.MapClient{mapClientTree}, "test-channel", tracer)
	assert.Nil(t, channel)
	assert.Nil(t, err)
}

func TestGetBadRes(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	get = getBadResMock
	trillMapWriteClient := test.NewTrillianMapMockClient(conn, false, false, false)
	assert.Equal(t, true, true)
	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapWriteClient}
	_, err := GetChannel(ctx, &client.MapClient{mapClientTree}, "test-channel", tracer)
	assert.Error(t, err)
}

func TestGetChannelClient(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	get = getMock
	assert.Equal(t, true, true)
	GetChannelClient(ctx, test.NewTrillianAdminMockClient(conn, false, false), test.NewTrillianMapMockClient(conn, false, false, false), 651, tracer)
}

func TestGetChannelClientError(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	get = getMock
	assert.Equal(t, true, true)
	_, err := GetChannelClient(ctx, test.NewTrillianAdminMockClient(conn, false, true), test.NewTrillianMapMockClient(conn, false, false, false), 651, tracer)
	assert.Error(t, err)
}

func addMock(c *client.Client, ctx context.Context, leaves []*trillian.MapLeaf, revision int64, tracer opentracing.Tracer) error {
	return nil
}

func addErrorMock(c *client.Client, ctx context.Context, leaves []*trillian.MapLeaf, revision int64, tracer opentracing.Tracer) error {
	return errors.New("Test Error")
}

func getMock(c *client.MapClient, ctx context.Context, indexes [][]byte, tracer opentracing.Tracer) ([]*trillian.MapLeafInclusion, *types.MapRootV1, error) {
	channel := models.Channel{
		ChannelID: "test-channel",
		MapID:     1654,
	}
	leafValue, _ := channel.MarshalBinary()
	mapLeaf := trillian.MapLeaf{LeafValue: leafValue}
	mapLeafInclusion := trillian.MapLeafInclusion{Leaf: &mapLeaf}
	inclusions := []*trillian.MapLeafInclusion{
		&mapLeafInclusion,
	}
	return inclusions, &types.MapRootV1{Revision: 1}, nil
}

func getErrorMock(c *client.MapClient, ctx context.Context, indexes [][]byte, tracer opentracing.Tracer) ([]*trillian.MapLeafInclusion, *types.MapRootV1, error) {
	return nil, nil, errors.New("Test Error")
}

func getNoResMock(c *client.MapClient, ctx context.Context, indexes [][]byte, tracer opentracing.Tracer) ([]*trillian.MapLeafInclusion, *types.MapRootV1, error) {
	mapLeaf := trillian.MapLeaf{LeafValue: []byte("")}
	mapLeafInclusion := trillian.MapLeafInclusion{Leaf: &mapLeaf}
	inclusions := []*trillian.MapLeafInclusion{
		&mapLeafInclusion,
	}
	return inclusions, &types.MapRootV1{Revision: 1}, nil
}

func getBadResMock(c *client.MapClient, ctx context.Context, indexes [][]byte, tracer opentracing.Tracer) ([]*trillian.MapLeafInclusion, *types.MapRootV1, error) {
	mapLeaf := trillian.MapLeaf{LeafValue: []byte("Test")}
	mapLeafInclusion := trillian.MapLeafInclusion{Leaf: &mapLeaf}
	inclusions := []*trillian.MapLeafInclusion{
		&mapLeafInclusion,
	}
	return inclusions, &types.MapRootV1{Revision: 1}, nil
}
