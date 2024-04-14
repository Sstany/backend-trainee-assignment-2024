package db

import "errors"

var (
	ErrConversionFailed = errors.New("type conversion failed")
	ErrBannerExists     = errors.New("banner already exists")
	ErrBannerNotExists  = errors.New("banner does not exist")
	ErrWrongFormat      = errors.New("wrong filter format")
)
