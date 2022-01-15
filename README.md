# go-rkvendorstorage

A pure Golang implementation of Rockchip `rknand` vendor storage interface.

[![Go Reference](https://pkg.go.dev/badge/github.com/jamesits/go-rkvendorstorage.svg)](https://pkg.go.dev/github.com/jamesits/go-rkvendorstorage)

## Usage

```go
package main

import (
	"fmt"
	"github.com/jamesits/go-rkvendorstorage"
)

func main() {
	size, data, err := rkvendorstorage.Read(rkvendorstorage.IDVendorSN)
	fmt.Printf("size=%d, data=%s, err=%s\n", size, string(data), err)
}
```

See [read_test.go](/read_test.go) for more examples.

## Acknowledgements

The following articles are referenced during the development of this library:

- [rktoolkit/vendor_storage.c](https://github.com/rockchip-linux/rktoolkit/blob/77fb41f99185a6f9dc2c9c69e099f417d8ae905d/vendor_storage.c)
- [RK Vendor Storage Application Note](https://github.com/liihag/RKDocs-1/blob/thzy_develop/Develop%20reference%20documents/Rockchip%20Vendor%20Storage%20Application%20Note.pdf)
- [Golang ioctl sample](https://gist.github.com/tetsu-koba/33b339d26ac9c730fb09773acf39eac5)
- [Another random ioctl sample](https://go.dev/play/p/rq8pJGL3ey)

Thank [Qingping](https://www.qingping.co/) for sponsoring some RK3128 development boards.
