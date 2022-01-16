package rkvendorstorage

import (
	"encoding/binary"
	"os"
	"syscall"
	"unsafe"
)

// Read reads the value corresponding to the given key from the vendor storage.
//
// If the on-disk data size exceeds RequestMaxSize, only the first [0:RequestMaxSize] bytes will be returned, and err
// will be set to ErrorDataTooLong. There is no way to read the trailing parts of the data.
func Read(id uint16) (size uint16, data []byte, err error) {
	f, err := os.OpenFile(DevicePath, syscall.O_RDWR, 0)
	if err != nil {
		return
	}
	defer f.Close()

	buf := make([]byte, RequestBufferAllocationSize)
	binary.LittleEndian.PutUint32(buf[0:4], RequestTag)
	binary.LittleEndian.PutUint16(buf[4:6], id)
	binary.LittleEndian.PutUint16(buf[6:8], RequestMaxSize)

	err = IOCtl(f.Fd(), IORequestRead, uintptr(unsafe.Pointer(&buf[0])))
	if err != nil {
		return
	}

	size = binary.LittleEndian.Uint16(buf[6:8])
	data = buf[8 : 8+size]
	if size > RequestMaxSize {
		err = ErrorDataTooLong
	}

	return
}
