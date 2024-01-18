package model

type ResPageBody struct {
	Object         string `json:"object"`
	Id             string `json:"id"`
	CreatedTime    string `json:"created_time"`
	LastEditedTime string `json:"last_edited_time"`
}

type ResQueryDatabaseBody struct {
	Object  string `json:"object"`
	Results []struct {
		Object         string `json:"object"`
		Id             string `json:"id"`
		CreatedTime    string `json:"created_time"`
		LastEditedTime string `json:"last_edited_time"`
		CreatedBy      struct {
			Object string `json:"object"`
			Id     string `json:"id"`
		} `json:"created_by"`
		LastEditedBy struct {
			Object string `json:"object"`
			Id     string `json:"id"`
		} `json:"last_edited_by"`
		Cover  interface{} `json:"cover"`
		Icon   interface{} `json:"icon"`
		Parent struct {
			Type       string `json:"type"`
			DatabaseId string `json:"database_id"`
		} `json:"parent"`
		Archived   bool        `json:"archived"`
		Properties interface{} `json:"properties"`
		URL        string      `json:"url"`
		PublicURL  interface{} `json:"public_url"`
	} `json:"results"`
}
