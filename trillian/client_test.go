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
	"trillian-agent/test"
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

// Test_logAgentConfig contains tests for the agent config logger
func TestAdd(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()

	trillMapWriteClient := test.NewTrillianMapWriteMockClient(conn, false, false)
	assert.Equal(t, true, true)
	client := NewClient(trillMapWriteClient, 6453)
	assert.Equal(t, true, true)
	ctx := context.Background()
	tracer, _, _ := tracing.SetupGlobalTracer()
	client.Add(ctx, nil, 1, tracer)
}

func TestAddError(t *testing.T) {
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()

	trillMapWriteClient := test.NewTrillianMapWriteMockClient(conn, false, true)
	assert.Equal(t, true, true)
	client := NewClient(trillMapWriteClient, 6453)
	assert.Equal(t, true, true)
	ctx := context.Background()
	tracer, _, _ := tracing.SetupGlobalTracer()
	assert.Error(t, client.Add(ctx, nil, 1, tracer))
}

func TestGet(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := test.NewTrillianMapMockClient(conn, false, false, false)
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
func TestGetError(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := test.NewTrillianMapMockClient(conn, true, false, false)
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
func TestGetErrorRoot(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := test.NewTrillianMapMockClient(conn, false, true, false)
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
func TestGetErrorVerify(t *testing.T) {
	verifyRootError = true
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := test.NewTrillianMapMockClient(conn, false, false, false)
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

func TestGetByRevision(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := test.NewTrillianMapMockClient(conn, false, false, false)
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
func TestGetByRevisionError(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := test.NewTrillianMapMockClient(conn, true, false, false)
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
func TestGetByRevisionErrorRoot(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := test.NewTrillianMapMockClient(conn, false, true, false)
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
func TestGetByRevisionErrorVerify(t *testing.T) {
	verifyRootError = true
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := test.NewTrillianMapMockClient(conn, false, false, false)
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
func TestGetRevision(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := test.NewTrillianMapMockClient(conn, false, false, false)
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
func TestGetRevisionErrorRoot(t *testing.T) {
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := test.NewTrillianMapMockClient(conn, false, true, false)
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
func TestGetRevisionErrorVerify(t *testing.T) {
	verifyRootError = true
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := test.NewTrillianMapMockClient(conn, false, false, false)
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
