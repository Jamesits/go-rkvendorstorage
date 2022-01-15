package rkvendorstorage

import "errors"

// ErrorDataTooLong is returned if the data read or to be written exceeds RequestMaxSize.
var ErrorDataTooLong = errors.New("data too long")
