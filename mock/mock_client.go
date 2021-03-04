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

package mock

import (
	"context"
	"errors"

	"github.com/google/trillian"
	"google.golang.org/grpc"
)

type trillianMapMockClient struct {
	cc             grpc.ClientConnInterface
	getLeavesError bool
	getRootError   bool
	initMapError   bool
}

func NewTrillianMapMockClient(cc grpc.ClientConnInterface, getLeavesError bool, getRootError bool, initMapError bool) trillian.TrillianMapClient {
	return &trillianMapMockClient{cc, getLeavesError, getRootError, initMapError}
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
	if c.getLeavesError {
		return nil, errors.New("Test Error")
	}
	return out, nil
}

func (c *trillianMapMockClient) GetLeavesByRevision(ctx context.Context, in *trillian.GetMapLeavesByRevisionRequest, opts ...grpc.CallOption) (*trillian.GetMapLeavesResponse, error) {
	out := new(trillian.GetMapLeavesResponse)
	if c.getLeavesError {
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
	if c.getRootError {
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
	if c.initMapError {
		return nil, errors.New("Init Map Error")
	}
	return nil, nil
}
