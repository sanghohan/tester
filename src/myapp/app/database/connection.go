package database

import (
	"database/sql"
	"fmt"
	"github.com/revel/revel"
)

var driver = "mysql"
var connectionString = "dev1:password123@tcp(localhost:3306)/testdb"

func TestSelect() (error,map[string]string) {

	fmt.Printf("db connection config : %s\n", revel.Config.StringDefault("db.connstring", "11111") )

	db, err := sql.Open(driver, revel.Config.StringDefault("db.connstring", connectionString))

	fmt.Println("db connection...")

	if err !=nil {
		fmt.Println("db connection error")
		return err, nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		return err, nil
	}


	rows, err := db.Query("select name, number from testTable" )
	if err != nil {

		fmt.Println("query execute error")
		return err, nil
	}
	defer rows.Close()


	nameMap := make(map[string]string)


	for rows.Next() {
		var name string
		var number string
		rows.Scan(&name, &number)
		nameMap[name] = number
	}

	if err != nil {
		return err, nil
	}

	defer db.Close()

	return nil, nameMap
}