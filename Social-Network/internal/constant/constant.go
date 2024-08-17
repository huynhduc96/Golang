package constant

import "time"

const (
	Ttl             = time.Minute * 15
	RedisTllMaxAge  = 60 * 15
	DOBLayout       = "2006-01-02"
	SaltSize        = 16
	DefaultBucket   = "images"
	ImageURLMaxLive = time.Minute * 30
)
