package parser

type ListParser struct {
	Size  int    `json:"Size" form:"Size" binding:"numeric,gt=0"`
	Page  int    `json:"Page" form:"Page" binding:"numeric,gt=0"`
	Order string `json:"Order" form:"Order"`
}
