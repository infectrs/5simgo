package main

import "errors"

// List of 5SimGo errors
var (
	missingApiKeyError = errors.New("5simgo: missing api key")
	missingArgsError   = errors.New("5simgo: missing args")
	httpError          = errors.New("5simgo: error while doing request")
	unmarshalError     = errors.New("5simgo: error while unmarshaling data")
)
