package repository

const (
	createComment = `INSERT INTO comments (message, user_id, post_id, created_at, updated_at) 
											VALUES ($1, $2, $3,now(), now()) 
											RETURNING *`
	getCommentsByPostID = `SELECT c.*, u.first_name, u.last_name
											FROM comments c
											INNER JOIN users u ON c.user_id = u.user_id
											WHERE c.post_id = $1`
	getCommentByID = `SELECT c.*, u.first_name, u.last_name
											FROM comments c
											INNER JOIN users u ON c.user_id = u.user_id
											WHERE c.comment_id = $1`
	updateCommentQuery = `UPDATE comments 
										SET message = COALESCE(NULLIF($1, ''), message),
												updated_at = now()
										WHERE Comment_id = $2
										RETURNING *`
	deleteCommentQuery = `DELETE FROM comments WHERE comment_id = $1`
)
