// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// assets/inject.php
// templates/overview.html
package main

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsInjectPhp = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x55\xdd\x6e\xda\x4c\x10\xbd\xe7\x29\xe6\x93\xd0\x67\x53\x11\x30\xed\x45\x25\xa8\x53\xb5\x69\x52\x55\x8a\x48\x04\xe4\x2a\x42\x68\xbb\x1e\xe2\x55\xed\x5d\x77\x77\xdd\x26\x8a\x78\xf7\x6a\xd6\x76\x88\xff\x08\x37\xce\x8e\xcf\x9c\x33\xeb\x39\x33\xf9\xf4\x39\x8b\xb3\xc1\x40\xec\xc1\xff\x8f\x27\xcc\x98\x1d\x3e\x0a\x63\x8d\xef\xdd\xc6\xd9\x57\x94\x3c\xf6\x46\x23\x78\x1e\x00\x00\xb8\xf7\x50\xc5\xe1\x79\xe0\xa2\xf4\x33\x96\x59\xc1\x61\x18\x21\xd7\xc8\x0c\x6e\x44\x8a\x10\x82\xcc\x93\x64\xd1\x02\x19\xb4\x4b\xdd\xfb\xd6\x8a\x14\xb5\x81\x10\xee\xb7\x8b\xa3\x40\x96\xff\x4c\x04\xaf\x40\xfb\x5c\x72\x2b\x94\xa4\xb3\xb6\x24\xa6\xfd\x22\xb1\x2a\xb5\xce\x3a\x9f\x97\xb4\xf7\x9e\xf5\x60\x52\x8a\x6c\x21\x84\x54\x70\xad\xe8\xe4\x5b\x9d\xe3\xe8\x58\xce\xe1\x4d\x6d\xca\xba\x50\x11\x96\xca\x63\x18\xee\x45\x82\x92\xa5\x38\x86\x61\x22\x24\x2e\x29\xc6\x55\x84\xcd\xa2\x86\x28\xa3\x2e\xed\x3a\xc8\xdd\x0d\xc2\xd3\x57\x58\xd4\x73\x28\xb8\xce\x50\x52\x9e\x53\x39\x2b\x79\xea\xb8\x5c\x1a\xb4\xfe\x49\xe2\x56\x39\x29\xf5\x44\xab\x5c\x46\xfe\x2b\x99\x77\x30\x0b\x82\xa0\x7c\x34\x73\x9c\xa9\x5e\x54\x5c\xdb\x9b\x9f\xa2\xd6\xa3\xca\x18\xfb\x44\x29\xed\x37\x3e\x4f\xa5\x01\x13\xf0\xbc\xfa\x75\x0e\xa7\x64\x5f\x5b\xf2\xa4\x7a\xc3\xbb\x45\x11\x65\x17\xa6\x4e\x3b\xa8\x8a\x08\x9a\xfa\x1d\x8d\x3b\x0b\xbb\xa9\x17\x6d\x2b\xf4\x43\x6b\xd8\xe9\x14\x90\xc7\xca\xb5\x62\x02\x5e\x6a\xbc\x36\xe0\x3b\x5a\xe0\x77\xab\x6b\xd0\x68\x54\xae\x39\xd6\xd5\x78\xae\x13\x08\x81\x1e\x3b\x21\x85\xf5\x47\x8b\x26\xc3\x1a\x2d\x18\x95\x22\xa8\x8c\x6c\x6e\xe0\x0c\xfe\x22\x30\x8d\x90\x31\x63\x84\x7c\x00\x21\x81\x41\x6e\x50\xb3\x07\xb2\x80\x55\x0a\x62\xd4\x75\x25\xa7\x60\xd0\xaa\xcc\xee\x98\xd6\xec\xc9\x77\xda\x63\xb8\x6f\x35\xe0\xe2\x6e\x75\x7d\x73\xbb\xd9\xad\x2e\x37\x77\xab\xe5\x66\xf5\x65\xb9\xbe\xba\x5c\x41\x78\x0e\xb3\x71\x2f\x98\xae\x18\x9e\x83\x17\x5b\x9b\xcd\xa7\xd3\xd9\xfb\x8f\x93\x60\x12\x4c\x66\xf3\x0f\x41\x30\x9b\x46\xcc\x32\xaf\x3f\xf9\xf6\x66\xbd\x79\x83\x9f\x20\x57\x3f\x2e\xaf\xbf\xad\x09\xd8\x2e\x9a\x7e\x5e\x35\xef\x1e\x61\x8e\xd3\xdf\x0d\x76\xe6\x76\xc8\xba\xdf\x7b\xe0\xbf\xf0\xa9\xa0\x2d\x96\x4b\x37\xa8\xd8\x32\x05\xae\xdc\x38\xdd\x40\xda\x42\x05\x8c\xfe\xea\x01\xa5\xa6\x80\xa4\xa6\xef\x0a\x64\x6c\x87\xa9\x0d\xc7\xeb\x1d\x30\xea\x49\x45\x19\xb9\xc4\x72\x81\x90\xeb\xdf\x48\xdb\xd6\x43\xdb\x2e\xa7\xca\x08\x6c\x8c\xa0\xf1\x77\x8e\xc6\xc2\xff\x60\xd8\x1f\x3a\x9a\x4c\x49\x83\x60\x15\x0c\xe9\x50\x9f\x01\x8a\x54\x33\x80\x8f\xc8\x0b\x67\xb6\xe9\x2f\x12\x65\x8e\xdc\x56\x01\x4f\x90\x69\xc8\xb3\x62\x3c\xaa\x01\x33\x6d\xdf\x73\xca\x7c\xa1\x6d\xfe\x4b\x39\x0c\x0e\x83\x7f\x01\x00\x00\xff\xff\xe3\x35\x4d\xf3\x76\x07\x00\x00")

