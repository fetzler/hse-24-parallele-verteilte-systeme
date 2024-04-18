-- table definitions go here
GRANT ALL PRIVILEGES ON DATABASE todo_list TO postgres;

CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY,
    title TEXT
);
