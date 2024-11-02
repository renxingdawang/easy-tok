package mysql

import (
	"context"
	"github.com/renxingdawang/easy-tok/model"
	"time"
)

// MGetVideos multiple get list of videos info
func MGetVideos(ctx context.Context, limit int, latestTime *int64) ([]*model.Video, error) {
	videos := make([]*model.Video, 0)

	if latestTime == nil || *latestTime == 0 {
		currentTime := int64(time.Now().UnixMilli())
		latestTime = &currentTime
	}
	conn := DB.WithContext(ctx)

	if err := conn.Limit(limit).Order("update_time desc").Find(&videos, "update_time < ?", time.UnixMilli(*latestTime)).Error; err != nil {
		return nil, err
	}
	return videos, nil
}
