package model

//harus pakai kapital di awal
type Users struct {
	Id   uint16 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
