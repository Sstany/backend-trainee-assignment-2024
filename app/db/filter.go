package db

import (
	"banney/sdk/models"
	"context"
	"database/sql"
	"fmt"
)

func (r *Client) FilterBanners(
	ctx context.Context,
	filter *models.BannerFilter,
) ([]*models.Banner, error) {
	var (
		banners []*models.Banner
		rows    *sql.Rows
		err     error
	)

	switch {
	case filter.FeatureID == 0 && filter.TagID > 0:
		rows, err = r.cli.QueryContext(
			ctx,
			queryGetBannersByTag,
			filter.TagID,
			filter.Limit,
			filter.Offset,
		)
		if err != nil {
			return nil, fmt.Errorf("filter by tag failed: %w", err)
		}
	case filter.TagID == 0 && filter.FeatureID > 0:
		rows, err = r.cli.QueryContext(
			ctx,
			queryGetBannersByFeature,
			filter.FeatureID,
			filter.Limit,
			filter.Offset,
		)
		if err != nil {
			return nil, fmt.Errorf("filter by tag failed: %w", err)
		}
	}

	for rows.Next() {
		tempBanner := models.Banner{}
		err = rows.Scan(
			&tempBanner.ID,
			&tempBanner.FeatureID,
			&tempBanner.Content,
			&tempBanner.IsActive,
		)
		if err != nil {
			return nil, fmt.Errorf("banners parse: %w", err)
		}

		banners = append(banners, &tempBanner)
	}

	for _, banner := range banners {
		banner.TagIDs, err = r.getTags(ctx, banner.ID)
		if err != nil {
			return nil, fmt.Errorf("list tags: %w", err)
		}
	}

	return banners, nil
}
