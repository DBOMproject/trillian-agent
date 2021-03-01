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

package test

import (
	"context"
	"errors"

	"github.com/google/trillian"
	"google.golang.org/grpc"
)

type trillianMapWriteClientMock struct {
	getLeavesError   bool
	writeLeavesError bool
}

func NewTrillianMapWriteMockClient(cc grpc.ClientConnInterface) trillian.TrillianMapWriteClient {
	return &trillianMapWriteClientMock{}
}

func (c *trillianMapWriteClientMock) GetLeavesByRevision(ctx context.Context, in *trillian.GetMapLeavesByRevisionRequest, opts ...grpc.CallOption) (*trillian.MapLeaves, error) {
	out := new(trillian.MapLeaves)
	if c.getLeavesError {
		return nil, errors.New("Test Error")
	}
	return out, nil
}

func (c *trillianMapWriteClientMock) WriteLeaves(ctx context.Context, in *trillian.WriteMapLeavesRequest, opts ...grpc.CallOption) (*trillian.WriteMapLeavesResponse, error) {
	out := new(trillian.WriteMapLeavesResponse)
	if c.writeLeavesError {
		return nil, errors.New("Test Error")
	}
	return out, nil
}
