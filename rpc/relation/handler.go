package main

import (
	"context"
	relation "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/relation"
)

// RelationSrvImpl implements the last service interface defined in the IDL.
type RelationSrvImpl struct{}

// RelationAction implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationAction(ctx context.Context, request *relation.TokRelationActionRequest) (resp *relation.TokRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowList implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationFollowList(ctx context.Context, request *relation.TokRelationFollowListRequest) (resp *relation.TokRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowerList implements the RelationSrvImpl interface.
func (s *RelationSrvImpl) RelationFollowerList(ctx context.Context, request *relation.TokRelationFollowerListRequest) (resp *relation.TokRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}
