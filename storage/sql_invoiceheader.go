package storage

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
	)`
	psqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES($1) RETURNING id, created_at`
)
