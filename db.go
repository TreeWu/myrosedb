package myrosedb

import (
	"errors"
	"math"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

const (
	logFileTypeNum   = 5
	encodeHeaderSize = 10
	initialListSeq   = math.MaxUint32 / 2
	discardFilePath  = "DISCARD"
	lockFileName     = "FLOCK"
)
