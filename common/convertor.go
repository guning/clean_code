package common

type convertor interface {
	ToIntNoError(obj interface{}) int
}

type CONVERTOR convertor
