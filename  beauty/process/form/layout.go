package form

import "demo/ beauty/process/field"

type Layout struct {
	FiledName      string
	FiledCode      string
	FiledType      string
	FieldAttribute map[string]interface{}
}

func (this Layout) ToField() *field.Field {
	return &field.Field{
		Code:      this.FiledCode,
		Name:      this.FiledName,
		Attribute: this.FieldAttribute,
	}
}
