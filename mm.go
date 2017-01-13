package main

import "github.com/arebus/mm/config"
import "fmt"

func main() {
	cnf := config.LoadConfig("config.json")
	dbConfig := &cnf.Database
	fmt.Println(dbConfig.Name)
	fmt.Println(dbConfig.User)
	fmt.Println(dbConfig.Port)
}
