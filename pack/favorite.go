package pack

import (
	"context"
	"github.com/renxingdawang/easy-tok/kitex_gen/easy/tok/feed"
	"github.com/renxingdawang/easy-tok/model"
)

// FavoriteVideos pack favoriteVideos info.
func FavoriteVideos(ctx context.Context, vs []model.Video, uid *int64) ([]*feed.TokVideo, error) {
	videos := make([]*model.Video, 0)
	for _, v := range vs {
		videos = append(videos, &v)
	}

	packVideos, err := Videos(ctx, videos, uid)
	if err != nil {
		return nil, err
	}

	return packVideos, nil
}
