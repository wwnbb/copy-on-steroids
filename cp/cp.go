package cp

import (
	"io"
	"os"

	prb "github.com/cheggaaa/pb/v3"
)

func CopyFile(from string, to string, limit int, offset int) error {
	src, err := os.OpenFile(from, os.O_RDONLY, 0o644)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
		if os.IsPermission(err) {
			return err
		}
	}
	defer src.Close()
	src_info, err := src.Stat()
	if err != nil {
		return err
	}
	fs := src_info.Size() - int64(offset)

	_, err = src.Seek(int64(offset), 0)
	if err != nil {
		return err
	}
	dst, err := os.Create(to)
	if err != nil {
		return err
	}

	bar := prb.Simple.Start64(fs)
	defer bar.Finish()
	barReader := bar.NewProxyReader(src)
	copy_ := func(offset int) (written int64, err error) {
		if offset > 0 {
			return io.CopyN(dst, barReader, int64(limit))
		} else {
			return io.Copy(dst, barReader)
		}
	}
	_, err = copy_(offset)
	if err != nil {
		return err
	}
	return nil
}
