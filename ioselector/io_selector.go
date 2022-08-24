package ioselector

import (
	"errors"
	"os"
)

var ErrInvalidFsize = errors.New("file size can`t be zero or negative")

// FilePerm default permission of the newly create log file
const FilePerm = 0644

// IOSelector io selector for fileio and mmp ,used by wal and value log right now
type IOSelector interface {
	// Write a slice to log file at offset
	// It returns the number of bytes written and an error, if any
	Write(b []byte, offset int64) (int, error)
	// Read a slice form offset
	// It Return the number of bytes read and any error encountered
	Read(b []byte, offset int64) (int, error)
	// Sync commits the current contents of the file to stable storage.
	// Typically, this means flushing the file system's in-memory copy
	// of recently written data to disk
	Sync() error
	// Close closes the File,rendering it unusable for I/O
	// It will return an error if it as already bean closed.
	Close() error
	// Delete delete the file
	// Must close it before delete and will unmap if in MMapSelector
	Delete() error
}

// open file and truncate it if necessary
// 如果打开的文件大小大于 fSize ,会通过Truncate函数对多于的部分进行一个截断
func openFile(fName string, fSize int64) (*os.File, error) {
	fd, err := os.OpenFile(fName, os.O_CREATE|os.O_RDWR, FilePerm)
	if err != nil {
		return nil, err
	}
	stat, err := fd.Stat()
	if err != nil {
		return nil, err
	}
	if stat.Size() < fSize {
		err := fd.Truncate(fSize)
		if err != nil {
			return nil, err
		}
	}
	return fd, nil
}
