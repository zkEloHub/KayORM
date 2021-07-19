package main

import (
	"KayORM"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := KayORM.NewEngine("sqlite3", "zk.db")
	defer engine.Close()

	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()

	result, _ := s.Raw("INSERT INTO User(`Name`) VALUES (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}