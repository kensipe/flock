package filelock

import (
	"io/fs"

	internal "github.com/kensipe/flock/internal/filelock"
)

// pulled from internal/filelock.go and modified to call the internal as needed.

// A File provides the minimal set of methods required to lock an open file.
// File implementations must be usable as map keys.
// The usual implementation is *os.File.
type File interface {
	// Name returns the name of the file.
	Name() string

	// Fd returns a valid file descriptor.
	// (If the File is an *os.File, it must not be closed.)
	Fd() uintptr

	// Stat returns the FileInfo structure describing file.
	Stat() (fs.FileInfo, error)
}

// Lock places an advisory write lock on the file, blocking until it can be
// locked.
//
// If Lock returns nil, no other process will be able to place a read or write
// lock on the file until this process exits, closes f, or calls Unlock on it.
//
// If f's descriptor is already read- or write-locked, the behavior of Lock is
// unspecified.
//
// Closing the file may or may not release the lock promptly. Callers should
// ensure that Unlock is always called when Lock succeeds.
func Lock(f File) error {
	return internal.Lock(f)
}

// RLock places an advisory read lock on the file, blocking until it can be locked.
//
// If RLock returns nil, no other process will be able to place a write lock on
// the file until this process exits, closes f, or calls Unlock on it.
//
// If f is already read- or write-locked, the behavior of RLock is unspecified.
//
// Closing the file may or may not release the lock promptly. Callers should
// ensure that Unlock is always called if RLock succeeds.
func RLock(f File) error {
	return internal.RLock(f)
}

// Unlock removes an advisory lock placed on f by this process.
//
// The caller must not attempt to unlock a file that is not locked.
func Unlock(f File) error {
	return internal.Unlock(f)
}
