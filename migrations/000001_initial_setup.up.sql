CREATE TABLE IF NOT EXISTS tasks(
    id uuid Primary Key,
    assignee VARCHAR(50),
    title VARCHAR(50),
    summary VARCHAR(50),
    deadline timestamp default current_timestamp,
    status VARCHAR(50),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);
