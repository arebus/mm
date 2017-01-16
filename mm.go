package main

import "github.com/arebus/mm/config"
import "github.com/arebus/mm/server"
import "fmt"
import "log"
import "database/sql"
import "github.com/go-sql-driver/mysql"
import "time"

func main() {
	cnf := config.LoadConfig("config.json")
	dbConfig := &cnf.Database
	fmt.Println(dbConfig.Name)
	fmt.Println(dbConfig.User)
	fmt.Println(dbConfig.Addr)

	var dbc = mysql.Config{
		User:   dbConfig.User,
		Passwd: dbConfig.Password,
		Net:    "tcp",
		Addr:   dbConfig.Addr,
	}

	dbDSN := dbc.FormatDSN()
	fmt.Println(dbDSN)
	db, err := sql.Open("mysql", dbDSN)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	for {
		row, err := db.Query("select variable_name, variable_value from information_schema.GLOBAL_STATUS")
		if err != nil {
			log.Print("Error:", err, " Trying to reconnect...")
			time.Sleep(3 * time.Second)
			continue
		}
		for row.Next() {
			var var_name, var_value string
			err := row.Scan(&var_name, &var_value)
			if err != nil {
				log.Fatal("Error on getting value:", err)
			}
			fmt.Printf("%s | %s\n", var_name, var_value)
		}
	}
	defer db.Close()
}
