package model

// Row represents data row.
type Row struct {
	Prefix string `db:"prefix"`
	Hash   string `db:"hash"`
	Count  int    `db:"count"`
}
