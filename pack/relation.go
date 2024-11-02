package pack

import (
	"context"
	"errors"
	"github.com/renxingdawang/easy-tok/dao/mysql"
	"github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/user"
	"github.com/renxingdawang/easy-tok/model"
	"gorm.io/gorm"
)

// FollowingList pack lists of following info.
func FollowingList(ctx context.Context, vs []*model.Relation, fromID int64) ([]*user.User, error) {
	users := make([]*model.User, 0)
	for _, v := range vs {
		user2, err := mysql.GetUserByID(ctx, int64(v.ToUserID))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		users = append(users, user2)
	}

	return Users(ctx, users, fromID)
}

// FollowerList pack lists of follower info.
func FollowerList(ctx context.Context, vs []*model.Relation, fromID int64) ([]*user.User, error) {
	users := make([]*model.User, 0)
	for _, v := range vs {
		user2, err := mysql.GetUserByID(ctx, int64(v.UserID))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		users = append(users, user2)
	}

	return Users(ctx, users, fromID)
}
