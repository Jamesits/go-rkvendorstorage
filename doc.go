// Package rkvendorstorage provides a read/write interface to Rockchip rknand vendor storage.
//
// All the operations are done with ioctl on the file `/dev/vendor_storage`.
// request buffer format:
//  struct __attribute__((packed)) rk_vendor_req {
//	    uint32 tag;
//	    uint16 id;
//	    uint16 len;
//	    uint8 data[1024];
//  };
//
package rkvendorstorage
