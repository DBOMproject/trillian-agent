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
	"crypto/sha256"
	"errors"
	"testing"
	"trillian-agent/mock"
	"trillian-agent/tracing"

	"github.com/google/trillian"
	tclient "github.com/google/trillian/client"
	"github.com/google/trillian/types"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

var writeLeavesError = false
var getLeavesError = false
var getRootError = false
var verifyRootError = false

//TestAdd tests successfully adding to the trillian map
func TestAdd(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()

	trillMapWriteClient := mock.NewTrillianMapWriteMockClient(conn, false, false)
	assert.Equal(t, true, true)
	client := NewClient(trillMapWriteClient, 6453)
	assert.Equal(t, true, true)
	ctx := context.Background()
	tracer, _, _ := tracing.SetupGlobalTracer()
	client.Add(ctx, nil, 1, tracer)
}

//TestAddError tests an error while adding to the trillian map
func TestAddError(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()

	trillMapWriteClient := mock.NewTrillianMapWriteMockClient(conn, false, true)
	assert.Equal(t, true, true)
	client := NewClient(trillMapWriteClient, 6453)
	assert.Equal(t, true, true)
	ctx := context.Background()
	tracer, _, _ := tracing.SetupGlobalTracer()
	assert.Error(t, client.Add(ctx, nil, 1, tracer))
}

//TestGet tests successfully getting from the trillian map
func TestGet(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := mock.NewTrillianMapMockClient(conn, false, false, false)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	index := hasher.Sum(nil)
	indexes := [][]byte{
		index,
	}
	client.Get(ctx, indexes, tracer)
}

//TestGetError tests an error while getting from the trillian map
func TestGetError(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := mock.NewTrillianMapMockClient(conn, true, false, false)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	index := hasher.Sum(nil)
	indexes := [][]byte{
		index,
	}
	_, _, err := client.Get(ctx, indexes, tracer)
	assert.Error(t, err)
}

//TestGetErrorRoot tests a root error while getting from the trillian map
func TestGetErrorRoot(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := mock.NewTrillianMapMockClient(conn, false, true, false)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	index := hasher.Sum(nil)
	indexes := [][]byte{
		index,
	}
	_, _, err := client.Get(ctx, indexes, tracer)
	assert.Error(t, err)
}

//TestGetErrorVerify tests a verify error while getting from the trillian map
func TestGetErrorVerify(t *testing.T) {
	verifyRootError = true
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := mock.NewTrillianMapMockClient(conn, false, false, false)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	index := hasher.Sum(nil)
	indexes := [][]byte{
		index,
	}
	_, _, err := client.Get(ctx, indexes, tracer)
	assert.Error(t, err)
}

//TestGetByRevision tests successfully getting from the trillian map by revision
func TestGetByRevision(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := mock.NewTrillianMapMockClient(conn, false, false, false)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	index := hasher.Sum(nil)
	indexes := [][]byte{
		index,
	}
	client.GetByRevision(ctx, indexes, 1, tracer)
}

//TestGetByRevisionError tests an error while getting from the trillian map by revision
func TestGetByRevisionError(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := mock.NewTrillianMapMockClient(conn, true, false, false)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	index := hasher.Sum(nil)
	indexes := [][]byte{
		index,
	}
	_, _, err := client.GetByRevision(ctx, indexes, 1, tracer)
	assert.Error(t, err)
}

//TestGetByRevisionErrorRoot tests a root error while getting from the trillian map by revision
func TestGetByRevisionErrorRoot(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := mock.NewTrillianMapMockClient(conn, false, true, false)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	index := hasher.Sum(nil)
	indexes := [][]byte{
		index,
	}
	_, _, err := client.GetByRevision(ctx, indexes, 1, tracer)
	assert.Error(t, err)
}

//TestGetByRevisionErrorVerify tests a verify error while getting from the trillian map by revision
func TestGetByRevisionErrorVerify(t *testing.T) {
	verifyRootError = true
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := mock.NewTrillianMapMockClient(conn, false, false, false)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	index := hasher.Sum(nil)
	indexes := [][]byte{
		index,
	}
	_, _, err := client.GetByRevision(ctx, indexes, 1, tracer)
	assert.Error(t, err)
}

//TestGetRevision tests successfully getting the revision from the trillian map
func TestGetRevision(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := mock.NewTrillianMapMockClient(conn, false, false, false)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	client.GetCurrentRevision(ctx, 1, tracer)
}

//TestGetRevisionErrorRoot tests a root error while getting the revision from the trillian map
func TestGetRevisionErrorRoot(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := mock.NewTrillianMapMockClient(conn, false, true, false)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	_, err := client.GetCurrentRevision(ctx, 1, tracer)
	assert.Error(t, err)
}

//TestGetRevisionErrorVerify tests a verify error while getting the revision from the trillian map
func TestGetRevisionErrorVerify(t *testing.T) {
	verifyRootError = true
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := mock.NewTrillianMapMockClient(conn, false, false, false)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	_, err := client.GetCurrentRevision(ctx, 1, tracer)
	assert.Error(t, err)
}

func verifyMock(c tclient.MapVerifier, smr *trillian.SignedMapRoot) (*types.MapRootV1, error) {
	if verifyRootError {
		return nil, errors.New("Test Error")
	}
	return &types.MapRootV1{Revision: 1}, nil
}
