package service

import (
	"context"
	"errors"
	"go-test-basic/common"
	"go-test-basic/model"

	"gorm.io/gorm"
)

func CreateItem(ctx context.Context, item *model.Item) error {
	err := common.GetDB().Create(item).Error
	switch {
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return common.ErrHasExists
	case err != nil:
		return common.ErrInternal
	}
	return nil
}

func GetItem(ctx context.Context, id int) (*model.Item, error) {
	item := &model.Item{}
	err := common.GetDB().First(item, id).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, common.ErrNotFound
	case err != nil:
		return nil, common.ErrInternal
	}

	return item, nil
}

func ListItems(ctx context.Context) ([]*model.Item, error) {
	items := make([]*model.Item, 0)
	err := common.GetDB().Find(&items).Error
	if err != nil {
		return nil, common.ErrInternal
	}
	return items, nil
}

func UpdateItem(ctx context.Context, id int, item *model.Item) error {
	res := common.GetDB().Model(&model.Item{}).Where("id = ?", id).Updates(item)
	switch {
	case res.RowsAffected == 0:
		return common.ErrNotFound
	case res.Error != nil:
		return common.ErrInternal
	}

	return nil
}

func DeleteItem(ctx context.Context, id int) error {
	res := common.GetDB().Delete(&model.Item{}, id)
	switch {
	case res.RowsAffected == 0:
		return common.ErrNotFound
	case res.Error != nil:
		return common.ErrInternal
	}
	return nil
}
