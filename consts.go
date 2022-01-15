package rkvendorstorage

// DevicePath contains the path to the default Linux device of RK Vendor Storage.
const DevicePath = "/dev/vendor_storage"

// request buffer format:
//   struct __attribute__((packed)) rk_vendor_req {
//	     uint32 tag;
//	     uint16 id;
//	     uint16 len;
//	     uint8 data[1024];
//   };

// RequestTag must appear at the first 4 bytes of the buffer.
// Contains the string "VREQ".
const RequestTag = 0x56524551

// RequestMaxSize is the max size of data you can read or write.
// Taken from https://github.com/rockchip-linux/rktoolkit/blob/77fb41f99185a6f9dc2c9c69e099f417d8ae905d/vendor_storage.c#L31
const RequestMaxSize = 1024
const RequestBufferMaxSize = RequestMaxSize + 8

// IORequestRead is used during a read request
const IORequestRead uintptr = 0x40047601

//var IORequestRead = _IOW('v', 0x01, 4)

// IORequestWrite is used during a write request
const IORequestWrite uintptr = 0x40047602

//var IORequestWrite = _IOW('v', 0x02, 4)

const (
	IDUnknown uint16 = iota
	IDVendorSN
	IDVendorWiFiMAC
	IDVendorLANMAC
	IDVendorBluetoothMAC
	IDVendorHDCP14HDMI
	IDVendorHDCP14DP
	IDVendorHDCP2x
	IDVendorDRMKey
	IDVendorPlayReadyCert
	IDVendorAttentionKey
	IDVendorPlayReadyRootKey0
	IDVendorPlayReadyRootKey1
	IDVendorSensorCalibration
	IDVendorReserve14
	IDVendorIMEI
	IDVendorCustom
)
