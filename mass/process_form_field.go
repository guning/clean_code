package mass

import (
	"bytes"
	"demo/common"
	"demo/common/table"
	"errors"
	"math/rand"
	"time"
)

type FeildInfo struct {
	typ       string
	Code      string
	Name      string
	Attribute string
}

func AddField(code, name, typ string, attributes map[string]interface{}) (int64, error) {
	// 没有code的情况下必须随机生成code
	if code == "" {
		rand.Seed(time.Now().Unix())
		b := &bytes.Buffer{}
		for i := 0; i < codeLength; i++ {
			b.Write([]byte{dict[rand.Intn(100)%len(dict)]})
		}
		code = b.String()
	}
	// 校验字段填值是否符合规则
	switch typ {
	case "text":
		if common.CONVERTOR.ToIntNoError(attributes["length_limit"]) == 0 {
			return 0, errors.New("文本类型不允许限制其长度为0")
		}
	case "option":
		if common.CONVERTOR.ToIntNoError(attributes["options"]) == 0 {
			return 0, errors.New("选项类型必须配置可选项")
		}
	case "table":
		if common.CONVERTOR.ToIntNoError(attributes["columns"]) == 0 {
			return 0, errors.New("选项类型必须配置可选项")
		}
	}

	return common.DB.Insert(&table.Field{
		Name:      name,
		Code:      code,
		Attribute: common.JSON.MustMarshalToString(attributes),
	})
}

func UpdateField(code, name, typ string, attributes map[string]interface{}) (int64, error) {
	// 没有code的情况下必须随机生成code
	if code == "" {
		rand.Seed(time.Now().Unix())
		b := &bytes.Buffer{}
		for i := 0; i < codeLength; i++ {
			b.Write([]byte{dict[rand.Intn(100)%len(dict)]})
		}
		code = b.String()
	}
	// 校验字段填值是否符合规则
	switch typ {
	case "text":
		if common.CONVERTOR.ToIntNoError(attributes["length_limit"]) == 0 {
			return 0, errors.New("文本类型不允许限制其长度为0")
		}
	case "option":
		if common.CONVERTOR.ToIntNoError(attributes["options"]) == 0 {
			return 0, errors.New("选项类型必须配置可选项")
		}
	case "table":
		if common.CONVERTOR.ToIntNoError(attributes["columns"]) == 0 {
			return 0, errors.New("选项类型必须配置可选项")
		}
	}

	return common.DB.Insert(&table.Field{
		Name:      name,
		Code:      code,
		Attribute: common.JSON.MustMarshalToString(attributes),
	})
}
