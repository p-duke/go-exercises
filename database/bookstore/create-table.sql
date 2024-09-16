/*
Key Changes:
- AUTO_INCREMENT to SERIAL: In PostgreSQL, SERIAL is used to create auto-incrementing integer columns.
- Primary Key Syntax: Primary keys are defined as part of the column definition or as a separate constraint.
- Foreign Key Constraints: Foreign key constraints are defined with CONSTRAINT
 keyword and proper syntax.
- CASCADE: Added CASCADE
 option to DROP TABLE to ensure associated constraints and dependencies are also dropped.
This PostgreSQL script should now create the same schema as your original script in a way that conforms to PostgreSQL's syntax and conventions.

*/
DROP TABLE IF EXISTS books CASCADE;
DROP TABLE IF EXISTS genres CASCADE;
DROP TABLE IF EXISTS customers CASCADE;
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS orders_books CASCADE;

CREATE TABLE genres (
  id SERIAL PRIMARY KEY, 
  name VARCHAR(128) NOT NULL
);

CREATE TABLE books (
  id SERIAL PRIMARY KEY, 
  title VARCHAR(128) NOT NULL, 
  description VARCHAR(255) NOT NULL, 
  author VARCHAR(128) NOT NULL, 
  price VARCHAR(128) NOT NULL, 
  isbn VARCHAR(255) NOT NULL, 
  publication_year DATE NOT NULL, 
  genre_id INT NOT NULL, 
  CONSTRAINT fk_genre 
    FOREIGN KEY(genre_id) REFERENCES genres(id)
);

CREATE TABLE customers (
  id SERIAL PRIMARY KEY, 
  username VARCHAR(128) NOT NULL, 
  email VARCHAR(128) NOT NULL, 
  password VARCHAR(128) NOT NULL, 
  shipping_address VARCHAR(255) NOT NULL
);

CREATE TABLE orders (
  id SERIAL PRIMARY KEY, 
  customer_id INT NOT NULL, 
  order_date DATE NOT NULL, 
  total_price VARCHAR(128) NOT NULL, 
  CONSTRAINT fk_customer 
    FOREIGN KEY(customer_id) REFERENCES customers(id)
);

CREATE TABLE orders_books (
  order_id INT NOT NULL,
  book_id INT NOT NULL,
  CONSTRAINT fk_order 
    FOREIGN KEY(order_id) REFERENCES orders(id),
  CONSTRAINT fk_book 
    FOREIGN KEY(book_id) REFERENCES books(id)
);


