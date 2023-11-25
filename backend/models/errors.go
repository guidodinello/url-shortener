package models

import "errors"

var ErrNotFound = errors.New("not found")

var ErrEmptyUrl = errors.New("empty")
var ErrInvalidUrl = errors.New("invalid")
