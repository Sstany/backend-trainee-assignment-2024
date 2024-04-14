package db

const (
	queryCreateBannerTable = `CREATE TABLE banner (
		banner_id serial PRIMARY KEY,
		feature_id int ,
		content JSONB,
		is_active bool NOT NULL)`

	queryCreateTagTable = `CREATE TABLE tag (tag_id serial PRIMARY KEY)`

	queryCreateBannerTagTable = `CREATE TABLE banner_tag (
		banner_id int references banner (banner_id) ON DELETE CASCADE,
		tag_id int references tag (tag_id),
		CONSTRAINT banner_tag_pkey PRIMARY KEY (banner_id, tag_id))`
)

const (
	queryInsertFeature   = `INSERT INTO feature (feature_id) VALUES ($1) ON CONFLICT DO NOTHING`
	queryInsertTag       = `INSERT INTO tag (tag_id) VALUES ($1) ON CONFLICT DO NOTHING`
	queryInsertBanner    = `INSERT INTO banner (feature_id,content,is_active) VALUES ($1,$2,$3) RETURNING banner_id`
	queryInsertBannerTag = `INSERT INTO banner_tag(banner_id, tag_id) VALUES($1,$2)`
)

const (
	querySelectBanner = `SELECT banner_id FROM banner WHERE banner_id = ($1)`
)

const (
	queryDeleteBanner = `DELETE FROM banner WHERE  banner_id = ($1)`
	queryDeleteTag    = `DELETE FROM banner_tag WHERE banner_id = $1`
	queryUpdateTag    = `UPDATE banner_tag SET tag_id = $2 WHERE banner_id = $1`
)

const (
	queryUpdateBanner = `UPDATE banner SET feature_id = $2, content = $3, is_active = $4 WHERE banner_id = $1`
)
const (
	queryTruncateAll = `TRUNCATE banner, tag, banner_tag`
)

const (
	queryGetBannerByTagAndFeature = `SELECT
    banner.banner_id,
    banner_tag.tag_id,
    banner.feature_id,
    banner.content,
    banner.is_active
FROM
    banner_tag
INNER JOIN
    banner
  ON banner_tag.banner_id = banner.banner_id
WHERE feature_id=$1
AND tag_id =$2`
)
