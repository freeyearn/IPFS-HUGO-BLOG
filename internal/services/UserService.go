package services

import (
	"IPFS-Blog-Hugo/internal/dao"
	"IPFS-Blog-Hugo/internal/database"
	"IPFS-Blog-Hugo/utils/parser"
)

type UserService struct {
	dao.UserDao
}

func (m *UserService) GetListByPage(p parser.ListParser) ([]UserService, int64, error) {
	db := database.GetMysqlClient()
	results := []UserService{}
	var count int64 = 0

	rows, err := db.Table(m.TableName()).Where("is_deleted = 0").Count(&count).Limit(p.Size).Offset((p.Page - 1) * p.Size).Order(p.Order).Rows()
	defer rows.Close()
	if err != nil {
		return results, 0, err
	}
	for rows.Next() {
		var result UserService
		err = db.ScanRows(rows, &result)
		if err != nil {
			return results, 0, err
		}
		results = append(results, result)
	}
	return results, count, nil
}
