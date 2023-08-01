package repository

import (
	"context"
	"database/sql"
	"go-friend-sphere/internal/comments"
	"go-friend-sphere/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type commentsRepo struct {
	db *sqlx.DB
}

func NewCommentsRepo(db *sqlx.DB) comments.Repository {
	return &commentsRepo{db: db}
}

func (r *commentsRepo) CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	p := &models.Comment{}
	if err := r.db.QueryRowxContext(ctx, createComment, comment.Message, comment.UserID, comment.PostID).StructScan(p); err != nil {
		return nil, errors.Wrap(err, "Comment repository, Create Comment")
	}

	return p, nil
}

func (r *commentsRepo) UpdateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	updatedComment := &models.Comment{}
	if err := r.db.GetContext(ctx, updatedComment, updateCommentQuery, &comment.Message, &comment.CommentID); err != nil {
		return nil, errors.Wrap(err, "Comment repository, Update Comment")
	}

	return updatedComment, nil
}

func (r *commentsRepo) DeleteComment(ctx context.Context, commentID int64) error {
	result, err := r.db.ExecContext(ctx, deleteCommentQuery, commentID)
	if err != nil {
		return errors.Wrap(err, "Comment repository, Delete Comment")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Comment repository, RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "Comment repository, rowsAffected")
	}

	return nil
}

func (r *commentsRepo) GetCommentByID(ctx context.Context, commentID int64) (*models.CommentWithUser, error) {
	com := &models.CommentWithUser{}
	if err := r.db.QueryRowxContext(ctx, getCommentByID, commentID).StructScan(com); err != nil {
		return nil, errors.Wrap(err, "Comment repository, Get Comment by id")
	}

	return com, nil
}

func (r *commentsRepo) GetCommentsByPostID(ctx context.Context, postID int64) ([]*models.CommentWithUser, error) {
	commentsList := []*models.CommentWithUser{}

	if err := r.db.SelectContext(ctx, &commentsList, getCommentsByPostID, postID); err != nil {
		return nil, errors.Wrap(err, "Comment repository, Get Comments")
	}

	return commentsList, nil
}
