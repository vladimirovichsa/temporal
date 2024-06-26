// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by cmd/tools/genfaultinjection. DO NOT EDIT.

package faultinjection

import (
	"context"

	"go.temporal.io/server/common/persistence"
)

type (
	faultInjectionShardStore struct {
		baseStore persistence.ShardStore
		generator faultGenerator
	}
)

func newFaultInjectionShardStore(
	baseStore persistence.ShardStore,
	generator faultGenerator,
) *faultInjectionShardStore {
	return &faultInjectionShardStore{
		baseStore: baseStore,
		generator: generator,
	}
}

func (c *faultInjectionShardStore) AssertShardOwnership(
	ctx context.Context,
	request *persistence.AssertShardOwnershipRequest,
) error {
	return inject0(c.generator.generate("AssertShardOwnership"), func() error {
		return c.baseStore.AssertShardOwnership(ctx, request)
	})
}

func (c *faultInjectionShardStore) Close() {
	c.baseStore.Close()
}

func (c *faultInjectionShardStore) GetClusterName() string {
	return c.baseStore.GetClusterName()
}

func (c *faultInjectionShardStore) GetName() string {
	return c.baseStore.GetName()
}

func (c *faultInjectionShardStore) GetOrCreateShard(
	ctx context.Context,
	request *persistence.InternalGetOrCreateShardRequest,
) (*persistence.InternalGetOrCreateShardResponse, error) {
	return inject1(c.generator.generate("GetOrCreateShard"), func() (*persistence.InternalGetOrCreateShardResponse, error) {
		return c.baseStore.GetOrCreateShard(ctx, request)
	})
}

func (c *faultInjectionShardStore) UpdateShard(
	ctx context.Context,
	request *persistence.InternalUpdateShardRequest,
) error {
	return inject0(c.generator.generate("UpdateShard"), func() error {
		return c.baseStore.UpdateShard(ctx, request)
	})
}
