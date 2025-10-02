package proxy

import (
	"fmt"
	"net/url"

	"github.com/NiceLabs/torrent-feed/internal/rss2"
)

func init() {
	hosts["share.dmhy.org"] = rewriteDMHY
}

func rewriteDMHY(item *rss2.ChannelItem) {
	baseUrl := &url.URL{Scheme: "https", Host: "dl.dmhy.org"}
	for _, enclosure := range item.Enclosures {
		if infoHash, err := parseMagnetLink(enclosure.URL.URL); err == nil {
			enclosure.URL.URL = baseUrl.JoinPath(
				item.PublishDate.Format("2006/01/02"),
				fmt.Sprintf("%02x.torrent", infoHash),
			)
		}
	}
}
