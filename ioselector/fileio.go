package ioselector

import "os"

//FileIOSelector 文件操作类标准
type FileIOSelector struct {
	fd *os.File // 系统文件描述符
}

func (f FileIOSelector) Write(b []byte, offset int64) (int, error) {

	return f.fd.WriteAt(b, offset)
}

func (f FileIOSelector) Read(b []byte, offset int64) (int, error) {
	return f.fd.ReadAt(b, offset)
}

func (f FileIOSelector) Sync() error {
	return f.fd.Sync()
}

func (f FileIOSelector) Close() error {
	return f.fd.Close()
}

func (f FileIOSelector) Delete() error {
	err := f.fd.Close()
	if err != nil {
		return err
	}
	return os.Remove(f.fd.Name())
}

//NewFileIOSelector 文件操作类
func NewFileIOSelector(fName string, fSize int64) (IOSelector, error) {
	if fSize <= 0 {
		return nil, ErrInvalidFsize
	}
	file, err := openFile(fName, fSize)
	if err != nil {
		return nil, err
	}

	return &FileIOSelector{fd: file}, nil
}
