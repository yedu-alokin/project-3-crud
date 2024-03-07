package objects

import "github.com/edgedb/edgedb-go"

type Student struct {
	Id     edgedb.OptionalUUID `edgedb:"id" json:"id" `
	Name   edgedb.OptionalStr  `edgedb:"name" json:"name"`
	Email  string  `edgedb:"email" json:"email" binding:"required"`
	Class  edgedb.OptionalStr  `edgedb:"class" json:"class"`
	School edgedb.OptionalStr  `edgedb:"school" json:"school"`
}

type CreateStudent struct {
	Id edgedb.OptionalUUID `edgedb:"id" json:"id"`
}
