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
	"trillian-agent/mock"
	"trillian-agent/models"
	"trillian-agent/tracing"
	client "trillian-agent/trillian"

	tclient "github.com/google/trillian/client"
	"github.com/google/trillian/types"

	"github.com/google/trillian"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

//TestCreateRecord tests creating a record successfully
func TestCreateRecord(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	add = addMock
	assert.Equal(t, true, true)

	client := client.NewClient(mock.NewTrillianMapWriteMockClient(conn, false, false), 1651)
	recID := "test-record"

	recordDef := &models.RecordDefinition{RecordID: &recID}

	CreateRecord(ctx, client, 2, 1, "test-channel", "CREATE", recordDef, tracer)
}

//TestCreateRecordError tests an error when creating a record
func TestCreateRecordError(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	add = addErrorMock

	client := client.NewClient(mock.NewTrillianMapWriteMockClient(conn, false, true), 1651)
	recID := "test-record"

	recordDef := &models.RecordDefinition{RecordID: &recID}

	assert.Error(t, CreateRecord(ctx, client, 2, 1, "test-channel", "CREATE", recordDef, tracer))
}

//TestGetRecord tests getting a record successfully
func TestGetRecord(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	get = getRecordMock
	trillMapWriteClient := mock.NewTrillianMapMockClient(conn, false, false, false)
	assert.Equal(t, true, true)
	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapWriteClient}
	record, err := GetRecord(ctx, &client.MapClient{mapClientTree}, "test-record", 0, tracer)
	assert.NotNil(t, record)
	assert.Equal(t, record.Revision, int64(2))
	assert.Equal(t, record.PreviousRevision, int64(1))
	assert.Nil(t, err)
}

//TestGetRecordByRevision tests getting a record by revision successfully
func TestGetRecordByRevision(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	getByRevision = getRecordByRevisionMock
	trillMapWriteClient := mock.NewTrillianMapMockClient(conn, false, false, false)
	assert.Equal(t, true, true)
	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapWriteClient}
	channel, err := GetRecord(ctx, &client.MapClient{mapClientTree}, "test-record", 2, tracer)
	assert.NotNil(t, channel)
	//assert.Equal(t, channel.ChannelID, "test-channel")
	//assert.Equal(t, channel.MapID, int64(1654))
	assert.Nil(t, err)
}

//TestGetRecordError tests an error when getting a record
func TestGetRecordError(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	get = getErrorMock
	trillMapWriteClient := mock.NewTrillianMapMockClient(conn, false, false, false)
	assert.Equal(t, true, true)
	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapWriteClient}
	_, err := GetRecord(ctx, &client.MapClient{mapClientTree}, "test-record", 0, tracer)
	assert.Error(t, err)
}

//TestGetRecordNoRes tests no response when getting a record
func TestGetRecordNoRes(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	get = getNoResMock
	trillMapWriteClient := mock.NewTrillianMapMockClient(conn, false, false, false)
	assert.Equal(t, true, true)
	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapWriteClient}
	channel, err := GetRecord(ctx, &client.MapClient{mapClientTree}, "test-record", 0, tracer)
	assert.Nil(t, channel)
	assert.Nil(t, err)
}

//TestGetRecordBadRes tests bad response when getting a record
func TestGetRecordBadRes(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	get = getBadResMock
	trillMapWriteClient := mock.NewTrillianMapMockClient(conn, false, false, false)
	assert.Equal(t, true, true)
	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapWriteClient}
	_, err := GetRecord(ctx, &client.MapClient{mapClientTree}, "test-record", 0, tracer)
	assert.Error(t, err)
}

//TestGetRecordChannelClient tests getting a record client successfully
func TestGetRecordChannelClient(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	get = getMock
	assert.Equal(t, true, true)
	GetChannelClient(ctx, mock.NewTrillianAdminMockClient(conn, false, false), mock.NewTrillianMapMockClient(conn, false, false, false), 651, tracer)
}

//TestGetRecordChannelClientError tests getting and error when getting a record client
func TestGetRecordChannelClientError(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	get = getMock
	assert.Equal(t, true, true)
	_, err := GetChannelClient(ctx, mock.NewTrillianAdminMockClient(conn, false, true), mock.NewTrillianMapMockClient(conn, false, false, false), 651, tracer)
	assert.Error(t, err)
}

func addRecordMock(c *client.Client, ctx context.Context, leaves []*trillian.MapLeaf, revision int64, tracer opentracing.Tracer) error {
	return nil
}

func getRecordMock(c *client.MapClient, ctx context.Context, indexes [][]byte, tracer opentracing.Tracer) ([]*trillian.MapLeafInclusion, *types.MapRootV1, error) {
	auditDefinition := models.AuditDefinition{}
	record := models.Record{
		Revision:         2,
		PreviousRevision: 1,
		AuditDefinition:  auditDefinition,
	}
	leafValue, _ := record.MarshalBinary()
	mapLeaf := trillian.MapLeaf{LeafValue: leafValue}
	mapLeafInclusion := trillian.MapLeafInclusion{Leaf: &mapLeaf}
	inclusions := []*trillian.MapLeafInclusion{
		&mapLeafInclusion,
	}
	return inclusions, &types.MapRootV1{Revision: 1}, nil
}

func getRecordByRevisionMock(c *client.MapClient, ctx context.Context, indexes [][]byte, revision int64, tracer opentracing.Tracer) ([]*trillian.MapLeafInclusion, *types.MapRootV1, error) {
	auditDefinition := models.AuditDefinition{}
	record := models.Record{
		Revision:         2,
		PreviousRevision: 1,
		AuditDefinition:  auditDefinition,
	}
	leafValue, _ := record.MarshalBinary()
	mapLeaf := trillian.MapLeaf{LeafValue: leafValue}
	mapLeafInclusion := trillian.MapLeafInclusion{Leaf: &mapLeaf}
	inclusions := []*trillian.MapLeafInclusion{
		&mapLeafInclusion,
	}
	return inclusions, &types.MapRootV1{Revision: 1}, nil
}

var verifyRootError = false

func verifyMock(c tclient.MapVerifier, smr *trillian.SignedMapRoot) (*types.MapRootV1, error) {
	if verifyRootError {
		return nil, errors.New("Test Error")
	}
	return &types.MapRootV1{Revision: 1}, nil
}
