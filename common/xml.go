package common

type xml interface {
	UnmarshalFromString(xml string) map[string]interface{}
}

var XML xml
