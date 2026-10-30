CREATE TABLE tasks (
    title VARCHAR PRIMARY KEY,
    description VARCHAR NOT NULL, 
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL,
    complete_at TIMESTAMPTZ
);