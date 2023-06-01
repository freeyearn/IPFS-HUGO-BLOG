package services

import (
	"IPFS-Blog-Hugo/internal/dao"
	"IPFS-Blog-Hugo/internal/database"
	"IPFS-Blog-Hugo/utils/message"
	"IPFS-Blog-Hugo/utils/parser"
	"database/sql"
	"encoding/json"
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
		if result.Category != "" {
			var category []string
			err := json.Unmarshal([]byte(result.Category), &category)
			if err != nil {
				message.PrintErr("category parser err: ", err)
				return nil, 0, err
			}
			categories := make([]string, len(category))
			for idx, value := range category {
				categoryService := CategoryService{}
				err := categoryService.Get(map[string]any{
					"category_id": value,
				})
				if err != nil {
					message.PrintErr("category query err: id:", value, " err:", err)
				}
				categories[idx] = categoryService.GetCategoryName()
			}
			res, err := json.Marshal(categories)
			if err != nil {
				message.PrintErr("category marshal err: ", err)
				return nil, 0, err
			}
			result.Category = string(res)
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

		if result.Category != "" {
			var category []string
			err := json.Unmarshal([]byte(result.Category), &category)
			if err != nil {
				message.PrintErr("category parser err: ", err)
				return nil, err
			}
			categories := make([]string, len(category))
			for idx, value := range category {
				categoryService := CategoryService{}
				err := categoryService.Get(map[string]any{
					"category_id": value,
				})
				if err != nil {
					message.PrintErr("category query err: id:", value, " err:", err)
				}
				categories[idx] = categoryService.GetCategoryName()
			}
			res, err := json.Marshal(category)
			if err != nil {
				message.PrintErr("category marshal err: ", err)
				return nil, err
			}
			result.Category = string(res)
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
