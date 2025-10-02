package proxy

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/NiceLabs/torrent-feed/internal/rss2"
)

var hosts = make(map[string]func(item *rss2.ChannelItem))

const (
	pathTorrentFile = "/torrent-file"
)

type Handler struct {
	HTTPClient *http.Client
	FileBase   fs.FS
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	switch r.URL.Path {
	case pathTorrentFile:
		h.downloadFile(w, r.URL.Query().Get("url"))
	case "/feed":
		query := r.URL.Query()
		limits, _ := strconv.Atoi(query.Get("limits"))
		h.rewriteFeed(w, r.Host, query.Get("url"), limits)
	case "/favicon.ico":
		http.Redirect(w, r, "/favicon.png", http.StatusPermanentRedirect)
	default:
		if h.FileBase != nil {
			http.ServeFileFS(w, r, h.FileBase, path.Clean(r.URL.Path))
		} else {
			http.NotFound(w, r)
		}
	}
}

func (h *Handler) rewriteFeed(w http.ResponseWriter, host string, feedUrl string, limits int) {
	channel, err := unmarshalChannel(h.HTTPClient, feedUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, item := range channel.Items {
		if rewrite, ok := hosts[item.Link.Host]; ok {
			rewrite(item)
		}
		if item.PublishDate == nil && item.Torrent != nil {
			item.PublishDate = item.Torrent.PublishDate
			item.Torrent = nil
		}
		hash := sha1.Sum([]byte(item.GUID))
		item.GUID = base64.StdEncoding.
			WithPadding(base64.NoPadding).
			EncodeToString(hash[:])
		for _, enclosure := range item.Enclosures {
			scheme := enclosure.URL.Scheme
			if !(scheme == "http" || scheme == "https") {
				continue
			}
			enclosure.URL.URL = &url.URL{
				Scheme:   "https",
				Host:     host,
				Path:     pathTorrentFile,
				RawQuery: "url=" + url.QueryEscape(enclosure.URL.String()),
			}
			if enclosure.Length < 1024 {
				enclosure.Length = 0
			}
		}
	}
	if limits > 0 && limits < len(channel.Items) {
		channel.Items = channel.Items[:limits]
	}
	if _, err = rss2.MarshalTo(w, channel); err != nil {
		panic(err)
	}
}

func (h *Handler) downloadFile(w http.ResponseWriter, torrentUrl string) {
	response, err := h.HTTPClient.Get(torrentUrl)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	header := w.Header()
	header.Set("Content-Type", "application/x-bittorrent")
	header.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", path.Base(torrentUrl)))
	if _, err = io.Copy(w, response.Body); err != nil {
		panic(err)
	}
}

func unmarshalChannel(client *http.Client, url string) (*rss2.Channel, error) {
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return rss2.ReadUnmarshal(response.Body)
}
