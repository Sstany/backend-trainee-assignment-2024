package db

import (
	"banney/sdk/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *Client) UpdateBanner(ctx context.Context, obj any) (int, error) {
	banner, ok := obj.(*models.Banner)
	if !ok {
		return -1, ErrConversionFailed
	}

	tx, err := r.cli.BeginTx(ctx, nil)
	if err != nil {
		return -1, fmt.Errorf("begin tx: %w", err)
	}

	defer tx.Rollback()
	var (
		foundBanner models.Banner
		foundtagID  int
	)

	err = tx.QueryRowContext(
		ctx,
		queryGetBannerByTagAndFeature,
		banner.FeatureID,
		banner.TagIDs[0],
	).Scan(&foundBanner.ID,
		&foundtagID,
		&foundBanner.FeatureID,
		&foundBanner.Content,
		&foundBanner.IsActive,
	)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return -1, fmt.Errorf("get banner: %w", err)
		}
	}
	if foundBanner.ID != 0 {
		return -1, ErrBannerExists
	}

	_, err = tx.ExecContext(ctx, queryDeleteTag, banner.ID)
	if err != nil {
		return -1, fmt.Errorf("delete tags: %w", err)
	}

	for _, tagID := range banner.TagIDs {
		_, err = tx.ExecContext(ctx, queryInsertTag, tagID)
		if err != nil {
			return -1, fmt.Errorf("insert tag: %w", err)
		}
		_, err = tx.ExecContext(ctx, queryInsertBannerTag, banner.ID, tagID)
		if err != nil {
			return -1, fmt.Errorf("insert banner_tag: %w", err)
		}
	}

	_, err = tx.ExecContext(ctx, queryUpdateBanner, banner.ID, banner.FeatureID, banner.Content, banner.IsActive)
	if err != nil {
		return -1, fmt.Errorf("insert banner: %w", err)
	}

	tx.Commit()

	return banner.ID, nil

}
