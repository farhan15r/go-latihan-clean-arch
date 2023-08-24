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