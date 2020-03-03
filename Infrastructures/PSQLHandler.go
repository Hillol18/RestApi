package Infrastructures

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/nafi/oj-testing/Interfaces"
)

type PSQLHandler struct {
	Conn *sql.DB
}

func (handler *PSQLHandler) Execute(statement string) error {
	handler.Conn.Exec(statement)
	_, err := handler.Conn.Query(statement)

	if err != nil {
		log.Println(err)
	}

	return err
}

func (handler *PSQLHandler) Query(statement string) (Interfaces.IRow, error) {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)

	if err != nil {
		fmt.Println(err)
		return new(PSQLRow), err
	}
	row := new(PSQLRow)
	row.Rows = rows

	return row, nil
}

type PSQLRow struct {
	Rows *sql.Rows
}

func (r PSQLRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r PSQLRow) Next() bool {
	return r.Rows.Next()
}
