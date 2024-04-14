package db

import (
	"banney/sdk/models"
	"context"
	"fmt"
	"strconv"
)

func (r *Client) GetBanner(
	ctx context.Context,
	featureIDStr,
	tagIDStr string,
) (*models.Banner, error) {
	tagID, err := strconv.Atoi(tagIDStr)
	if err != nil {
		return nil, fmt.Errorf("tagID to string: %w", err)
	}

	featureID, err := strconv.Atoi(featureIDStr)
	if err != nil {
		return nil, fmt.Errorf("featureID to string: %w", err)
	}

	var banner models.Banner
	var foundTagID int

	err = r.cli.QueryRowContext(
		ctx,
		queryGetBannerByTagAndFeature,
		featureID,
		tagID,
	).Scan(
		&banner.ID,
		&foundTagID,
		&banner.FeatureID,
		&banner.Content,
		&banner.IsActive,
	)
	if err != nil {
		return nil, fmt.Errorf("get banner: %w", err)
	}

	rows, err := r.cli.QueryContext(ctx, queryGetTagsByBannerID, banner.ID)
	if err != nil {
		return nil, fmt.Errorf("get tags: %w", err)
	}

	var tempTag int

	for rows.Next() {
		if err = rows.Scan(&tempTag); err != nil {
			return nil, fmt.Errorf("parse tag: %w", err)
		}

		banner.TagIDs = append(banner.TagIDs, tempTag)
	}

	return &banner, nil
}
