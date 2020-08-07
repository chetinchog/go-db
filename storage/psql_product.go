package storage

import (
	"database/sql"
	"fmt"

	"github.com/chetinchog/go-db/pkg/product"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

func scanRowProduct(s scanner) (*product.Model, error) {
	prod := &product.Model{}
	observationNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}
	if err := s.Scan(
		&prod.ID,
		&prod.Name,
		&observationNull,
		&prod.Price,
		&prod.CreatedAt,
		&updatedAtNull,
	); err != nil {
		return &product.Model{}, err
	}
	prod.Observations = observationNull.String
	prod.UpdatedAt = updatedAtNull.Time
	return prod, nil
}

// PsqlProduct used to work with postgres - product
type PsqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct returns new pointer to PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Migrate implements interface product.Storage
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
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
func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	).Scan(&m.ID)
	if err != nil {
		return err
	}
	fmt.Println("Product creation Succeeded!")

	return nil
}

// GetAll implements interface product.Storage
func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
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
func (p *PsqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProdcutByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

// Update implements interface product.Storage
func (p *PsqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
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
