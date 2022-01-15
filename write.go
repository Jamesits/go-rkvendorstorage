package rkvendorstorage

import (
	"encoding/binary"
	"os"
	"syscall"
	"unsafe"
)

// Write writes a value into the vendor storage with the given ID as the key.
func Write(id uint16, data []byte) (err error) {
	if len(data) > RequestMaxSize {
		return ErrorDataTooLong
	}

	f, err := os.OpenFile(DevicePath, syscall.O_RDWR, 0)
	if err != nil {
		return
	}
	defer f.Close()

	buf := make([]byte, RequestBufferMaxSize)
	binary.LittleEndian.PutUint32(buf[0:4], RequestTag)
	binary.LittleEndian.PutUint16(buf[4:6], id)
	binary.LittleEndian.PutUint16(buf[6:8], uint16(len(data)))
	copy(buf[8:], data)

	err = IOCtl(f.Fd(), IORequestWrite, uintptr(unsafe.Pointer(&buf[0])))
	return
}
