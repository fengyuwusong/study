package main

import (
	"fmt"
	"geeorm"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, err := geeorm.NewEngine("sqlite3", "gee.db")
	if err != nil {
		fmt.Printf("new engine error, err: %v\n", err)
		return
	}
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User (Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User (Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User (`Name`) VALUES (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
