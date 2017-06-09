// Code generated by go-bindata.
// sources:
// report/content.tmpl
// report/footer.tmpl
// report/header.tmpl
// report/test.tmpl
// DO NOT EDIT!

package perfTestUtils

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

var _reportContentTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x58\xdb\x6e\xdb\x38\x13\xbe\x76\x9f\x62\xa0\x3f\x41\x12\xa0\xb5\x93\xa6\x09\xf0\xab\xb2\x01\x27\xdb\x6d\xd3\x6d\xba\x46\x9d\xed\x4d\xd1\x0b\x5a\x9a\xd8\xdc\xc8\xa4\x96\xa4\x9d\xb8\xaa\xde\x7d\x41\x9d\x0f\x94\xec\xb4\x0d\x56\x80\x2f\x2c\xce\x89\x33\x1f\x87\xf3\x29\x0c\x3d\xbc\xa5\x0c\xc1\x72\x39\x53\xc8\x94\x15\x45\xcf\x00\x1c\x8f\xae\xc1\xf5\x89\x94\x43\x4b\xf1\xe0\x82\x08\x6b\xf4\x0c\x4a\x8f\xb3\x38\xc9\xd6\x03\xe2\x79\x94\xcd\xad\x51\x18\xf6\x2f\x39\xbb\xa5\xf3\xfe\x78\x72\xf5\x91\x2c\x31\x8a\xc0\xb6\x61\xbc\x52\x7c\x49\x14\x7a\x30\x41\x71\xcb\xc5\x92\x30\x17\xe1\x06\xa5\x82\x4f\x18\x70\xa1\xb4\xd0\x61\x18\xf6\xf5\xf2\x54\x11\x25\xfb\x6f\x51\xe9\xf5\x1b\xaa\x6d\x1c\x39\x83\xc5\x49\xe1\xdd\x19\x78\x74\x5d\xfa\x5b\x8a\xd4\xa3\xeb\x77\x48\x92\x60\x8a\x68\x1d\x45\x66\x3e\x1a\x64\x60\xc6\x85\x87\x62\x68\x1d\x5b\x70\x4f\x3d\xb5\x18\x5a\xff\x3f\xde\x2f\xa9\x3a\x4a\x54\x77\x5d\xc9\x80\xf2\x32\xad\x33\xad\xe5\x2c\xce\x1b\x19\x79\xc7\xa5\x82\x15\xf3\x50\x80\x42\xa9\x6c\x28\x52\x74\x43\xc4\x1c\x95\x16\x88\x22\xbb\xfe\x7a\xc2\x85\x8a\x22\x67\xb0\x38\x1f\x39\x03\xe5\xb5\x07\xd1\x11\xd4\xcb\xb3\x96\xa0\xa6\x28\xd6\xd4\x45\x59\x0b\xcc\x47\x06\xa5\x12\xa4\x52\x9f\x50\x06\x9c\x49\xd4\xa5\x90\x4f\x16\xd2\x16\xb3\xce\xa0\xad\x10\xce\x20\x2e\x6e\xdb\xe2\x2e\x48\x01\xce\x5c\x9f\xba\x77\x43\x6b\x41\x3d\xbc\xc6\x25\x17\x9b\x31\x23\xfe\x46\x52\x79\x78\x54\x87\xfd\x0f\x63\xa9\xb0\xd0\xb6\x95\x06\x9e\x4e\x1b\x79\x4a\xa2\x83\x2c\x3c\x67\xb0\x38\xed\x4a\xdb\xf6\xcc\x83\x54\x1b\x1f\x87\xd6\xfd\x82\x2a\x7c\x21\x03\xe2\xa2\xcd\xf8\xbd\x20\x81\x35\x1a\xfb\x3e\xbf\x47\x0f\x3e\x13\x41\xe3\x23\x5b\x86\x6f\xbc\xa8\x73\x31\x41\x72\x97\x84\x95\xcb\x7d\x87\x40\x50\xa6\x6e\xc1\xda\x7f\xd5\x7f\x79\x6b\x45\xd1\xfe\xb6\x02\xef\x80\x91\x30\xa4\xb7\xd0\xbf\x92\x89\xb3\x09\x91\x12\xa2\xc8\xb9\xe5\x4c\x81\xcb\x7d\x2e\x86\xd6\x5c\x20\x32\x6b\x34\x19\x4f\xa7\xce\x40\x2f\x8c\xc2\x10\x7d\x89\x35\x31\x81\x9e\x35\xfa\x7d\x7c\xf5\xa1\x10\x62\x5e\x27\xb4\x9b\xf8\x6b\xe0\xce\x84\x35\xea\x0d\xad\x65\x1c\xed\x25\x67\x8a\x50\x86\x8d\x2e\x5a\x6e\xb3\x71\x36\xb3\xdd\x1a\x70\x53\x46\x5e\x5e\xbf\x4e\xac\x75\xe2\x2d\xcb\x7b\x66\xe2\xe4\xfc\xd8\x1a\x39\x17\xa3\x0b\x22\x11\x74\x55\x21\xc9\xb4\xed\x0c\x2e\xb6\x1c\x79\x47\x79\xba\xf9\x6b\xcd\xa2\x81\x24\xff\x32\x70\xc0\x77\x58\xe2\xf2\x86\x5f\x5f\xc0\x77\x88\x2f\x01\x75\x8d\xcb\x28\xba\xbe\xd8\x6a\x3a\x0f\xf0\x4c\x07\x38\x1b\xc5\x37\x47\x35\xc0\xd9\x6e\x01\x16\xc1\xfd\xda\xc0\x4e\x92\xc0\xf6\xf3\xa3\xb2\x5b\x48\x10\x43\x9a\x71\x55\x85\x75\x14\xa5\x87\x32\xc6\xab\xad\xe1\x9a\x42\x34\xd9\x43\xfd\xbc\x4d\x50\xb8\xc8\x14\x99\x63\x75\x07\xfb\x8f\x6d\xa8\xc6\x66\x9a\x02\xbb\x0d\xb6\x07\x6e\x06\xed\x03\x83\xc1\xb2\xdc\x82\x08\x65\x90\x81\xd2\x69\x39\xf8\x40\x19\x5e\x26\x82\xb5\x03\x55\x0b\x67\xdb\x2b\xe9\x0a\x1a\xa8\xa6\xfa\x9a\x08\xc8\x9d\xbc\x9f\xc2\x10\xdc\xd3\xfe\x1c\x19\x0a\xa2\xf0\x30\x6c\xc8\x7b\x44\x11\x1b\x9a\xef\xf5\xe3\x72\x7f\xb5\x64\xd2\x86\x2f\x6d\x45\x0e\xc3\xbf\x25\x67\xd7\xb8\x04\x4b\x9f\x06\x0b\x6a\x47\x24\xbd\x6c\x56\x1e\x55\x51\xf4\x7c\x07\x2b\x1a\xfa\x56\xf9\x96\xae\x58\x30\x1a\xf8\xda\x78\x6b\xf0\xf4\x8d\xf3\x65\xdb\x36\x91\x69\x50\x78\x36\x28\xb1\xc2\x5d\x8c\x81\x5c\xcd\xe2\x6a\xb7\x59\x94\x0b\x7e\xbf\xbb\x39\xf2\x40\x65\x9b\xa5\x4d\xdb\x82\x7e\x7c\x32\x43\xdf\x86\x83\xf4\x9c\x1f\xfe\x71\x71\x74\x60\x14\x36\x38\x35\xbc\x9a\x0b\xea\xb5\xba\x7b\xe8\x0c\x84\x32\xec\x82\x49\xe5\x09\xc3\xfe\x7b\xc9\x99\xae\xf4\x84\x08\x45\x15\xe5\x4c\xb6\xd4\xb6\xfc\x34\xeb\x9c\x3f\x66\xe5\xea\xdb\xe8\xe8\x75\x55\x6a\xef\xd0\xfa\x5f\x7e\x52\xac\xa3\x3e\x09\x02\x64\xde\x61\xe9\xf0\xf4\xd1\xc7\x25\x32\x55\xd3\x74\x06\xf5\xc3\x57\x6f\x21\x3b\x0d\x62\xe9\xfc\xf9\x1f\x4d\x62\xc6\xe9\x2b\x0d\x09\xb2\x99\x18\xf4\x50\xfc\x88\x61\xec\x69\x06\x30\xc3\xa0\xfe\x73\x93\x98\x79\x42\x2f\x4f\x50\xd9\x28\xa6\x7d\x25\x37\x56\x69\xf4\xca\x66\xae\x7c\xcc\xca\xe7\xab\x78\xda\x7a\x9a\x31\x4b\x26\x49\x30\xcd\x59\x3b\xce\x58\x29\x9e\x2a\x88\xe9\xf5\x7a\xbd\x78\xaf\xf8\x0f\xf4\xf5\x89\x9c\x2a\x7d\x4f\xcc\x37\x60\x4d\x57\x54\xa1\xee\xe6\x9e\x26\xcc\x5a\xb0\xe7\x28\x91\x55\x73\x46\xdc\xbb\xb9\xe0\x2b\xe6\xd9\x1f\xe8\x7c\xa1\xde\x0a\xb2\x79\x0d\x0a\x1f\xd4\x0b\xe2\xd3\x39\xb3\x85\x7e\x9b\x7a\xd0\x9a\x9e\x4e\xad\x0c\x08\x1b\x5a\xaf\x72\x4c\xe8\x7c\xbd\x90\xf4\x1b\xda\x72\x49\x7c\x1f\xc5\x6b\x30\xc1\xe4\xcf\x35\x8a\xb1\xef\xc3\x25\x5f\x31\x25\xed\x24\xb7\x85\xe1\xc7\x19\xbb\x11\x84\x49\xe2\xc6\x5d\x07\xbe\x54\x66\xa6\xd4\x4f\x2c\x11\xfb\x8a\xa2\xaf\x3f\xe7\xec\x8d\x10\x5c\xb4\xb8\x89\xd7\x7e\x8d\x9b\x9b\xc9\xb4\x65\x2b\x93\xa9\xe1\x88\x94\xbd\x25\x90\x4c\x70\x10\xa3\xd8\xd4\x45\xba\xaa\xde\x3a\x95\xd7\x08\x4f\x32\xd6\x7e\x24\x4b\xec\x1e\x1d\xcb\x63\xe7\x69\xa2\xa8\x51\x18\x77\xa2\xc3\x6b\xea\xfb\xf4\xe8\xd1\x06\xb2\x4f\x2d\x3f\x6c\x60\x7f\x9d\x36\x9c\x42\xb3\xb7\xfb\xd9\xd9\xe6\xe4\x65\x1a\x65\x8e\xbb\x6a\x7c\x0d\x7d\x83\x6e\x01\xa6\x5a\x84\x3d\x93\xa3\xc9\xb4\xb1\x8f\xa4\xf6\x2d\x1d\x2b\xc5\x87\x20\x6c\x8e\xb0\x77\x87\x9b\xe7\xb0\x37\xd3\x2c\xca\x1e\xc2\x9e\x81\x15\x99\x3f\xad\x64\xae\xf6\xc8\x7a\xae\x35\x29\xf3\xf0\x01\xf6\xb6\x7c\x92\x89\xfd\x95\x94\x55\x20\x3b\x95\x35\xe4\xeb\x2a\xc2\xed\x56\xc9\x13\x5f\xd7\xc4\x2d\x9a\x45\xda\x7f\x3e\x4c\xcd\x94\x92\xec\x82\x35\x55\x44\xa8\xbf\x82\xbc\xf3\xc6\xad\x77\x81\xfa\xcc\x0d\x4f\x8e\x83\x87\xac\xb8\xbd\x84\xf7\x25\x86\xca\x55\x4f\x17\xf4\xcd\x90\x54\xea\x04\xcf\xab\xe4\xa9\x4d\x5a\x17\xa7\x5b\x38\x03\x7d\x2c\x7a\x9c\x87\x58\x69\x5c\x05\xb3\x8b\x6f\xc8\x37\xbf\xd5\x2c\x24\x37\x68\x45\x33\x36\xbb\xd7\xbf\x92\x59\x86\xd2\xbb\x37\x4d\x53\xa6\xd3\x41\x1d\xc9\x7a\xfe\x99\x88\x24\xae\x64\xd7\xc6\xf1\xa0\x1a\x48\x8c\xfb\xfa\xce\xb6\xdf\x85\x45\xee\x95\x70\x6b\x09\xca\x56\xb0\x75\x45\xa3\xa3\x19\x5b\x47\x68\x45\x9b\xae\x2c\x98\x7a\xf6\x4e\x23\xc5\x16\x6e\xbb\x0b\xaf\xcd\x39\xed\x8c\x88\x76\x4a\x5b\x77\x5e\xfb\x6b\xa0\xb1\x9a\xc1\x66\x26\x9f\x9c\xc0\xe6\x6c\x84\x2e\x71\x2c\x04\xd9\xb4\x91\x4c\x33\x79\x55\x9b\x00\x6d\xd0\x19\x68\x12\xaf\xa7\x66\xa1\x33\x22\xda\x6c\xc5\xfd\xbe\x8b\xb0\x09\xa2\x28\xb7\xe1\xb8\x7f\xb6\x03\x73\x6a\xf1\xdf\x46\x5b\x5b\x29\x6b\x46\x57\xc7\xeb\x79\xcc\x30\xa0\x74\x25\x4f\xd1\xe5\xcc\x93\x06\xfe\xda\xf2\xd9\xa0\x93\x8f\xa6\x65\x71\xf5\xf1\xe5\x62\x73\xd0\xfa\xe5\x01\x52\x11\xaa\xe9\x6b\x09\x0a\x69\x0b\xd2\xd3\x4a\x1b\x35\x35\xe4\xa8\x93\x6f\x6a\xba\x99\xc1\xba\x60\x9b\x05\xd0\x0d\x64\xb3\x95\x68\xe6\x7f\x67\x62\xd4\xf5\x7b\x96\xf6\x87\x7f\x03\x00\x00\xff\xff\xe5\x75\xba\xdd\x04\x1b\x00\x00")

