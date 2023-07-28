package repository

const (
	findUserByEmail = `SELECT * FROM users WHERE email = $1`
	createUser      = `INSERT INTO users (email, password, first_name, last_name, created_at, updated_at, last_login_at) VALUES ($1, $2, $3, $4, now(), now(), now()) RETURNING *`
)
