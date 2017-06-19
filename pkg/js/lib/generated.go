// Code generated by go-bindata.
// sources:
// js/runner.js
// DO NOT EDIT!

package lib

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

var _jsRunnerJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x59\xfb\x6f\x1b\x37\xf2\xff\x5d\x7f\xc5\x7c\x17\x45\x23\x25\xca\x2a\x41\xbe\x77\x28\x94\x13\x0e\xae\xed\x24\x4e\x6a\xc7\x67\xa7\xe9\x1d\x82\xc0\xa0\x76\x47\x5a\x5a\x5c\x72\x4b\x72\xa5\x08\x85\xff\xf7\xc3\x0c\xb9\x2f\x3f\xe2\xe0\x0a\x34\xb6\x49\xce\x67\xde\x0f\x72\x67\x4f\x61\xad\xcc\x52\x28\xa8\x6a\x57\x5c\x60\x66\x6c\x0e\x57\xe0\x14\x62\x05\x9b\x7a\x89\x56\xa3\x47\x07\x0e\xb3\x33\x51\x22\x58\x5c\x4b\xe7\xd1\x1e\x6f\x51\x7b\x07\x48\x3f\x78\xc3\xb9\xe2\x03\xee\x21\x33\x7a\x25\xd7\xbc\xf2\x74\x36\x9a\x3d\x05\xfc\x56\x19\xeb\x31\x87\xf7\x66\x09\x7f\x08\xe9\xdf\x5a\x53\x57\x60\x6b\x4d\x07\x46\xb3\x19\x7c\x2a\xa4\x03\xe9\xc0\x17\x48\xcb\x1a\x2d\xec\xac\xa8\x2a\xa9\xd7\xe0\x32\x2b\x2b\x9f\x8e\x32\xa3\x9d\x51\x98\x2a\xb3\x1e\x27\xbf\x19\x91\xd3\xe6\xc1\xe1\xc9\x11\x64\xc6\x62\x32\x19\x1c\xe8\x34\x49\x2d\x56\xc6\x49\x6f\xec\x3e\xd5\xa2\xc4\x49\x64\x88\x90\xe3\x4a\xd4\xca\x83\x2c\xc5\x1a\x89\xbb\xf3\x26\xdb\x40\xbd\xac\xb5\xaf\xe1\xe5\xdf\xd3\x17\xff\x0f\xcf\xa0\x14\x1b\x04\xa1\x73\x58\x4b\x9f\x8e\xb6\xc2\x82\xc8\x64\x7e\xc2\x34\x0b\x48\xe8\x8f\xe7\x81\x64\xae\x84\x47\xe7\x13\x66\xc0\xc6\x79\x27\x74\xae\xd0\x42\x8e\xa4\xc4\x12\x83\x82\x4a\x3a\x0f\x66\x15\x0c\x47\x4b\xc2\xc3\x41\x26\x73\x12\x41\xec\x84\x45\x30\xab\x74\xb4\xaa\x75\xe6\xa5\xd1\x03\xa4\xf1\x04\xfe\x1a\x01\x04\x7c\xbb\x0f\x10\x50\x44\x36\x6b\xf4\x81\x43\x25\xac\x28\xe1\x49\x2e\xbc\x78\x32\x85\x5d\x21\xb3\xa2\xb1\xee\xd2\xe4\x7b\x62\xce\x96\xc6\x3f\x6b\x74\x3e\x1d\x01\xf8\x42\xba\x94\x8c\x06\x0b\x68\x38\x13\xb3\x9b\x6e\x4f\xa9\x8b\x70\xfe\xce\x91\x9b\x4e\xe3\x9e\xaa\x42\x47\xf1\x58\xc1\x9d\x70\xe0\xad\x5c\xaf\xd1\x62\x9e\xd2\xf9\x83\x66\x5f\x3a\xa8\x84\x73\x98\x83\x37\x1d\x55\x21\xb4\xca\xd1\xa6\xec\xaa\xa1\x9e\x99\xd0\x50\x3b\x64\xc1\x08\xc9\x2c\xaf\x31\xf3\x4c\x9d\xfd\x59\x4b\x8b\xac\x1d\x69\x1f\x78\xb7\x7c\x79\x9d\xb1\x6e\x99\x77\xdc\x06\xf1\x94\xac\xd2\x5a\x99\x02\xa6\xb1\x5c\x90\x81\x57\xc6\x98\xae\x53\x78\x42\xf6\x7a\x32\x69\x2c\xc4\x3b\x8b\x2e\x1d\x02\x42\xb4\xf1\x10\x84\x45\x5b\x59\x53\xde\x76\xc3\x6c\x06\x6f\x8c\x05\x01\x6f\xa5\x7f\x57\x2f\xa1\x30\x66\x33\x05\xe9\x9f\x34\x7e\xdd\x2b\x23\x72\xb0\x98\xa1\xdc\x62\x1e\x30\xc2\xd9\x46\x0c\xdb\x3a\xc9\xe2\x9f\x01\x32\x24\x24\xfd\xf0\x42\xea\x00\x15\xd6\x6a\x2b\xd8\x08\x2b\x63\x99\xfa\x96\x2c\x4d\x5a\x0a\xc8\x25\x5b\x4b\xd8\x7d\x1b\x2d\x11\x75\xd1\x45\xc0\xb9\x35\xde\xf8\x7d\x85\x8c\xf7\xde\x2c\x7b\x51\xfc\xde\x2c\xc7\x9a\xed\xeb\x85\xdb\xb8\x60\x61\xca\xa6\x72\x0f\x8b\xe0\xc9\xc0\x93\xeb\xc6\x4e\x2a\x05\x4b\xcc\x4c\x19\x7c\x59\x59\x5c\xc9\x6f\x51\x4c\x84\xca\xe4\xb3\xc0\xbe\x14\x15\xbb\xc4\xa5\xb7\xdc\x40\x3f\x5e\x47\x2d\x88\x61\xe3\x80\x26\xf9\x58\x0a\x0a\x19\x5b\x6b\x8e\xb1\x3d\x50\xe6\xe1\x37\xcc\x6a\x2a\x55\x52\x83\x23\x5b\xe8\x0c\x41\x6a\x27\x73\xca\xca\x80\x27\xc0\x15\xa8\x14\x8c\x67\x4b\xa9\x67\xae\x98\xb4\xbc\x03\xe8\x22\x80\xbf\x8e\x0a\x1d\x40\x66\x94\xc2\x60\x06\xb3\x62\xc9\x66\x5b\xa1\x6a\xf2\xa7\xb4\x8e\x4b\x81\xde\x4a\x6b\x74\x49\xc1\xb1\x15\x56\x8a\xa5\xea\xa9\x84\x7a\xcb\x66\x6e\x10\x29\x23\x42\xd1\xa2\xba\x24\x34\x98\x8a\x9d\xa3\xc0\x8b\x75\x4b\x25\x63\x89\x6a\xcb\x55\x27\x10\x59\xcd\xac\xe0\xf8\xec\xf3\xd5\xe7\x83\x8b\x60\x40\xc6\xba\xc4\xcc\xa2\x8f\x16\x0d\xee\x67\x57\x54\xa6\xaa\xa9\xba\x85\x18\xee\x84\x0d\x80\x8d\xc4\xb0\x93\xbe\xe0\x23\x41\xbd\x95\xa9\x35\x5b\x92\x96\x1c\x43\xf7\x03\x8b\x91\xcd\x16\xad\x25\xeb\x0a\x28\x85\xcf\x0a\x2a\xea\xa4\x2f\x85\x46\x9b\x21\xb4\x50\x8a\xaa\x55\x2d\x60\xb9\x10\x7b\x01\xb1\x32\xf9\x59\x4c\x56\x87\x1e\x96\x7b\x72\xec\x78\x32\x65\xb5\x06\x91\xcf\x01\x12\x2b\x60\x65\x72\xc8\x2c\x0a\x4f\x35\xa9\xa9\x74\x01\x29\xe2\x52\x93\x72\xa8\x73\x17\xf2\xe3\xda\x2c\x29\x68\x3e\xb4\x5d\xb1\x25\xa3\x83\xc3\xc2\x38\x02\x88\x7b\x4b\x91\x6d\xd6\x96\xac\xd1\x6b\x4c\x93\x6e\x7f\x27\xa4\x1f\x4f\x46\xbc\x60\xd1\xd7\x56\x87\x8c\x00\x08\x3e\xbf\x05\xf2\x20\x9f\xc6\x08\x0b\x56\x9e\x96\xa6\xbd\x9e\x3e\x79\xdd\x02\xce\x66\x40\x3c\x7f\xd7\x5e\xaa\x23\xa3\xd9\x6c\x05\xda\x90\xb8\xc4\x68\x27\x6c\xee\x20\x33\x65\x25\xbc\x5c\x4a\x25\xfd\x7e\x0a\xcb\xda\x43\x6e\xd0\x81\x36\x9e\x1c\x15\x9d\x79\x74\x7c\x7e\x71\x7c\x78\xf0\xe9\xf8\x68\x0e\x7f\x84\xc4\x05\x8b\xa5\xa1\x02\x95\xd7\x96\xdb\xb4\xaa\x0a\x31\xea\x69\xdb\x71\xbe\xa5\xcb\x7d\x36\xe8\x49\xcc\xff\x38\xa8\x89\x1c\x04\xfb\xaf\xa0\x85\xe4\xb2\xce\x32\xc4\x1c\xf3\x84\x0f\x77\x14\xe3\x09\xf7\x8b\x25\x42\x26\x94\xc2\x1c\x8c\x86\xce\x94\xf4\x37\x77\x0f\x97\xde\xa5\xa3\xca\x57\x7b\x53\x0a\x2f\x89\x76\xdf\x20\x2c\xf7\xad\xcb\xfb\x54\x27\x2b\x96\x27\x08\xef\x20\x79\x23\xa4\xc2\x3c\x99\x86\xc8\xf1\x85\x35\xbb\xd0\x15\xbf\x65\xc8\x09\x9b\x0e\xe9\x6a\xed\xd8\xfe\xa5\xe1\x0e\x26\x34\xbc\xfc\x1b\x94\x52\xd7\x34\x7c\x8d\x5f\xbd\x78\x01\x4f\xe1\xd5\x73\x87\x99\xe1\xac\xf2\x68\xb7\x42\xb9\xc9\xb4\x85\x06\x2f\x4b\x34\xb5\x1f\x72\x68\x6d\x7e\x5f\xd8\x10\xbf\x31\xe5\x9a\x84\x05\xbc\x78\x0d\x12\xfe\x01\xaf\x5e\xd0\x2f\xcf\x9e\x35\x67\x00\x06\xa3\x57\x56\x60\xb6\xe1\xc1\xcc\x0b\x5f\x73\xf9\x4a\x68\x4e\xda\x37\xe1\x37\x89\x54\x04\xbb\x81\x45\x6f\x84\x4c\xa9\x40\x9c\x5d\x8e\x93\x38\x7a\x25\xfd\xa3\xe5\x9e\xac\xb0\x80\x4d\x4a\xf3\xdc\xe7\x97\x04\x97\xae\xd1\x8f\xfb\xc8\xf7\x08\xf4\xfe\xf2\xe3\x59\xea\x3c\xc5\x99\x5c\xed\xc7\x0c\x33\x99\xdc\x27\xf9\xb9\xc9\x6f\x89\x0a\xcf\x20\x21\x37\x53\xb5\xf7\x54\xdf\xc2\x36\x71\x0e\xda\xa5\x55\x21\x5c\xc7\x57\xae\x60\x7c\x77\x1b\x16\x8b\xd6\xdb\x9d\xd1\x20\xf8\x05\x1e\x62\xbb\x62\x82\xd8\x82\xe8\x07\x65\x9b\x42\x72\x4f\xf2\x3a\x62\xdc\x3c\xce\xb7\x0b\xfc\x3e\xeb\x26\x83\x6c\x8d\xb7\xa0\x66\x33\xb8\xe4\xc9\x7e\xc5\x83\x46\x8e\x2b\xa9\x31\x07\x51\x9a\x5a\x87\xce\x28\x4b\x4c\xe3\x61\xbe\x03\x8c\x5f\x05\x6b\xde\x8c\x7a\x4a\xd1\xa9\x1c\x28\xd8\x28\xb4\x28\x1c\x08\xaf\xba\x5f\xd5\xa0\x63\x12\xca\x4f\x98\x15\xba\x6b\x40\xc8\x68\x0e\xfc\x5a\x79\x59\x29\xa4\x2a\xcb\xbd\x79\x25\xb5\x74\x45\x0a\x27\x3e\x34\x8b\xc0\x9b\x52\xc8\x5a\x63\x09\x46\x38\x70\xc6\x68\xfa\x29\xb8\x38\xd3\xb4\x6f\xbd\x6b\x0f\xf5\x46\x90\x96\x65\x8c\x7e\x4e\x0c\x66\xb5\x80\x2f\x5f\x63\x91\x11\x79\x4e\xff\x13\x9c\xc6\x5d\x53\xef\xa9\x57\x90\x9c\x54\x34\xaa\x86\x94\x8e\xf6\x52\xea\xda\x2c\x07\xc5\x98\x90\x79\xa4\xe6\x9d\x5e\x1d\x23\x77\x73\xb2\x23\xcf\xf0\xc4\x22\xf6\x48\x86\x0f\x2d\xcb\x17\xa8\x7b\xa6\xf1\x05\x96\xbd\x18\xc1\x1f\x6a\x3b\x2c\xc1\xca\xd8\x63\x91\x15\xe3\xd6\x0c\xe3\xeb\x2e\x50\xae\xfb\xad\x29\x7a\xf9\x6e\x5f\xfa\x4e\x09\x66\xf3\xd1\xe0\xd4\x0a\x06\x67\x86\xe7\x04\x9a\xbb\xa9\xee\x35\x0d\x23\x44\x79\x11\x9c\x1b\x26\xae\x80\x59\xd6\xce\x53\x79\x76\x5e\xf0\x25\x11\xbf\x79\xb4\x9a\xea\x6d\x0a\xe3\x4b\xc4\xce\x6f\xac\x6c\x9c\x2c\x3b\xb9\x27\x8f\xd5\xb8\x1f\xb6\x45\xab\x6e\xb4\x42\x33\xd7\xb6\x0e\x0b\x21\x16\x12\xa7\x77\xc3\xa4\xf8\x6c\xd5\x64\x29\x9a\x41\x35\x86\x4f\xe7\xb5\xc1\x50\xad\xcc\xee\xb9\xc2\x2d\x2a\xa8\xac\x2c\xa5\x97\x5b\xec\x85\x2b\xb5\xf0\x6b\xb3\x1c\x74\xf0\xe6\x5e\xf2\xd3\xfb\x8f\xbf\x3e\xff\xe9\xd3\xc9\xe9\xf1\xf3\x9f\xde\x9e\x7c\xba\x7c\x77\x10\xa7\xe9\xcd\x2f\x2e\x4e\x00\xd7\x66\x19\x66\xe1\x67\x90\x3c\xa7\x9c\x3c\x12\x1e\x53\x6d\x76\xe3\x49\xbb\xd4\xd3\xa1\x40\x91\x5f\x65\xa6\x2c\xa5\x4f\x65\x9e\xba\x7a\x19\x4a\xea\xf8\xc5\x14\x7e\xe1\xc9\x81\xe0\xb3\x32\xa2\x47\x3e\x71\x39\x5c\xd9\xef\xdd\xca\x4a\x1a\xc6\x71\x77\x78\x3a\x0e\xc4\x93\x01\x4d\xd8\xbc\xe0\xdf\xcf\x4d\x3e\xee\x90\xb8\xe6\x86\x3f\xd3\x12\xbd\xa0\x3b\x53\xaa\xc4\x12\x15\x3b\x53\x0f\xb5\x7c\xf8\xec\x12\x95\xd1\x6b\xe7\x0d\x2c\xe0\xfe\x57\x01\xb3\x23\xba\x81\xad\x1e\x7e\x3e\x78\x98\x51\x30\xde\x90\xcb\xd0\xaa\xcd\xfc\x9d\xe7\xcd\x88\xcb\x53\x07\x59\x03\xf5\xf6\xb3\xb0\x83\x4a\x74\x95\x22\xc5\x2a\x69\x88\x7a\x3b\xed\xe2\x7a\x2b\xd4\x14\x36\xb8\x9f\x82\x0a\x01\xf1\xe8\xc9\x26\xc8\x23\x97\x50\x95\xfe\x22\x75\xe6\x01\x88\xe7\xf5\x39\xfd\xb8\x09\x53\xe2\xe4\x75\x27\x6b\x33\x6d\x0b\x37\x10\xbb\xc7\x35\x9e\x78\x80\xf3\xa8\x69\x3b\x74\x9b\x6b\xc0\x76\x74\xd7\xaa\x94\xcc\xa4\x57\x7b\x58\x2a\x93\x6d\x42\xd3\xa1\x56\x77\x95\x36\x13\xfb\xf8\x4b\x12\x28\x92\xaf\x2c\xe5\xa4\xcb\xd7\xd0\xe2\x62\x67\x6a\x78\xbc\x45\xdf\xb2\x68\xaf\x0f\x6b\xb9\x45\xdd\x3c\x5f\xdd\x63\x88\x08\xd8\xd9\xa3\x9d\x4a\x54\x8d\x6f\xac\x29\xe7\xbd\xce\x1a\xd0\x3f\xe0\xfe\x02\x57\x73\x88\x46\x8c\xd8\xac\x72\x30\xe3\xa0\xeb\x76\x46\x1d\x05\x05\xc3\x4b\xd9\xfd\x6e\x19\x48\x93\x1c\x1c\x9e\x1c\x5d\x5d\x1c\x9f\x7f\xbc\xfa\x70\xfc\x9f\x64\x20\xd9\x3c\x3e\xb8\xf5\x0b\x56\xe7\x35\x6f\xaa\x58\x5c\x5a\xa7\x51\xd1\x75\x18\x0a\x6e\x77\xef\xd2\x7b\x10\xde\x63\x59\xf1\x0b\x0a\x5d\x9e\xda\x1b\x9c\x0b\x70\xb4\x6c\x4a\xf4\xe1\x72\xa6\x1c\x77\x9f\xa1\xd0\x8d\xb8\x87\xbf\x7d\x3c\x3b\xbe\xfa\xfd\xe2\xb7\xa4\x0d\xaa\xfb\x73\x29\x53\x46\xe3\x55\x6d\x55\x90\xfc\x7e\xb4\xcb\xcb\x77\x3f\x82\xe5\x5c\xf1\x18\xd2\xdb\x93\x4f\x3f\x82\xb4\x96\xfe\x31\xa4\x77\xc7\x07\x47\x57\x87\x1f\x4f\x4f\x4f\x3e\x5d\x9d\x1c\xdd\x0b\x38\xcc\xfa\x00\x16\x8b\x86\xab\x30\x6b\x82\x1b\xad\xfb\xf2\xe2\x6b\xbc\xdd\x47\x66\x3d\x07\xc6\xf7\x15\xba\xab\x6f\x8d\xaa\xfb\xa5\x87\x51\xc2\x22\x97\x0c\x8e\x80\x46\xc2\x50\x66\xa7\x91\xfe\x54\x54\x6d\x94\xc6\xea\x7d\x73\xd3\x64\xcb\xb4\x25\x4a\x64\x6e\x9d\x48\xa6\x31\xbc\xe7\xf0\x57\xfc\xed\xac\x1f\xdf\x81\xf4\xeb\xeb\xef\xea\x13\xe4\x3a\xa5\x71\xf2\x61\xe1\x78\xda\x3c\x17\xbe\x98\x43\x32\x2b\x8c\xd9\xcc\xa8\x94\x26\x37\xd3\xef\x88\x76\x97\xc6\xb9\x22\x99\x82\x45\x91\x7f\xd4\x6a\x3f\xe7\x71\x37\x4a\x18\xec\xf8\xb1\x89\x72\xdf\x3e\x99\x18\xad\xf6\x94\x83\xb4\x52\x3b\xb4\x14\xf1\x0e\x8c\xe6\xa0\xa6\xdc\xa4\x82\xc6\x47\x9b\xf4\x7c\x58\xd5\xe6\x89\xa5\x25\xe9\xe5\xe0\x7b\x13\x87\xba\xf6\xa5\x29\xbc\x24\xf3\x58\x49\xc1\x21\x74\x3e\x8f\xd5\x9f\xfa\x63\x99\x47\x20\x3e\x9f\x5e\x1b\xa9\xc7\x09\xfc\xfc\x33\x24\xdc\x09\xb3\x32\x25\x0b\x7d\x49\x4a\x21\x75\xea\x8a\xe4\x6b\xec\xab\x25\xb7\x95\x1f\xba\x6b\x8d\x6e\xdd\x87\x0e\x2d\x0a\x1e\xdd\xbb\xb7\x34\x6a\x7e\x59\xd9\xf5\x36\x1d\x3b\xf6\x77\x2e\x5c\x59\xc9\xb7\xad\x4d\x4a\x83\x9b\x76\xd2\xe8\xe6\x6d\xb0\x14\x55\x1a\x5e\x57\xe8\xd0\x83\xdc\x9b\x4b\xc3\xc3\x76\xfe\x01\x29\x02\x71\x94\xa4\x77\x81\x8c\xfc\xe3\xf6\x6d\x19\x68\x59\xea\x75\x9a\xa6\xc1\x3c\xf1\xda\xd4\x4d\x21\x7c\x5d\x69\x47\xb2\xc1\xa8\x52\x99\x9c\xe5\xe2\x28\x89\x84\x21\x60\x92\x8d\xd4\x79\x32\xe7\x9b\x5f\x2c\xd9\x89\xa8\xe4\x67\xb4\x64\x1d\xda\xd8\xbe\x6c\xd6\x1b\x4b\x27\x5d\x8f\x49\x08\x36\x99\x43\x64\x30\xed\x2f\xbb\x4a\x64\xb4\xd7\x7a\xb5\xdd\x0d\x53\x48\xd2\x6f\x55\x49\x81\x56\x7a\xb1\x66\x82\x7f\xd5\x66\xb3\x11\xed\x79\x62\x2d\xb4\x58\x63\xfe\xeb\x9e\xb6\x45\x26\xf9\xe9\xa4\xd7\xb4\xa2\x84\xe4\x8f\xbe\x74\x16\x79\x52\x3f\x37\x4a\x66\x4c\x7a\x46\xf7\x98\x4e\x90\xce\x75\xc9\x3c\x66\x3f\x57\x80\xf6\xb7\x4e\x43\xe6\x4a\x37\xc3\x69\x7f\x93\x33\x29\x99\x77\x0f\x97\x83\xdd\x98\x3a\x03\x6c\xde\xe0\x92\x40\x99\xd1\x5b\xfe\xda\x27\x9d\xcd\xe0\xcd\xc9\xbf\x4f\x8f\xe7\x70\x58\x08\xbd\x46\xca\xc9\xe4\x64\x75\x66\xfc\xb9\x45\x87\xda\x27\x77\xa4\x38\xaf\x95\xea\x14\x3d\x50\x3b\xb1\x77\xdd\xa9\xa6\xc9\x7f\x6d\xaf\xc7\xe1\x7a\x3b\x08\x98\xc3\xd3\xf1\x63\x61\x72\xd8\x54\xea\xff\x21\x58\x3a\x63\xf6\x63\xe5\xb1\x68\xb9\x3f\x5e\x1e\x8f\x99\x18\x13\x4d\x68\xdc\x95\xa5\xa9\x4e\x73\x48\x30\x2b\x0c\x14\xa8\x94\xa1\x2a\xc6\x7f\xad\x8d\xc9\x97\x7b\x4c\x3a\x90\xf6\x3d\x60\xf1\xf0\x7f\xb4\x7d\x2a\xa4\x06\x65\x44\x8e\xf6\xb1\xd3\x77\x3f\x14\x62\xce\xdf\x09\x29\xc5\xa9\xc0\xff\xdf\xf0\x03\x66\xf3\x18\xc0\xcf\x1a\x67\x66\xf8\xad\xc9\x35\x0f\x24\x09\x3b\x96\x47\xf5\x50\x78\xef\x7e\x38\xea\x3f\xed\x8e\xb0\xf7\x7d\x24\x34\x84\xd3\xf3\x5f\xe7\x70\x02\xb9\xd1\x4f\xf8\x6a\xac\x37\xd4\x1e\x2c\xbf\xbd\xd2\x10\x66\x51\x38\xc3\x0f\x41\x95\x70\xf1\xb9\x59\xea\x69\xfc\x70\x61\xf1\x9f\x01\x26\x09\x43\x5f\xd2\x4c\x7f\x64\xc5\xa4\xfb\xe4\x9a\xcc\x7b\xdf\x5f\x49\x66\x16\xb9\xff\x31\xb2\x27\x7d\xfb\x55\x71\x68\xb4\x5b\xa5\xb5\x4f\x3d\x99\x8c\x86\xd6\x1b\xee\xde\x32\x7e\xfc\xc2\xc9\x7e\xcb\x53\x78\x23\xf9\x31\x98\x6a\x7d\x6b\xb6\xf0\x39\x76\xa5\x61\x01\xca\x98\x4d\x5d\xdd\x31\xeb\x90\x01\x1f\x1e\xbc\xa2\xbb\x9d\xf4\x59\x01\x1d\x49\x70\x68\x26\x1c\x42\x42\x07\x93\x79\x98\xb2\x7b\x38\x61\xa6\x8b\x2f\x95\x4b\x8b\x62\x33\x82\xe6\x7b\xf0\xbc\xff\xcc\xf5\xbb\xde\x68\xb3\x6b\xbe\x49\x0e\x24\x27\xe3\xfe\x37\x00\x00\xff\xff\x46\x30\xc4\x64\x3b\x1f\x00\x00")

func jsRunnerJsBytes() ([]byte, error) {
	return bindataRead(
		_jsRunnerJs,
		"js/runner.js",
	)
}

func jsRunnerJs() (*asset, error) {
	bytes, err := jsRunnerJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/runner.js", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"js/runner.js": jsRunnerJs,
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
	"js": &bintree{nil, map[string]*bintree{
		"runner.js": &bintree{jsRunnerJs, map[string]*bintree{}},
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
