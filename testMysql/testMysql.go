package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	query()
}

func query()  {
	db, err := sql.Open("mysql", "root:1a2b3c4d@tcp(127.0.0.1:3306)/ddx_test")
	check(err)
	rows, err := db.Query("SELECT * FROM cus_user")
	defer rows.Close()
	check(err)
	for rows.Next() {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}

	}
}



func check(err error) {
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}