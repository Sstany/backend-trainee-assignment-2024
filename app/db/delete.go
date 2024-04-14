package db

import (
	"banney/sdk/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *Client) DeleteBanner(ctx context.Context, obj any) (int, error) {
	banner, ok := obj.(*models.Banner)
	if !ok {
		return -1, errConversionFailed
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
	if foundBanner.ID == 0 {
		return -1, errBannerNotExists
	}
	TagID := banner.TagIDs[0]
	rows, err := tx.QueryContext(ctx, queryDeleteBanner, banner.FeatureID, TagID)
	if err != nil {
		return -1, fmt.Errorf(" banner: %w", err)
	}
	var bannerID int
	if rows.Next() {
		if err = rows.Scan(&bannerID); err != nil {
			return -1, fmt.Errorf("decode bannerID: %w", err)
		}
		rows.Close()
	}

	tx.Commit()

	return bannerID, err
}
