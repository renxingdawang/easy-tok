package main

import (
	"context"
	feed "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/feed"
)

// FeedSrvImpl implements the last service interface defined in the IDL.
type FeedSrvImpl struct{}

// GetUserFeed implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetUserFeed(ctx context.Context, request *feed.TokFeedRequest) (resp *feed.TokFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// GetVideoById implements the FeedSrvImpl interface.
func (s *FeedSrvImpl) GetVideoById(ctx context.Context, request *feed.TokFeedRequest) (resp *feed.TokVideo, err error) {
	// TODO: Your code here...
	return
}
