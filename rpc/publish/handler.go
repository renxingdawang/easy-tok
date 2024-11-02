package main

import (
	"context"
	publish "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/publish"
)

// PublishSrvImpl implements the last service interface defined in the IDL.
type PublishSrvImpl struct{}

// PublishAction implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishAction(ctx context.Context, request *publish.TokPublishActionRequest) (resp *publish.TokPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) PublishList(ctx context.Context, request *publish.TokPublishListRequest) (resp *publish.TokPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// DeletePublish implements the PublishSrvImpl interface.
func (s *PublishSrvImpl) DeletePublish(ctx context.Context, request *publish.TokDeletePublishRequest) (resp *publish.TokDeletePublishResponse, err error) {
	// TODO: Your code here...
	return
}
