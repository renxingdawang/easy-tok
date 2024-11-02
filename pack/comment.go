package pack

import (
	"context"
	"errors"
	"github.com/renxingdawang/easy-tok/dao/mysql"
	"github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/comment"
	"github.com/renxingdawang/easy-tok/model"
	"gorm.io/gorm"
)

// Comments Comment pack Comments info.
func Comments(ctx context.Context, vs []*model.Comment, fromID int64) ([]*comment.Comment, error) {
	comments := make([]*comment.Comment, 0)
	for _, v := range vs {
		user, err := mysql.GetUserByID(ctx, int64(v.UserID))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		packUser, err := User(ctx, user, fromID)
		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment.Comment{
			Id:         int64(v.ID),
			User:       packUser,
			Content:    v.Content,
			CreateDate: v.CreatedAt.Format("01-02"),
		})
	}
	return comments, nil
}
