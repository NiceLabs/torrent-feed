# Torrent Feed Proxy

This is a simple web service that proxies torrent RSS feeds to make them compatible with BitTorrent clients that support
RSS feeds, such as qBittorrent and uTorrent. It fetches the original RSS feed, modifies the torrent links to use magnet
links, and serves the modified feed.

## Usage

```
/feed?url=<original_rss_feed_url>
```

## References

- [BEP 0009](https://www.bittorrent.org/beps/bep_0009.html)
- [BEP 0036](https://www.bittorrent.org/beps/bep_0036.html)
- <https://www.ghostchu.com/bt-tracker-bep-0036-torrent-rss-feed/>
- <https://www.rssboard.org/rss-specification>

## LICENSE

[MIT LICENSE](LICENSE.txt)