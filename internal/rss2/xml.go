package rss2

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

func MarshalTo(w io.Writer, channel *Channel) (_ int64, err error) {
	var root Root
	root.Version = "2.0"
	root.Channel = channel

	var buf bytes.Buffer
	buf.WriteString(xml.Header)
	_, _ = fmt.Fprintf(&buf, "<?xml-stylesheet href=%q type=%q?>\n", "/feed.xsl", "text/xsl")
	encoder := xml.NewEncoder(&buf)
	if err = encoder.Encode(&root); err != nil {
		return
	}
	return buf.WriteTo(w)
}

func ReadUnmarshal(r io.Reader) (channel *Channel, err error) {
	var root Root
	err = xml.NewDecoder(r).Decode(&root)
	channel = root.Channel
	return
}
