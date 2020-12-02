CREATE TABLE IF NOT EXISTS dishes
(
    id        SERIAL PRIMARY KEY,
    title     VARCHAR(128) UNIQUE NOT NULL,
    portion   INT                 NOT NULL,
    calories  INT                 NOT NULL,
    ingestion VARCHAR(256)        NOT NULL
);

CREATE TABLE IF NOT EXISTS programs
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(128) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS programs_dishes
(
    id         SERIAL PRIMARY KEY,
    program_id INT REFERENCES programs,
    dish_id    INT REFERENCES dishes
)