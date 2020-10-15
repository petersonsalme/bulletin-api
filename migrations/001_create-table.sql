CREATE TABLE IF NOT EXISTS bulletins (
    id serial PRIMARY KEY,
    author text NOT NULL,
    content text NOT NULL,
    created_at timestamp with time zone DEFAULT current_timestamp
)