// Code generated by go-bindata.
// sources:
// tmpl/css/style.css
// tmpl/index.html
// tmpl/js/mapping.js
// tmpl/js/request_logger.js
// tmpl/js/util.js
// DO NOT EDIT!

package console

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _tmplCssStyleCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\xcf\x6e\xa3\x30\x10\x87\xcf\xeb\xa7\xb0\x94\xeb\x82\x94\x28\xd2\xae\x8c\xb4\x52\xb6\x6d\x5e\xa3\xf2\x9f\x01\x2c\x26\x1e\x34\x38\x4d\xa2\x28\xef\xde\x81\xd0\xa6\x87\x52\x2e\xc0\xef\x1b\x7f\xe3\x19\x63\x8a\x13\xb8\x2e\xe6\x62\xf0\x4c\x88\xce\xb2\xbe\x2a\x2d\xcf\x29\x86\xdc\x1a\xbd\xde\xf4\xe7\x4a\xdd\xd4\x37\x85\x45\x66\xeb\xbb\xb9\xdc\xc9\x67\xc3\x74\x4c\xa1\xf0\x84\xc4\x46\xaf\x9e\x76\x2f\xdb\xfd\xae\xba\x63\xe2\x00\x5c\x20\xd4\x59\x9c\xfd\x59\x0f\x84\x31\xe8\x95\xf7\x7e\xd1\xde\x1e\x0f\x6e\xd9\xbe\xdd\xfc\x75\xde\xfe\x78\xd8\xb4\xf4\x06\xbc\xac\x58\x3f\xff\xd9\xfe\xdf\x8f\x0a\xa5\xca\x01\x10\x7c\x86\xf0\xca\x74\xd2\x57\x5d\x53\xca\xa2\x8d\x4d\x2b\x17\x76\x84\xa1\xba\x49\x4d\xe6\x98\x1a\xa1\xb3\xa1\x61\x80\x54\x69\x21\x49\xda\x8d\xad\x3e\x48\xb0\xdc\x11\xdb\xd4\xc0\x84\x1d\x11\x82\x4d\x0f\xee\xf0\x08\xf3\x41\xc4\x47\x7c\xb0\x0d\xa4\x6c\x27\xd2\xc1\xe5\x01\x18\xc2\x18\xaa\x6c\x1d\x82\x2e\x25\x45\xdb\x0f\x50\x46\x71\xaa\x5f\x21\x0e\x3d\xda\x8b\x99\x68\x21\x03\x4c\x33\x95\xf7\xdf\xfb\x12\xfe\xe9\xec\x28\x5c\xc6\x37\xcf\xb3\x8e\x85\xe6\x93\x86\xdf\xea\x4b\x3e\x25\xcb\x9b\xab\xeb\x5a\x5a\xbc\x07\x00\x00\xff\xff\xc8\xe4\x9f\xe6\x3d\x02\x00\x00")

func tmplCssStyleCssBytes() ([]byte, error) {
	return bindataRead(
		_tmplCssStyleCss,
		"tmpl/css/style.css",
	)
}

