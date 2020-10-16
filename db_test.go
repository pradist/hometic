package main

import (
	"database/sql"
	"testing"
)

type mockDB struct {
	q string
}

func (md *mockDB) Exec(query string, args ...interface{}) (sql.Result, error)  {
	md.q = query
	return nil, nil
}

func TestInsertPairDevice(t *testing.T) {

	mock := &mockDB{}
	insert := NewCreatePairDevice(mock)
	insert(Pair{DeviceID: 123, UserID: 456})

	//if mock.q != "XXX" {
	//	t.Error("Expect xxx but got", mock.q)
	//}
}
