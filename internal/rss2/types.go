package rss2

import "encoding/xml"

type Root struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title       string         `xml:"title"`
	Link        *URL           `xml:"link"`
	Description string         `xml:"description"`
	Items       []*ChannelItem `xml:"item"`
}

type ChannelItem struct {
	GUID        string       `xml:"guid,omitempty"`
	Title       string       `xml:"title,omitempty"`
	Link        *URL         `xml:"link,omitempty"`
	Author      string       `xml:"author,omitempty"`
	Categories  []string     `xml:"category,omitempty"`
	Enclosures  []*Enclosure `xml:"enclosure,omitempty"`
	Torrent     *Torrent     `xml:"torrent,omitempty"`
	PublishDate *DateTime    `xml:"pubDate,omitempty"`
}

type Enclosure struct {
	URL    *URL   `xml:"url,omitempty,attr"`
	Length uint64 `xml:"length,omitempty,attr"`
	Type   string `xml:"type,omitempty,attr"`
}

type Torrent struct {
	URL         *URL      `xml:"link,omitempty"`
	Length      uint64    `xml:"contentLength,omitempty"`
	PublishDate *DateTime `xml:"pubDate,omitempty"`
}
