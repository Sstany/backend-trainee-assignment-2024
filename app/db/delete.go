package db

import (
	"context"
	"fmt"
	"strconv"
)

func (r *Client) DeleteBanner(ctx context.Context, obj any) error {
	bannerIDString, ok := obj.(string)
	if !ok {
		return ErrConversionFailed
	}

	bannerID, err := strconv.Atoi(bannerIDString)
	if err != nil {
		return fmt.Errorf("error banner input: %w", err)
	}
	tx, err := r.cli.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer tx.Rollback()
	var foundBannerID int
	err = tx.QueryRowContext(ctx, querySelectBanner, bannerID).Scan(&foundBannerID)
	if foundBannerID == 0 {
		return ErrBannerNotExists
	}
	if err != nil {
		return fmt.Errorf("get banner: %w", err)

	}

	_, err = tx.ExecContext(ctx, queryDeleteBanner, bannerID)
	if err != nil {
		return fmt.Errorf(" banner: %w", err)
	}

	tx.Commit()

	return nil
}
