package model

type ResRowsBody struct {
	List     []map[string]interface{} `json:"list"`
	PageInfo struct {
		TotalRows   int  `json:"totalRows"`
		Page        int  `json:"page"`
		PageSize    int  `json:"pageSize"`
		IsFirstPage bool `json:"isFirstPage"`
		IsLastPage  bool `json:"isLastPage"`
	} `json:"pageInfo"`
}
