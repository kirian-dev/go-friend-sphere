package repository

import (
	"context"
	"database/sql"
	"go-friend-sphere/internal/friendships"
	"go-friend-sphere/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type friendshipsRepo struct {
	db *sqlx.DB
}

func NewFriendshipRepo(db *sqlx.DB) friendships.Repository {
	return &friendshipsRepo{db: db}
}

func (r *friendshipsRepo) CreateFriendship(ctx context.Context, friendship *models.Friendship) (*models.Friendship, error) {
	f := &models.Friendship{}
	if err := r.db.QueryRowxContext(ctx, createFriendship, friendship.Status, friendship.UserID, friendship.FriendID).StructScan(f); err != nil {
		return nil, errors.Wrap(err, "Friendship repository, Create Friendship")
	}

	return f, nil
}

func (r *friendshipsRepo) UpdateFriendship(ctx context.Context, friendship *models.Friendship) (*models.Friendship, error) {
	updatedFriendship := &models.Friendship{}
	if friendship.Status == "reject" {
		err := r.DeleteFriendship(ctx, friendship.FriendshipID)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	if err := r.db.GetContext(ctx, updatedFriendship, updateFriendshipQuery, &friendship.Status, &friendship.FriendshipID); err != nil {
		return nil, errors.Wrap(err, "Friendship repository, Update Friendship")
	}

	return updatedFriendship, nil
}

func (r *friendshipsRepo) DeleteFriendship(ctx context.Context, FriendshipID int64) error {
	result, err := r.db.ExecContext(ctx, deleteFriendshipQuery, FriendshipID)
	if err != nil {
		return errors.Wrap(err, "Friendship repository, Delete Friendship")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Friendship repository, RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "Friendship repository, rowsAffected")
	}

	return nil
}

func (r *friendshipsRepo) GetFriendshipByID(ctx context.Context, friendshipID int64) (*models.FriendshipWithFriend, error) {
	// ... (your implementation)
	// Handle the error from the scan operation
	friendship := &models.FriendshipWithFriend{}
	err := r.db.QueryRowContext(ctx, getFriendshipByID, friendshipID).Scan(
		&friendship.FriendshipID,
		&friendship.UserID,
		&friendship.FriendID,
		&friendship.Status,
		&friendship.CreatedAt,
		&friendship.UpdatedAt,
		&friendship.FirstName,
		&friendship.LastName,
	)
	if err != nil {
		return nil, err
	}
	return friendship, nil
}

func (r *friendshipsRepo) GetFriendshipsByUserID(ctx context.Context, userID int64) ([]*models.FriendshipWithFriend, error) {
	// ... (your implementation)
	// Handle the error from the rows.Scan operation
	friendshipsList := []*models.FriendshipWithFriend{}
	rows, err := r.db.QueryContext(ctx, getFriendshipsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		friendship := &models.FriendshipWithFriend{}
		err := rows.Scan(
			&friendship.FriendshipID,
			&friendship.UserID,
			&friendship.FriendID,
			&friendship.Status,
			&friendship.CreatedAt,
			&friendship.UpdatedAt,
			&friendship.FirstName,
			&friendship.LastName,
		)
		if err != nil {
			return nil, err
		}
		friendshipsList = append(friendshipsList, friendship)
	}

	return friendshipsList, nil
}
