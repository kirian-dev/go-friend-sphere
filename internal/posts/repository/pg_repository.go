package repository

import (
	"context"
	"database/sql"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/internal/posts"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type postsRepo struct {
	db *sqlx.DB
}

func NewPostsRepo(db *sqlx.DB) posts.Repository {
	return &postsRepo{db: db}
}

func (r *postsRepo) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	p := &models.Post{}
	if err := r.db.QueryRowxContext(ctx, createPost, post.Content, post.UserId, post.LikesCount, post.ImageUrl).StructScan(p); err != nil {
		return nil, errors.Wrap(err, "post repository, Create post")
	}

	return p, nil
}

func (r *postsRepo) UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	updatedPost := &models.Post{}
	if err := r.db.GetContext(ctx, updatedPost, updatePostQuery, &post.Content, &post.ImageUrl, &post.PostID); err != nil {
		return nil, errors.Wrap(err, "post repository, Update post")
	}

	return updatedPost, nil
}

func (r *postsRepo) DeletePost(ctx context.Context, postId int64) error {
	result, err := r.db.ExecContext(ctx, deletePostQuery, postId)
	if err != nil {
		return errors.Wrap(err, "post repository, Delete post")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "post repository, RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "post repository, rowsAffected")
	}

	return nil
}

func (r *postsRepo) GetPostById(ctx context.Context, postId int64) (*models.Post, error) {
	p := &models.Post{}
	if err := r.db.QueryRowxContext(ctx, getPostById, postId).StructScan(p); err != nil {
		return nil, errors.Wrap(err, "post repository, Get post by id")
	}

	return p, nil
}

func (r *postsRepo) GetPosts(ctx context.Context) ([]*models.Post, error) {
	postsList := []*models.Post{}

	if err := r.db.SelectContext(ctx, &postsList, getPosts); err != nil {
		return nil, errors.Wrap(err, "post repository, Get posts")
	}

	return postsList, nil
}

func (r *postsRepo) HasLikedPost(ctx context.Context, postId, userId int64) (bool, error) {
	var exists bool
	if err := r.db.QueryRowContext(ctx, hasLikedPost, postId, userId).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}
func (r *postsRepo) LikePost(ctx context.Context, postId, userId int64) error {
	_, err := r.db.ExecContext(ctx, createLike, postId, userId)
	return err
}

func (r *postsRepo) RemoveLike(ctx context.Context, postId, userId int64) error {
	_, err := r.db.ExecContext(ctx, removeLike, postId, userId)
	return err
}

func (r *postsRepo) UpdateLikesCount(ctx context.Context, postId int64, count int) error {
	_, err := r.db.ExecContext(ctx, updateLikesCount, count, postId)
	return err
}
