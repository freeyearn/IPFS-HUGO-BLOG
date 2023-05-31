package dao

import (
	"IPFS-Blog-Hugo/internal/database"
	"IPFS-Blog-Hugo/internal/models"
	"errors"
)

type ArticleDao struct {
	models.Article
}

func (dao *ArticleDao) Get(args map[string]any) error {
	db := database.GetMysqlClient()
	return db.Where(args).Take(&dao).Error
}

func (dao *ArticleDao) Add(args map[string]any) error {
	db := database.GetMysqlClient()
	err := dao.Get(args)
	if err == nil {
		return errors.New("数据已存在")
	}
	return db.Create(&dao).Error
}

func (dao *ArticleDao) Update(query map[string]any, args map[string]any) error {
	db := database.GetMysqlClient()
	err := dao.Get(query)
	if err != nil {
		return err
	}
	return db.Model(&dao).Updates(args).Error
}

func (dao *ArticleDao) Delete(args map[string]any) error {
	db := database.GetMysqlClient()
	err := dao.Get(args)
	if err != nil {
		return err
	}
	return db.Model(&dao).Delete(&dao).Error
}

func (dao *ArticleDao) FindById(id string) (*ArticleDao, error) {
	db := database.GetMysqlClient()
	var user ArticleDao
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *ArticleDao) FindByAuthorId(authorId string) ([]ArticleDao, error) {
	db := database.GetMysqlClient()
	var users []ArticleDao
	err := db.Where("author_id = ?", authorId).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dao *ArticleDao) FindAll() ([]ArticleDao, error) {
	db := database.GetMysqlClient()
	var users []ArticleDao
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dao *ArticleDao) GetList() ([]ArticleDao, error) {
	db := database.GetMysqlClient()
	var results []ArticleDao

	rows, err := db.Table(dao.TableName()).Where("is_deleted", false).Rows()
	defer rows.Close()
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var result ArticleDao
		err = db.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}