func assetsInjectPhpBytes() ([]byte, error) {
	return bindataRead(
		_assetsInjectPhp,
		"assets/inject.php",
	)
}

func assetsInjectPhp() (*asset, error) {
	bytes, err := assetsInjectPhpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/inject.php", size: 1910, mode: os.FileMode(511), modTime: time.Unix(1563249143, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOverviewHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x6d\x6f\xe3\xb8\xf1\x7f\x9f\x4f\x31\x7f\x2d\xfe\xb5\x73\xb5\x25\x39\xc9\x06\x59\xc5\x76\x71\x68\xef\x80\xde\x76\x6f\x81\x66\xd1\x37\xbe\xbc\xa0\xa4\x91\xc5\x35\x4d\x0a\x24\x9d\xc4\xe7\xf3\x77\x2f\x48\xc9\xb6\x9e\xed\x2d\x0e\x7d\xd0\x62\x23\x89\x9c\x19\xce\xc3\x6f\x7e\xa4\x35\x65\x94\xaf\x20\x95\x98\xcc\x9c\x54\xeb\x4c\x05\x9e\x97\x08\xae\x95\xbb\x14\x62\xc9\x90\x64\x54\xb9\x91\x58\x7b\x91\x52\x7f\x4a\xc8\x9a\xb2\xed\xec\x73\x86\xfc\x8f\x4f\x84\xab\x3f\xc4\x54\x65\x8c\x6c\x67\xea\x95\x64\x0e\x48\x64\x33\x47\xe9\x2d\x43\x95\x22\x6a\x67\x7e\x75\x35\xb5\xaf\xf3\x2b\x00\x80\x50\xc4\x5b\xd8\xd9\x47\x73\x99\x55\xc6\xb9\xc5\x00\x06\xc6\x26\x18\x9b\x83\x47\x2b\xb1\xbf\xb2\x37\x57\xd1\x18\x43\x22\x4f\x6a\x99\x50\x54\x53\xc1\x03\x48\xe8\x1b\xc6\x8f\xc7\x89\x5f\xc7\x94\xc7\xf8\x16\xdc\x9e\x86\xb4\xc8\x02\xff\xf4\xca\x30\xd1\xe5\xf7\x50\x68\x2d\xd6\xe5\x91\x57\x1a\xeb\x34\xb8\xf5\xfd\xec\xad\x24\x46\xa2\xd5\x52\x8a\x0d\x8f\xc7\x91\x60\x42\x06\xef\x10\xb1\x6c\xe5\x6d\xac\x52\x12\x8b\xd7\x00\xfc\xec\xcd\xfe\x9f\xbc\xcf\xde\x40\x2e\x43\x32\xf4\x47\xf6\x9f\x3b\xb9\x6e\x8d\xcb\x4d\x28\x43\xb5\x6b\x04\x71\x53\xf3\xba\xe4\x52\xcd\x02\xb8\x29\x92\x18\x4b\x19\x6a\xf1\x37\xb9\xf7\xef\x4b\x61\x66\x24\x8e\x29\x5f\x06\x37\x07\x5f\x4f\x53\x85\x46\x92\x24\x8f\xd5\x4a\x29\xfa\x2b\x06\x93\xfb\xb2\xac\x1d\x7f\x45\xba\x4c\x75\x00\xa1\x60\xa5\x62\x78\xdf\x41\x28\x64\x8c\x72\x2c\xed\xf4\x24\x7b\x03\x25\x18\x8d\xe1\xdd\x03\x26\x93\x24\x79\x84\xef\xbc\xf6\x70\xa8\xc6\xf5\xae\xe1\xaa\xcd\x68\xd5\xd5\x62\x81\xa2\x8a\xa5\x15\xa2\x28\x2a\x05\xb4\x91\x4a\xc8\x00\x32\x41\xb9\x46\xd9\x53\x85\xfa\xd2\xaf\x42\xc6\xe3\x50\x22\x59\x05\x60\x6f\x63\xc2\x58\x47\x0d\x8c\xa6\x4b\x22\x4d\x5f\xb0\xaf\x10\x89\xbd\x2e\x08\xe1\xc1\x5e\x95\x7c\x96\x2a\xd3\x9b\x3b\x97\x2e\xb9\x90\x25\x37\x34\xbe\xe9\x71\x8c\x91\x90\x24\x6f\x1c\x46\x39\x8e\x75\x2a\xc5\x66\x99\x36\x2a\x4f\x08\xa9\xc6\x18\x09\xae\x91\xeb\xb6\x0e\x24\xa1\x12\x6c\xa3\xb1\xd9\x84\x93\xfe\x26\xbc\xaf\x76\x58\x8e\x11\xbf\xba\xae\x0e\xd9\xae\x96\xa8\x00\xb8\xe0\x58\x6f\xd7\x89\xef\xff\xff\x69\x6c\x4d\xe4\x92\xf2\x63\x4a\x9b\x7d\xa3\x43\x06\xda\xf4\x0c\xe8\xb4\x89\xb3\xbc\x25\xce\xf4\xff\xcd\x07\xf4\x49\xd8\x28\x63\x03\xea\x37\x61\x98\x7c\xb8\xeb\x6e\xae\x8a\x4b\x96\x1e\x75\xdc\x02\xfd\x66\x97\x76\x2d\x18\xc7\x71\x63\xb5\x23\x8e\xda\x56\x73\xa3\x1a\x4e\x08\xa3\x4b\x1e\x40\x84\xcd\x66\x29\xe9\xc9\x80\xeb\x74\x1c\xa5\x94\xc5\x43\x11\xc7\xd7\x3d\x98\x87\x13\x5b\x1e\x0d\x09\x4d\xd8\x38\x14\x6f\xe7\xd7\x86\x1e\xfa\xb1\x3a\x5a\x12\xae\x12\x21\xd7\x01\x6c\xb2\x0c\x65\x44\x14\x5e\xc4\x50\x67\xf7\x91\xf7\x75\xcc\xde\xbc\x6f\x62\xf6\xa6\x56\x16\x0b\x3a\xa8\x8e\x96\x29\xec\x5b\xb7\x8d\x9b\xeb\x1e\x1c\xc2\xbb\x36\x32\x91\x24\xa6\x1b\x15\x3c\x9c\x70\x3f\xf5\x8a\x5d\xf8\x6a\x1a\xd3\x17\xa0\xf1\xcc\x21\x59\xe6\xe4\xbb\xb2\x1d\x8a\x18\x51\x6a\xe6\x14\x3c\x52\xcc\xd4\x67\xf3\x7d\xc6\x99\x7f\x41\xa5\xd5\xd4\x8b\xe9\x4b\xbb\x9c\x21\x21\x07\x5e\xc6\x89\x90\x33\x67\x85\x5b\xa0\x1c\x14\xea\x8f\xb8\x55\x0e\x04\x85\xd0\x0e\x72\xb6\x0c\x0c\x41\x4b\xe4\xfa\xc9\x4a\xc0\x6c\x06\x46\x65\x6f\x0c\x08\x1e\x44\x8c\x46\xab\x99\xa3\xb0\x98\x1f\xae\x70\x7b\xed\xcc\x17\x0b\x2b\xf5\xfc\x5c\x72\xa3\x78\xec\x0a\x0a\x2c\xc7\x9f\x09\xed\x47\x23\x73\x41\x68\x27\xcf\x8c\x55\x4e\xd6\xa8\xfe\x6a\x59\x37\x86\x19\x2c\x9e\x1f\xdb\xe2\xac\x4b\xba\x0c\xf9\x52\xa7\x26\x62\x1f\xf6\xce\xfc\x7b\xc6\x72\x1f\x2f\x4d\xad\x11\x36\xb9\x3d\x1a\x2e\xaf\x9a\x6f\x02\x2d\xab\x5a\x70\x7f\x4e\x86\x66\xe2\x1a\xe6\x30\x9e\xd4\x93\x9d\xab\x9a\x54\xe4\x42\x36\xdd\x76\xb1\x8b\xf2\x5d\xec\x17\xe5\x4c\x6b\x12\x32\x3c\xcc\xeb\x90\x39\x10\x21\x63\x2a\x23\x11\xe5\xcb\x99\xe3\x97\x64\x73\x79\x53\x8f\xea\x58\x3e\x2e\x9b\x83\x85\xc2\xfc\xcf\x22\xc6\xa9\xa7\xd3\x6e\x89\xb5\xea\x9f\xff\x42\xd7\xa8\x20\x22\x8c\x61\xdc\x2f\xf9\xfd\x0b\x4a\xb2\x44\xe8\xb2\x38\xf5\xea\x9e\x1a\xb9\x46\x4c\x53\xcb\xa6\xf5\x31\x5c\x67\x8c\x68\x84\x97\x31\x4d\x66\xce\xa9\x3b\x9c\xd6\x84\x1c\xc0\x10\x99\x4d\xdd\xa0\x41\xa2\xda\x30\xad\x5a\xa4\x73\x8d\xd8\x54\xd3\x4a\xbb\x26\x65\xb6\xa6\xba\x25\xd9\x0d\xe9\x4f\xea\x2f\xf9\x69\x1f\x9e\x9f\x6d\xe0\x3d\x4a\xfe\x85\x36\x8b\x44\x5e\x62\xba\x35\xa7\x45\xae\x1a\xe3\xd5\xbc\x4e\x3d\x0b\xc0\x2a\x6c\x4f\xb8\x6d\x24\xfa\x88\xd4\xc3\x3e\x55\x4a\xe6\x17\x33\x16\xc0\x62\x01\x76\xd6\x3a\x5c\xb1\x7b\xb4\x3f\x55\x91\xa4\x99\x06\x25\xa3\xd3\x4f\xab\x28\xe6\xee\x57\x15\x23\xa3\x2f\xd2\xe5\xa8\x3d\x9e\xad\xbd\x97\x0d\x7a\x31\x55\xda\x3c\xb8\x5f\x95\x33\x9f\x7a\xb9\xee\xfc\x68\x25\x77\xe0\x85\xc8\xa2\xbe\x4f\xa8\x15\xcc\xe0\xa7\xa7\xcf\x3f\xbb\x19\x91\x0a\x87\xbb\x1d\xb8\x7f\x3f\xce\xfd\xa4\x04\x87\xfd\xbe\xd8\x3d\x3c\xaf\xa9\xba\x73\x26\xef\xef\x6f\x27\x1f\xee\xfc\xdb\xbb\x87\x87\x5b\x27\xd8\x39\x06\x0f\x7f\xa3\x1c\x95\x79\x99\x98\x3f\x3f\x16\xec\xe1\x04\x8e\xb7\xe6\xda\x8b\xbc\xd7\xd7\x57\x2f\x4b\xb3\x10\x79\x94\x7a\x31\x86\x9b\xa5\xa7\x51\x69\xf7\x30\x66\x1e\x9c\x91\x35\xe5\x04\x8e\x62\x88\xd9\xf0\xe6\xfa\xd1\x19\x39\xc6\xf2\xcf\xd2\x09\x9c\x3b\x67\xe4\xfc\xc0\xb5\xa4\x66\x21\xbe\x61\x6c\xe4\x7c\x52\x4e\x70\x63\x6e\x05\x0e\xec\x5b\x1d\x1c\x4e\xe0\xef\x47\xce\xcd\xef\xe1\x17\x46\xa9\x80\x5f\x9c\x58\x70\xfc\xc5\xa9\x38\x77\xdf\xee\xdc\x5d\xc5\xb9\xbb\x2e\xe7\x6e\x7f\x0f\xe7\xcc\xd4\xb0\x9a\xb2\x0f\xed\x5e\x4d\x2a\x5e\x4d\xda\xbd\xda\x8f\x4e\x85\x7d\x12\x52\x63\x7c\x30\x61\xb1\x6c\xec\x18\xd7\x0b\x34\x4c\xfc\xbb\xbb\xbb\x87\xff\x3c\x1a\xfc\x4a\x68\xfe\x7f\x17\x1a\x2e\x74\xee\xdf\x8c\x86\x8b\xbc\xba\x08\x0d\xfb\xc7\xab\x23\xe1\x19\xe2\x20\x59\x06\x33\xe0\xf8\x0a\xff\xd8\xe0\x70\x57\x61\x5b\xc3\x65\x6b\xaa\x51\xaa\x00\x16\x83\xc5\x62\x30\x82\xc1\xf3\xf3\xe0\x79\x54\x91\x42\x16\xc0\xe0\x1d\xc9\xb2\x41\x75\x3c\x26\x9a\x04\xb0\x6b\xb0\x7d\x71\x5e\x0c\x60\x51\x33\x64\xae\x13\x8f\x05\xa5\xe7\xa6\xdc\x89\xd2\x03\xb0\x31\xf6\x48\x7c\xc4\x6d\x97\xd0\xf1\x08\xd5\xee\x4d\xfd\x84\xd5\x94\xda\x57\x5f\x23\xb1\xce\x36\xda\x08\x36\xe3\x2e\x76\xef\x00\x92\x0d\x8f\xcc\xcf\x93\xe1\x75\x53\xc8\x5c\x34\x19\xfe\x9f\x4e\xa9\x72\x4f\x21\x74\x48\xe6\x66\xf5\x46\x72\x73\x32\x6d\x15\xd9\x5f\xb5\x0e\x9f\xf6\x8c\xfc\x54\xdb\x29\xc4\xa8\x32\x22\x35\x7f\xdc\x23\xd0\x1e\xdb\xed\x27\x42\x0e\x8d\xfa\xca\x9c\x5b\x8c\x8d\x9e\x08\x8c\x1c\x85\x99\x15\x5b\xac\x9e\x3b\x2c\x1e\x24\x55\xc6\xa8\xf1\x88\xba\x87\xf6\x73\xed\xd0\x70\xe0\x0d\xae\xdb\xe3\x38\xa8\x26\x1c\x66\xb9\xfe\xc2\xfe\x3d\x1c\xd3\xc7\x30\xe9\xc8\x80\x0d\xc5\x68\x25\xdc\x95\x98\x31\x12\xe1\x70\x50\xe9\x66\xd3\x13\xf6\xde\xb3\x36\x4d\x86\x36\x7f\xdd\x07\x76\x9e\x1f\xd7\x7b\x92\x04\x16\x5d\x5c\x53\xbe\xc1\xee\x95\x3a\xaa\x0d\x47\xf8\xb9\xd9\x46\xa5\x43\xda\xe1\x6c\x87\x7a\xa1\xaa\x84\xd4\x79\x20\xe6\xe9\xba\xa3\x4e\x05\x1e\x73\x9d\xe6\x32\xfb\x66\x97\xe9\xfc\xdc\x75\xae\x2d\x6c\xf1\x35\x91\xa6\xf8\xa6\x9b\xbb\x31\x8b\xdc\xfc\x56\xf3\xdb\x05\x2a\xc8\xb4\xd1\x14\x8d\x79\x01\x42\xcb\xe2\xfd\x48\xa5\xc9\xb0\xf0\x75\x96\x7b\x0b\xbf\xfd\x06\xd4\x7d\xb2\x63\xd3\x3c\x8e\x33\xd5\x3e\xc4\x5a\x68\xfd\x4b\x45\xa7\xc9\x90\xba\x3f\xf0\x18\xe6\x26\x29\x67\x16\xcc\xd3\x66\xe5\xfb\x16\xbb\x7c\xb4\x80\xc2\x10\x79\x3c\xce\x23\x06\x0f\x26\x7e\x4b\x61\xf6\x7d\xa4\xba\x46\x9d\x8a\x58\x75\xec\x25\x07\x82\x3f\xc2\x67\x85\xdb\x8e\x40\x6b\x24\x56\x2d\xa8\xd9\x67\x16\x2b\xdc\x76\x30\x41\x4d\xd7\x7e\xcf\x80\x15\x6e\x7b\xa4\x8f\xed\xde\x43\xb1\x4d\x34\xb6\x91\xec\x19\x68\x46\xac\x8f\xa1\x2d\x50\xfb\xd4\x0f\x8c\x1a\xb1\xff\x61\x4a\xad\x70\xe9\x6c\x76\x9e\x4c\x6b\xea\x96\x17\x13\xde\xb3\xe4\xa5\xc8\x6f\x61\xb8\xd3\x67\x97\x12\x4e\xed\x07\x98\x6e\xaa\xb3\xf1\x1c\xea\xda\xff\xa9\xa7\xdd\x67\xd3\xfb\xd6\xc6\x99\x9d\xa5\x7d\x01\x53\xcd\x08\x73\x0b\x23\x98\x74\x6d\x18\xc8\x14\x7e\xab\xe9\x3c\xd5\xdd\x8e\x5f\x94\x52\xb3\x05\x95\x92\x49\x46\x10\x76\x1e\xa7\x60\x08\xc4\xfd\xa4\x60\x0a\xa1\xb9\x75\xe6\xa2\xe0\xab\xc9\xa5\x6e\x55\xad\xcf\x2f\xb3\x3e\xfe\x26\xf3\x85\xd2\xb7\x52\x66\x24\x91\xd8\x63\x68\xdf\xa6\x6a\x79\x07\xb7\x86\x9e\x3e\x87\x5f\x31\xd2\xae\x79\x1b\xd6\x38\xb1\xa5\x48\x46\x2c\x3f\x0d\x74\x4d\x4a\x7c\x41\xa9\xb0\x6d\x9e\x26\x86\xa1\xd5\xc2\x7f\xee\x63\xe9\xca\x57\x61\x2b\x7c\x2e\x03\x65\xd5\x8f\x79\x54\x46\xb5\xaa\x76\x52\xd9\x9b\xf3\x4b\xe9\x63\xcc\x3f\x03\x00\x00\xff\xff\x60\x7d\x55\x48\x2d\x1f\x00\x00")

func templatesOverviewHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesOverviewHtml,
		"templates/overview.html",
	)
}

func templatesOverviewHtml() (*asset, error) {
	bytes, err := templatesOverviewHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/overview.html", size: 7981, mode: os.FileMode(511), modTime: time.Unix(1563247125, 0)}
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
	"assets/inject.php":       assetsInjectPhp,
	"templates/overview.html": templatesOverviewHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"inject.php": &bintree{assetsInjectPhp, map[string]*bintree{}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"overview.html": &bintree{templatesOverviewHtml, map[string]*bintree{}},
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
