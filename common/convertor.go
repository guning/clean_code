package common

type convertor interface {
	ToIntNoError(obj interface{}) int
}

var CONVERTOR convertor
