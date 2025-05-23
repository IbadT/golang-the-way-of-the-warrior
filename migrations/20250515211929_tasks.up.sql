CREATE TABLE tasks(
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    is_done BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL
);

ALTER TABLE tasks ADD COLUMN user_id UUID REFERENCES users(id) ON DELETE CASCADE;