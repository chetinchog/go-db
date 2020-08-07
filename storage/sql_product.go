package storage

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(100) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
	)`
	psqlCreateProduct  = `INSERT INTO products(name, observations, price, created_at) VALUES($1, $2, $3, $4) RETURNING id`
	psqlGetAllProduct  = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	psqlGetProdcutByID = psqlGetAllProduct + ` WHERE id = $1`
	psqlUpdateProduct  = `UPDATE products SET name = $1, observations = $2, price = $3, updated_at = $4 WHERE id = $5`
)
