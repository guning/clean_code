package field

import (
	"bytes"
	"demo/common"
	"errors"
	"math/rand"
	"time"
)

type Field struct {
	typ       string
	Code      string
	Name      string
	Attribute map[string]interface{}
}

func (this *Field) Validate() error {
	switch this.typ {
	case "text":
		if common.CONVERTOR.ToIntNoError(this.Attribute["length_limit"]) == 0 {
			return errors.New("文本类型不允许限制其长度为0")
		}
	case "option":
		if common.CONVERTOR.ToIntNoError(this.Attribute["options"]) == 0 {
			return errors.New("选项类型必须配置可选项")
		}
	case "table":
		if common.CONVERTOR.ToIntNoError(this.Attribute["columns"]) == 0 {
			return errors.New("选项类型必须配置可选项")
		}
	}
	return nil
}

const (
	dict       = "abcdefghijklmnopqrstuvwxyz"
	codeLength = 8
)

func (this *Field) ensureCode() {
	b := &bytes.Buffer{}
	for i := 0; i < codeLength; i++ {
		b.Write([]byte{dict[rand.Intn(100)%len(dict)]})
	}
	this.Code = b.String()
}

func init() {
	rand.Seed(time.Now().Unix())
}
