package dao

import (
	"IPFS-Blog-Hugo/internal/database"
	"IPFS-Blog-Hugo/internal/models"
	"errors"
)

type CategoryDao struct {
	models.Category
}

func (dao *CategoryDao) Get(args map[string]any) error {
	db := database.GetMysqlClient()
	return db.Where(args).Take(&dao).Error
}

func (dao *CategoryDao) Add(args map[string]any) error {
	db := database.GetMysqlClient()
	err := dao.Get(args)
	if err == nil {
		return errors.New("数据已存在")
	}
	return db.Create(&dao).Error
}

func (dao *CategoryDao) Update(query map[string]any, args map[string]any) error {
	db := database.GetMysqlClient()
	err := dao.Get(query)
	if err != nil {
		return err
	}
	return db.Model(&dao).Updates(args).Error
}

func (dao *CategoryDao) Delete(args map[string]any) error {
	db := database.GetMysqlClient()
	err := dao.Get(args)
	if err != nil {
		return err
	}
	return db.Model(&dao).Delete(&dao).Error
}

func (dao *CategoryDao) FindById(id string) (*CategoryDao, error) {
	db := database.GetMysqlClient()
	var user CategoryDao
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *CategoryDao) FindByAuthorId(authorId string) ([]CategoryDao, error) {
	db := database.GetMysqlClient()
	var users []CategoryDao
	err := db.Where("author_id = ?", authorId).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dao *CategoryDao) FindAll() ([]CategoryDao, error) {
	db := database.GetMysqlClient()
	var users []CategoryDao
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dao *CategoryDao) GetList() ([]CategoryDao, error) {
	db := database.GetMysqlClient()
	var results []CategoryDao

	rows, err := db.Table(dao.TableName()).Where("is_deleted", false).Rows()
	defer rows.Close()
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var result CategoryDao
		err = db.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}
