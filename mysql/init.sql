CREATE DATABASE IF NOT EXISTS wallet;

USE wallet;

CREATE TABLE IF NOT EXISTS clients (
    id varchar(255),
    name varchar(255),
    email varchar(255),
    created_at date
);

CREATE TABLE IF NOT EXISTS accounts (
    id varchar(255),
    client_id varchar(255),
    balance float,
    created_at date
);

CREATE TABLE IF NOT EXISTS transactions (
    id varchar(255),
    account_id_from varchar(255),
    account_id_to varchar(255),
    amount float,
    created_at date
);


INSERT INTO clients (id, name, email, created_at) VALUES
('123x431ffg', 'John', 'john@email.com', CURDATE()),
('123x4141fa', 'Jane', 'jane@email.com', CURDATE());

INSERT INTO accounts (id, client_id, balance, created_at) VALUES
('acc1', '123x431ffg', 1000.00, CURDATE()),
('acc2', '123x4141fa', 500.00, CURDATE());