func tmplCssStyleCss() (*asset, error) {
	bytes, err := tmplCssStyleCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/css/style.css", size: 573, mode: os.FileMode(420), modTime: time.Unix(1485530015, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x1a\xeb\x52\xdb\xca\xf9\x3f\x4f\xb1\x47\xbd\x0c\x4c\x22\x09\x63\x08\x04\x6c\xcf\x10\xf0\x69\xc2\xe0\x84\x80\x03\x3d\xe7\x4f\x66\x6d\xad\xad\x35\x2b\xad\xd0\xae\x7c\x49\x86\x99\x3e\x41\xdf\xa3\x3f\xfa\x12\xed\x9b\xf4\x49\xfa\xed\xea\x62\x59\xb6\x6c\x01\x93\xb4\xcc\x1c\xc7\x5a\x7d\xd7\xfd\xee\x9f\x4f\xe3\x97\xf3\x4f\x67\xdd\xdf\xae\xda\xe8\x7d\xb7\x73\xd9\xda\x6a\xb8\xd2\x63\xad\x2d\xf8\x97\x60\xa7\xb5\x85\xe0\xaf\xe1\x11\x89\x91\x2b\x65\x60\x92\x87\x88\x8e\x9b\xc6\x19\xf7\x25\xf1\xa5\xd9\x9d\x05\xc4\x40\xfd\xf8\xa9\x69\x48\x32\x95\xb6\xc2\x3f\x41\x7d\x17\x87\x82\xc8\xe6\x97\xee\xaf\xe6\x91\x81\xec\x84\x92\xa4\x92\x91\x56\xa7\xc3\xfb\xf7\x08\x88\x08\xce\x48\xc3\x8e\x0f\x63\x00\xd1\x0f\x69\x20\x91\x08\xfb\x4d\x43\x71\x14\xc7\xb6\xdd\xe7\x0e\xb1\x46\x0f\x11\x09\x67\x56\x9f\x7b\x76\xfc\xd5\xdc\xb3\x6a\xd6\xbe\xe5\x51\xdf\x1a\x09\xa3\xd5\xb0\x63\xd4\x75\x74\x1c\x7f\x24\xac\x3e\xe3\x91\x33\x60\x38\x24\x9a\x18\x1e\xe1\xa9\xcd\x68\x4f\xd8\x2e\xf6\x1d\x46\x7a\x20\x37\x10\xb4\xf7\xac\x5d\x6b\x77\xf1\x2c\xcf\x24\xe6\xc2\xa8\x7f\x8f\x42\xc2\x9a\x86\x90\x33\x46\x84\x4b\x88\x34\x90\x1b\x92\xc1\x9c\xab\x87\xa7\xc0\xd8\xea\x71\x2e\x85\x0c\x71\xa0\x1e\x14\xe3\xec\xc0\xae\x5b\x75\xeb\x8d\xdd\x17\x62\x7e\xa6\xd5\x82\x13\x03\x51\xb8\xdb\x61\x48\xe5\x0c\x78\xb8\xb8\x7e\xb4\x6f\xd6\x1e\x8e\xbc\xee\xc5\xa7\xd3\x9b\xe9\xd1\xa8\x76\x1a\xbd\xc2\x07\x77\xe7\xb7\xfe\x15\xdd\x63\xf7\xbf\x0e\x26\x93\xf6\x29\x3e\x72\xcf\xcf\x9d\xd1\xef\x2c\xb8\x24\xc3\xa9\x3b\xba\xed\xb4\x6b\x83\xe1\xe8\xee\xea\x2f\xde\xfd\x37\x71\x08\x06\x0b\xb9\x10\x3c\xa4\x43\xea\x37\x0d\xec\x73\x7f\xe6\xf1\x08\xd4\xfb\xc1\x4a\x99\xd2\x25\x1e\x59\xa7\xda\xe0\xf2\x6e\xef\xe3\x6e\x8d\x75\x1e\x46\xf8\xfe\xdd\xfd\xb4\xce\xec\xce\xdb\x36\x76\xa3\x49\x70\x33\x20\x1f\xc7\xb7\x6f\xea\x17\x07\xe4\x9b\x5f\x8f\x7e\xff\x86\x83\xee\x6e\x74\xd8\xfe\x4d\xfc\xb5\x33\xfa\x7c\xfb\x6a\xb7\xed\x1f\x84\x9b\x54\x5b\xe5\x15\x55\x55\x19\x15\xcd\x33\x5a\xa9\xc2\xae\x77\xd3\xbb\x38\x6f\xbf\xa7\x98\x0d\xbc\xe8\xdd\xbb\xcf\x57\x6f\x4e\xf7\x3f\x87\x41\xf8\x70\xf0\xe9\x76\x70\x57\x3f\xbc\xba\xbe\xae\x8f\x0e\xda\x97\x0f\x53\x21\x6a\xb3\xdb\x87\x4f\xd2\x27\x81\xff\xfe\xf6\xea\x2d\xbe\x38\x9c\xde\x94\xab\xb0\xd2\xf9\x16\xcd\xb2\xc1\xc3\x33\xf9\xeb\xa6\x03\xf2\xf1\xa1\x5d\xb3\xea\xfb\xd6\xdb\x82\x99\xe2\x77\x73\x3b\x2d\xf9\x82\x84\xb0\x4f\xa2\x5d\x01\x3c\x3f\xe4\x4a\x05\x1a\x95\xc8\xb3\x2a\x08\x13\xbe\x39\xa1\x2e\xf0\x18\xdf\xe8\x53\x23\x16\x07\xc8\x45\x92\xb2\x35\x89\x62\x03\x76\x08\x89\x8f\x08\xf9\x15\xe4\x18\x92\xf0\xf9\x74\x3c\x1c\x04\xd4\x1f\xae\x22\xb0\x3a\xec\x0a\x57\x9d\xd8\x5b\xd9\x4b\x43\x59\xb9\xfb\xd7\x59\xda\xc7\x1e\xc0\x8f\x29\x99\x04\x3c\x94\xb9\xdc\x3c\xa1\x8e\x74\x9b\x0e\x19\xd3\x3e\x31\xf5\xc3\x6b\xf0\x5e\x2a\xe1\x6e\x4d\xd1\xc7\x8c\x34\x6b\x46\x15\x55\x62\x18\xf5\xf7\xc7\x6d\xe4\xf0\x7e\xe4\x01\x79\xb4\x63\x85\x50\x2e\x66\xdb\x83\xc8\xef\x4b\xca\xfd\xed\x1d\xf4\x1d\x65\x90\xea\x6f\xe1\x61\x8c\x43\x94\x5c\xe9\xa5\xbe\x51\xd4\x44\x3e\x99\xa0\xeb\xfc\xd9\xf6\xce\xc9\x32\xd6\x44\x24\xa0\x77\xa4\x77\x03\x55\x84\xc8\x6d\x63\xa2\x3c\xcd\x40\xaf\x10\xe3\x7d\xac\xb8\x5b\x2e\x17\x12\x9e\x0d\x9b\xf4\x5d\x6e\xec\x9c\x2c\x90\x99\x08\x8b\xfb\x1e\x11\x02\x0f\x09\x10\xcb\x44\x26\x63\xd0\x64\x49\xee\x94\xf1\x1c\xe1\xe2\xe6\xd3\x47\x2b\x50\xe5\x2d\x46\xb1\x1c\x2c\x71\x81\x87\xfa\x5b\x50\xd0\x02\xcf\x69\xfb\x32\x9c\x6d\x27\x84\x0a\x08\x8f\x27\x5b\x5b\x4b\xca\x26\xde\x92\x68\xdc\x89\x9f\xb6\x0d\x50\xf6\xd5\x82\xaa\xaf\x0c\x1b\x07\xb4\xa8\x67\xea\x6b\x03\xca\xd8\x25\x15\x72\x7b\x27\x7b\xfd\x08\xa0\xb1\xad\x33\x1f\x6c\xd8\x71\xc1\xdf\x6a\xf4\xb8\x33\x4b\x3c\xc1\x0d\x93\x2f\x0e\x1d\xa3\x3e\xc3\x42\x80\xef\x81\x47\x61\xea\x93\xd0\x1c\xb0\x88\x3a\x39\x87\x68\xb8\xb5\x62\x69\x87\x93\xf9\xeb\x1c\x91\x90\x4f\x72\x88\xfa\x2d\x16\xd4\x21\x73\x26\xcc\xf4\x1c\x73\x0f\x05\xd8\x31\x21\x15\xba\xd2\xdc\x2d\x20\x68\xa4\x88\xa5\x18\x3e\x1e\x83\xf3\x8f\xcd\x00\xb4\x15\xfa\x9b\x90\x18\xfc\xc3\x59\x81\xa6\x51\x19\x4d\x51\x31\xd8\x7f\x4c\x20\x20\x31\x52\xa6\x34\x25\x58\x8c\x29\xd7\xc7\xbd\x34\xe0\xfe\x90\x18\xd3\x68\x25\x2e\x2a\x1a\x36\x2e\x21\x0c\xa9\xad\x94\xe5\x5a\x1e\x89\xc1\x8c\x56\x62\xe9\xe7\xb3\x88\x09\x8e\x20\x6a\x63\xf3\x1e\x9f\x18\x8b\x6c\x3d\xee\x60\x96\x9e\xe1\x70\x08\x9d\x1a\xf0\x9f\x75\xf4\x71\xeb\xb4\xc7\x23\xf9\x44\xee\x0d\x3b\x62\x05\x8b\xda\xda\xa4\x85\xc3\x05\x4f\xd2\x46\xae\x29\xcb\x2e\xd3\xcb\x01\xc2\x25\x99\x49\x22\x2b\xb3\xe6\x2f\xa6\xd9\xc5\x3d\x54\x33\xcd\x12\x00\x45\x8e\x3a\xe0\x79\x89\x21\xf3\xb4\x03\xec\x13\x34\xc0\xe0\x7e\xd4\x47\xa9\x37\xac\x24\x93\xf2\xa2\x7e\x8f\x4f\x91\xe4\x9c\x41\x6f\x58\xc6\xb3\xa8\xc6\xb2\xd3\xaf\x83\x56\xb7\x33\x15\x66\x6d\x6f\x03\x4e\x11\x2f\x88\x18\x8b\x43\xa6\x02\xa2\x46\xc6\x29\x6a\x4f\xfa\x08\xfe\x33\x1d\x32\xc0\x11\x93\xfa\xfb\x14\x8a\x8e\x6e\xd0\x9b\xc6\x35\x19\x84\x50\x97\x0c\x7d\x8f\xf0\xee\x8c\x11\x1c\x42\x7a\xab\xc0\x47\x83\xae\x57\xbd\xcc\xdd\x16\x81\x40\xd1\x0d\x57\x58\x01\x04\x0c\x68\xab\xeb\x5d\x63\x37\x0d\xb7\xca\x1a\x48\x04\xb8\x4f\xc2\x03\x55\xbf\xd7\x72\xda\xf4\x5a\x09\x01\x1e\xb1\xd6\x79\x14\xcc\xb3\x5c\x4d\x79\x34\x43\xfa\x33\xb3\xe6\xbc\x1e\x19\xf9\x48\x80\xe7\x2b\x05\xb7\xc9\x35\x41\x96\xb4\x06\x32\xa8\x26\x95\x2e\x2f\xc7\xc6\x54\xb9\x26\x1f\x75\x8c\x98\xe0\x4d\x01\x54\x8a\xf5\xe1\x96\x11\xd4\x38\x0b\x24\xf5\xc9\x22\x4d\x14\x53\x76\xf9\x98\xa8\x16\x43\x3b\xac\x99\x95\x2c\xa4\x83\xc3\xe1\x13\x1f\x5a\x32\xd5\x3f\xa5\x92\x30\x3c\x83\x94\x77\x8c\x06\x74\x4a\x9c\xb8\x8c\x4e\x78\xe8\x98\x13\xe8\x3f\x8f\x51\x0f\x9a\x9b\x7b\x53\x1d\x9c\x54\x8d\x29\xa9\xaa\x69\x2a\x99\x4b\x1d\x87\xe8\x50\xaa\x86\x1d\x53\x08\xab\x03\x27\x2c\xf3\xce\x2a\x3c\xb3\x0e\x3e\x0a\x1d\x34\xf7\x87\xad\x73\x2c\xa1\x1a\x27\x0f\x4f\x23\x6b\x4b\xf7\xa5\x82\xd4\xe6\x82\x9c\xc1\x2c\xff\x7f\x21\x48\x87\x48\x97\x3b\xff\x3b\x51\x0e\xe7\xa2\x5c\x61\xe9\xfe\x0c\x41\x00\xba\xa2\x4f\x29\xba\xd9\xfa\x67\x23\xb0\x9c\x77\x8c\x15\x08\x57\x03\x06\x40\x15\x98\x2f\xca\xf5\x55\x32\x70\x9c\x5d\xe3\x44\x59\x31\xb7\x4e\x08\x63\x48\x7d\x98\x02\xa9\xe9\x28\x2d\xb6\x0d\xe2\xb5\x2e\xb1\xc8\x12\x2d\x8a\x02\xc8\x79\xc4\x39\x86\x89\x0a\x18\xe8\xcc\x05\x14\xe4\xd7\xe4\xdc\x68\xdd\x61\x98\xbf\xa0\x43\xb7\x2c\x30\x3f\x80\x40\x49\x01\x1a\xcf\xd1\x67\xdd\x2b\x50\x13\xa9\xde\x68\x6f\x63\x6f\x94\x36\xa0\x2b\x7b\xa3\x0d\x1d\xd1\xc6\x5a\x56\x60\xf2\xa2\x7a\x90\xab\x05\x29\xb9\xf2\x5a\x00\x81\x45\x03\xe2\xe4\x2b\x43\xa5\x72\xf3\x94\x08\x78\x42\xae\x56\xa9\x20\x1e\xc6\x8d\xda\xc1\x9f\x8c\xd6\x97\x90\x3e\x31\x8c\xe7\x04\xea\xbb\x40\xe0\x9c\xc4\x4d\x3e\x4c\x83\xcf\x26\xa4\x04\x49\x33\xe2\x4b\x64\x89\x53\xd9\x0b\x84\xb8\x26\x02\x7a\x95\x27\x93\x68\xfd\x48\x8c\x6a\x79\x33\xcb\x99\x15\x5c\xab\x5a\x0e\xac\xc2\x72\x33\xa5\x8d\x99\xf4\x19\x69\x65\xc5\x79\xe1\x28\xf7\x98\x7c\x8d\xbf\xab\x64\xa4\xa7\x4c\xa4\x87\xcc\x34\x63\xe4\xf3\xab\x9e\x4d\xe3\x94\x13\x47\x78\x32\x96\xa2\x90\xab\x86\x2d\xde\x06\x1a\xab\xf7\x0a\x1a\x37\x59\x18\xa2\xf8\x41\x78\xc5\x4d\xc3\x12\x7c\xf9\x6c\xb9\x0c\xab\x6c\x5c\x9a\x40\x1a\xbd\x48\x4a\xee\x27\xab\xb3\xf8\x21\xcb\x49\x7d\xc6\x05\x49\x66\x6e\x87\x0a\x8f\x66\x44\x8d\xd6\x9f\x25\x85\xe6\xfa\xa4\x61\xc7\x38\x25\xd4\xdd\xfd\x45\x59\xf4\x5c\x66\x24\x3b\x97\x3b\xd2\xcb\xed\x5d\xf6\x57\x4d\xe8\x2b\x8d\xb9\xac\xa2\xf2\xa9\x32\x05\x83\x56\xc7\x53\xdc\xa8\x40\x50\xd2\xde\x77\xbb\x57\x48\x3d\xab\x2d\x15\x24\x62\x46\xe3\xad\x14\x1a\xf0\x10\x6a\xa3\x90\xfa\xdc\x77\xc0\x9a\x50\x0f\x83\x90\x4b\x0e\x57\xa3\x57\x1b\x41\x29\x83\x4b\xda\x27\x7e\x9f\x80\x32\xc1\x4c\x97\x56\xf4\xaf\x7f\xfc\xe7\x6f\x7f\x47\x7b\xbb\xb5\xc3\xd7\xe8\x02\xba\x70\x8a\x3a\x38\x94\xff\xfe\xa7\x8f\xb6\xb3\x65\x87\xda\x38\x1f\xdb\xf6\x48\xbd\xb6\x28\x87\xa1\x35\x59\x6b\x7c\xed\x31\xec\xdf\x1b\xad\x02\x80\x9a\x37\x77\xd6\x89\x71\x4d\x60\x66\x15\x50\x37\x22\x1f\x0c\x8e\x3a\x1f\xba\x30\xf7\x80\x64\x82\xbc\x46\x82\x10\xb4\xc0\x59\x6d\x20\x87\x54\xba\x51\x2f\xfe\x85\xca\x03\xf9\xa8\x7f\xb4\x67\x7b\xea\x76\xec\x1e\xe3\x3d\xdb\x83\x3b\x20\xa1\x7d\xf9\xe1\xac\xfd\xf1\xa6\x6d\x79\xce\xb2\x8c\xc9\xbb\x35\xab\x97\xa0\xa5\xef\xd6\x21\x30\xd0\x30\x61\x29\x8b\x43\x79\x03\xaf\x21\xc7\x4f\x91\x68\xf5\xf5\xac\xc7\x59\x29\x55\x65\xa7\x1a\x70\x2e\x9f\x17\x37\x85\xad\x44\x49\x04\x9d\xa9\xe8\x2a\x8f\x9f\x97\x65\xac\x6e\xbb\x73\x75\x79\xda\x6d\xdf\xa0\x2c\x61\x25\x4b\xf2\xfc\x38\x4a\xd4\x7e\x77\x61\x65\x3f\x35\xe7\x3f\x20\x9a\x92\x78\x01\x83\xc6\x2f\xbf\xed\x82\xce\x61\x9e\x1e\x68\xff\x5e\xb7\x2d\xdf\xbf\xa7\x3f\x37\xc0\xac\xc0\xc3\xc7\xc7\xc2\xa2\x2e\x9b\xa8\x63\x27\x30\xe7\xf0\x7e\xe4\x01\x74\x41\x4d\xe9\x2c\xcf\x85\x73\x14\xd5\x8a\x3e\x3e\x42\x91\x70\x36\xa1\xa9\xe1\xa9\xd7\xca\x0b\xe7\x68\xcc\x5e\xf1\x5a\x2b\x91\x9a\xd3\xf1\x74\xdf\x51\x4d\x86\xc3\x3c\x62\x00\xcd\x46\x11\x6d\xb1\x56\xab\xeb\x5d\xd8\x41\x94\xdc\x58\xca\x45\x40\xcc\xf7\xa1\xa1\xe7\x93\x6c\x3f\x00\x7e\x06\x56\x9b\x1d\x23\x9f\xfb\xc5\xb6\x54\xcb\xc7\x99\x6a\xdd\x9b\xc6\x7e\x59\xf9\x48\x08\x05\xd8\x71\x20\xf3\x1d\xa3\xdd\x60\x8a\x6a\x07\xf9\x8f\xb2\xa8\x70\xeb\xd0\xdc\x69\x89\x21\xa3\xd7\xcb\x81\xe2\xf9\x22\xd1\x81\xe1\x1e\x61\x48\x7f\x9a\x41\x48\x21\x86\x67\xd9\xd6\x3a\x1b\x33\x4a\xa9\xad\x9f\x98\x82\x30\xdb\xd0\xc3\x57\x13\x62\x80\x33\xa6\x1b\xef\x56\x43\xb9\x03\x18\x27\xbd\xd9\x47\x65\x19\x7d\x06\x49\x2b\x2c\x69\x3f\xd6\xb5\x1e\xeb\xf4\x5a\x8e\x10\xd5\x38\xaa\xd9\x81\xfc\x1c\x15\x63\x5e\x3f\x54\xc7\x34\xdd\xb5\x2e\xf9\xf0\xa7\x28\x05\x2d\xd3\x53\x14\xaa\x96\x62\x57\xc7\x66\xee\x27\xdc\xa5\x6c\x9a\x0e\x74\xd5\xb3\x69\x46\xbf\xd8\xa6\x43\x7c\xc6\x29\xeb\xcb\xf5\x87\x38\x53\xad\xcc\x31\x00\xe0\xcc\x67\xa8\x92\x44\xb4\xce\x19\xe3\x0c\x36\xf7\xc5\xcc\x3d\xad\x79\x6e\x4b\x0c\xb8\x9a\x7d\x06\x1f\xa7\xb4\xa7\x4b\x20\x24\x96\x91\x58\x94\x20\xf6\x51\x2b\x7e\x75\x96\x24\xeb\x35\x62\x6c\xfa\x1d\x40\xff\xa3\x7e\x59\x36\xb3\x3d\x81\x2e\x4b\x51\x48\x9b\x46\x72\xc7\x46\xeb\x16\x00\x54\xa7\xf0\x02\x1e\xc4\xa1\x72\x2d\x8f\x36\x00\x2c\xf3\x58\xed\x5e\x60\x75\x3d\x22\xa9\x1f\x38\xd5\xff\xd9\xf4\xdf\x00\x00\x00\xff\xff\x80\xd7\x7d\xbf\xf0\x24\x00\x00")

func tmplIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_tmplIndexHtml,
		"tmpl/index.html",
	)
}

func tmplIndexHtml() (*asset, error) {
	bytes, err := tmplIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/index.html", size: 9456, mode: os.FileMode(420), modTime: time.Unix(1495125698, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsMappingJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x55\x6d\x8b\xdb\x46\x10\xfe\x9e\x5f\xb1\x6c\x0f\x56\x26\xb6\x14\x28\xa5\xe0\x7b\xf9\x90\x6b\x21\x94\xa6\x2d\x5c\xfa\xe9\x7a\x84\xb5\x34\x67\x6d\xb2\xd6\xaa\xbb\xa3\xbb\x13\xc5\xff\xbd\x33\xb2\xe4\xc8\x96\xec\x73\xc0\x03\x06\xed\xce\xcc\x33\x2f\xcf\xcc\xfa\xb1\x2a\x52\x34\xae\x10\x1f\x75\x59\x9a\x62\x19\x65\x6e\xa5\x4d\x31\x11\xff\xbd\x79\x23\x48\x30\x37\x21\xde\xdc\x89\x6b\xb1\xf9\x68\x14\x49\x92\xba\x55\x69\x2c\x90\x09\xfd\x60\x55\x5a\x8d\xd0\xa8\x9e\xb4\x17\xab\x0d\xdc\xe7\xe0\x2a\x9f\x02\xb9\x5e\x44\xf2\x87\xf6\x72\x06\x05\xfa\x5a\x4e\xe2\x1c\x57\x36\x9a\x5c\x0e\x9c\x3a\x34\x72\xfb\xa0\x8b\xcc\xc2\x42\xfb\x10\xb7\xf1\xa2\x5d\x68\x72\xdf\xfa\x3f\x1a\x6b\x7f\x37\x01\xc9\xaf\xab\x2b\xe2\x4a\x44\x2b\x17\xf1\x12\xf0\xb7\xbb\x3f\xff\x88\x64\x8e\x58\xce\x93\x44\x8a\xb7\x6d\x51\xf4\x21\x13\x5d\x9a\xa4\x85\x97\xd3\x6f\x18\x99\x46\xdd\xc7\xe9\xe2\x19\x4a\x34\x50\xb0\xfb\x87\xcb\x1d\x5d\xbf\x56\xd4\x0b\xee\xd1\xc2\x65\x5c\x31\x15\x86\x75\x57\xf2\xb7\xb4\x40\xa7\x79\x13\xa5\x17\xf4\x2b\xd4\xd3\xae\x23\xfb\xc1\x59\x5a\xd5\xbd\x0a\xa8\xb1\x0a\x9f\x53\x67\x9d\x57\x0f\x94\x8e\xa0\x2a\x6f\xf9\xf4\xbe\xbe\x6b\x74\x5d\xcb\x62\x0f\xa1\x74\x45\x80\x78\xe3\x73\xeb\x32\xd8\xcb\x65\x07\x79\x05\x98\xbb\x6c\x1c\xf9\x63\xa3\xeb\x21\xff\x5b\x41\xc0\x78\xe3\x32\x02\xca\xfd\x62\xbe\x09\x65\x9f\xe6\x0e\x64\xc4\xeb\x70\x27\xe9\x12\x8a\x2c\x62\xc8\x3d\xbf\x75\xef\xbc\xee\xc6\x63\x7d\xd9\x9b\xe7\xfe\x9c\xb4\x9f\xad\xfa\x22\x52\xdb\x78\xcc\x87\x9a\xc4\x44\x85\x4a\xad\x49\xbf\xaa\xa9\x50\xf1\x02\x8b\xd9\x93\x81\xe7\x59\x6b\xa5\xa6\xe3\xa3\xc6\xe5\x56\xde\x34\x83\xcf\x31\x27\x31\xc3\x45\x92\xee\x64\x2f\xbf\xef\x99\xc8\x46\x45\xfe\x27\x0c\x66\xea\x0a\xa4\x25\xa3\xe8\x8c\x4d\x6c\x7b\xf2\x37\x8f\x75\x3b\x63\x45\x65\xed\x54\xc8\x7f\x50\x4e\x76\x5c\xdf\x3b\x87\x64\xab\xcb\x5f\x8c\xb6\x6e\x19\x87\xdc\x3d\x47\xc3\xc9\x43\x83\x16\xe6\x42\xb5\x4f\x86\xc8\xe0\xd1\x14\x86\x13\x52\xd3\xe1\x30\x41\x08\x7a\xc9\xe6\x57\xa5\x87\x1b\x45\x35\x84\xba\x40\xfd\xf2\xc1\x2c\x73\x4b\x3f\x8c\xda\x6c\x27\xa4\x52\x57\x49\x63\x75\x94\xd0\x2d\xb1\xa7\x53\x06\x99\xc1\x33\x51\xc6\x66\x34\x79\xa5\x33\x4d\x83\x4f\xe7\x6d\x8c\xf5\x0e\xe8\x4c\x9c\x0e\x29\x65\xcf\x0b\x84\x17\x6c\xca\x52\x57\xfc\xa9\x3d\x68\x61\xb2\x6b\xc9\x87\x59\x55\x12\x00\x74\xcd\x91\x22\x60\x6d\xe1\x5a\xae\x4c\x31\x7b\x36\x19\xe6\xf3\x9f\x7e\x7e\x57\xbe\x5c\xf2\x39\x07\xe6\x6b\xfe\xe3\x3b\xba\x90\x37\x57\x49\x07\x76\xa3\xf6\x5f\x34\xd6\xc4\x4f\xda\x6e\xb9\xdd\x51\x9f\x63\xe8\x98\xd1\x63\xe3\xd6\xa4\x30\xd4\x2f\x2a\x44\x7a\x01\xe7\xe2\x7e\x18\x82\xc5\xea\x05\x58\x0a\x73\xab\x8b\x14\xec\x48\x00\x16\xdd\x30\x35\xef\x71\xd6\xa4\x3e\xf6\x4a\x77\xb2\xb1\x88\x53\xeb\x02\xec\xbf\xff\x9d\xac\x07\xb7\xeb\xe9\x01\xc8\x2e\xcd\x3b\xfd\x04\x07\x92\x4c\x43\xb8\xb5\x3a\x50\xa9\x8a\x37\xa0\xf4\x66\xa5\x7d\x7d\xb6\x8a\x76\x47\x92\x97\x70\x64\x9a\x68\x17\x79\x08\x0e\xd4\x2b\x78\x13\xf4\x17\xfd\x32\x42\x77\x5f\xb0\x2e\x99\xf9\xbf\xfe\xfe\x74\x20\xfb\x4e\x2a\x4f\x3d\xd9\xae\xd4\x51\x53\x5e\x9a\x79\x57\xc0\x71\xd3\x50\xa5\x29\x70\x1f\xf7\x57\x74\xc0\xea\xf6\xdf\x84\xde\x95\xf5\x71\x50\xf0\xde\xf9\x11\x48\x6d\xc1\x63\x24\x7f\xdd\xa8\xe5\xdb\x91\x5d\x9f\x50\xa0\xfd\xc0\xaf\x44\x6b\xeb\xfc\xd4\x34\x52\x12\x37\xf4\x38\x6a\x8e\x9b\x7c\x09\xae\x90\xaf\xb7\x6a\xe3\xa9\xd8\x5a\x1d\x34\x5e\x9f\x3e\xd7\x0f\x27\x3e\xf0\x24\xeb\xff\x03\x00\x00\xff\xff\x3c\x1c\xde\x73\xa3\x0a\x00\x00")

func tmplJsMappingJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsMappingJs,
		"tmpl/js/mapping.js",
	)
}

func tmplJsMappingJs() (*asset, error) {
	bytes, err := tmplJsMappingJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/mapping.js", size: 2723, mode: os.FileMode(420), modTime: time.Unix(1493733009, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsRequest_loggerJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x55\xdd\x6a\xdb\x4c\x10\xbd\xf7\x53\x6c\x94\x80\x57\x9f\xfd\x29\x4e\xe9\x95\x42\xa1\x34\x75\x09\x25\x4e\x69\x9c\x52\x7a\x15\xd6\xd2\xc4\x16\x91\x76\xd5\xd5\xa8\x8d\x09\x7e\xf7\xce\xee\x6a\xad\x1f\x9c\x1f\x81\x8d\x59\x9d\x39\x7b\xe6\xcc\x78\xe6\xbe\x96\x09\x66\x4a\xb2\x1b\xf8\x5d\x43\x85\x57\x6a\xbd\x06\xcd\x43\xf6\x34\x62\xf4\x9c\x9e\x26\xaa\x28\xb3\x1c\x18\x6e\xe8\x03\x45\x99\x0b\x04\xfb\xea\x8f\xd0\xac\x52\xb5\x4e\x80\x7d\x60\x27\x3c\x38\xd6\x8e\xe1\x7f\x90\xa8\xb7\x41\x18\x6d\xb0\xc8\x79\x78\xbe\x07\xfb\x68\x82\x5f\x0a\x99\xe6\xb0\x12\xba\x8a\x1a\x7e\xee\xa8\x08\x6e\xf1\xb8\xc9\xaa\x48\xd6\x05\x61\x67\xee\x64\x2f\xb4\x2e\x53\x22\xb9\xcd\x90\x82\xbc\x4c\xf3\x9c\xf0\x54\x25\x75\x41\xb7\x87\x91\x40\xd4\x3c\x40\x83\x09\xa6\x2c\xb8\x9e\xff\x64\x37\xf3\xef\x3f\xe6\xcb\xdb\xa3\xa3\xa0\x91\x64\x9e\x0a\xf0\x36\x2b\x40\xd5\xc8\xf7\xfc\x86\xf4\x25\xb2\xc5\x42\x25\x0f\xec\x42\xc9\x4a\xd1\x41\x78\xce\x76\x53\xf6\x6e\x36\x9b\x35\xb4\xbb\x81\xdc\x35\x60\x63\xad\xb9\x89\x23\x7d\x55\x28\x8a\xb2\x2b\xdd\xb8\xa3\x5b\x10\x25\x2d\xe1\x2f\xfb\x4c\x69\xb6\xf8\xff\xce\xda\x3b\x7c\x8c\x31\x02\x5d\x40\x27\x3c\xa2\x2b\x6d\x6c\xc8\x26\x2c\x38\x0d\xd8\x64\x1f\x64\x1e\x3e\x80\x2e\x94\xc4\x8d\xc5\x9e\x1d\x0e\x18\xe0\xbf\xd4\x79\xfe\x0b\x84\x76\xf4\xec\x23\x7b\x05\x7f\x49\x85\xad\x1c\x38\x7e\x05\xba\xc8\x64\x8d\xf0\x46\xf0\x12\x12\x25\xd3\x8a\x77\x3c\xd1\x80\xb5\x96\x7b\x5b\x0e\x17\xc4\xf5\xcf\x95\xa8\x7c\x5d\xfa\x3e\xf7\x5b\x6a\x7c\x9c\x13\xf0\xce\xc5\xa4\xe3\x30\x42\x78\x44\xfe\x6c\x4d\x9f\xef\x01\xea\x17\x1b\x4a\x3d\x3d\x35\x02\xc5\xb0\xfe\x44\x80\x75\x45\x95\x34\x2f\x23\x0d\x55\x49\x1d\x06\x91\x3b\xbe\x50\x29\xf4\x4b\x5f\x00\x6e\x54\xda\xc2\xad\x9a\xc8\x9d\xf6\x91\xa5\xc0\xcd\x10\x67\xce\xce\x0f\xb5\x1f\x01\xab\xad\x44\xf1\x78\x99\xad\x37\x39\x7d\x90\x7f\x5d\x7e\xbb\x26\x19\x3a\x93\xeb\xec\x7e\xcb\xbb\x3c\x53\x56\xcb\x14\xee\x33\x09\xe9\x94\xbd\x0f\xc3\x21\xa5\xcb\xe1\xcd\x9c\x0e\xfe\x32\x69\xae\xd6\x6f\xe7\xab\xf3\x57\x24\x26\x2a\x57\x9a\xf8\x6c\x81\xe8\xe7\xa7\xed\xd2\xfa\xcd\x9d\xed\xe1\x68\xd8\x5a\x4f\xde\xa8\x3b\x2a\x64\xcc\x6c\x35\x9b\x93\x98\xed\x5d\xf1\xa9\xc4\xac\x4d\x4a\x93\xf2\xd8\xc8\xdf\x07\xdc\x99\x9e\x8a\x87\x03\xc2\x6a\x37\x1d\x15\xb6\xc0\x84\xca\x1f\x37\x1d\xd2\x9e\xba\x62\xc7\x4d\x2b\xb4\xe7\xa6\xb8\xb1\x2d\x7b\x97\x81\xb2\x8b\x9b\x7c\x77\xbd\x2e\xb5\xa3\x96\x74\xcd\xcd\xd8\x26\x2f\xda\x41\x38\x6c\x53\x3f\x94\x27\x93\xa1\x8b\xb6\xb7\xbd\x8f\xae\xd1\x3d\xb8\xe9\xf6\x7e\x84\x59\x0d\x04\xf7\x1b\x81\x37\x0c\x1d\x54\x77\xa1\xa0\x58\x99\x05\xb4\x52\xa9\x59\x2b\xa5\x86\x12\x64\xca\x0d\x47\x27\xa0\xb7\x16\x86\xc7\xc3\x7f\x7b\x6b\x72\xb3\x6f\xc8\x12\xb2\xa3\x3b\xf5\x35\x88\x74\xdb\xdf\x0a\x23\x2f\x6c\x85\xf2\x22\xa7\x11\x48\xcb\x92\x04\x25\x79\x96\x3c\x1c\x40\xbe\x9c\x06\xa5\x8e\x5b\xaf\x74\xe7\x75\x98\x89\xe3\xf1\x46\x24\x4d\x1c\x25\xf9\xd8\x5e\x31\x9e\xb2\x71\xc3\xa1\xe9\xf7\xe1\x0b\x8d\xbb\x59\x6a\x37\xb2\xa9\x40\x18\x19\x16\xda\x5e\x42\x53\x6d\x82\x81\xc1\x34\x61\x09\x4c\x43\x8d\x96\x7e\x6b\x9b\x15\x33\x32\xdf\xff\x02\x00\x00\xff\xff\x02\x4a\x2c\xba\x1d\x08\x00\x00")

