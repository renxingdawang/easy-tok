package pack

import (
	"context"
	"errors"
	"github.com/renxingdawang/easy-tok/dao/mysql"
	"github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/feed"
	"github.com/renxingdawang/easy-tok/model"
	"gorm.io/gorm"
)

// Video pack feed info
func Video(ctx context.Context, v *model.Video, fromID int64) (*feed.TokVideo, error) {
	if v == nil {
		return nil, nil
	}
	user, err := mysql.GetUserByID(ctx, int64(v.AuthorID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	author, err := User(ctx, user, fromID)
	if err != nil {
		return nil, err
	}
	favoriteCount := int64(v.FavoriteCount)
	commentCount := int64(v.CommentCount)

	return &feed.TokVideo{
		Id:            int64(v.ID),
		Author:        author,
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		Title:         v.Title,
	}, nil
}

// Videos pack list of video info
func Videos(ctx context.Context, vs []*model.Video, fromID *int64) ([]*feed.TokVideo, error) {
	videos := make([]*feed.TokVideo, 0)
	for _, v := range vs {
		video2, err := Video(ctx, v, *fromID)
		if err != nil {
			return nil, err
		}

		if video2 != nil {
			flag := false
			if *fromID != 0 {
				results, err := mysql.GetFavoriteRelation(ctx, *fromID, int64(v.ID))
				if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, err
				} else if errors.Is(err, gorm.ErrRecordNotFound) {
					flag = false
				} else if results != nil && results.AuthorID != 0 {
					flag = true
				}
			}
			video2.IsFavorite = flag
			videos = append(videos, video2)
		}
	}
	return videos, nil
}
