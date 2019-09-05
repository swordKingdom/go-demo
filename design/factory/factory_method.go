package main

import (
	"fmt"
	"io"
)

//方法工厂
type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type MemoryStorageImpl struct {
}

func (v *MemoryStorageImpl) Open(string) (io.ReadWriteCloser, error) {
	fmt.Println("get MemoryStorageImpl")
	return nil, nil
}

func newMemoryStorage() *MemoryStorageImpl {
	return &MemoryStorageImpl{}
}

type DiskStorageImpl struct {
}

func (v *DiskStorageImpl) Open(string) (io.ReadWriteCloser, error) {
	fmt.Println("get DiskStorageImpl")
	return nil, nil
}

func newDiskStorage() *DiskStorageImpl {
	return &DiskStorageImpl{}
}

type TempStorageImpl struct {
}

func (v *TempStorageImpl) Open(string) (io.ReadWriteCloser, error) {
	fmt.Println("get TempStorageImpl")
	return nil, nil
}

func newTempStorage() *TempStorageImpl {
	return &TempStorageImpl{}
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage()
	case DiskStorage:
		return newDiskStorage()
	default:
		return newTempStorage()
	}
}

func main() {
	a := NewStore(DiskStorage)
	a.Open("")
	a = NewStore(TempStorage)
	a.Open("")
	a = NewStore(MemoryStorage)
	a.Open("")
}
