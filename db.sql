-- Active: 1692371058716@@127.0.0.1@5432@db_clean_arch_hicoll
CREATE TABLE students (
    id BIGSERIAL PRIMARY KEY,
    fullname VARCHAR,
    address VARCHAR,
    birthdate DATE,
    class VARCHAR,
    batch INT,
    school_name VARCHAR
)

SELECT * FROM students

DROP TABLE IF EXISTS students

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    address TEXT NULL
)

DROP TABLE IF EXISTS users

SELECT * FROM users 

CREATE TABLE authentications (
    refresh_token TEXT NOT NULL
)

DROP TABLE IF EXISTS authentications

SELECT * FROM authentications

TRUNCATE TABLE authentications