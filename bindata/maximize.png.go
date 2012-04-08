package bindata

import (
	"bytes"
	"compress/gzip"
	"io"
)

// MaximizePng returns the decompressed binary data.
// It panics if an error occurred.
func MaximizePng() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xea, 0x0c,
		0xf0, 0x73, 0xe7, 0xe5, 0x92, 0xe2, 0x62, 0x60, 0x60, 0xe0, 0xf5, 0xf4,
		0x70, 0x09, 0x02, 0xd2, 0x0d, 0x20, 0xcc, 0xc1, 0x06, 0x24, 0x0f, 0xdb,
		0x25, 0x9e, 0x06, 0x52, 0x8c, 0xc5, 0x41, 0xee, 0x4e, 0x0c, 0xeb, 0xce,
		0xc9, 0xbc, 0x04, 0x72, 0xd8, 0x92, 0xbc, 0xdd, 0x5d, 0x18, 0xfe, 0x83,
		0xe0, 0x82, 0xbd, 0xcb, 0x27, 0x03, 0x45, 0x38, 0x0b, 0x3c, 0x22, 0x8b,
		0x19, 0x18, 0xb8, 0x85, 0x41, 0x98, 0x91, 0x61, 0xd6, 0x1c, 0x09, 0xa0,
		0x20, 0x7b, 0x89, 0xa7, 0xaf, 0x2b, 0xfb, 0x1d, 0x66, 0x69, 0x66, 0x43,
		0xdd, 0x17, 0x21, 0x96, 0x2a, 0x40, 0x21, 0xd9, 0xcc, 0x90, 0x88, 0x12,
		0xe7, 0xfc, 0xdc, 0xdc, 0xd4, 0xbc, 0x12, 0x06, 0x10, 0x70, 0x2e, 0x4a,
		0x4d, 0x2c, 0x49, 0x4d, 0x51, 0x28, 0xcf, 0x2c, 0xc9, 0x50, 0x70, 0xf7,
		0xf4, 0x0d, 0x48, 0xd1, 0x4b, 0x65, 0x07, 0x8a, 0xaf, 0xf1, 0x74, 0x71,
		0x0c, 0xa9, 0xb8, 0xf5, 0xf6, 0xce, 0x45, 0x4e, 0x06, 0x05, 0x0e, 0x07,
		0xc7, 0xec, 0xff, 0x73, 0x73, 0x1b, 0xf7, 0x31, 0x39, 0x7e, 0xca, 0x6b,
		0x7e, 0x12, 0x6a, 0x1f, 0xc8, 0xc2, 0x80, 0x09, 0x7a, 0xd7, 0x28, 0xbe,
		0xf6, 0x31, 0xde, 0x75, 0xfb, 0xdb, 0x81, 0x1c, 0x26, 0x01, 0x06, 0x26,
		0x07, 0x06, 0x0e, 0x06, 0x46, 0x05, 0x06, 0x96, 0x06, 0x06, 0x24, 0xce,
		0x81, 0xfa, 0x5d, 0xfb, 0xcb, 0x3c, 0x37, 0x97, 0x31, 0xa1, 0x4b, 0x80,
		0x39, 0x0f, 0x5a, 0xf9, 0x95, 0xda, 0xce, 0x9f, 0xe0, 0x7c, 0xc0, 0x85,
		0x55, 0x33, 0x88, 0x83, 0x05, 0x2c, 0xe5, 0x7b, 0xc3, 0xc2, 0x9c, 0xa2,
		0x26, 0xf1, 0x4a, 0x11, 0xc4, 0xf3, 0x74, 0xf5, 0x73, 0x59, 0xe7, 0x94,
		0xd0, 0x04, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x02, 0x3b, 0x5b, 0x55,
		0x01, 0x00, 0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}