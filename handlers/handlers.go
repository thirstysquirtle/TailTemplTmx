package handlers

import (
	"bjj_cms/view"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type handlers struct {
	db *sqlx.DB
}

func New() handlers {
	db, err := sqlx.Connect("sqlite3", "file:bjj_cms.db")
	if err != nil {
		log.Fatal(err)
	}
	db.MustExec(`CREATE TABLE IF NOT EXISTS collections 
		(id INTEGER PRIMARY KEY, name TEXT NOT NULL);`)

	db.MustExec(`CREATE TABLE IF NOT EXISTS columns 
		(id INTEGER PRIMARY KEY, name TEXT NOT NULL);`)

	db.MustExec(`CREATE TABLE IF NOT EXISTS collections_columns (
		collection_id INTEGER NOT NULL,
		column_id INTEGER,
		FOREIGN KEY (collection_id) REFERENCES collections (id),
		FOREIGN KEY (column_id) REFERENCES columns (id) 
		);`)

	db.MustExec(`CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		collection_id INTEGER,
		FOREIGN KEY (collection_id) REFERENCES collections (id)
	);`)

	db.MustExec(`CREATE TABLE IF NOT EXISTS item_column_values (
		item_id INTEGER NOT NULL,
		column_id INTEGER NOT NULL,
		column_value NOT NULL,
		FOREIGN KEY (item_id) REFERENCES items (id),
		FOREIGN KEY (column_id) REFERENCES columns (id)
	);`)

	return handlers{db}
}

func (h handlers) ViewCollections(c echo.Context) error {
	var collections []string
	err := h.db.Select(&collections, `SELECT name FROM collections ORDER BY name`)
	if err != nil {
		log.Fatal(err)
	}
	return view.ViewCollections(collections).Render(c.Request().Context(), c.Response().Writer)
}
