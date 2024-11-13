package tar

import (
	"archive/tar"
	"bytes"
	"io"
	"strconv"
)

type Reader struct {
	tr   *tar.Reader
	r    io.Reader
	size int64
}

func NewReader(r io.Reader) *Reader {
	if Reference {
		return &Reader{tr: tar.NewReader(r)}
	}
	return &Reader{r: r}
}
func (tr *Reader) Next() (*tar.Header, error) {
	if Reference {
		return tr.tr.Next()
	}
	var blk = make([]byte, 512)
	io.ReadFull(tr.r, blk)
	if bytes.Equal(blk[:], make([]byte, 512)) {
		return nil, io.EOF
	}
	var hdr = &tar.Header{}
	hdr.Name = string(blk[:bytes.IndexByte(blk, 0)])
	hdr.Mode, _ = strconv.ParseInt(string(blk[108:][:bytes.IndexByte(blk[108:], 0)]), 8, 64)
	hdr.Size, _ = strconv.ParseInt(string(blk[124:][:bytes.IndexByte(blk[124:], 0)]), 8, 64)
	tr.size = hdr.Size
	return hdr, nil
}
func (tr *Reader) Read(b []byte) (int, error) {
	if Reference {
		return tr.tr.Read(b)
	}
	return tr.r.Read(b)
}
