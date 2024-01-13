package model

type InsertBody struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Data struct {
		TableId   string        `json:"table_id"`
		TableName string        `json:"table_name"`
		ViewId    string        `json:"view_id"`
		ViewName  string        `json:"view_name"`
		Rows      []interface{} `json:"rows"`
	} `json:"data"`
}

type UpdateBody struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Data struct {
		TableId      string        `json:"table_id"`
		TableName    string        `json:"table_name"`
		ViewId       string        `json:"view_id"`
		ViewName     string        `json:"view_name"`
		PreviousRows []interface{} `json:"previous_rows"`
		Rows         []interface{} `json:"rows"`
	} `json:"data"`
}
