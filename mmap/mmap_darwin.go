package mmap

import (
	"golang.org/x/sys/unix"
	"os"
	"syscall"
	"unsafe"
)

func mmap(fd *os.File, writable bool, size int64) ([]byte, error) {
	mtype := unix.PROT_READ
	if writable {
		mtype |= unix.PROT_WRITE
	}
	return unix.Mmap(int(fd.Fd()), 0, int(size), mtype, unix.MAP_SHARED)
}

func munmap(b []byte) error {
	return unix.Munmap(b)
}

func madvise(b []byte, readahead bool) error {
	advice := unix.MADV_NORMAL
	if !readahead {
		advice = unix.MADV_RANDOM
	}
	_, _, err := syscall.Syscall(syscall.SYS_MADVISE, uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)), uintptr(advice))
	if err != 0 {
		return err
	}
	return nil
}

func msync(b []byte) error {
	return unix.Msync(b, unix.MS_SYNC)
}
