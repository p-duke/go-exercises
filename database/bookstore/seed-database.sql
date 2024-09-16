-- Insert data into genres
INSERT INTO genres (name) VALUES
('Science Fiction'),
('Fantasy'),
('Mystery'),
('Romance'),
('Non-Fiction');

-- Insert data into books
INSERT INTO books (title, description, author, price, isbn, publication_year, genre_id) VALUES
('Dune', 'A science fiction novel', 'Frank Herbert', '9.99', '9780441013593', '1965-08-01', 1),
('Neuromancer', 'A science fiction novel', 'William Gibson', '8.99', '9780441569595', '1984-07-01', 1),
('Foundation', 'A science fiction novel', 'Isaac Asimov', '7.99', '9780553293357', '1951-01-01', 1),
('The Hobbit', 'A fantasy novel', 'J.R.R. Tolkien', '10.99', '9780547928227', '1937-09-21', 2),
('Harry Potter and the Sorcerers Stone', 'A fantasy novel', 'J.K. Rowling', '12.99', '9780439708180', '1997-06-26', 2),
('The Name of the Wind', 'A fantasy novel', 'Patrick Rothfuss', '11.99', '9780756404741', '2007-03-27', 2),
('The Da Vinci Code', 'A mystery thriller novel', 'Dan Brown', '12.99', '9780307474278', '2003-03-18', 3),
('Gone Girl', 'A mystery thriller novel', 'Gillian Flynn', '9.99', '9780307588371', '2012-06-05', 3),
('The Girl with the Dragon Tattoo', 'A mystery novel', 'Stieg Larsson', '10.99', '9780307454546', '2005-08-01', 3),
('Pride and Prejudice', 'A romance novel', 'Jane Austen', '8.99', '9781503290563', '1813-01-28', 4),
('Outlander', 'A romance novel', 'Diana Gabaldon', '11.99', '9780440212560', '1991-06-01', 4),
('Me Before You', 'A romance novel', 'Jojo Moyes', '9.99', '9780143124542', '2012-01-05', 4),
('Sapiens', 'A non-fiction book', 'Yuval Noah Harari', '14.99', '9780062316097', '2011-01-01', 5),
('Educated', 'A non-fiction memoir', 'Tara Westover', '13.99', '9780399590504', '2018-02-20', 5),
('The Immortal Life of Henrietta Lacks', 'A non-fiction book', 'Rebecca Skloot', '12.99', '9781400052189', '2010-02-02', 5);

-- Insert data into customers
INSERT INTO customers (username, email, password, shipping_address) VALUES
('john_doe', 'john@example.com', 'password123', '123 Main St, Anytown, USA'),
('jane_smith', 'jane@example.com', 'securepassword', '456 Elm St, Othertown, USA'),
('alice_jones', 'alice@example.com', 'alicepassword', '789 Oak St, Sometown, USA'),
('bob_brown', 'bob@example.com', 'bobbypassword', '101 Pine St, Randomtown, USA'),
('carol_white', 'carol@example.com', 'carolspassword', '202 Maple St, Anycity, USA');

-- Insert data into orders
INSERT INTO orders (customer_id, order_date, total_price) VALUES
(1, '2024-09-01', '22.98'),
(1, '2024-09-02', '14.99'),
(2, '2024-09-03', '10.99'),
(3, '2024-09-04', '12.99'),
(4, '2024-09-05', '17.98'),
(5, '2024-09-06', '20.98');

-- Insert data into orders_books
INSERT INTO orders_books (order_id, book_id) VALUES
(1, 1),
(1, 2),
(2, 5),
(3, 2),
(4, 3),
(4, 9),
(5, 4),
(5, 11),
(6, 13),
(6, 14);