func tmplJsRequest_loggerJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsRequest_loggerJs,
		"tmpl/js/request_logger.js",
	)
}

func tmplJsRequest_loggerJs() (*asset, error) {
	bytes, err := tmplJsRequest_loggerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/request_logger.js", size: 2077, mode: os.FileMode(420), modTime: time.Unix(1493733009, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplJsUtilJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x92\xdf\xaf\x93\x30\x14\xc7\xdf\xf9\x2b\x1a\x62\x2e\xb0\xc9\x98\x66\x2f\x6e\x63\x44\xaf\x44\x1f\x34\x9a\xdc\xeb\x8b\x94\x25\x1d\x74\x0c\xed\xda\xa5\x3f\xd4\x79\xd9\xff\x2e\x65\x64\x63\x5c\x70\x3e\x78\x9b\x6c\xa5\xe7\xc7\xe7\x7c\x7b\x4e\x41\xb9\xd6\x8a\x26\x32\x67\x14\x64\x58\xde\x32\xc2\xf8\x9b\xfd\x9d\x44\x52\x09\x5b\x54\xdb\x2d\x4b\xb1\x03\x1e\x0c\x50\xaf\x7c\x0d\x1a\x1e\xe0\xfb\x3e\x78\x39\x1e\x83\xa2\x00\x8f\xac\x2f\x9a\x79\x7a\x71\x2c\x15\xa7\xc0\x14\x2a\x49\xb0\x10\xe6\xec\xe4\x3d\x00\x4c\x04\xee\x82\x4f\xc6\x93\x3e\x4c\x8a\x68\x86\xf9\x63\x4a\x77\xf4\x4f\xc4\x69\x4e\xb3\x66\xb8\x71\xfc\x37\x7a\xfa\xf0\x11\xcb\x0d\x4b\xed\x6d\xb5\xb5\x7b\x70\xb4\x56\x12\xad\x77\xe1\xbd\xd5\x27\x32\xa7\x6b\xd6\x7d\xd1\x26\xe1\xf3\xa7\xbb\x7e\xc4\x5f\xdb\x75\x41\xf9\xf2\x1f\x20\x6f\xc3\x0f\xe1\x7d\xd8\xcb\x79\x9a\xa6\x8b\x3d\x95\xe8\xd7\xfb\x3c\xdb\x90\xf2\x27\xed\x6f\x82\xd1\xa6\x02\x7d\x06\x7e\xb5\x8d\x38\xde\x11\x94\x60\xdb\xbb\xf1\xb2\xe7\xc0\xba\x41\xdb\xdd\xcc\x72\xce\xe6\xf9\xd1\x4c\xe4\x85\x75\x71\xb4\x66\xda\x7a\x56\x53\xeb\xbc\xe4\xda\xa6\x0d\xa1\x8a\x90\xfb\xfb\xb5\xfb\x75\xec\xbe\x8a\x1f\x26\x87\x02\xc2\x68\xa9\xe2\x22\x5a\x42\x68\xc6\xce\xa0\x0c\x11\x83\xa9\x13\x14\x70\x65\x4b\xae\x70\xb1\x46\x65\x17\x0a\xaa\x08\x71\xe0\xaa\x70\x03\x98\x0e\xed\x60\x0a\x47\x30\x1d\x38\x41\xf9\x15\xe1\x30\x8e\x86\xd0\x8d\xb5\xc7\x09\x1c\x2d\xe7\x74\x7d\x7b\x8b\x64\xb2\x69\xb7\xfc\x07\xe2\x20\x21\xa2\xbc\xb7\x45\xd5\x76\x85\xb9\x35\xbb\xf0\xeb\xc9\x79\x4b\xd3\x1b\x49\x2c\x64\x8d\x68\x33\x4e\x71\xd3\x67\x57\xe3\xf4\xaa\xeb\x7d\xc7\xfb\x56\xb1\xde\x51\xb7\x32\x85\xe4\xe5\xc0\xbb\x92\x8d\x0e\x54\x25\xed\xdc\xbf\xab\x12\xeb\x22\x2b\xc6\x08\x46\xb4\x55\xa5\xc9\xd4\x83\xf8\x57\x9a\x8e\x6d\xa3\xba\x5e\xb3\x35\x17\x3b\x44\xcb\x2c\x24\x84\x6f\x5a\x60\x58\x01\x86\xc0\x32\x17\xfa\x50\xd5\xd1\xc7\xb9\xa7\xe3\x16\x0d\xe4\xa1\x7e\x72\x07\xe3\x4f\x00\x00\x00\xff\xff\x60\x64\x0c\xd1\x73\x05\x00\x00")

