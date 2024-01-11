package notion

type PageBody struct {
	Parent struct {
		DatabaseId string `json:"database_id"`
	} `json:"parent"`
	Properties struct {
		Title []struct {
			Text struct {
				Content string `json:"content"`
			} `json:"text"`
		} `json:"title"`
		Description []struct {
			Text struct {
				Content string `json:"content"`
			} `json:"text"`
		} `json:"description"`
	} `json:"properties"`
}
