package notion

import (
	"encoding/json"

	"github.com/esc-chula/esc-docsync/platform/notion/model"
)

func (r *NotionService) GetPropertyValue(property interface{}, dataType string) (any, error) {
	encodedProperty, err := json.Marshal(property.(map[string]interface{}))
	if err != nil {
		return nil, err
	}

	switch dataType {
	case "title":
		title, err := titleValue(encodedProperty)
		if err != nil {
			return nil, err
		}
		return title, nil
	case "rich_text":
		richText, err := richTextValue(encodedProperty)
		if err != nil {
			return nil, err
		}
		return richText, nil
	default:
		return nil, nil
	}
}

func titleValue(property []byte) (string, error) {
	var title model.Title

	if err := json.Unmarshal(property, &title); err != nil {
		return "", err
	}

	return title.PlainText, nil
}

func richTextValue(property []byte) (string, error) {
	var richText model.RichText

	if err := json.Unmarshal(property, &richText); err != nil {
		return "", err
	}

	return richText.PlainText, nil
}
