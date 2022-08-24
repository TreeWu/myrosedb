package ioselector

import (
	"io"
	"os"
	"wushr.cn/myrosedb/mmap"
)

type MMapSelector struct {
	fd     *os.File
	buf    []byte
	bufLen int64
}

func NewMMapSelector(fName string, fSize int64) (IOSelector, error) {
	if fSize <= 0 {
		return nil, ErrInvalidFsize
	}
	file, err := openFile(fName, fSize)
	if err != nil {
		return nil, err
	}
	buf, err := mmap.Mmap(file, true, fSize)
	if err != nil {
		return nil, err
	}
	return &MMapSelector{
		fd:     file,
		buf:    buf,
		bufLen: int64(len(buf)),
	}, nil
}

func (lm *MMapSelector) Write(b []byte, offset int64) (int, error) {
	length := int64(len(b))
	if length <= 0 {
		return 0, nil
	}
	if offset < 0 || length+offset > lm.bufLen {
		return 0, io.EOF
	}
	return copy(lm.buf[offset:], b), nil
}

func (lm *MMapSelector) Read(b []byte, offset int64) (int, error) {
	if offset < 0 || offset >= lm.bufLen {
		return 0, io.EOF
	}
	if offset+int64(len(b)) >= lm.bufLen {
		return 0, io.EOF
	}
	return copy(b, lm.buf[offset:]), nil
}

func (lm *MMapSelector) Sync() error {
	return mmap.Msync(lm.buf)
}

func (lm *MMapSelector) Close() error {
	if err := mmap.Msync(lm.buf); err != nil {
		return err
	}

	if err := mmap.Munmap(lm.buf); err != nil {
		return err
	}

	return lm.fd.Close()
}

func (lm *MMapSelector) Delete() error {
	if err := mmap.Munmap(lm.buf); err != nil {
		return nil
	}

	if err := lm.fd.Truncate(0); err != nil {
		return nil
	}

	if err := lm.fd.Close(); err != nil {
		return nil
	}
	return os.Remove(lm.fd.Name())

}
