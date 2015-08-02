package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)


var Datab *sql.DB

var openingStatements = [...]string{
	`CREATE TABLE IF NOT EXISTS student (
    id CHAR(36) NOT NULL PRIMARY KEY,
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL
)`, `CREATE TABLE IF NOT EXISTS note (
    id CHAR(36) NOT NULL PRIMARY KEY,
    student CHAR(36) NOT NULL,
    createTime VARCHAR(50) NOT NULL,
    note TEXT NOT NULL,
    title TEXT NOT NULL
)`, `CREATE TABLE IF NOT EXISTS fieldDef (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    fieldType VARCHAR(100) NOT NULL,
    defaultValue TEXT NOT NULL
)`, `CREATE TABLE IF NOT EXISTS customField (
    fieldId VARCHAR(255) NOT NULL,
    student CHAR(36) NOT NULL,
    fieldValue TEXT NOT NULL
)`, `CREATE TABLE IF NOT EXISTS profile (
    id VARCHAR(255) NOT NULL,
    val TEXT NOT NULL
)`,
}

func InitDB(logger *log.Logger) {
	var err interface{}
	Datab, err = sql.Open("sqlite3", "C:\\Users\\Layne\\Documents\\code\\test\\blub.db")
	if err != nil {
		logger.Fatalln(err)
	}

	for _, stmt := range openingStatements {
		Datab.Exec(stmt)
	}
}
