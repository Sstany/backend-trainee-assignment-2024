package db

import "errors"

var (
	errConversionFailed = errors.New("type conversion failed")
	errBannerExists     = errors.New("banner exists")
)
