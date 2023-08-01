package repository

const (
	createMessage = `INSERT INTO messages (message, sender_id, recipient_id, sent_at, read_at, updated_at) 
											VALUES ($1, $2, $3, now(), NULL) 
											RETURNING *`
	getMessagesByUserID = `
											SELECT m.message_id, m.sender_id, m.recipient_id, m.status, m.created_at, m.updated_at, u.first_name AS recipient_first_name, u.last_name AS recipient_last_name
											FROM messages m
											INNER JOIN users u ON m.recipient = u.user_id
											WHERE m.user_id = $1
											`
	getMessageByID = `SELECT m.*, u.first_name AS recipient_first_name, u.last_name AS recipient_last_name
											FROM messages m
											INNER JOIN users u ON m.recipient_id = u.user_id
											WHERE m.message_id = $1`
	updateMessageQuery = `UPDATE messages 
										SET message = COALESCE(NULLIF($1, ''), message),
												updated_at = now()
										WHERE message_id = $2
										RETURNING *`
	deleteMessageQuery = `DELETE FROM messages WHERE message_id = $1`
)
