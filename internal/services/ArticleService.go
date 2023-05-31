package services

import (
	"IPFS-Blog-Hugo/internal/dao"
	"IPFS-Blog-Hugo/internal/database"
	"IPFS-Blog-Hugo/utils/parser"
	"database/sql"
)

type ArticleService struct {
	dao.ArticleDao
}

func (m *ArticleService) GetListByPage(p parser.ListParser) ([]ArticleService, int64, error) {
	db := database.GetMysqlClient()
	results := []ArticleService{}
	var count int64 = 0
	var rows *sql.Rows
	var err error

	query := db.Table(m.TableName()).Where("is_deleted = 0")
	// filter by title, if title is not empty then fuzzy query title.
	if m.GetTitle() != "" {
		query = query.Where("title like ?", "%"+m.GetTitle()+"%")
	}
	if m.GetCategory() != "" {
		query = query.Where("category like ?", "%"+m.GetCategory()+"%")
	}
	if m.GetAuthorId() != "" {
		query = query.Where("author_id = ?", m.GetAuthorId())
	}
	rows, err = query.Count(&count).Limit(p.Size).Offset((p.Page-1)*p.Size).Order(p.Order, true).Rows()
	defer rows.Close()
	if err != nil {
		return results, 0, err
	}
	for rows.Next() {
		var result ArticleService
		err = db.ScanRows(rows, &result)
		if err != nil {
			return results, 0, err
		}
		results = append(results, result)
	}
	return results, count, nil
}

func (m *ArticleService) GetList() ([]ArticleService, error) {
	db := database.GetMysqlClient()
	results := []ArticleService{}
	rows, err := db.Table(m.TableName()).Where("is_deleted = 0").Where(m).Rows()
	defer rows.Close()
	if err != nil {
		return results, err
	}
	for rows.Next() {
		var result ArticleService
		err = db.ScanRows(rows, &result)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return results, nil
}

// Count article number.
func (m *ArticleService) Count() int64 {
	db := database.GetMysqlClient()
	var count int64 = 0
	db.Table(m.TableName()).Where("is_deleted = 0").Where(m).Count(&count)
	return count
}
