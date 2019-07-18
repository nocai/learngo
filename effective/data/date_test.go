package data

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func TestSyncedBuffer(t *testing.T) {
	p := new(SyncedBuffer)
	var v SyncedBuffer

	fmt.Println(p)
	fmt.Println(v)
}

type File struct {
	fd      int
	name    string
	dirinfo *string
	nepipe  int
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	f := new(File)
	f.fd = fd
	f.name = name
	f.dirinfo = nil
	f.nepipe = 0
	return f
}

func NewFile2(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	f := File{fd, name, nil, 0}
	return &f
}

func NewFile3(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd, name, nil, 0}
}

func NewFile4(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd: fd, name: name}
}

const (
	Enone int = 1
)
func TestA(t *testing.T) {
	var o struct{}
	fmt.Println(o)

	s := []string{Enone: "no error"}
	fmt.Println(s)
}
