package storage

import (
	"database/sql"
	"fmt"

	"github.com/chetinchog/go-db/pkg/product"
)

// MySQLProduct used to work with MySQL - product
type MySQLProduct struct {
	db *sql.DB
}

// NewMySQLProduct returns new pointer to MySQLProduct
func NewMySQLProduct(db *sql.DB) *MySQLProduct {
	return &MySQLProduct{db}
}

// Migrate implements interface product.Storage
func (p *MySQLProduct) Migrate() error {
	stmt, err := p.db.Prepare(mySQLMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Product migration Succeeded!")

	return nil
}

// Create implements interface product.Storage
func (p *MySQLProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(mySQLCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = uint(id)
	fmt.Printf("Product creation Succeeded with ID: %v\n", id)

	return nil
}

// GetAll implements interface product.Storage
func (p *MySQLProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(mySQLGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	listProduct := make(product.Models, 0)
	for rows.Next() {
		prod, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		listProduct = append(listProduct, prod)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listProduct, nil
}

// GetByID implements interface product.Storage
func (p *MySQLProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(mySQLGetProdcutByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

// Update implements interface product.Storage
func (p *MySQLProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(mySQLUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("Doesn't exists product with ID: %v", m.ID)
	}

	fmt.Println("Product update Succeeded!")
	return nil
}

// Delete implements interface product.Storage
func (p *MySQLProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(mySQLDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return fmt.Errorf("Doesn't exists product with ID: %v", id)
	}

	fmt.Println("Product delete Succeeded!")
	return nil
}
