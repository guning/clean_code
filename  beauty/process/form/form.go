package form

import (
	"demo/ beauty/process/field"
	"demo/common"
	"demo/common/table"
)

func Add(layouts []*Layout) (int64, error) {
	formId, _ := common.DB.Insert(&table.Form{})

	for _, layout := range layouts {
		layout.FiledCode, _ = field.MustSetField(layout.ToField())
		common.DB.Insert(&table.FormLayout{
			FormId:    formId,
			FieldCode: layout.FiledCode,
		})
	}
	return formId, nil
}
