package model

type InsertBody struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Data struct {
		TableId   string `json:"table_id"`
		TableName string `json:"table_name"`
		ViewId    string `json:"view_id"`
		ViewName  string `json:"view_name"`
		Rows      []struct {
			Id          uint   `json:"Id"`
			Title       string `json:"Title"`
			Description string `json:"Description"`
			CreatedAt   string `json:"CreatedAt"`
			UpdatedAt   string `json:"UpdatedAt"`
		} `json:"rows"`
	} `json:"data"`
}
