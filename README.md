# File Lock library (flock)

Flock is a project which provides a Go solution for system level file locks for all platforms Golang supports.

## Purpose

The purpose of this project is to provide an implementation example for the Go proposal to [add public lockedfile pkg](https://github.com/golang/proposal/blob/master/design/33974-add-public-lockedfile-pkg.md) to Go.  The proposal was pre-covid, starting Oct 2019 and merged into the proposal/design on Mar 2020.   The proposal marks as it's point of reference [07b4abd](https://github.com/golang/go/tree/07b4abd62e450f19c47266b3a526df49c01ba425/src/cmd/go/internal/lockedfile).

It turns out, I initially misunderstood the file to be exposed... it was supposed to be `lockedfile.go` in cmd/go/internal/lockedfile and not `filelocked.go` which is in cmd/go/internal/lockedfile/internal/filelock.  However I believe I like this level of exposure of just `Lock`, `RLock` and `Unlock`.

**note:** This project will refactor to match the proposal, but will maintain the lower level API for consideration.

## Usage

```go
package example

import "github.com/kensipe/flock/internal/filelock"

def m() {
	name := "filename"

	// will write lock the file, with an os level lock which will be removed on Unlock or will be removed by
	// OS if process is killed
	filelock.Lock(name)
	// do stuff
	filelock.Unlock(name)

	// also available is RLock which puts a read lock on the file
	filelock.RLock(name)
}

```
