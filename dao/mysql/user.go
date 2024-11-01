package mysql

import (
	"context"
	"github.com/renxingdawang/easy-tok/model"
)

func CreateUser(ctx context.Context, user []*model.User) error {
	return DB.WithContext(ctx).Create(user).Error
}
func GetUserByID(ctx context.Context, userID int64) (*model.User, error) {
	res := new(model.User)

	if err := DB.WithContext(ctx).First(&res, userID).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateUser(ctx context.Context, user *model.User) error {
	return DB.WithContext(ctx).Updates(user).Error
}

// QueryUserByName query list of user info
func QueryUserByName(ctx context.Context, userName string, page, pageSize int64) ([]*model.User, int64, error) {
	db := DB.WithContext(ctx).Model(&model.User{})
	if userName != " " {
		db = db.Where("name LIKE ?", "%"+userName+"%")

	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []*model.User
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}
