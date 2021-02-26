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

	writeLeavesError = false
	getLeavesError = false
	trillMapWriteClient := NewTrillianMapWriteMockClient(conn)
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

	writeLeavesError = true
	getLeavesError = false
	trillMapWriteClient := NewTrillianMapWriteMockClient(conn)
	assert.Equal(t, true, true)
	client := NewClient(trillMapWriteClient, 6453)
	assert.Equal(t, true, true)
	ctx := context.Background()
	tracer, _, _ := tracing.SetupGlobalTracer()
	assert.Error(t, client.Add(ctx, nil, 1, tracer))
}

func TestGet(t *testing.T) {
	getLeavesError = false
	getRootError = false
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := NewTrillianMapMockClient(conn)
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
	getLeavesError = true
	getRootError = false
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := NewTrillianMapMockClient(conn)
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
	getLeavesError = false
	getRootError = true
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := NewTrillianMapMockClient(conn)
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
	getLeavesError = false
	getRootError = false
	verifyRootError = true
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := NewTrillianMapMockClient(conn)
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
	getLeavesError = false
	getRootError = false
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := NewTrillianMapMockClient(conn)
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
	client.GetByRevison(ctx, indexes, 1, tracer)
}
func TestGetByRevisionError(t *testing.T) {
	getLeavesError = true
	getRootError = false
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := NewTrillianMapMockClient(conn)
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
	_, _, err := client.GetByRevison(ctx, indexes, 1, tracer)
	assert.Error(t, err)
}
func TestGetByRevisionErrorRoot(t *testing.T) {
	getLeavesError = false
	getRootError = true
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := NewTrillianMapMockClient(conn)
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
	_, _, err := client.GetByRevison(ctx, indexes, 1, tracer)
	assert.Error(t, err)
}
func TestGetByRevisionErrorVerify(t *testing.T) {
	getLeavesError = false
	getRootError = false
	verifyRootError = true
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := NewTrillianMapMockClient(conn)
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
	_, _, err := client.GetByRevison(ctx, indexes, 1, tracer)
	assert.Error(t, err)
}
func TestGetRevision(t *testing.T) {
	getLeavesError = false
	getRootError = false
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := NewTrillianMapMockClient(conn)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	client.GetCurrentRevison(ctx, 1, tracer)
}
func TestGetRevisionErrorRoot(t *testing.T) {
	getLeavesError = false
	getRootError = true
	verifyRootError = false
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := NewTrillianMapMockClient(conn)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	_, err := client.GetCurrentRevison(ctx, 1, tracer)
	assert.Error(t, err)
}
func TestGetRevisionErrorVerify(t *testing.T) {
	getLeavesError = false
	getRootError = false
	verifyRootError = true
	conn, _ := grpc.Dial("localhost:3000", grpc.WithInsecure())
	defer conn.Close()
	trillMapClient := NewTrillianMapMockClient(conn)
	tracer, _, _ := tracing.SetupGlobalTracer()
	ctx := context.Background()

	verifySignedMapRoot = verifyMock

	mapClientTree := &tclient.MapClient{MapVerifier: &tclient.MapVerifier{}, MapID: 1, Conn: trillMapClient}
	assert.Equal(t, true, true)
	client := MapClient{MapClient: mapClientTree}
	assert.Equal(t, true, true)
	hasher := sha256.New()
	hasher.Write([]byte("1"))
	_, err := client.GetCurrentRevison(ctx, 1, tracer)
	assert.Error(t, err)
}

type trillianMapWriteClientMock struct {
}

func NewTrillianMapWriteMockClient(cc grpc.ClientConnInterface) trillian.TrillianMapWriteClient {
	return &trillianMapWriteClientMock{}
}

func (c *trillianMapWriteClientMock) GetLeavesByRevision(ctx context.Context, in *trillian.GetMapLeavesByRevisionRequest, opts ...grpc.CallOption) (*trillian.MapLeaves, error) {
	out := new(trillian.MapLeaves)
	if getLeavesError {
		return nil, errors.New("Test Error")
	}
	return out, nil
}

func (c *trillianMapWriteClientMock) WriteLeaves(ctx context.Context, in *trillian.WriteMapLeavesRequest, opts ...grpc.CallOption) (*trillian.WriteMapLeavesResponse, error) {
	out := new(trillian.WriteMapLeavesResponse)
	if writeLeavesError {
		return nil, errors.New("Test Error")
	}
	return out, nil
}

