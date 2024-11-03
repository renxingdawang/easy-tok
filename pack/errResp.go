package pack

import (
	"github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/comment"
	"github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/favorite"
	"github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/feed"
	"github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/publish"
	"github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/relation"
	"github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/user"
)

const (
	SuccessCode         = 0
	UnknownErrorCode    = -1
	SuccessMessage      = "Success"
	UnknownErrorMessage = "Unknown error occurred"
)

func userRegisterResp(code int32, msg string) *user.TokUserRegisterResponse {
	return &user.TokUserRegisterResponse{StatusCode: code, StatusMsg: &msg}
}
func BuildUserRegisterResp(err error) *user.TokUserRegisterResponse {
	if err == nil {
		return userRegisterResp(SuccessCode, SuccessMessage)
	}
	return userRegisterResp(UnknownErrorCode, err.Error())
}

// BuildUserResp 构建用户响应
func BuildUserResp(err error) *user.TokUserResponse {
	if err == nil {
		return userResp(SuccessCode, SuccessMessage)
	}
	return userResp(UnknownErrorCode, err.Error())
}

func userResp(code int32, msg string) *user.TokUserResponse {
	return &user.TokUserResponse{StatusCode: code, StatusMsg: &msg}
}

// BuildVideoResp 构建视频响应
func BuildVideoResp(err error) *feed.TokFeedResponse {
	if err == nil {
		return videoResp(SuccessCode, SuccessMessage)
	}
	return videoResp(UnknownErrorCode, err.Error())
}

func videoResp(code int32, msg string) *feed.TokFeedResponse {
	return &feed.TokFeedResponse{StatusCode: code, StatusMsg: &msg}
}

// BuildPublishResp 构建发布响应
func BuildPublishResp(err error) *publish.TokPublishActionResponse {
	if err == nil {
		return publishResp(SuccessCode, SuccessMessage)
	}
	return publishResp(UnknownErrorCode, err.Error())
}

func publishResp(code int32, msg string) *publish.TokPublishActionResponse {
	return &publish.TokPublishActionResponse{StatusCode: code, StatusMsg: &msg}
}

// BuildDeleteResp  构建删除响应
func BuildDeleteResp(err error) *publish.TokDeletePublishResponse {
	if err == nil {
		return deleteResp(SuccessCode, SuccessMessage)
	}
	return deleteResp(UnknownErrorCode, err.Error())
}

func deleteResp(code int32, msg string) *publish.TokDeletePublishResponse {
	return &publish.TokDeletePublishResponse{StatusCode: code, StatusMsg: &msg}
}

// BuildPublishListResp 构建发布列表响应
func BuildPublishListResp(err error) *publish.TokPublishListResponse {
	if err == nil {
		return publishListResp(SuccessCode, SuccessMessage)
	}
	return publishListResp(UnknownErrorCode, err.Error())
}

func publishListResp(code int32, msg string) *publish.TokPublishListResponse {
	return &publish.TokPublishListResponse{StatusCode: code, StatusMsg: &msg}
}

// BuildFavoriteActionResp 构建点赞操作响应
func BuildFavoriteActionResp(err error) *favorite.TokFavoriteActionResponse {
	if err == nil {
		return favoriteActionResp(SuccessCode, SuccessMessage)
	}
	return favoriteActionResp(UnknownErrorCode, err.Error())
}

func favoriteActionResp(code int32, msg string) *favorite.TokFavoriteActionResponse {
	return &favorite.TokFavoriteActionResponse{StatusCode: code, StatusMsg: &msg}
}

// BuildFavoriteListResp 构建点赞列表响应
func BuildFavoriteListResp(err error) *favorite.TokFavoriteListResponse {
	if err == nil {
		return favoriteListResp(SuccessCode, SuccessMessage)
	}
	return favoriteListResp(UnknownErrorCode, err.Error())
}

func favoriteListResp(code int32, msg string) *favorite.TokFavoriteListResponse {
	return &favorite.TokFavoriteListResponse{StatusCode: code, StatusMsg: &msg}
}

// BuildRelationActionResp 构建关系操作响应
//func BuildRelationActionResp(err error) *relation.TokRelationActionResponse {
//	if err == nil {
//		return relationActionResp(SuccessCode, SuccessMessage)
//	}
//	return relationActionResp(UnknownErrorCode, err.Error())
//}
//
//func relationActionResp(code int32, msg string) *relation.TokRelationActionResponse {
//	return &relation.TokRelationActionResponse{StatusCode: code, StatusMsg: &msg}
//}

// BuildFollowingListResp 构建关注列表响应
func BuildFollowingListResp(err error) *relation.TokRelationFollowListResponse {
	if err == nil {
		return followingListResp(SuccessCode, SuccessMessage)
	}
	return followingListResp(UnknownErrorCode, err.Error())
}

func followingListResp(code int32, msg string) *relation.TokRelationFollowListResponse {
	return &relation.TokRelationFollowListResponse{StatusCode: code, StatusMsg: &msg}
}

// BuildFollowerListResp 构建粉丝列表响应
func BuildFollowerListResp(err error) *relation.TokRelationFollowerListResponse {
	if err == nil {
		return followerListResp(SuccessCode, SuccessMessage)
	}
	return followerListResp(UnknownErrorCode, err.Error())
}

func followerListResp(code int32, msg string) *relation.TokRelationFollowerListResponse {
	return &relation.TokRelationFollowerListResponse{StatusCode: code, StatusMsg: &msg}
}

// BuildCommentActionResp 构建评论操作响应
func BuildCommentActionResp(err error) *comment.TokCommentActionResponse {
	if err == nil {
		return commentActionResp(SuccessCode, SuccessMessage)
	}
	return commentActionResp(UnknownErrorCode, err.Error())
}

func commentActionResp(code int32, msg string) *comment.TokCommentActionResponse {
	return &comment.TokCommentActionResponse{StatusCode: code, StatusMsg: &msg}
}

// BuildCommentListResp 构建评论列表响应
func BuildCommentListResp(err error) *comment.TokCommentListResponse {
	if err == nil {
		return commentListResp(SuccessCode, SuccessMessage)
	}
	return commentListResp(UnknownErrorCode, err.Error())
}

func commentListResp(code int32, msg string) *comment.TokCommentListResponse {
	return &comment.TokCommentListResponse{StatusCode: code, StatusMsg: &msg}
}
