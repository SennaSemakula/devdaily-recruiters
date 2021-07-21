CREATE TABLE developers (
    id BIGSERIAL CONSTRAINT firstkey PRIMARY KEY,
    first_name varchar(40) NOT NULL,
    age INTEGER NOT NULL,
    CONSTRAINT age check (age >= 0),
    created_at timestamp
);

