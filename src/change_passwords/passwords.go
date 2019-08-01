package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var password string

func main() {
	password := os.Args[1] //getting new password to set

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/mysql") //git
	logger(err)
	defer db.Close()

	rows, err := db.Query("select user from user where user not in ('some_user', 'some_other_user') and host = '%' order by user desc;")
	logger(err)
	defer rows.Close()

	tx, err := db.Begin()
	logger(err)
	defer tx.Rollback()

	var username string

	for rows.Next() {
		err := rows.Scan(&username)
		logger(err)

		// updatestmt, err := tx.Prepare("set password for ? = ?")
		// logger(err)
		// _, err = updatestmt.Exec(username, password)
		_, err = tx.Exec("set password for '" + username + "' = '" + password + "'")
		logger(err)
		log.Println("User " + username + " processed.")
	}
	err = rows.Err()
	logger(err)
	err = tx.Commit()
	logger(err)
}
func logger(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
