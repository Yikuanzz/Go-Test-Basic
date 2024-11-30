package service

import (
	"go-test-basic/common"
	"go-test-basic/model"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Initialize
	common.SetEnv()
	m.Run()
	// Cleanup
	common.TeardownEnv()
}

func TestCreateItem(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		// common.InitTestDBWithContainer(t)
		// common.InitTestDB(t)
		common.InitTestDBTwo(t)
		item := &model.Item{
			Name:        "test name",
			Description: "test desc",
		}

		err := CreateItem(nil, item)
		require.NoError(t, err)
		require.NotEqual(t, 0, item.ID)

		var itemInDB = &model.Item{}
		err = common.GetDB().First(itemInDB, item.ID).Error
		require.NoError(t, err)
		require.Equal(t, item.Name, itemInDB.Name)
		require.Equal(t, item.Description, itemInDB.Description)
	})

	t.Run("duplicate", func(t *testing.T) {
		// common.InitTestDBWithContainer(t)
		// common.InitTestDB(t)
		common.InitTestDBTwo(t)
		item := &model.Item{
			Name:        "test name",
			Description: "test desc",
		}

		err := CreateItem(nil, item)
		require.NoError(t, err)

		err = CreateItem(nil, item)
		require.Error(t, err)
		require.Equal(t, common.ErrHasExists, err)
	})

}

func TestGetItem(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		common.InitTestDB(t)
		item := &model.Item{
			Name:        "test name",
			Description: "test desc",
		}

		err := common.GetDB().Create(item).Error
		require.NoError(t, err)

		itemInDB, err := GetItem(nil, item.ID)
		require.NoError(t, err)
		require.Equal(t, item.Name, itemInDB.Name)
		require.Equal(t, item.Description, itemInDB.Description)
	})

	t.Run("no found", func(t *testing.T) {
		common.InitTestDB(t)
		itemInDB, err := GetItem(nil, 1)
		require.Error(t, err)
		require.Equal(t, common.ErrNotFound, err)
		require.Nil(t, itemInDB)
	})
}

func TestListItems(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		common.InitTestDB(t)
		items := []*model.Item{
			{
				Name: "test name 1",
			},
			{
				Name: "test name 2",
			},
		}
		err := common.GetDB().Create(items).Error
		require.NoError(t, err)

		itemsInDB, err := ListItems(nil)
		require.NoError(t, err)
		require.Equal(t, len(items), len(itemsInDB))
		for i, item := range items {
			require.Equal(t, item.Name, itemsInDB[i].Name)
			require.Equal(t, item.Description, itemsInDB[i].Description)
		}
	})
}

func TestUpdateItem(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		common.InitTestDB(t)
		item := &model.Item{
			Name:        "test name",
			Description: "test desc",
		}
		err := common.GetDB().Create(item).Error
		require.NoError(t, err)

		item.Name = "test name 2"
		item.Description = "test desc 2"
		err = UpdateItem(nil, item.ID, item)
		require.NoError(t, err)
		itemInDB, err := GetItem(nil, item.ID)
		require.NoError(t, err)
		require.Equal(t, item.Name, itemInDB.Name)
		require.Equal(t, item.Description, itemInDB.Description)
	})

	t.Run("not found", func(t *testing.T) {
		common.InitTestDB(t)
		item := &model.Item{
			Name:        "test name",
			Description: "test desc",
		}

		err := UpdateItem(nil, 1, item)
		require.Error(t, err)
		require.Equal(t, common.ErrNotFound, err)
	})
}

func TestDeleteItem(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		common.InitTestDB(t)
		item := &model.Item{
			Name:        "test name",
			Description: "test desc",
		}
		err := common.GetDB().Create(item).Error
		require.NoError(t, err)

		err = DeleteItem(nil, item.ID)
		require.NoError(t, err)

		itemInDB, err := GetItem(nil, item.ID)
		require.Error(t, err)
		require.Equal(t, common.ErrNotFound, err)
		require.Nil(t, itemInDB)
	})

	t.Run("not found", func(t *testing.T) {
		common.InitTestDB(t)
		err := DeleteItem(nil, 1)
		require.Error(t, err)
		require.Equal(t, common.ErrNotFound, err)
	})
}
