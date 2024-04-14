package db

import (
	"banney/sdk/models"
	"context"
)

func (r *Client) GetUser(ctx context.Context, obj any) (*models.User, error) {
	loginForm, ok := obj.(*models.LoginForm)
	if !ok {
		return nil, ErrConversionFailed
	}

	row := r.cli.QueryRowContext(
		ctx,
		queryGetUser,
		loginForm.Login,
	)

	var user models.User

	err := row.Scan(&user.ID, &user.Login, &user.Hash, &user.IsAdmin)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
