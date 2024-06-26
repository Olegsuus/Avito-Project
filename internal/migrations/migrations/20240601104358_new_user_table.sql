-- +goose Up
-- +goose StatementBegin
CREATE TABLE Access_levels(
    level INT PRIMARY KEY,
    job_title VARCHAR(64) NOT NULL
);

CREATE TABLE Users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    access_levels INT REFERENCES Access_levels(level),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    token VARCHAR(64) NOT NULL
);

CREATE TABLE Banners (
    id SERIAL PRIMARY KEY NOT NULL,
    title VARCHAR(255) NOT NULL,
    text VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    owner_id INT REFERENCES Users(id),
    f_id INT
);

CREATE TABLE Tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL ,
    banner_id INT REFERENCES Banners(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Tags;
DROP TABLE Banners;
DROP TABLE Users;
DROP TABLE Access_levels;
-- +goose StatementEnd
