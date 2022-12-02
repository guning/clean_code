package field

import (
	"demo/common"
	"demo/common/table"
)

func MustSetField(field *Field) (code string, err error) {
	if err = field.Validate(); err != nil {
		return
	}
	field.ensureCode()

	exists, _ := common.DB.Get(&table.Field{Code: field.Code})
	if !exists {
		add(field)
	} else {
		update(field)
	}

	return field.Code, nil
}

func update(field *Field) error {
	_, err := common.DB.Where("code = " + field.Code).Update(&table.Field{
		Name:      field.Name,
		Attribute: common.JSON.MustMarshalToString(field.Attribute),
	})
	return err
}

func add(field *Field) error {
	_, err := common.DB.Insert(&table.Field{
		Name:      field.Name,
		Code:      field.Code,
		Attribute: common.JSON.MustMarshalToString(field.Attribute),
	})
	return err
}
