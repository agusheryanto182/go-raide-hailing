CREATE INDEX IF NOT EXISTS users_id ON users (user_id);
CREATE INDEX IF NOT EXISTS users_username ON users (username);
CREATE INDEX IF NOT EXISTS users_email ON users (email);
CREATE INDEX IF NOT EXISTS users_role_admin ON users (role) WHERE role = 'admin';
CREATE INDEX IF NOT EXISTS users_role_user ON users (role) WHERE role = 'user';
