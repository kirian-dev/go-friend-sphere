package repository

const (
	createFriendship = `INSERT INTO friendships (status, user_id, friend_id, created_at, updated_at) 
											VALUES ($1, $2, $3,now(), now()) 
											RETURNING *`
	getFriendshipsByUserID = `
											SELECT f.friendship_id, f.user_id, f.friend_id, f.status, f.created_at, f.updated_at, u.first_name AS friend_first_name, u.last_name AS friend_last_name
											FROM friendships f
											INNER JOIN users u ON f.friend_id = u.user_id
											WHERE f.user_id = $1
											`
	getFriendshipByID = `SELECT f.*, u.first_name AS friend_first_name, u.last_name AS friend_last_name
											FROM friendships f
											INNER JOIN users u ON c.friend_id = u.user_id
											WHERE f.friendship_id = $1`
	updateFriendshipQuery = `UPDATE friendships 
										SET status = COALESCE(NULLIF($1, ''), status),
												updated_at = now()
										WHERE friendship_id = $2
										RETURNING *`
	deleteFriendshipQuery = `DELETE FROM friendships WHERE friendship_id = $1`
)
