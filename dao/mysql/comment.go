package mysql

import (
	"context"
	"github.com/renxingdawang/easy-tok/model"
	"gorm.io/gorm"
)

// NewComment creates a new Comment
func NewComment(ctx context.Context, comment *model.Comment) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		// 1. 新增评论数据
		err := tx.Create(comment).Error
		if err != nil {
			return err
		}

		// 2.改变 video 表中的 comment count
		res := tx.Model(&model.Video{}).Where("ID = ?", comment.VideoID).Update("comment_count", gorm.Expr("comment_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return ErrDatabase
		}

		return nil
	})
	return err
}

// DelComment deletes a comment from the database.
func DelComment(ctx context.Context, commentID int64, vid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		comment := new(model.Comment)
		if err := tx.First(&comment, commentID).Error; err != nil {
			return err
		}

		// 1. 删除评论数据
		// 因为 Comment中包含了gorm.Model所以拥有软删除能力
		// 而tx.Unscoped().Delete()将永久删除记录
		err := tx.Unscoped().Delete(&comment).Error
		if err != nil {
			return err
		}

		// 2.改变 video 表中的 comment count
		res := tx.Model(&model.Video{}).Where("ID = ?", vid).Update("comment_count", gorm.Expr("comment_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return ErrDatabase
		}

		return nil
	})
	return err
}

// GetVideoComments returns a list of video comments.
func GetVideoComments(ctx context.Context, vid int64) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := DB.WithContext(ctx).Model(&model.Comment{}).Where(&model.Comment{VideoID: int(vid)}).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
