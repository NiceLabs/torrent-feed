package proxy

import (
	"encoding/base32"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"
)

func parseMagnetLink(u *url.URL) ([]byte, error) {
	if u.Scheme != "magnet" {
		return nil, fmt.Errorf("unsupported scheme: %q", u.Scheme)
	}
	xt := u.Query().Get("xt")
	infoHash, ok := strings.CutPrefix(xt, "urn:btih:")
	if !ok {
		return nil, fmt.Errorf("invalid xt parameter: %q", xt)
	}
	if len(infoHash) == 40 {
		return hex.DecodeString(infoHash)
	}
	return base32.StdEncoding.DecodeString(infoHash)
}
