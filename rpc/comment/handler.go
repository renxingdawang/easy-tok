package main

import (
	"context"
	comment "github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/comment"
)

// CommentSrvImpl implements the last service interface defined in the IDL.
type CommentSrvImpl struct{}

// CommentAction implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentAction(ctx context.Context, request *comment.TokCommentActionRequest) (resp *comment.TokCommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the CommentSrvImpl interface.
func (s *CommentSrvImpl) CommentList(ctx context.Context, request *comment.TokCommentActionRequest) (resp *comment.TokCommentListResponse, err error) {
	// TODO: Your code here...
	return
}
