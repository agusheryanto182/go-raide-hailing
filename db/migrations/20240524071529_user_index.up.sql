CREATE INDEX IF NOT EXISTS users_id_idx ON users (id);
CREATE INDEX IF NOT EXISTS users_username_idx ON users (username);
CREATE INDEX IF NOT EXISTS users_email_idx ON users (email);
CREATE INDEX IF NOT EXISTS users_role_admin_idx ON users (role) WHERE role = 'admin';
CREATE INDEX IF NOT EXISTS users_role_user_idx ON users (role) WHERE role = 'user';
