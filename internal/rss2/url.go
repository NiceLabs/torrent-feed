package rss2

import (
	"errors"
	"net/url"
)

type URL struct {
	*url.URL
}

func (u *URL) MarshalText() (text []byte, err error) {
	if u == nil || u.URL == nil {
		return nil, errors.New("invalid URL")
	}
	return []byte(u.URL.String()), nil
}

func (u *URL) UnmarshalText(text []byte) (err error) {
	u.URL, err = url.Parse(string(text))
	return
}
