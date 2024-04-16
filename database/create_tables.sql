-- table definitions go here
GRANT ALL PRIVILEGES ON DATABASE todo_list TO postgres;
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE tasks (
    task_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    task_name VARCHAR(100) NOT NULL,
    description TEXT,
    due_date DATE,
    completed BOOLEAN NOT NULL DEFAULT FALSE
);
