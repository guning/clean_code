package common

type db interface {
	TransactionBegin()
	TransactionCommit()
	Where(cond interface{}) db
	Get(ptrRow interface{}) (bool, error)
	Find(ptrSlice interface{}) error
	Insert(ptrRow interface{}) (int64, error)
	Update(ptrRow interface{}) (int64, error)
}

var DB db