CREATE TABLE shortener
(
    id       serial PRIMARY KEY,
    short    VARCHAR(100)  NOT NULL,
    original VARCHAR(200) NOT NULL
)