func reportContentTmplBytes() ([]byte, error) {
	return bindataRead(
		_reportContentTmpl,
		"report/content.tmpl",
	)
}

func reportContentTmpl() (*asset, error) {
	bytes, err := reportContentTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "report/content.tmpl", size: 6916, mode: os.FileMode(420), modTime: time.Unix(1497024437, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _reportFooterTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xcb\xcf\x2f\x49\x2d\x52\xaa\xad\xe5\x52\xb0\xd1\x4f\xca\x4f\xa9\xb4\xe3\xb2\xd1\xcf\x28\xc9\xcd\xb1\xe3\xaa\xae\x4e\xcd\x4b\xa9\xad\x05\x04\x00\x00\xff\xff\x70\x9a\x96\xda\x2c\x00\x00\x00")

func reportFooterTmplBytes() ([]byte, error) {
	return bindataRead(
		_reportFooterTmpl,
		"report/footer.tmpl",
	)
}

func reportFooterTmpl() (*asset, error) {
	bytes, err := reportFooterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "report/footer.tmpl", size: 44, mode: os.FileMode(420), modTime: time.Unix(1493141754, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _reportHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x97\x5d\x6f\xe2\x38\x17\xc7\xaf\x87\x4f\xe1\xc9\x73\xd3\x91\x1e\xe7\xa5\x01\x26\x61\x01\xa9\x6f\xda\x56\xea\xee\x54\xdb\xde\x54\x5a\x69\xe5\xd8\xc7\xc4\xad\x63\x67\x6d\x43\xcb\x20\xbe\xfb\x2a\xa1\x2b\x94\x40\x5b\xe8\x8b\xb4\x91\x90\x62\x38\x3e\xff\xdf\x39\xff\x83\x21\x8b\x05\x03\x2e\x14\x20\x2f\x07\xc2\xc0\x78\xcb\x65\x67\xf8\xf5\xf4\xc7\xc9\xcd\xed\xd5\x19\xca\x5d\x21\xc7\x9d\x0e\x42\x08\x0d\xcf\xcf\x8e\x4e\xc7\x9d\x2f\x5f\x86\x05\x38\x82\x68\x4e\x8c\x05\x37\xf2\xa6\x8e\xe3\xc4\xfb\x73\x5c\x07\xd5\x81\x37\x17\x37\x97\x67\xe3\xc5\xc2\x3f\xd1\x8a\x8b\x89\x7f\x74\x75\xf1\x3b\x29\x60\xb9\x44\x47\x53\xa7\x0b\xe2\x80\xa1\x2b\x30\x5c\x9b\x82\x28\x0a\xe8\x06\xac\x43\x7f\x40\xa9\x8d\x1b\x06\xab\xcd\x9d\xce\x3a\x9d\xa5\x46\x94\x0e\x59\x43\x47\x5e\xee\x5c\x39\x08\x02\x1a\xdf\x59\x5f\x9b\x49\x70\x67\x83\xbb\xbf\xa7\x60\xe6\x38\xf2\xa3\xc8\x0f\xfd\x42\x28\x9c\x46\x61\xd8\xef\xf3\xcc\xbf\xb3\x1e\x72\xf3\x12\x46\x9e\x83\x47\x17\xdc\x91\x19\x59\x25\xf3\xc6\xc3\x60\x75\x37\xde\x59\x87\xc5\x38\xf6\x7b\x4f\x12\x71\xd6\xeb\x77\xe3\xf4\xf0\x63\x25\x68\x5c\x27\xef\xd2\x5e\x06\x3c\xe1\x7b\x26\xff\x8a\xf1\x6b\x02\x5c\x4f\x15\x23\x4e\x68\x55\x0b\x45\x8c\x43\x12\x45\xe1\x4e\x42\x18\xef\xa7\x95\x8b\x49\x2e\xc5\x24\x77\x7e\x49\xe8\x3d\xee\x12\xde\x0b\xc3\x2e\xfb\x14\xad\x19\x28\xa6\x4d\x50\x68\x06\x46\x89\x9f\x06\x1f\xfa\x7d\x3f\xaa\x6b\xec\x27\x9c\x51\x9a\xa6\x3b\xeb\xae\x85\xa5\x50\xf7\x28\x37\xc0\x37\x35\xa9\xb5\x81\xaa\xe6\x57\x8a\x9f\x80\xa1\xdb\xef\xd1\x2c\xe9\xfb\xd4\x5a\x0f\x15\xc0\x04\x19\x79\x96\x1a\x00\xe5\x21\x03\x72\xe4\x59\x37\x97\x60\x73\x00\xd7\x80\xa8\xe3\x83\xf1\xce\x92\x2d\xff\xd2\xef\x09\xeb\x52\x48\x3e\x5d\x97\xc6\x38\x0b\xe3\xe8\xb0\xc7\xc9\xa7\x6b\xd5\x69\x70\x9a\xf2\x2c\x49\x93\xf4\xfd\x72\xd5\xf8\xbc\x26\x49\x8a\x52\x82\x0d\xaa\x33\xcd\xfd\x65\x4b\x29\x14\x60\x46\xe2\x94\xc4\x00\x6f\x25\xd8\x98\xe1\x57\x20\x9c\x2e\xb4\x31\xfa\x01\xb3\xef\x94\x87\xe9\x61\xf4\x1e\xe1\xb5\xf2\xf5\xcd\xed\xe5\x59\x3b\x6a\x4d\x56\x5d\xc7\x3f\x4e\x6f\xd1\xa2\xf1\x56\x75\x65\x84\xde\x4f\x4c\x35\x71\x03\xf4\x90\x0b\x07\xbf\x6c\x84\x50\x2d\xb5\x19\xa0\x4c\x12\x7a\xdf\xf8\x70\xd9\x69\x2c\x7d\xa7\xcb\x63\x62\x36\x35\x4a\xc2\x98\x50\x13\xec\x74\x39\x40\x38\x0a\xcb\xc7\x4d\x91\x35\x07\x7e\xd2\xbb\xac\x8e\x95\x6b\x49\x1c\xfc\x6a\xc8\x7c\xcb\x0e\x6d\x18\x18\x5c\xb7\x69\x80\x26\x46\xeb\x59\x0b\xbe\xc5\x97\x47\xfe\x13\xc9\x96\x3e\xfc\xcb\x28\x81\xbb\x01\x4a\xb6\x32\x72\xad\xdc\x00\xc5\x61\xf9\x88\x88\x11\x44\xfe\x1f\x59\xa2\x2c\xb6\x60\x04\x7f\xb6\x6d\xcf\x34\xb5\x72\x09\xdb\x9c\x30\xfd\x30\x40\x51\xf9\x58\xbf\xea\x92\x37\xab\x6d\xd7\x11\xbf\xb3\x8e\x8f\x24\x6b\x80\xf5\xff\x43\x60\xcd\xd1\x64\x62\x76\x0e\xa4\x02\x78\xe9\x2b\xd0\x18\xbd\x63\x39\xdd\x02\xd1\x9c\x3a\xa5\xd5\xcb\x33\xe7\x3b\x92\x49\xb8\x7a\x73\x57\x5a\xe9\xfe\x57\x40\xa1\xcd\xfc\x44\x2b\x47\x84\x02\xb3\x25\x63\x0e\x15\xfc\xa0\x1b\xbe\x94\x6b\x18\xd4\x27\xc6\xc6\x9f\x96\xe6\x91\x31\x23\x06\x09\xfb\x5b\x2d\x79\xa4\x88\x9c\x5b\x61\xcf\x05\x63\xa0\xd0\x08\x71\x22\x6d\xab\x76\x3e\x55\xb4\xfa\xd9\x42\xb9\x60\xd0\xdc\x76\xf0\x6d\x93\x54\xf0\x83\xed\xd9\xb7\xc4\x56\x17\xd3\x74\x5a\x80\x72\xfe\x04\xdc\x99\x84\xea\xf6\x78\x7e\xc1\x0e\xbc\x56\x53\xbc\x6f\x7e\x6d\x8f\xcf\x84\x2d\x25\x99\xa3\x11\xf2\x84\xaa\x8e\x7b\x6f\xd3\xcf\x1a\x64\x8f\x1a\xeb\x46\x82\xb4\xf0\xf1\x8c\xd5\x30\xed\x4b\xe8\xcc\xb6\x21\x6d\x4d\x4d\x6b\xb9\xb2\xf5\x1a\xcc\x4c\x50\xd8\xdb\xd7\xd6\xbe\x67\x8d\xdd\x9a\x7f\x5f\x67\xed\x2a\xc9\x7b\xac\xdd\xa3\x4e\xf4\x46\x6f\x77\xa1\x7c\xd1\xdc\xe7\x18\xf7\x75\xb7\xf9\xbc\x30\x0c\x56\x0f\x73\xab\x45\xa6\xd9\x7c\xdc\x59\x2c\x40\xb1\xe5\xf2\x9f\x00\x00\x00\xff\xff\xfa\x02\x30\xe2\x10\x0e\x00\x00")

func reportHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_reportHeaderTmpl,
		"report/header.tmpl",
	)
}

func reportHeaderTmpl() (*asset, error) {
	bytes, err := reportHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "report/header.tmpl", size: 3600, mode: os.FileMode(420), modTime: time.Unix(1497026095, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _reportTestTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x97\x5d\x6f\xf2\x36\x14\xc7\xef\xf9\x14\x7e\xb2\x9b\x3e\xd2\x9c\x17\x02\x3c\x84\x01\x52\x5f\xd0\x5a\xa9\xdb\x53\xad\xdc\xf4\x6a\x72\xec\x63\x62\xea\xd8\x99\x6d\x68\x69\xd5\xef\x3e\x25\x74\x42\x09\xb4\x85\xbe\x48\x8b\x84\x64\x83\x7d\xfe\xbf\x73\xfe\x07\x83\x87\xdf\xce\x7e\x9e\x4e\x6f\xae\x26\x28\x73\xb9\x1c\xb7\x5a\x08\x21\x34\x3c\x9f\x1c\x9f\x8d\xab\x61\x35\x9d\x5e\x4c\x2f\x27\xe3\xc7\x47\xff\x54\x2b\x2e\x66\xfe\xf1\xd5\xc5\x9f\x24\x87\xa7\x27\x74\xbc\x70\x3a\x27\x0e\x18\xba\x02\xc3\xb5\xc9\x89\xa2\x80\xa6\x60\x1d\xfa\x0b\x0a\x6d\xdc\x30\x58\x6f\x6e\xb5\x36\xe1\x2c\x35\xa2\x70\xc8\x1a\x3a\xf2\x32\xe7\x8a\x41\x10\xd0\x78\x6e\x7d\x6d\x66\xc1\xdc\x06\xf3\x7f\x16\x60\x56\x38\xf2\xa3\xc8\x0f\xfd\x5c\x28\x9c\x44\x61\xd8\xeb\xf1\xd4\x9f\x5b\x0f\xb9\x55\x01\x23\xcf\xc1\xbd\x0b\xe6\x64\x49\xd6\xc1\xbc\xf1\x30\x58\x8f\xc6\x7b\xeb\xb0\x18\xc7\x7e\xf7\x59\x22\x4e\xbb\xbd\x4e\x9c\xb4\x3f\x57\x82\xc6\x55\xf0\xa8\x9d\x44\x6d\x9e\xf6\x0e\x0c\xfe\x0d\xe3\xb7\x04\xb8\x5e\x28\x46\x9c\xd0\x6a\x2d\xc4\x38\xf4\xa3\x28\xdc\x4b\x08\xe3\xc3\xb4\x32\x31\xcb\xa4\x98\x65\xce\x2f\x08\xbd\xc5\x1d\xc2\xbb\x61\xd8\x61\x5f\xa2\xb5\x04\xc5\xb4\x09\x72\xcd\xc0\x28\xf1\x60\x70\xdb\xef\xf9\x51\x95\x63\xaf\xcf\x19\xa5\x49\xb2\xb7\xee\x46\x58\x0a\x75\x8b\x32\x03\x7c\x5b\x93\x5a\x1b\xa8\xb2\x7f\xa5\x78\x00\x0c\x9d\x5e\x97\xa6\xfd\x9e\x4f\xad\xf5\x50\x0e\x4c\x90\x91\x67\xa9\x01\x50\x1e\x32\x20\x47\x9e\x75\x2b\x09\x36\x03\x70\x35\x88\x6a\x7d\x30\xde\x5b\xb2\xe1\x5f\xf2\xa3\xcf\x3a\x14\xfa\x5f\xae\x4b\x63\x9c\x86\x71\xd4\xee\x72\xf2\xe5\x5a\x55\x18\x9c\x24\x3c\xed\x27\xfd\xe4\xe3\x72\x65\xfb\xbc\x25\x49\xf2\x42\x82\x0d\x68\x46\x8c\xfb\xdb\x16\x52\x28\xc0\x8c\xc4\x09\x89\x01\xde\x4b\xb0\xd5\xc3\x6f\x40\x38\x9d\x6b\x63\xf4\x1d\x66\x3f\x28\x0f\x93\x76\xf4\x11\xe1\x8d\xf2\xf5\xf4\xe6\x72\xd2\x5c\xb5\x21\x2b\x9f\x93\x9f\x67\x37\xe8\xb1\xf6\x56\xf9\xa4\x84\xde\xce\x4c\xd9\x71\x03\x74\x97\x09\x07\xbf\x6d\x2d\xa1\x5a\x6a\x33\x40\xa9\x24\xf4\xb6\xf6\xe1\x53\xab\x36\xf5\x9d\x2e\x4e\x88\xd9\xd6\x28\x08\x63\x42\xcd\xb0\xd3\xc5\x00\xe1\x28\x2c\xee\xb7\x45\x36\x1c\xf8\x59\xef\xb2\x3c\x56\xae\x25\x71\xf0\xbb\x21\xab\x1d\x3b\xb4\x61\x60\x70\x55\xa6\x01\x9a\x19\xad\x97\x0d\xf8\x06\x5f\x16\xf9\xcf\x24\x3b\xea\xf0\x1f\xa3\x04\xee\x06\xa8\xbf\x93\x91\x6b\xe5\x06\x28\x0e\x8b\x7b\x44\x8c\x20\xf2\x57\x64\x89\xb2\xd8\x82\x11\xfc\xc5\xb2\xbd\x50\xd4\xd2\x25\x6c\x33\xc2\xf4\xdd\x00\x45\xc5\x7d\xf5\xaa\x52\xde\xce\xb6\x99\x47\xfc\xc1\x3c\x3e\x93\xac\x06\xd6\xfb\x1f\x81\xd5\x5b\x93\x89\xe5\x39\x90\x12\xe0\xb5\xaf\x40\xad\xf5\x4e\xe4\x62\x07\x44\xbd\xeb\x94\x56\xaf\xf7\x9c\xef\x48\x2a\xe1\xea\xdd\x55\x69\x84\xfb\x25\x87\x5c\x9b\xd5\xa9\x56\x8e\x08\x05\x66\x47\xc4\x0c\x4a\xf8\x41\x27\x7c\x2d\xd6\x30\xa8\x4e\x8c\xad\x3f\x2d\xf5\x23\x63\x49\x0c\x12\xf6\x8f\x4a\xf2\x58\x11\xb9\xb2\xc2\x9e\x0b\xc6\x40\xa1\x11\xe2\x44\xda\x46\xee\x7c\xa1\x68\xf9\xb3\x85\x32\xc1\xa0\xbe\xed\xe8\xfb\x36\xa9\xe0\x47\xbb\xa3\xef\x58\x5b\x3e\x4c\xd3\x45\x0e\xca\xf9\x33\x70\x13\x09\xe5\xf0\x64\x75\xc1\x8e\xbc\x46\x51\xbc\xef\x7e\x65\x8f\xcf\x84\x2d\x24\x59\xa1\x11\xf2\x84\x2a\x8f\x7b\x6f\xdb\xcf\x0a\xe4\x80\x1c\xab\x42\x82\xb4\xf0\xf9\x8c\x65\x33\x1d\x4a\xe8\xcc\xae\x26\x6d\x74\x4d\x63\xba\xb6\xf5\x1a\xcc\x52\x50\x38\xd8\xd7\xc6\xbe\x17\x8d\xdd\x19\xff\x50\x67\xed\x3a\xc8\x47\xac\x3d\x20\x4f\xf4\x4e\x6f\xf7\xa1\x7c\xd5\xdc\x97\x18\x0f\x75\x77\x73\x5f\x58\x5f\xd9\x82\xf5\x9d\x6d\x3d\x49\x35\x5b\x8d\xff\x0d\x00\x00\xff\xff\x1f\x48\x7c\x53\xdb\x0d\x00\x00")

func reportTestTmplBytes() ([]byte, error) {
	return bindataRead(
		_reportTestTmpl,
		"report/test.tmpl",
	)
}

func reportTestTmpl() (*asset, error) {
	bytes, err := reportTestTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "report/test.tmpl", size: 3547, mode: os.FileMode(420), modTime: time.Unix(1493141754, 0)}
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
	"report/content.tmpl": reportContentTmpl,
	"report/footer.tmpl": reportFooterTmpl,
	"report/header.tmpl": reportHeaderTmpl,
	"report/test.tmpl": reportTestTmpl,
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
	"report": &bintree{nil, map[string]*bintree{
		"content.tmpl": &bintree{reportContentTmpl, map[string]*bintree{}},
		"footer.tmpl": &bintree{reportFooterTmpl, map[string]*bintree{}},
		"header.tmpl": &bintree{reportHeaderTmpl, map[string]*bintree{}},
		"test.tmpl": &bintree{reportTestTmpl, map[string]*bintree{}},
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

