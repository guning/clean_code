package process

import "uglydemo/form"

type ProcessInfo struct {
	Name string
	BPMNXml string
	Form []*form.FormLayout
}