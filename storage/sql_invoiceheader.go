package storage

const (
	// Postgres
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
	)`
	psqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES($1) RETURNING id, created_at`

	// MySQL
	mySQLMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
	mySQLCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES(?)`
)
