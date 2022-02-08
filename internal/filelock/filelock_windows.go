// inspired by https://raw.githubusercontent.com/kensipe/go/master/src/cmd/go/internal/lockedfile/internal/filelock/filelock_windows.go
// modified to reduce needs on internal Go code.  Maintaining internal Go interfaces for filelock.

//go:build windows

package filelock

import (
	"io/fs"
	"syscall"

	"golang.org/x/sys/windows"
)

type lockType uint32

const (
	readLock lockType = 0
	//writeLock lockType = windows.LOCKFILE_EXCLUSIVE_LOCK   // originally
	writeLock lockType = 0x00000002

	// from https://github.com/golang/go/blob/f229e7031a6efb2f23241b5da000c3b3203081d6/src/internal/syscall/windows/syscall_windows.go#L41-L42
	ERROR_NOT_SUPPORTED        syscall.Errno = 50
	ERROR_CALL_NOT_IMPLEMENTED syscall.Errno = 120
)

const (
	reserved = 0
	allBytes = ^uint32(0)
)

func lock(f File, lt lockType) error {
	// Per https://golang.org/issue/19098, “Programs currently expect the Fd
	// method to return a handle that uses ordinary synchronous I/O.”
	// However, LockFileEx still requires an OVERLAPPED structure,
	// which contains the file offset of the beginning of the lock range.
	// We want to lock the entire file, so we leave the offset as zero.
	//ol := new(syscall.Overlapped)  // originally
	ol := new(windows.Overlapped)

	//err := windows.LockFileEx(syscall.Handle(f.Fd()), uint32(lt), reserved, allBytes, allBytes, ol)  // originally
	err := windows.LockFileEx(windows.Handle(f.Fd()), uint32(lt), reserved, allBytes, allBytes, ol)
	if err != nil {
		return &fs.PathError{
			Op:   lt.String(),
			Path: f.Name(),
			Err:  err,
		}
	}
	return nil
}

func unlock(f File) error {
	//ol := new(syscall.Overlapped)						// originally
	//err := windows.UnlockFileEx(syscall.Handle(f.Fd()), reserved, allBytes, allBytes, ol)   // originally
	ol := new(windows.Overlapped)
	err := windows.UnlockFileEx(windows.Handle(f.Fd()), reserved, allBytes, allBytes, ol)
	if err != nil {
		return &fs.PathError{
			Op:   "Unlock",
			Path: f.Name(),
			Err:  err,
		}
	}
	return nil
}

func isNotSupported(err error) bool {
	switch err {
	case ERROR_NOT_SUPPORTED, ERROR_CALL_NOT_IMPLEMENTED, ErrNotSupported:
		return true
	default:
		return false
	}
}
