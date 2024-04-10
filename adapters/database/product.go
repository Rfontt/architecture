package database

import (
	"database/sql"

	"github.com/Rfontt/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}


func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb { db: db }
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID,&product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare(
		`insert into products(id, name, price, status) values(?,?,?,?)`
	)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetID(),
		product.getName(),
		product.getPrice(),
		product.getStatus()
	)

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}