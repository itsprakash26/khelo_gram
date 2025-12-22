CREATE TABLE tasks(
    Id SERIAL PRIMARY KEY,
    task_name TEXT NOT NULL,
    task  TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    user_id INTEGER NOT NULL,

    CONSTRAINT fk_user_tasks
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE

);