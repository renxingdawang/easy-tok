package mysql

import (
	"context"
	"github.com/renxingdawang/easy-tok/model"
	"gorm.io/gorm"
)

func GetRelation(ctx context.Context, uid int64, tid int64) (*model.Relation, error) {
	relation := new(model.Relation)

	if err := DB.WithContext(ctx).First(&relation, "user_id = ? and to_user_id = ?", uid, tid).Error; err != nil {
		return nil, err
	}
	return relation, nil
}

// NewRelation creates a new Relation
// uid关注tid，所以uid的关注人数加一，tid的粉丝数加一
func NewRelation(ctx context.Context, uid int64, tid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		// 1. 新增关注数据
		err := tx.Create(&model.Relation{UserID: int(uid), ToUserID: int(tid)}).Error
		if err != nil {
			return err
		}
		// 2.改变 user 表中的 following count
		res := tx.Model(new(model.User)).Where("ID = ?", uid).Update("following_count", gorm.Expr("following_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return ErrDatabase
		}

		// 3.改变 user 表中的 follower count
		res = tx.Model(new(model.User)).Where("ID = ?", tid).Update("follower_count", gorm.Expr("follower_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return err
		}
		return nil
	})
	return err
}

// DisRelation deletes a relation from the database.
func DisRelation(ctx context.Context, uid int64, tid int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		relation := new(model.Relation)
		if err := tx.Where("user_id = ? AND to_user_id=?", uid, tid).First(&relation).Error; err != nil {
			return err
		}

		// 1. 删除关注数据
		err := tx.Unscoped().Delete(&relation).Error
		if err != nil {
			return err
		}
		// 2.改变 user 表中的 following count
		res := tx.Model(new(model.User)).Where("ID = ?", uid).Update("following_count", gorm.Expr("following_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return ErrDatabase
		}

		// 3.改变 user 表中的 follower count
		res = tx.Model(new(model.User)).Where("ID = ?", tid).Update("follower_count", gorm.Expr("follower_count - ?", 1))
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

// FollowingList returns the Following List.
func FollowingList(ctx context.Context, uid int64) ([]*model.Relation, error) {
	var RelationList []*model.Relation
	err := DB.WithContext(ctx).Where("user_id = ?", uid).Find(&RelationList).Error
	if err != nil {
		return nil, err
	}
	return RelationList, nil
}

// FollowerList returns the Follower List.
func FollowerList(ctx context.Context, tid int64) ([]*model.Relation, error) {
	var RelationList []*model.Relation
	err := DB.WithContext(ctx).Where("to_user_id = ?", tid).Find(&RelationList).Error
	if err != nil {
		return nil, err
	}
	return RelationList, nil
}
