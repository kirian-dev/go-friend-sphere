package repository

const (
	findUserByEmail = `SELECT * 
											FROM users 
											WHERE email = $1`
	createUser = `INSERT INTO users (email, password, first_name, last_name, created_at, updated_at, last_login_at, role) 
											VALUES ($1, $2, $3, $4, now(), now(), now(), $5) 
											RETURNING *`
	getUsers = `SELECT user_id, email, first_name, last_name, created_at, updated_at, last_login_at role
											FROM users`
	getUsersById = `SELECT user_id, email, first_name, last_name, created_at, updated_at, last_login_at role
									FROM users
									WHERE user_id = $1`
	updateUserQuery = `UPDATE users 
										SET email = COALESCE(NULLIF($1, ''), email),
												first_name = COALESCE(NULLIF($2, ''), first_name),
												last_name = COALESCE(NULLIF($3, ''), last_name),
												updated_at = now()
										WHERE user_id = $4
										RETURNING *`
	deleteUserQuery = `DELETE FROM users WHERE user_id = $1`
)
