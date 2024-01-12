package model

type CreatedPageResponse struct {
	Object         string `json:"object"`
	Id             string `json:"id"`
	CreatedTime    string `json:"created_time"`
	LastEditedTime string `json:"last_edited_time"`
}
