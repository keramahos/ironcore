// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package bucketowner

import (
	"context"

	api "github.com/ironcore-dev/ironcore/iri/apis/bucketowner/v1alpha1"
)

type RuntimeService interface {
	ListEvents(context.Context, *api.ListEventsRequest) (*api.ListEventsResponse, error)
	ListBuckets(context.Context, *api.ListBucketsRequest) (*api.ListBucketsResponse, error)
	CreateBucket(context.Context, *api.CreateBucketRequest) (*api.CreateBucketResponse, error)
	ListBucketClasses(ctx context.Context, request *api.ListBucketClassesRequest) (*api.ListBucketClassesResponse, error)
	DeleteBucket(context.Context, *api.DeleteBucketRequest) (*api.DeleteBucketResponse, error)
}
