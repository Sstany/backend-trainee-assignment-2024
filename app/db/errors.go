package db

import "errors"

var (
	errConversionFailed = errors.New("type conversion failed")
	errBannerExists     = errors.New("banner already exists")
	errBannerNotExists  = errors.New("banner does not exist")
)