func tmplJsUtilJsBytes() ([]byte, error) {
	return bindataRead(
		_tmplJsUtilJs,
		"tmpl/js/util.js",
	)
}

func tmplJsUtilJs() (*asset, error) {
	bytes, err := tmplJsUtilJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/js/util.js", size: 1395, mode: os.FileMode(420), modTime: time.Unix(1493733009, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"tmpl/css/style.css": tmplCssStyleCss,
	"tmpl/index.html": tmplIndexHtml,
	"tmpl/js/mapping.js": tmplJsMappingJs,
	"tmpl/js/request_logger.js": tmplJsRequest_loggerJs,
	"tmpl/js/util.js": tmplJsUtilJs,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"tmpl": &bintree{nil, map[string]*bintree{
		"css": &bintree{nil, map[string]*bintree{
			"style.css": &bintree{tmplCssStyleCss, map[string]*bintree{}},
		}},
		"index.html": &bintree{tmplIndexHtml, map[string]*bintree{}},
		"js": &bintree{nil, map[string]*bintree{
			"mapping.js": &bintree{tmplJsMappingJs, map[string]*bintree{}},
			"request_logger.js": &bintree{tmplJsRequest_loggerJs, map[string]*bintree{}},
			"util.js": &bintree{tmplJsUtilJs, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

