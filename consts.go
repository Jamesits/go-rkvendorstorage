package rkvendorstorage

// DevicePath contains the path to the default Linux device of RK Vendor Storage.
const DevicePath = "/dev/vendor_storage"

// RequestTag must appear at the first 4 bytes of the buffer.
// Probably means "VREQ", but characters are reversed in the actual buffer as these chips are of little endian.
const RequestTag = 0x56524551

// RequestMaxSize is the max size of data you can read or write.
// The value is hardcoded differently in different documentations, so a reasonable default is selected.
const RequestMaxSize = 1024

// RequestBufferMaxSize is the size of the buffer used for ioctl.
const RequestBufferMaxSize = RequestMaxSize + 8

// IORequestRead is used during a read request.
// To calculate on the fly:
//  var IORequestRead = _IOW('v', 0x01, 4)
const IORequestRead uintptr = 0x40047601

// IORequestWrite is used during a write request.
// To calculate on the fly:
//  var IORequestWrite = _IOW('v', 0x02, 4)
const IORequestWrite uintptr = 0x40047602

// These are pre-defined keys.
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
