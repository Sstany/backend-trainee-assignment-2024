package db

import (
	"banney/sdk/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *Client) CreateBanner(ctx context.Context, obj any) (int, error) {
	banner, ok := obj.(*models.Banner)
	if !ok {
		return -1, errConversionFailed
	}

	tx, err := r.cli.BeginTx(ctx, nil)
	if err != nil {
		return -1, fmt.Errorf("begin tx: %w", err)
	}

	defer tx.Rollback()
	var foundBanner models.Banner
	var tagID int

	err = tx.QueryRowContext(
		ctx,
		queryGetBannerByTagAndFeature,
		banner.FeatureID,
		banner.TagIDs[0],
	).Scan(&foundBanner.ID,
		&tagID,
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
		return -1, errBannerExists
	}

	_, err = tx.ExecContext(ctx, queryInsertFeature, banner.FeatureID)
	if err != nil {
		return -1, fmt.Errorf("insert feature: %w", err)
	}

	rows, err := tx.QueryContext(ctx, queryInsertBanner, banner.FeatureID, banner.Content, banner.IsActive)
	if err != nil {
		return -1, fmt.Errorf("insert banner: %w", err)
	}

	var bannerID int
	if rows.Next() {
		if err = rows.Scan(&bannerID); err != nil {
			return -1, fmt.Errorf("decode bannerID: %w", err)
		}
		rows.Close()
	}

	for _, tagID := range banner.TagIDs {
		_, err = tx.ExecContext(ctx, queryInsertTag, tagID)
		if err != nil {
			return -1, fmt.Errorf("insert tag: %w", err)
		}
		_, err = tx.ExecContext(ctx, queryInsertBannerTag, bannerID, tagID)
		if err != nil {
			return -1, fmt.Errorf("insert banner_tag: %w", err)
		}
	}

	tx.Commit()

	return bannerID, nil

}
