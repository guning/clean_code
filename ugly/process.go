package ugly

import (
	"bytes"
	"demo/common"
	"demo/common/table"
	"errors"
	"math/rand"
	"time"
)

type ProcessInfo struct {
	Name string
	BPMNXml string
	Form []*FormLayout
}

const (
	dict       = "abcdefghijklmnopqrstuvwxyz"
	codeLength = 8
)

func AddProcess(name, bpmnxml string, form []*FormLayout) (int64, error) {
	common.DB.TransactionBegin()

	formId, _ := common.DB.Insert(&table.Form{})

	for _, layout := range form {
		// 没有code的情况下必须随机生成code
		if layout.FiledCode == "" {
			rand.Seed(time.Now().Unix())
			b := &bytes.Buffer{}
			for i := 0; i < codeLength; i++ {
				b.Write([]byte{dict[rand.Intn(100)%len(dict)]})
			}
			layout.FiledCode = b.String()
		}
		// 校验字段填值是否符合规则
		switch layout.FiledType {
		case "text":
			if common.CONVERTOR.ToIntNoError(layout.FieldAttribute["length_limit"]) == 0 {
				return 0, errors.New("文本类型不允许限制其长度为0")
			}
		case "option":
			if common.CONVERTOR.ToIntNoError(layout.FieldAttribute["options"]) == 0 {
				return 0, errors.New("选项类型必须配置可选项")
			}
		case "table":
			if common.CONVERTOR.ToIntNoError(layout.FieldAttribute["columns"]) == 0 {
				return 0, errors.New("选项类型必须配置可选项")
			}
		}
		// 新增或更新该字段
		tmp, _ := common.DB.Get(&table.Field{Code: layout.FiledCode})
		if !tmp {
			common.DB.Insert(&table.Field{
				Name:      layout.FiledName,
				Code:      layout.FiledCode,
				Attribute: common.JSON.MustMarshalToString(layout.FieldAttribute),
			})
		} else {
			common.DB.Where("code = " + layout.FiledCode).Update(&table.Field{
				Name:      layout.FiledName,
				Attribute: common.JSON.MustMarshalToString(layout.FieldAttribute),
			})
		}
		// 写入formLayout表
		common.DB.Insert(&table.FormLayout{
			FormId:    formId,
			FieldCode: layout.FiledCode,
		})
	}

	common.DB.Insert(&table.Process{
		Name:    name,
		BPMNXml: bpmnxml,
		FormId:  formId,
	})

	common.DB.TransactionCommit()
	return 0, nil
}