type trillianMapMockClient struct {
	cc grpc.ClientConnInterface
}

func NewTrillianMapMockClient(cc grpc.ClientConnInterface) trillian.TrillianMapClient {
	return &trillianMapMockClient{cc}
}

func (c *trillianMapMockClient) GetLeaf(ctx context.Context, in *trillian.GetMapLeafRequest, opts ...grpc.CallOption) (*trillian.GetMapLeafResponse, error) {
	out := new(trillian.GetMapLeafResponse)
	/*err := c.cc.Invoke(ctx, "/trillian.TrillianMap/GetLeaf", in, out, opts...)
	if err != nil {
		return nil, err
	}*/
	return out, nil
}

func (c *trillianMapMockClient) GetLeafByRevision(ctx context.Context, in *trillian.GetMapLeafByRevisionRequest, opts ...grpc.CallOption) (*trillian.GetMapLeafResponse, error) {
	out := new(trillian.GetMapLeafResponse)
	err := c.cc.Invoke(ctx, "/trillian.TrillianMap/GetLeafByRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapMockClient) GetLeaves(ctx context.Context, in *trillian.GetMapLeavesRequest, opts ...grpc.CallOption) (*trillian.GetMapLeavesResponse, error) {
	out := new(trillian.GetMapLeavesResponse)
	if getLeavesError {
		return nil, errors.New("Test Error")
	}
	return out, nil
}

func (c *trillianMapMockClient) GetLeavesByRevision(ctx context.Context, in *trillian.GetMapLeavesByRevisionRequest, opts ...grpc.CallOption) (*trillian.GetMapLeavesResponse, error) {
	out := new(trillian.GetMapLeavesResponse)
	if getLeavesError {
		return nil, errors.New("Test Error")
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *trillianMapMockClient) GetLeavesByRevisionNoProof(ctx context.Context, in *trillian.GetMapLeavesByRevisionRequest, opts ...grpc.CallOption) (*trillian.MapLeaves, error) {
	out := new(trillian.MapLeaves)
	err := c.cc.Invoke(ctx, "/trillian.TrillianMap/GetLeavesByRevisionNoProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapMockClient) GetLastInRangeByRevision(ctx context.Context, in *trillian.GetLastInRangeByRevisionRequest, opts ...grpc.CallOption) (*trillian.MapLeaf, error) {
	out := new(trillian.MapLeaf)
	err := c.cc.Invoke(ctx, "/trillian.TrillianMap/GetLastInRangeByRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *trillianMapMockClient) SetLeaves(ctx context.Context, in *trillian.SetMapLeavesRequest, opts ...grpc.CallOption) (*trillian.SetMapLeavesResponse, error) {
	out := new(trillian.SetMapLeavesResponse)
	err := c.cc.Invoke(ctx, "/trillian.TrillianMap/SetLeaves", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapMockClient) GetSignedMapRoot(ctx context.Context, in *trillian.GetSignedMapRootRequest, opts ...grpc.CallOption) (*trillian.GetSignedMapRootResponse, error) {
	out := new(trillian.GetSignedMapRootResponse)
	out.MapRoot = &trillian.SignedMapRoot{}
	if getRootError {
		return nil, errors.New("Test Error")
	}
	return out, nil
}

func (c *trillianMapMockClient) GetSignedMapRootByRevision(ctx context.Context, in *trillian.GetSignedMapRootByRevisionRequest, opts ...grpc.CallOption) (*trillian.GetSignedMapRootResponse, error) {
	out := new(trillian.GetSignedMapRootResponse)
	err := c.cc.Invoke(ctx, "/trillian.TrillianMap/GetSignedMapRootByRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianMapMockClient) InitMap(ctx context.Context, in *trillian.InitMapRequest, opts ...grpc.CallOption) (*trillian.InitMapResponse, error) {
	out := new(trillian.InitMapResponse)
	err := c.cc.Invoke(ctx, "/trillian.TrillianMap/InitMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func verifyMock(c tclient.MapVerifier, smr *trillian.SignedMapRoot) (*types.MapRootV1, error) {
	if verifyRootError {
		return nil, errors.New("Test Error")
	}
	return &types.MapRootV1{Revision: 1}, nil
}
