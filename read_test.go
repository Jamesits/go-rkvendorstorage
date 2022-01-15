package rkvendorstorage

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestReadSN(t *testing.T) {
	// test read VENDOR_SN_ID
	size, data, err := Read(IDVendorSN)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}
	fmt.Printf("size=%d, data=%s\n", size, string(data))
}

func TestReadWiFiMAC(t *testing.T) {
	// test read VENDOR_WIFI_MAC_ID
	size, data, err := Read(IDVendorWiFiMAC)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		t.Fail()
	}
	fmt.Printf("size=%d, data=%s\n", size, hex.EncodeToString(data))
}

func TestReadNonExistent(t *testing.T) {
	// make sure it is not set on your platform
	size, data, err := Read(IDVendorCustom + 1)
	if err == nil {
		t.Fail()
	}
	fmt.Printf("err: %v\n", err)
	fmt.Printf("size=%d, data=%s\n", size, hex.EncodeToString(data))
}
