package database

import (
	"cfasuite/internal/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DB interface {
	RegisterModel(model DataModel) error
	CreateTables() error
	CreateTable(sqlCommand string) error
	PrintTables() error
}

type SQLiteDB struct {
	DB *sqlx.DB
}

func (s SQLiteDB) RegisterModel(model DataModel) error {
	err := s.CreateTable(model.Table())
	if err != nil {
		return err
	}
	return nil
}

func (s SQLiteDB) CreateTable(sqlCommand string) error {
	tx := s.DB.MustBegin()
	tx.MustExec(sqlCommand)
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s SQLiteDB) PrintTables() error {
	var tables []string
	query := `
		SELECT name
		FROM sqlite_master
		WHERE type='table'
		ORDER BY name;
	`
	err := s.DB.Select(&tables, query)
	if err != nil {
		return err
	}

	fmt.Println("Tables in the database:")
	for _, table := range tables {
		fmt.Println(table)
	}
	return nil
}

func (s SQLiteDB) CreateTables() error {
	if err := s.RegisterModel(model.User{}); err != nil {
		return err
	}
	if err := s.RegisterModel(model.Session{}); err != nil {
		return err
	}
	if err := s.RegisterModel(model.CFA{}); err != nil {
		return err
	}
	return nil
}
