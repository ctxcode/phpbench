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

var _assetsInjectPhp = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x94\x51\x6f\xda\x30\x10\xc7\xdf\xf9\x14\xff\x49\x68\x49\x27\x1a\xc2\xf6\x30\x09\x96\x4e\x5b\xd7\x4e\x93\x2a\x8a\x02\x3c\x21\x84\x3c\xe7\x68\xac\x25\x76\x66\x3b\xdb\xaa\x8a\xef\x3e\x39\x21\x84\x14\x04\x2f\xe1\x72\xbf\xfb\xdf\xd9\x77\x97\x4f\x9f\x8b\xb4\xe8\xf5\x78\xc6\x8c\xc1\x2c\x2d\xbe\x92\xe4\x29\x5e\x7a\x3d\x00\x30\x96\x59\xc1\xd1\x37\x64\xa7\x1a\x11\x64\x99\x65\x93\x8e\xc7\x8a\x9c\xb4\x41\x84\xd5\x7a\x52\xc7\x14\xe5\xcf\x4c\xf0\x06\xd8\x96\x92\x5b\xa1\xa4\xb3\xb5\x5d\x38\xda\xaf\x83\xae\xf0\x52\xf1\xad\xda\x78\xbc\x97\x5b\x79\xd6\x43\xb0\x17\x5f\x23\x42\x2e\xb8\x56\xce\xf2\xad\x2e\xe9\xaa\x2e\x61\x77\x31\x9f\xa3\x6f\x55\x42\xfb\x6c\x03\xf4\xb7\x22\x23\xc9\x72\x1a\xa0\x9f\x09\x49\x53\xf7\x8e\xab\x84\x8e\x0b\xe9\x93\x4c\xce\xe5\x6b\x01\xf7\x76\x5e\x90\xb4\x88\x2a\xfa\xfa\x62\xed\x93\x43\x60\x29\x0d\x59\xff\x22\xdc\xc9\x93\xbb\x4b\xd5\xaa\x94\x89\x7f\x94\xf3\x1d\x46\x61\x18\xee\x1f\xc7\xbc\xd8\xc2\x7f\x73\x50\xaf\xfa\x75\x7c\xae\xce\x25\x37\xdd\xdc\x66\x4a\x69\xff\xd5\x59\x1b\x6d\x04\xf0\xbc\xb6\xfc\x5d\x9b\x6a\x38\x04\xf1\x54\x55\x25\x06\xf0\x72\xe3\x4d\x3a\xce\xef\x64\xc1\x97\xf1\x03\x34\x19\x55\x6a\x4e\xed\xa1\x78\xa9\x33\x44\x70\x8f\x8d\x90\xc2\xfa\x57\x93\xe3\xc8\x39\x59\x18\x95\x13\x54\xe1\x7a\x68\x70\x8d\xbf\x04\xa6\x09\x05\x33\x46\xc8\x27\x08\x09\x86\xd2\x90\x66\x4f\xee\x3a\xac\x52\x48\x49\xb7\x19\x2a\x65\x43\x56\x15\x76\xc3\xb4\x66\xcf\x7e\x95\x73\x80\x55\xe7\x2a\x6e\x97\xf1\xc3\xe3\x6c\xb1\x89\xef\x16\xcb\x78\xba\x88\xbf\x4c\xe7\xf7\x77\x31\xa2\x1b\x8c\x06\x67\x41\x77\x9c\xe8\x06\x5e\x6a\x6d\x31\x1e\x0e\x47\xef\x3f\x06\x61\x10\x06\xa3\xf1\x87\x30\x1c\x0d\x13\x66\x99\x77\x3e\x70\xf6\x38\x5f\x5c\xd0\x75\xee\xfb\x1f\x77\x0f\xdf\xe6\x0e\xea\x16\xe9\x7e\x5e\x33\xb4\x9e\xf3\xb7\x23\x7c\x0a\x56\x4d\xad\xa8\x6e\x9f\xcf\xa0\xbf\xe8\xb9\x96\xab\x37\xe3\x14\xa8\xd7\xa3\x66\xf6\xab\x72\x0a\xb9\xd5\xa9\x11\xf7\xef\x0c\x90\x9b\xda\x9d\x9b\xae\x73\xdd\x9a\xeb\xd7\xfd\x97\x09\x6c\x4a\xd0\xf4\xbb\x24\x63\xf1\x16\x86\xfd\x71\xa6\x29\x94\x34\x04\xab\xd0\x77\x46\x3b\x51\xce\x6a\x26\x8a\xfe\x11\xaf\xfb\xdd\x95\xbd\xcd\x94\x69\x35\xad\x02\xcf\x88\x69\x94\x45\x3d\x6c\xcd\x98\x9a\xee\x14\x71\x17\x75\x90\x6b\xbe\x38\xbb\xde\xff\x00\x00\x00\xff\xff\xe3\x71\xad\x23\x31\x05\x00\x00")

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

	info := bindataFileInfo{name: "assets/inject.php", size: 1329, mode: os.FileMode(511), modTime: time.Unix(1563239455, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOverviewHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x6d\x6f\xe3\xb8\x11\xfe\x9e\x5f\x31\xd5\xa2\xb5\x73\x67\x4b\xb2\x93\x0d\xb2\x8a\x9d\xe2\xd0\xf6\x80\xde\x75\x6f\x81\xee\xa2\x5f\x7c\xfe\x40\x89\x23\x8b\x1b\x9a\x14\x48\x3a\x89\xcf\xf0\x7f\x2f\x48\xc9\xb6\x64\xbd\xc4\x0b\x1c\xfa\xe2\xc5\x46\x12\x39\x33\x9c\x97\x67\x1e\x52\x9a\x71\x26\x9e\x20\x53\x98\xce\xbd\xcc\x98\x5c\x47\x41\x90\x4a\x61\xb4\xbf\x92\x72\xc5\x91\xe4\x4c\xfb\x89\x5c\x07\x89\xd6\x7f\x4e\xc9\x9a\xf1\xed\xfc\x53\x8e\xe2\xfb\xcf\x44\xe8\x3f\x51\xa6\x73\x4e\xb6\x73\xfd\x42\x72\x0f\x14\xf2\xb9\xa7\xcd\x96\xa3\xce\x10\x8d\xf7\x78\x75\x35\x73\x8f\x8f\x57\x00\x00\xb1\xa4\x5b\xd8\xb9\x5b\xfb\xb3\xab\x8c\x0b\x8b\x11\x0c\xac\x4d\xb0\x36\x07\x0f\x4e\x62\x7f\xe5\x2e\xbe\x66\x14\x63\xa2\x4e\x6a\xb9\xd4\xcc\x30\x29\x22\x48\xd9\x2b\xd2\x87\xe3\xc4\x6f\x63\x26\x28\xbe\x46\x37\xa7\x21\x23\xf3\x28\x3c\x3d\x72\x4c\x4d\xf5\x39\x96\xc6\xc8\x75\x75\xe4\x85\x51\x93\x45\x37\x61\x98\xbf\x56\xc4\x48\xf2\xb4\x52\x72\x23\xe8\x38\x91\x5c\xaa\xe8\x1d\x22\x56\xad\xbc\x8e\x75\x46\xa8\x7c\x89\x20\xcc\x5f\xdd\xff\xc9\xfb\xfc\x15\xd4\x2a\x26\xc3\x70\xe4\xfe\xf9\x93\xeb\xd6\xb8\xfc\x94\x71\xd4\xbb\x46\x10\xd3\x33\xaf\x2b\x2e\x9d\x59\x00\x3f\x43\x42\xb1\x92\xa1\x16\x7f\xd3\xbb\xf0\xae\x12\x66\x4e\x28\x65\x62\x15\x4d\x0f\xbe\x9e\xa6\x4a\x8d\x34\x4d\x1f\xea\x95\xd2\xec\x37\x8c\x26\x77\x55\x59\x37\xfe\x82\x6c\x95\x99\x08\x62\xc9\x2b\xc5\x08\xbe\x83\x58\x2a\x8a\x6a\xac\xdc\xf4\x24\x7f\x05\x2d\x39\xa3\xf0\xee\x1e\xd3\x49\x9a\x3e\xc0\x77\x41\x7b\x38\xcc\xe0\x7a\xd7\x70\xd5\x65\xb4\xee\x6a\xb9\x40\x59\xc5\xca\x0a\x49\x92\x54\x02\xda\x28\x2d\x55\x04\xb9\x64\xc2\xa0\xea\xa9\xc2\xf9\xd2\x2f\x52\xd1\x71\xac\x90\x3c\x45\xe0\x2e\x63\xc2\x79\x47\x0d\xac\xa6\x4f\x12\xc3\x9e\xb1\xaf\x10\xa9\xfb\x5d\x10\xc2\xbd\xfb\xd5\xf2\x59\xa9\x4c\x6f\xee\x7c\xb6\x12\x52\x55\xdc\x30\xf8\x6a\xc6\x14\x13\xa9\x48\xd1\x38\x9c\x09\x1c\x9b\x4c\xc9\xcd\x2a\x6b\x54\x9e\x10\x52\x8f\x31\x91\xc2\xa0\x30\x6d\x1d\x48\x62\x2d\xf9\xc6\x60\xb3\x09\x27\xfd\x4d\x78\x57\xef\xb0\x02\x23\x61\x7d\x5d\x13\xf3\xdd\x59\xa2\x22\x10\x52\xe0\x79\xbb\x4e\xc2\xf0\x8f\xa7\xb1\x35\x51\x2b\x26\x8e\x29\x6d\xf6\x8d\x89\x39\x18\xdb\x33\x60\xb2\x26\xce\x8a\x96\x78\xa3\xff\xa7\x1f\x30\x24\x71\xa3\x8c\x0d\xa8\x4f\xe3\x38\xfd\x70\xdb\xdd\x5c\x35\x97\x1c\x3d\x1a\xda\x02\xfd\x66\x97\x76\x2d\x48\x29\x6d\xac\x76\xc4\x51\xdb\x6a\x7e\x72\x86\x13\xc2\xd9\x4a\x44\x90\x60\xb3\x59\x2a\x7a\x2a\x12\x26\x1b\x27\x19\xe3\x74\x28\x29\xbd\xee\xc1\x3c\x9c\xd8\xf2\x68\x48\x1a\xc2\xc7\xb1\x7c\x7d\x7b\x6d\xe8\xa1\x1f\xa7\x63\x14\x11\x3a\x95\x6a\x1d\xc1\x26\xcf\x51\x25\x44\xe3\x45\x0c\xd5\xb9\x8f\x14\x00\x9d\xbe\x6f\x02\x74\x7a\x56\x03\x87\x30\xa8\x8f\x56\xf9\xea\x5b\xf7\x88\xe9\x75\x0f\xe8\xa0\xce\xc9\x47\x82\x8d\xa6\xf5\xf2\x1f\xb8\xa1\x26\x74\x68\x07\xd7\x3e\x55\x89\x03\x90\x08\x65\x1b\x1d\xdd\x9f\x7a\x65\x16\x94\x3b\xf7\xd5\x8c\xb2\x67\x60\x74\xee\x91\x3c\xf7\x8a\x9d\xdc\x0d\x25\x9c\x68\x3d\xf7\x4a\xee\x29\x67\xce\x67\x8b\xbd\xc9\x7b\xfc\x82\xda\xe8\x59\x40\xd9\x73\xbb\x9c\x25\x2e\x0f\x9e\xc7\xa9\x54\x73\x4f\xa3\x19\x3d\xe1\x16\x98\x00\x85\x7a\xc3\xcd\x67\x34\xda\x83\xa8\x94\xdd\x41\x41\xb4\x91\xe5\x76\x85\xc2\x4e\xff\x8c\x5b\x98\xcf\xc1\x6a\xed\xad\x1d\x29\xa2\x84\xb3\xe4\xc9\x19\x2b\xe6\x87\x4f\xb8\xbd\xf6\x1e\x17\x0b\x27\xb5\x5c\x56\xbc\x29\x6f\xbb\x62\x03\xb7\x3d\xbc\x11\xe1\x8f\x56\xe6\x82\x08\x4f\x9e\x59\xab\x82\xac\x51\xff\xdd\x11\x36\x85\x39\x2c\x96\x0f\x6d\x71\x9e\x4b\xfa\x1c\xc5\xca\x64\x36\xe2\x10\xf6\xde\xe3\x0f\x9c\x17\x3e\x5e\x9a\x61\x2b\x6c\xd3\x7b\x34\x5c\x5d\xb5\xd8\x3f\x5a\x56\x75\xd4\xfe\x29\x1d\xda\x89\x6b\x78\x84\xf1\xe4\x3c\xd9\x85\xaa\x4d\x45\x21\xe4\xd2\xed\x16\xbb\x28\xdf\xe5\x56\x53\xcd\xb4\x21\x31\xc7\xc3\xbc\x89\xb9\x07\x09\x72\xae\x73\x92\x30\xb1\x9a\x7b\x61\x45\xb6\x90\xb7\xf5\xa8\x8f\x15\xe3\xaa\x39\x58\x2a\x3c\xfe\x45\x52\x9c\x05\x26\xeb\x96\x58\xeb\xfe\xf9\x2f\x6c\x8d\x1a\x12\xc2\x39\xd2\x7e\xc9\x1f\x9e\x51\x91\x15\x42\x97\xc5\x59\x70\xee\xa9\x95\x6b\xc4\x34\x73\x44\x7c\x3e\x86\xeb\x9c\x13\x83\xf0\x3c\x66\xe9\xdc\x3b\x75\x87\xd7\x9a\x90\x03\x18\x12\x7b\x1e\x38\x35\x9b\x6e\x91\x2e\x34\xa8\xad\xa6\x93\xf6\x6d\xca\x5c\x4d\x4d\x4b\xb2\x1b\xd2\x1f\xf5\x5f\x8b\x17\x05\x58\x2e\x5d\xe0\x3d\x4a\xe1\x85\x36\xcb\x44\x5e\x62\xba\x35\xa7\x65\xae\x1a\xe3\xf5\xbc\xce\x02\x07\xc0\x3a\x6c\x4f\xb8\x6d\x24\xfa\x88\xd4\xc3\x16\x57\x49\xe6\x17\x3b\x16\xc1\x62\x01\x6e\xd6\x39\x5c\xb3\x7b\xb4\x3f\xd3\x89\x62\xb9\x01\xad\x92\xd3\x5b\x59\x42\x85\xff\x55\x53\xe4\xec\x59\xf9\x02\x4d\x20\xf2\x75\xf0\xbc\xc1\x80\x32\x6d\xec\x8d\xff\x55\x7b\x8f\xb3\xa0\xd0\x7d\x3c\x5a\x29\x1c\x78\x26\xaa\x42\xa6\x30\x87\x9f\x3e\x7f\xfa\xc5\xcf\x89\xd2\x38\xdc\xed\xc0\xff\xe7\x71\xee\x27\x2d\x05\xec\xf7\xe5\x5e\x14\x04\x4d\xd5\x9d\x37\x79\x7f\x77\x33\xf9\x70\x1b\xde\xdc\xde\xdf\xdf\x78\xd1\xce\xb3\x78\xf8\x07\x13\xa8\xed\xc3\xc4\xfe\xf9\xb1\x64\x0f\x2f\xf2\x82\xb5\x30\x41\x12\xbc\xbc\xbc\x04\x79\x96\xc7\x28\x92\x2c\xa0\x18\x6f\x56\x81\x41\x6d\xfc\xc3\x98\xbd\xf1\x46\xce\x94\x17\x79\x9a\x23\xe6\xc3\xe9\xf5\x83\x37\xf2\xac\xe5\x5f\x94\x17\x79\xb7\xde\xc8\xfb\x9b\x30\x8a\xd9\x85\xc4\x86\xf3\x91\xf7\x51\x7b\xd1\xd4\x5e\x4a\x1c\xb8\xa7\x73\x70\x78\x51\xb8\x1f\x79\xd3\xdf\xc3\x2f\x4c\x32\x09\xbf\x7a\x54\x0a\xfc\xd5\xab\x39\x77\xd7\xee\xdc\x6d\xcd\xb9\xdb\x2e\xe7\x6e\x7e\x0f\xe7\xec\xd4\xb0\x9e\xb2\x0f\xed\x5e\x4d\x6a\x5e\x4d\xda\xbd\xda\x8f\x4e\x85\xfd\x2c\x95\x41\x7a\x30\xe1\xb0\x6c\xed\x58\xd7\x4b\x34\x4c\xc2\xdb\xdb\xdb\xfb\xff\x3e\x1a\xc2\x5a\x68\xe1\xff\x16\x1a\x2e\x74\xee\x3f\x8c\x86\x8b\xbc\xba\x08\x0d\xfb\x87\xab\x23\xe1\x59\xe2\x20\x79\x0e\x73\x10\xf8\x02\xff\xda\xe0\x70\x57\x63\x5b\xcb\x65\x6b\x66\x50\xe9\x08\x16\x83\xc5\x62\x30\x82\xc1\x72\x39\x58\x8e\x6a\x52\xc8\x23\x18\xbc\x23\x79\x3e\xa8\x8f\x53\x62\x48\x04\xbb\x06\xdb\x9f\xa8\x2a\xaa\xdc\x8f\x1a\x72\x27\xd6\x8e\xc0\x85\xd1\x23\xf1\x33\x6e\xbb\x84\x8e\xa7\xa4\x08\x16\xcb\x9e\xe9\xf2\x10\xd5\x94\xda\xd7\x1f\x13\xb9\xce\x37\xc6\x0a\x76\x85\xa6\x23\x48\x37\x22\xb1\x2f\x2f\xc3\xeb\xa6\x90\xfd\xb1\x74\xf8\x07\x93\x31\xed\x9f\x42\xe8\x90\x2c\xcc\x9a\x8d\x12\xf6\xf0\xd9\x2a\xb2\xbf\x6a\x1d\x3e\x6d\x0b\xc5\xc1\xb5\x53\x88\x33\x6d\x45\xce\xfc\xf1\x8f\x58\x7a\x68\xb7\x9f\x4a\x35\xb4\xea\x4f\xf6\x68\x62\x6d\xf4\x44\x60\xe5\x18\xcc\x9d\xd8\xe2\x69\xd9\x61\xf1\x20\xa9\x73\xce\xac\x47\xcc\x3f\x74\x98\xef\x86\x86\x83\x60\x70\xdd\x1e\xc7\x41\x35\x15\x30\x2f\xf4\x17\xee\xef\xe1\x24\x3e\x86\x49\x47\x06\x5c\x28\x56\x2b\x15\xbe\xc2\x9c\x93\x04\x87\x83\x5a\xc3\x5a\xd8\xbb\x6b\xcf\xda\x2c\x1d\xba\xfc\x75\x9f\xc9\x45\x71\x22\xef\x49\x12\x38\x74\x09\xc3\xc4\x06\xbb\x57\xea\xa8\x36\x1c\xe1\xe7\xe7\x1b\x9d\x0d\x59\x87\xb3\x1d\xea\xa5\xaa\x96\xca\x14\x81\xd8\xbb\xeb\x8e\x3a\x95\x78\x2c\x74\x9a\xcb\xec\x9b\x5d\x66\x8a\xa3\xd5\x5b\x6d\x61\x2b\x58\x9c\xbd\xe6\x10\xb6\xfb\x5f\x43\x9d\xf3\xb4\x6c\xba\x0b\xd0\x57\x15\x77\x28\xec\x52\x28\x7c\xf8\xde\x22\xf0\xc8\xb0\x5d\xe9\x6c\x1b\x2d\x13\xe4\xcc\xb4\xe4\xa7\x8f\x5e\xd6\x68\x32\x49\x75\x1b\xbb\x1c\xdf\x94\x2b\x89\xb4\xef\xcc\xed\x81\x9f\xb5\x73\x3d\x7c\xcb\xb8\x8b\x27\xdc\x76\xe4\xe0\x4c\xd7\xbd\xbc\xdb\xb7\xf2\x1e\xe9\x23\xf0\x7b\xc8\xa6\x59\xbb\x36\xba\x79\xa3\x90\x09\xef\xe3\xaa\xde\xb2\x56\xb9\x25\xe1\xff\xc7\xe4\x52\x63\x95\xf9\xfc\x6d\x5a\x39\x53\x77\x0c\x91\x8a\x9e\x25\xdb\x71\xdd\x1c\x6d\xe9\xf5\xd3\x37\x86\x0a\x4e\xdd\xd7\x86\xee\xa6\x77\xf1\x1c\xea\xda\xff\x5d\xa3\xdd\x67\x96\x0e\x0b\x1b\x6f\x70\x6c\xfb\x02\xb6\x9a\x09\x16\x16\x46\x30\xe9\xa2\x4e\xe4\x1a\xbf\xd5\x74\x91\xea\x6e\xc7\x2f\x4a\xa9\x25\xe3\x4a\x32\xc9\x08\xe2\xce\x83\x05\x0c\x81\xf8\x1f\x35\xcc\x20\xb6\x97\xce\x5c\x94\x1c\x35\xf9\x16\x5e\x3b\x59\x7f\xbc\xcc\xfa\xf8\x9b\xcc\x97\x4a\x2d\xcc\xdf\x4b\x99\x89\x42\xe2\x0e\x64\x7d\xdb\x8b\xe3\x1d\xdc\x5a\x7a\xfa\x14\x7f\xc5\xc4\xf8\xf6\x69\x78\xc6\x89\x2d\x45\x62\xa9\x25\x59\xbd\x08\x97\x7d\x44\x5b\xfb\x8a\xe9\x84\xdf\x0c\xe2\xf8\x64\x5f\xe5\xaf\x2a\xdf\x05\xfe\x1d\x00\x00\xff\xff\xa3\x29\x1e\x89\xf3\x1d\x00\x00")

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

	info := bindataFileInfo{name: "templates/overview.html", size: 7667, mode: os.FileMode(511), modTime: time.Unix(1563244174, 0)}
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
