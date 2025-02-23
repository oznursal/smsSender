package db

const (
	host               = "127.0.0.1"
	port               = 5432
	user               = "postgres"
	password           = "postgres123"
	dbname             = "postgres"
	create_table_query = "CREATE TABLE IF NOT EXISTS messages (\n  id integer NOT NULL DEFAULT 1,\n  content varchar(160) NOT NULL,\n  phone_number varchar(12) NOT NULL,\n  sending_status boolean NOT NULL DEFAULT false,\n  PRIMARY KEY (id)\n)"
	insert_query       = "INSERT INTO messages(id, content, phone_number, sending_status) VALUES (1, 'Hello', '905555555555', false);\nINSERT INTO messages(id, content, phone_number, sending_status) VALUES (2, 'Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source.', '905555555555', false);\nINSERT INTO messages(id, content, phone_number, sending_status) VALUES (3, '', '905555555555', false);\nINSERT INTO messages(id, content, phone_number, sending_status) VALUES (4, ' ', '905555555555', false);\nINSERT INTO messages(id, content, phone_number, sending_status) VALUES (5, '   ', '905555555555', false);\nINSERT INTO messages(id, content, phone_number, sending_status) VALUES (6, '232', '905555555555', false);"
)
