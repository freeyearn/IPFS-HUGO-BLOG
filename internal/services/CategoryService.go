package services

import (
	"IPFS-Blog-Hugo/internal/dao"
	"IPFS-Blog-Hugo/internal/database"
	"IPFS-Blog-Hugo/utils/parser"
)

type CategoryService struct {
	dao.CategoryDao
}

func (m *CategoryService) GetListByPage(p parser.ListParser) ([]CategoryService, int64, error) {
	db := database.GetMysqlClient()
	results := []CategoryService{}
	var count int64 = 0

	rows, err := db.Table(m.TableName()).Where("is_deleted = 0").Where(m).Count(&count).Limit(p.Size).Offset((p.Page - 1) * p.Size).Order(p.Order).Rows()
	defer rows.Close()
	if err != nil {
		return results, 0, err
	}
	for rows.Next() {
		var result CategoryService
		err = db.ScanRows(rows, &result)
		if err != nil {
			return results, 0, err
		}
		results = append(results, result)
	}
	return results, count, nil
}

func (m *CategoryService) GetCategories() ([]CategoryService, error) {
	db := database.GetMysqlClient()
	results := []CategoryService{}
	rows, err := db.Table(m.TableName()).Where("is_deleted = 0").Where(m).Rows()
	defer rows.Close()
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var result CategoryService
		err = db.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}
