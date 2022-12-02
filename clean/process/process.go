package process

import (
	"demo/clean/process/form"
	"demo/common"
	"demo/common/table"
	"errors"
)

type Process struct {
	Name        string
	BPMNXml     string
	FormLayouts []*form.Layout
}

func (this Process) Validate() error {
	nodes := common.XML.UnmarshalFromString(this.BPMNXml)
	if len(nodes) == 0 {
		return errors.New("流程未配置")
	}
	if len(this.Name) == 0 {
		return errors.New("未配置名称")
	}
	if len(this.FormLayouts) == 0 {
		return errors.New("未配置表单")
	}
	return nil
}

func Add(process *Process) (int64, error) {
	if err := process.Validate(); err != nil {
		return 0, err
	}
	common.DB.TransactionBegin()

	formId, _ := form.Add(process.FormLayouts)

	common.DB.Insert(&table.Process{
		Name:    process.Name,
		BPMNXml: process.BPMNXml,
		FormId:  formId,
	})

	common.DB.TransactionCommit()
	return 0, nil
}
