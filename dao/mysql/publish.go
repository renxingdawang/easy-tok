package mysql

import (
	"context"
	"errors"
	"fmt"
	"github.com/renxingdawang/easy-tok/model"
	"gorm.io/gorm"
)

// CreateVideo creates a new video
func CreateVideo(ctx context.Context, video *model.Video) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(video).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// PublishList returns a list of videos with AuthorID.
func PublishList(ctx context.Context, authorId int64) ([]*model.Video, error) {
	var pubList []*model.Video
	err := DB.WithContext(ctx).Model(&model.Video{}).Where(&model.Video{AuthorID: int(authorId)}).Find(&pubList).Error
	if err != nil {
		return nil, err
	}
	return pubList, nil
}

func DeleteVideo(ctx context.Context, videoId int64, authorId int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var video model.Video
		if err := tx.Where("id = ? AND author_id = ?", videoId, authorId).First(&video).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("video not found or does not belong to the author")
			}
			return err
		}

		// Delete the video
		if err := tx.Delete(&video).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
