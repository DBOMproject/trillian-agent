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

type trillianAdminMockClient struct {
	cc              grpc.ClientConnInterface
	createTreeError bool
	getTreeError    bool
}

func NewTrillianAdminMockClient(cc grpc.ClientConnInterface, createTreeError bool, getTreeError bool) trillian.TrillianAdminClient {
	return &trillianAdminMockClient{cc, createTreeError, getTreeError}
}

func (c *trillianAdminMockClient) ListTrees(ctx context.Context, in *trillian.ListTreesRequest, opts ...grpc.CallOption) (*trillian.ListTreesResponse, error) {
	out := new(trillian.ListTreesResponse)
	err := c.cc.Invoke(ctx, "/trillian.TrillianAdmin/ListTrees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianAdminMockClient) GetTree(ctx context.Context, in *trillian.GetTreeRequest, opts ...grpc.CallOption) (*trillian.Tree, error) {
	if c.getTreeError {
		return nil, errors.New("Get Tree Error")
	}
	return &trillian.Tree{TreeId: 651}, nil
}

func (c *trillianAdminMockClient) CreateTree(ctx context.Context, in *trillian.CreateTreeRequest, opts ...grpc.CallOption) (*trillian.Tree, error) {
	if c.createTreeError {
		return nil, errors.New("Create Tree Error")
	}
	return &trillian.Tree{TreeId: 3513}, nil
}

func (c *trillianAdminMockClient) UpdateTree(ctx context.Context, in *trillian.UpdateTreeRequest, opts ...grpc.CallOption) (*trillian.Tree, error) {
	out := new(trillian.Tree)
	err := c.cc.Invoke(ctx, "/trillian.TrillianAdmin/UpdateTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianAdminMockClient) DeleteTree(ctx context.Context, in *trillian.DeleteTreeRequest, opts ...grpc.CallOption) (*trillian.Tree, error) {
	out := new(trillian.Tree)
	err := c.cc.Invoke(ctx, "/trillian.TrillianAdmin/DeleteTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trillianAdminMockClient) UndeleteTree(ctx context.Context, in *trillian.UndeleteTreeRequest, opts ...grpc.CallOption) (*trillian.Tree, error) {
	out := new(trillian.Tree)
	err := c.cc.Invoke(ctx, "/trillian.TrillianAdmin/UndeleteTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
