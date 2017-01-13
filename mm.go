package main

import "github.com/arebus/mm/config"
import "fmt"
import "log"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
	cnf := config.LoadConfig("config.json")
	dbConfig := &cnf.Database
	fmt.Println(dbConfig.Name)
	fmt.Println(dbConfig.User)
	fmt.Println(dbConfig.Port)
	
	db, err := sql.Open("mysql", dbConfig.User+":"+dbConfig.Password+"@"+dbConfig.Host+"/")
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	defer db.Close()
}
