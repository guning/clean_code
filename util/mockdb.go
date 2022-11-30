package util

type DB interface {
	Get(ptrRow interface{}) (bool, error)
	Find(ptrSlice interface{}) error
	Insert(ptrRow interface{}) (int64, error)
	Update(ptrRow interface{}) (int64, error)
}