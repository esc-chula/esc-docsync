package model

type Annotation struct {
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Strikethrough bool   `json:"strikethrough"`
	Underline     bool   `json:"underline"`
	Code          bool   `json:"code"`
	Color         string `json:"color"`
}

type Title struct {
	Type string `json:"type"`
	Text struct {
		Content string      `json:"content"`
		Link    interface{} `json:"link"`
	} `json:"text"`
	Annotations Annotation  `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	Href        interface{} `json:"href"`
}

type RichText struct {
	Type string `json:"type"`
	Text struct {
		Content string      `json:"content"`
		Link    interface{} `json:"link"`
	} `json:"text"`
	Annotations Annotation  `json:"annotations"`
	PlainText   string      `json:"plain_text"`
	Href        interface{} `json:"href"`
}
