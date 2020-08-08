package storage

const (
	// Postgres
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id SERIAL NOT NULL,
		invoice_header_id SERIAL NOT NULL,
		product_id SERIAL NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
	psqlCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES($1, $2) RETURNING id, created_at`

	// MySQL
	mySQLMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`
	mySQLCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES(?, ?)`
)
