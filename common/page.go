package common

type PageResult struct {
	PageIndex int64       `json:"page_index"`
	PageSize  int64       `json:"page_size"`
	Total     int64       `json:"total"`
	Data      interface{} `json:"data"`
}

func NewPageResult(pageIndex int64, pageSize int64, total int64, data interface{}) *PageResult {
	return &PageResult{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Total:     total,
		Data:      data,
	}
}
