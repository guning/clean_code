package table

type Form struct {
	Id int64
}

type FormLayout struct {
	Id        int64
	FormId    int64
	FieldCode string
}
