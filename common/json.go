package common


type json interface {
	MustMarshalToString(obj interface{}) string
}

var JSON json