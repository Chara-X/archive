package tar

import (
	"archive/tar"
	"fmt"
	"io"
	"strings"
)

type Writer struct {
	tw   *tar.Writer
	w    io.Writer
	size int64
}

func NewWriter(w io.Writer) *Writer {
	if Reference {
		return &Writer{tw: tar.NewWriter(w)}
	}
	return &Writer{w: w}
}
func (tw *Writer) WriteHeader(hdr *tar.Header) error {
	if Reference {
		return tw.tw.WriteHeader(hdr)
	}
	tw.Flush()
	var blk = make([]byte, 512)
	copy(blk[:100], hdr.Name)
	copy(blk[100:108], fmt.Sprintf("%07o", hdr.Mode))
	copy(blk[124:136], fmt.Sprintf("%011o", hdr.Size))
	copy(blk[148:156], strings.Repeat(" ", 8))
	var sum int64
	for _, v := range blk {
		sum += int64(v)
	}
	copy(blk[148:156], fmt.Sprintf("%06o", sum))
	tw.w.Write(blk[:])
	tw.size = hdr.Size
	return nil
}
func (tw *Writer) Write(b []byte) (n int, err error) {
	if Reference {
		return tw.tw.Write(b)
	}
	return tw.w.Write(b)
}
func (tw *Writer) Flush() error {
	if Reference {
		return tw.tw.Flush()
	}
	if n := tw.size % 512; n > 0 {
		tw.w.Write(make([]byte, 512-n))
	}
	return nil
}
func (tw *Writer) Close() error {
	if Reference {
		return tw.tw.Close()
	}
	tw.Flush()
	tw.w.Write(make([]byte, 1024))
	return nil
}
