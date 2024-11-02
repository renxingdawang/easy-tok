package pack

import (
	"context"
	"errors"
	"github.com/renxingdawang/easy-tok/dao/mysql"
	"github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/user"
	"github.com/renxingdawang/easy-tok/model"
	"gorm.io/gorm"
)

func User(ctx context.Context, u *model.User, fromID int64) (*user.User, error) {
	if u == nil {
		return &user.User{
			Name: "已注销用户",
		}, nil
	}

	followCount := int64(u.FollowingCount)
	followerCount := int64(u.FollowerCount)

	// true->fromID已关注u.ID，false-fromID未关注u.ID
	isFollow := false
	relation, err := mysql.GetRelation(ctx, fromID, int64(u.ID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if relation != nil {
		isFollow = true
	}
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.Username,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      isFollow,
	}, nil
}

func Users(ctx context.Context, us []*model.User, fromID int64) ([]*user.User, error) {
	users := make([]*user.User, 0)
	for _, u := range us {
		user2, err := User(ctx, u, fromID)
		if err != nil {
			return nil, err
		}

		if user2 != nil {
			users = append(users, user2)
		}
	}
	return users, nil
}
