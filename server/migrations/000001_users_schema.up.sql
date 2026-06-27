CREATE SCHEMA IF NOT EXISTS users_schema; 

SET search_path TO users_schema;


CREATE TABLE IF NOT EXISTS users (
    id          SERIAL          PRIMARY KEY
    ,user_name  VARCHAR(40)     NOT NULL
    ,email      VARCHAR(150)    NOT NULL
    ,password   VARCHAR(100)    NOT NULL
);