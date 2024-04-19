package models

import "time"

type Content struct {
	Body   []byte
	Expiry time.Duration
}
