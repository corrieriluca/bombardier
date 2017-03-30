package main

import (
	"errors"
	"sort"
	"time"
)

const (
	decBase = 10

	maxRps            = 10000000
	rateLimitInterval = 10 * time.Millisecond
	oneSecond         = 1 * time.Second

	exitFailure = 1
)

var (
	version = "unspecified"

	emptyConf = config{}
	parser    = newKingpinParser()

	defaultTestDuration  = 10 * time.Second
	defaultNumberOfConns = uint64(125)
	defaultTimeout       = 2 * time.Second

	httpMethods = []string{
		"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS",
	}
	cantHaveBody = []string{"GET", "HEAD"}

	errInvalidURL = errors.New(
		"No hostname or invalid scheme")
	errInvalidNumberOfConns = errors.New(
		"Invalid number of connections(must be > 0)")
	errInvalidNumberOfRequests = errors.New(
		"Invalid number of requests(must be > 0)")
	errInvalidTestDuration = errors.New(
		"Invalid test duration(must be >= 1s)")
	errNegativeTimeout = errors.New(
		"Timeout can't be negative")
	errLargeTimeout = errors.New(
		"Timeout is too big(more that 10s)")
	errBodyNotAllowed = errors.New(
		"GET and HEAD requests cannot have body")
	errNoPathToCert = errors.New(
		"No Path to TLS Client Certificate")
	errNoPathToKey = errors.New(
		"No Path to TLS Client Certificate Private Key")
	errZeroRate = errors.New(
		"Rate can't be less than 1")

	errInvalidHeaderFormat = errors.New("invalid header format")
)

func init() {
	sort.Strings(httpMethods)
	sort.Strings(cantHaveBody)
}