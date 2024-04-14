package db

import (
	"banney/sdk/models"
	"context"
	"fmt"
)

func (r *Client) CreateUser(ctx context.Context, obj any) (int, error) {
	user, ok := obj.(*models.RegisterForm)
	if !ok {
		return -1, ErrConversionFailed
	}

	tx, err := r.cli.BeginTx(ctx, nil)
	if err != nil {
		return -1, fmt.Errorf("begin tx: %w", err)
	}

	defer tx.Rollback()

	rows, err := tx.QueryContext(
		ctx,
		queryInsertUser,
		user.Login,
		user.Password, //TODO: add password hashing
		user.IsAdmin,
	)
	if err != nil {
		return -1, fmt.Errorf("insert user: %w", err)
	}

	var userID int
	if rows.Next() {
		if err = rows.Scan(&userID); err != nil {
			return -1, fmt.Errorf("decode bannerID: %w", err)
		}
		rows.Close()
	}

	tx.Commit()

	return userID, nil
}
