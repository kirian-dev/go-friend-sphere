package repository

const (
	createPost = `INSERT INTO posts (content, user_id, likes_count, image_url, created_at, updated_at) 
											VALUES ($1, $2, $3, $4, now(), now()) 
											RETURNING *`
	getPosts = `SELECT content, user_id, image_url, likes_count, created_at, updated_at , post_id
											FROM posts`
	getPostById = `SELECT content, user_id, image_url, likes_count, created_at, updated_at 
									FROM posts
									WHERE post_id = $1`
	updatePostQuery = `UPDATE posts 
										SET content = COALESCE(NULLIF($1, ''), content),
												image_url = COALESCE(NULLIF($2, ''), image_url),
												updated_at = now()
										WHERE post_id = $3
										RETURNING *`
	deletePostQuery  = `DELETE FROM posts WHERE post_id = $1`
	hasLikedPost     = `SELECT EXISTS (SELECT 1 FROM post_likes WHERE post_id = $1 AND user_id = $2)`
	createLike       = `INSERT INTO post_likes (post_id, user_id, created_at) VALUES ($1, $2, now())`
	removeLike       = `DELETE FROM post_likes WHERE post_id = $1 AND user_id = $2`
	updateLikesCount = `UPDATE posts SET likes_count = likes_count + $1 WHERE post_id = $2`
)
