package dao

import (
	"IPFS-Blog-Hugo/internal/database"
	"IPFS-Blog-Hugo/internal/models"
	"errors"
)

type UserDao struct {
	models.User
}

func (dao *UserDao) Get(args map[string]any) error {
	db := database.GetMysqlClient()
	return db.Where(args).Take(&dao).Error
}

func (dao *UserDao) Add(args map[string]any) error {
	db := database.GetMysqlClient()
	err := dao.Get(args)
	if err == nil {
		return errors.New("数据已存在")
	}
	return db.Create(&dao).Error
}

func (dao *UserDao) Update(query map[string]any, args map[string]any) error {
	db := database.GetMysqlClient()
	err := dao.Get(query)
	if err != nil {
		return err
	}
	return db.Model(&dao).Updates(args).Error
}

func (dao *UserDao) Delete(args map[string]any) error {
	db := database.GetMysqlClient()
	err := dao.Get(args)
	if err != nil {
		return err
	}
	return db.Model(&dao).Delete(&dao).Error
}

func (dao *UserDao) FindById(id string) (*UserDao, error) {
	db := database.GetMysqlClient()
	var user UserDao
	err := db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDao) FindByWallet(wallet string) (*UserDao, error) {
	db := database.GetMysqlClient()
	var user UserDao
	err := db.Where("wallet = ?", wallet).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDao) FindAll() ([]*UserDao, error) {
	db := database.GetMysqlClient()
	var users []*UserDao
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dao *UserDao) GetList() ([]UserDao, error) {
	db := database.GetMysqlClient()
	results := []UserDao{}

	rows, err := db.Model(dao).Where("is_deleted = 0").Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var result UserDao
		err = db.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}
