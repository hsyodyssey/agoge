package dao

type Testtable struct {
	UserId     int
	UserName   string
	UserAge    int
	UserSchool string
}

func (b *Testtable) TableName() string {
	return "testTable"
}
