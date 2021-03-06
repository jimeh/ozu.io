// Code generated by go-bindata.
// sources:
// static/furtive.min.css
// static/main.css
// templates/_foot.html
// templates/_head.html
// templates/index.html
// templates/redirect.html
// DO NOT EDIT!

package web

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

var _staticFurtiveMinCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x5a\xef\x8e\xe4\xa8\x11\x7f\x95\xd6\x9c\x4e\xba\x8d\xc0\xb2\xdd\xff\xa6\xdd\x4a\x94\xe8\xbe\xe4\xc3\x25\x2f\x70\xda\x0f\xd8\xe0\x36\x59\x0c\x16\xe0\x99\xee\xb3\xfc\xee\x11\xe0\x7f\xb8\xe9\xb9\x89\x26\x3b\x7b\x37\x33\x6b\x17\x55\xbf\x2a\x8a\xa2\x5c\xa6\x9c\x0b\x7c\xeb\x6a\x24\x2f\x94\x67\x71\x8f\x14\xc5\x04\x60\xa2\x11\x65\x0a\xd4\x88\x72\xa0\xda\xba\x46\xf2\xd6\x61\xaa\x1a\x86\x6e\x59\xce\x44\xf1\xad\xd7\xa4\x6e\x18\xd2\x64\x22\x73\xc1\x49\x8f\x32\x54\x68\xfa\x42\x00\xca\x2a\xf1\x42\x64\x27\x5a\xcd\x28\x27\x06\x39\xcf\xe5\xef\x9a\x6a\x46\xbe\x76\xb9\x90\x98\x48\x98\x0b\xad\x45\x9d\x25\xcd\x75\x83\x85\xd6\x04\xf7\x39\x50\x5a\x0a\x7e\xe9\x4a\xc1\x35\x7c\x25\xf4\x52\xe9\xec\x18\xc7\x3d\x2e\xb9\xa3\x29\x7d\x63\x24\xa3\x1a\x31\x5a\xf4\xaa\xcd\x81\x6a\x9b\x61\x84\xfe\x41\xb2\xe3\xfe\xe7\xb3\x51\x08\x2b\x27\x1b\x9f\x1b\xa1\xa8\xa6\x82\x67\x92\x30\x64\x6c\x3b\xbf\x10\xa9\x69\x81\x18\x44\x8c\x5e\x78\x96\x23\x45\x8c\x48\x6f\x90\xb4\x68\x32\x18\xed\x49\x6d\xb0\xbb\xc1\x40\x18\xa5\x96\xf2\x72\xc9\xb8\xd0\xbf\x64\x52\x08\xfd\xa5\x33\xf3\x2b\x99\x78\xcd\x2a\x8a\x31\xe1\x7d\x49\x2f\xad\x24\xa3\x2b\x13\x52\x6f\x76\x71\x73\xed\x2b\xd9\xe5\xe2\x6a\xac\xa3\xfc\x92\x15\x82\x6b\xc2\x35\xcc\xc5\xf5\x3c\x9a\xd8\x37\x92\xcc\x68\xa8\xd5\xa2\x2f\x04\x26\xe0\x5b\x8e\x41\x23\x09\x50\xa8\x1e\xa6\x58\xa2\x9a\xb2\x5b\x56\x0b\x2e\x54\x83\x0a\x02\xa6\xab\xf3\xec\x82\x84\xd4\x7d\xde\x6a\x2d\x38\xa0\xbc\x69\x35\x10\x8d\xbe\x48\xd1\x36\x40\x11\x46\x0a\x0d\x34\xb9\x6a\x24\x09\xea\x0a\xc1\x84\xcc\x28\xaf\x88\xa4\xda\x22\x4c\x37\x53\x40\x38\xa4\xd9\xbc\x17\xaa\x68\xce\xc8\xa8\xc1\x41\x76\x06\x12\x6a\x89\xb8\x2a\x85\xac\x5d\x2c\x0c\x1c\x95\xae\xd9\xc6\x1a\xf2\xbb\xbe\x35\xe4\xaf\x8e\xfc\x15\x2c\x48\x92\x28\xa2\x3d\x8a\x6a\xf3\x9a\xea\xaf\x1d\x7c\x25\xf9\x37\xaa\x21\x6a\x1a\x82\x24\xe2\x05\xc9\x9c\xfc\xb9\x68\xa5\x12\x32\x6b\x04\xe5\x9a\xc8\x41\xd9\xef\x98\x2a\x94\x33\x82\xbf\x2e\xd5\x4e\xc4\x6e\x10\xc2\xa4\x44\x2d\xd3\x83\x50\x96\xc1\x5a\xfc\x01\x4b\x51\xb4\x0a\x52\xce\x89\x74\x96\xdc\xd3\x87\xb8\x35\x41\x85\x30\x36\xcb\x19\xf7\x96\xb5\x5b\x86\x1c\x17\xb2\x46\xac\x5f\xcc\x86\xb7\x75\x4e\xe4\xd7\x2c\x1b\xa7\x63\xd1\xa0\x6a\x28\x87\xcb\x95\x7a\xc0\x2d\x5a\xed\x73\x77\x83\x26\x1b\x2a\x4b\xaf\x11\x24\x8b\x2a\xe8\x35\xb3\x40\x25\x25\x0c\x9f\xc3\xc1\x18\x80\x99\x0d\x70\x04\x58\x18\x24\x16\xb0\xf8\x91\x00\x26\x85\x90\xc8\x6c\xbe\x90\x49\x36\x48\xac\x4d\x8a\xe8\xd1\xb5\x26\x17\x28\xc1\x28\xde\x28\xca\x5e\x88\x9c\x02\x71\x93\x36\xd7\xc9\xed\xd1\x76\x4f\xea\x4d\x74\x48\xed\x3f\x47\xb3\x39\x19\xb9\x10\x8e\x43\x2b\x34\x85\xbb\xbf\xc5\xc6\x5d\x71\x97\x6a\xb4\x09\x96\x31\x45\x15\x82\x31\xd4\x28\x92\x8d\x17\xe7\x61\xc0\xec\xba\x01\x1f\x03\x5d\x75\xb3\xbe\xe8\x22\xf1\x78\x0b\x19\x29\x75\x96\x48\x52\x8f\x06\x41\x69\x15\x19\x92\xe5\x84\x52\xbc\x2e\x13\xc4\x94\x1a\xaf\xe7\x31\xb3\x8e\xbe\x2b\x19\x59\x10\x6b\x65\x09\x4b\x46\xcb\xb0\xe4\x86\x98\x4a\x52\xb8\xec\x27\x5e\xcf\xa3\xcc\x8a\x1c\xe2\x5c\x82\xbc\x4a\xd4\x64\xe6\xcf\x0c\x30\x93\x56\xb7\x35\xba\xc2\x57\x8a\x75\x95\x25\x71\xfc\xf3\x34\x43\xe3\x47\x98\xc0\x43\xb7\x18\x3f\x44\x87\xc3\xe1\x70\xfc\xd9\x57\x96\x23\x45\xd5\x72\x70\x54\xd9\x48\x52\x12\x29\x09\x1e\xb2\xdc\xc4\x11\x12\xf3\x15\xa7\x9e\xe2\xed\x36\xda\x9a\xff\x82\x8a\x17\x83\x0f\x14\xcf\x1c\x21\x31\x5f\xf1\xd6\x53\xbc\x8f\x83\x2a\x2d\xf9\x81\x32\x33\xe6\xb3\xfa\x0a\x76\x9e\x82\xc3\x5b\x2e\x3d\xfc\xa9\x4b\x0f\x21\x97\x1e\xc2\x2e\xdd\x7b\x8a\x9f\xdf\x72\xe9\xf3\x9f\xba\xf4\x39\xe4\xd2\xe7\xb0\x4b\xbd\x10\x8a\xc3\x1e\x75\xf4\x47\x81\x13\xfb\x3e\x0d\xc6\x29\x58\x13\x20\x64\x97\x00\xb1\xc6\x01\xa2\xaa\xc1\x3a\xfc\xee\x08\xf7\x78\x69\x08\x2f\x0d\xe1\x6d\xd7\x78\xdb\x10\xde\x36\x84\xb7\x0d\xe1\xed\xd6\x78\xbb\x10\xde\x2e\x84\xb7\x0b\xe1\xed\xd7\x78\xfb\x10\xde\x3e\x84\xb7\x0f\xe1\xad\xd0\x02\x58\x01\x24\x8b\xd3\x2d\x83\x23\x4b\xa6\x98\xc8\x92\xf3\x48\x09\x67\xcb\x42\xb0\xb6\xe6\xa1\x84\x39\x8c\x84\xa9\xc1\x3c\x7e\x17\x5c\x1f\x8c\xa5\x0f\x86\xce\x07\x23\xe5\x83\x81\xf1\xa1\x38\xf8\x1f\x16\xfe\x7b\x24\x89\xbf\xd7\x04\x53\xb4\x51\x85\x24\x84\x6f\x10\xc7\x9b\x5f\x6a\xca\xc7\xa7\x4a\x2a\x49\xfd\xa5\x0b\xad\xe0\x0f\x7b\xe8\xad\x94\x7f\xee\x83\x6f\xa5\xfc\xfb\x3c\xfc\x56\x4a\x3e\xf7\x01\xb8\x52\xfe\x99\x0f\xc1\xef\x16\xe3\x6f\x07\xf9\xee\xf9\x41\x90\xd7\xf8\x07\x06\xb9\xa7\xfc\xb3\x83\xdc\x53\xfe\xbd\x82\xdc\x53\xf2\xd9\x41\xee\x29\xff\xdc\x20\xf7\xc3\xea\x93\x82\xfc\xb0\x7b\x10\xe4\xec\xf2\x03\x83\xdc\x53\xfe\xd9\x41\xee\x29\xff\x5e\x41\xee\x29\xf9\xec\x20\xf7\x94\x7f\x6e\x90\xfb\x61\xf5\xff\x0b\xf2\xbf\x80\x2c\x43\xa5\x26\x12\x64\x59\x4e\x4a\x21\x49\xf8\x9c\xa1\x8f\x6a\x82\x94\x3b\xba\xbc\x2e\x13\xfd\x7c\xf6\x62\x0f\x4c\xc6\xf3\xc2\x9f\xb6\xbb\xd3\x33\xce\x97\x67\xbd\xf6\x8c\x6c\x3a\xf3\x1d\xd8\xd2\xe4\x88\xf3\xbc\x8f\xf2\x0b\x84\x92\xe0\x2e\x47\xc5\xb7\x8b\x14\x2d\xc7\x70\xe0\x20\xc7\x5d\xb1\x2d\x1c\x87\x90\x88\x5f\x48\x80\xa9\xdc\x9e\x8a\x24\x75\x4c\x39\x6b\x43\x2c\x83\x41\x96\xe5\x62\xb6\x74\x80\x27\xdd\xe7\xe8\x79\xe7\x78\x5e\x2b\xaa\x83\xaa\xca\xd2\x31\x30\x7a\xa9\x34\xbc\x48\x74\xbb\xe7\x92\x97\x1c\xfd\x92\x26\x07\x30\xfe\x46\xa7\xd3\x17\x27\x56\x53\xfc\x96\x50\xb2\xdb\x81\xf1\x77\x16\xc2\x48\x7e\x7b\x4b\xea\x98\x02\xf7\x33\x8b\x88\xb2\x7c\x34\x07\x67\xdd\x3e\x06\xe3\xaf\x93\x2a\xb9\x76\x8b\xb0\xf2\xbc\xa5\x0f\xae\x5f\xf9\xdb\x0e\x59\x87\xaf\xbc\x6c\x07\x9c\x9b\x57\xbe\xb5\x23\xce\xb0\xa5\x47\x2d\x79\xe1\xd2\x37\xfd\x68\x99\x27\x47\xbe\xe9\x3d\xcb\x3a\xbb\xef\xb1\xcf\xdc\x2c\x27\xa7\xbd\xe9\xa9\x3a\x9e\x9b\x21\x51\x7d\x1d\xef\xdc\x41\x5e\x3c\x6c\x89\xe1\x10\xcf\x70\xdc\x26\x0e\x2d\x9a\x99\x61\xe8\x1c\x18\x8e\x64\x6a\x09\xd8\x33\xbf\xfa\x9a\x78\x98\xc9\xbc\xd3\xbc\xb3\xc1\xfa\x96\x2c\x91\x97\x6c\x63\xdf\xc4\xf1\xa5\x23\x7e\x3a\xe0\xa7\x1e\x7e\x7a\x87\x3f\xf0\xdd\xd2\x25\x7e\x7a\x8f\xef\xf8\x9a\x78\x79\xb6\xd9\x5c\x63\xff\x6c\x33\x5e\x1d\x6c\x1a\x9e\xdb\xcc\xe3\x9c\x32\xde\xcd\x5e\x69\x92\x09\xd5\x4d\xa3\xb9\x26\xef\x3c\x34\x6d\x6e\x89\x07\xef\x31\x7a\xae\x69\xd2\x49\xc9\x30\x97\x6b\xea\x2b\x49\xef\x95\x0c\x9c\xb7\xd4\x53\x92\x06\x94\x58\xce\x4a\xd7\x6c\xd1\x7f\x4a\xd2\xe6\xfa\x9e\x17\xc5\x81\xea\x27\xdc\x2f\xdd\x1a\x6d\xdf\x5c\xdf\x57\x92\xaf\x05\x0f\x46\xd0\xb6\xf6\x96\x8d\xa3\x7f\x12\xf6\x42\x34\x2d\xd0\xbf\x49\x4b\xe0\x6f\x66\xbe\xe0\x69\x22\x6e\x0c\x75\x63\xa9\x4f\x6b\xf2\x13\x98\xee\xc1\x3f\x24\x45\x0c\x28\xc4\x15\x54\x44\xd2\xf2\xbc\x3c\x41\xdf\xc5\xb1\xd7\x80\x4b\xa2\xe7\xfd\x79\xcc\x20\xdb\x6d\x1f\x35\x60\xd9\xaf\x0b\x85\x75\xb4\xb5\x4b\x50\x25\x20\xaa\x52\x10\x55\x5b\x10\x55\x3b\x50\x25\xa0\x4a\x41\xb5\x05\xd5\x6e\xda\x50\xd1\x2e\x31\xc5\xda\x26\xde\x44\x7b\x03\xb4\x34\x64\xec\x64\xf9\xc6\xec\x52\x0b\x5c\x25\xfe\xae\x9d\x0d\xda\x46\xa7\xd3\xb3\xd3\x9f\x82\x2a\x5d\x98\x9a\x46\xcf\xe9\xd1\x8d\x6c\x41\xb5\x5d\x4e\x22\x3a\x9d\x4e\x6e\x64\x67\xcc\x5b\x8e\x38\x0b\xfb\xa8\xda\x83\x6a\x0f\x6c\x6f\x61\xa3\x2b\x8f\x25\x49\x5d\xb8\x56\x07\x50\x1d\x16\x23\xd1\xb3\x33\x24\xd7\x1c\x42\x05\x22\x55\x23\xc6\x80\xfd\xbb\xe4\x3a\xc6\x47\x52\xbb\x8e\x61\x23\x89\xb7\xde\xff\x12\x1c\x15\x02\x3c\xfd\xd6\x16\x14\xa3\xcd\xaf\x82\x2b\xc1\xc8\x13\xf8\x55\xb4\x92\x12\x39\xb7\x0f\xfb\xc8\x75\x54\x03\x3d\xd6\x48\x57\x94\x7b\x4d\x92\x24\x8e\xfb\x28\x17\x0c\xdf\xb5\x4e\xa2\x02\x35\x6a\x35\xc9\x17\x24\x29\xe2\x3a\xb3\x76\x43\xc3\xd0\x47\xff\x69\x95\xa6\xe5\xcd\xb5\x0d\x5d\x17\x76\x20\xf5\x11\x61\x8c\x36\x8a\x2a\x37\x38\xf5\x6c\x46\xf2\xd9\xe6\x70\xdb\x7e\x21\x19\x17\xb6\xe9\xb0\xee\xc4\x46\x75\xab\x09\xee\x44\x83\x0a\xaa\x6f\x59\x74\xec\xa3\x82\x11\x24\x4b\x7a\x75\x15\x51\x37\x74\xbe\xb2\xa7\xa7\xa9\x79\x62\xad\x3e\x5b\xbe\x2c\x17\xba\xea\xa3\x92\x99\xe7\x16\x29\x75\x57\x32\x81\x74\x66\x2e\x07\xaa\x4d\x15\x03\xd9\x5e\x0f\x74\x2e\x38\x19\xc8\xb6\xbd\x15\x15\x84\x1b\x7d\x43\xc0\xda\x3a\x2a\xd2\x57\x0d\xe1\x30\xb0\xf0\x80\xa3\x0c\xc3\x4e\xc1\x62\x74\xd0\x62\x07\x8d\x21\xb3\x97\x17\x4c\xce\x42\x61\x7c\x01\x6f\x10\xaa\x42\x0a\xc6\xa6\xbe\x17\xbc\x65\x8e\xd2\x47\xaf\x49\x1c\x03\xfb\x17\x42\xd5\x2d\x4a\x4f\xbf\xc9\xff\xae\x13\x2f\x1f\xc5\x4c\x71\x42\xa1\xdc\xee\x3e\x07\xd6\x17\x88\xbf\x20\x05\x68\x29\x51\x4d\x00\xad\x2f\x63\x43\x5a\xbd\x5c\xa6\xa6\x34\x78\xa1\x98\x88\x6e\xdd\x47\xca\x25\x96\x10\xda\x4a\x87\x60\x60\xb6\x04\x70\xfb\xc2\x54\x28\xc3\xa5\xa9\x02\xa6\x4b\x42\x46\x0e\x46\xf9\xb7\xe1\xd2\x96\xa1\xae\xf0\x95\x08\xd3\x56\x65\xdb\xe6\x3a\x82\x2f\x8b\x3e\xd7\x45\x34\x91\x3f\x74\x23\x1f\x55\x7d\x56\x72\xae\xfb\xc2\x72\x81\xc2\xcf\xca\x2d\x4a\xbf\xb0\xe0\xba\xf6\xb3\x52\x8b\xea\xef\x81\x99\x77\x45\xcd\x07\xdc\x35\xb5\x5c\x4d\x82\xdd\xd8\x74\x3d\xa8\x4d\xa7\x56\xed\x4f\x45\x51\x2c\x72\xfc\xd9\x46\xe4\xdc\xfc\xb5\x5b\xe1\x1c\x88\xf4\x5c\xf3\xf1\x1d\xc2\x5e\xbb\xd7\x08\x7b\xe9\xde\x24\xc2\xe5\x6d\x0c\xcc\xff\x51\xbc\xff\x32\x2a\x8d\x6d\x46\x1a\xad\x5f\xf4\x6e\x85\xcc\xec\x57\x09\x0d\x92\x84\xeb\xf3\xfa\x7d\x61\x94\x78\xa4\x2e\x20\x9b\xe0\x43\x89\xf6\xfd\xec\x4d\x5f\xdb\x4f\xe9\xf1\x94\xe2\xd3\xf9\xd1\x4b\xca\xd9\x7b\xd5\x18\x31\x96\x5e\x18\x28\xb3\x2f\x06\x82\xf7\x6e\x55\x96\x65\x40\xc5\xf4\xc6\x35\xad\xe9\xda\xb8\x14\x9d\x8e\xc7\x90\xa4\xad\xe0\xef\x8d\xb3\x20\xbe\x75\x8e\xb4\x34\xcf\x51\xde\x61\x5f\x82\x9f\xcb\xc3\xbe\x9f\xc3\xcf\x37\xcf\x2e\xee\xfe\x04\xdc\x8f\x89\xdc\x7b\x8c\xfb\x6d\x11\x32\x1a\xdd\xd6\x36\xa3\xdb\xca\x64\x74\x7b\x87\xc5\x56\xdb\xee\x00\xdc\xcf\xb4\x97\xbc\x44\x32\xbe\x51\xed\xb6\xfb\x74\x1f\x98\xb4\x7b\xd5\xba\xb7\x52\x12\xec\x1b\x69\x08\x4b\x1b\xcd\xfd\x3b\x9c\x8a\x0f\x29\x4a\xd0\x58\x25\xf8\xfb\xb5\x8f\x18\x55\x1a\xc2\x96\xdb\x47\x3a\xee\xec\xad\xbd\x86\xfa\xd6\x0c\xdf\x60\xb8\xaf\x84\xa6\xcf\x24\x3c\x84\x55\x6d\xe6\x68\xde\x43\xe2\xbc\x78\x7a\xd8\xf3\x85\x0a\x61\xf1\xea\xf6\xfc\x7b\xf3\x68\xe0\xcb\x9f\x45\x79\x16\x3d\xbb\xb9\xd8\x87\x51\xf7\xb1\xaf\x22\x6c\x02\x82\x54\x93\x5a\xd9\x01\xa8\x34\x92\x7a\x3e\x59\x71\x09\xca\x11\xc3\xbc\x83\x1d\x1b\xf7\x0f\xf4\xbe\xfa\x1a\xdf\x53\x26\x6b\x47\x2e\x5b\x89\xfb\xcd\xc8\x4d\xbc\x89\x17\x0d\x49\x7b\x3b\x5f\xba\xaf\x50\xfe\xa6\x2b\x82\xf0\xea\x73\xb9\x3f\x4d\xf7\x43\x65\x80\xe7\x12\x61\x5a\xd2\xd4\xae\xdf\xab\x90\xd8\x7d\xb3\xe1\xbe\x55\xf2\x6b\xe4\xfe\xbf\x01\x00\x00\xff\xff\xe5\x79\x61\x74\x12\x28\x00\x00")

func staticFurtiveMinCssBytes() ([]byte, error) {
	return bindataRead(
		_staticFurtiveMinCss,
		"static/furtive.min.css",
	)
}

func staticFurtiveMinCss() (*asset, error) {
	bytes, err := staticFurtiveMinCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/furtive.min.css", size: 10258, mode: os.FileMode(420), modTime: time.Unix(1468778746, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticMainCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x54\x4f\x6f\xdb\x3e\x0c\xbd\xe7\x53\x10\xfd\xa1\xc0\x6f\x5d\x9c\xda\x49\x93\x76\xea\x69\x1b\x3a\x6c\x18\x86\x1d\x76\xd8\x59\x96\x68\x47\xa8\x2c\x1a\x12\xb3\x24\x2d\xf6\xdd\x07\x29\xce\xbf\x26\x5d\x03\x6c\xd8\xc5\xb0\xa8\xc7\x47\xf2\x91\x54\x49\x7a\x09\x8f\x3d\x80\x8a\x1c\x67\x95\x6c\x8c\x5d\x0a\x38\xfb\xda\xa2\x83\x6f\xd2\x85\xb3\x3e\x7c\x44\xfb\x03\xd9\x28\xd9\x87\xb7\xde\x48\xdb\x87\x20\x5d\xc8\x02\x7a\x53\xdd\xae\x3d\xe7\x68\xea\x29\x0b\x70\xe4\x1b\x69\xa3\xd9\x1a\x87\xd9\xb4\x33\x17\x83\xf1\x6d\xef\x67\xaf\x37\x98\x93\xb7\x3a\x05\x6c\x29\x18\x36\xe4\x04\x54\x66\x81\x3a\xba\x30\xb5\x02\xf2\x76\x91\xfe\x71\xc1\x99\xb4\xa6\x76\x02\x14\x3a\x46\x1f\xad\x25\x31\x53\xb3\x01\xcd\x8d\xe6\xa9\x80\x22\xcf\xcf\x77\xe9\x07\x8a\x1c\xa3\xe3\x14\xa7\xc3\x8c\x6f\xae\x56\x3e\xdb\xb8\xb2\x0c\x64\x67\x8c\x9b\xd0\xe3\x48\x03\x60\xb1\xe2\xcd\x21\x6b\xe8\x21\x63\x2f\x5d\xa8\xc8\x37\x02\xd2\xaf\x95\x8c\xff\x67\xe3\xfc\xbc\x0f\xf1\xfb\x2a\x01\xe9\x34\x58\x13\x4e\xc3\xcd\xb1\xbc\x37\x7c\x12\xf6\x65\x4c\xd4\xc6\x52\x4d\xd9\xdc\xcb\xb6\x45\xff\xbb\x96\xc3\x7b\x72\x1a\x5d\x40\xfd\x62\xf3\xd7\xb4\x5b\xba\x60\x1e\x50\xc0\xa4\xd3\x7a\xdd\xfe\xeb\xae\x5f\x8d\xf4\xb5\x71\x59\x12\x3b\x1b\x8e\xa3\x31\x52\x04\x4b\xb5\x74\x4f\x49\x8a\xcd\x7d\x83\x21\xc8\x1a\x13\x60\x97\x62\x98\x3f\x41\x0c\xd0\x7b\x5a\x15\xa7\xc8\x92\x17\xf0\x1f\x5e\x5f\xa9\x91\xda\xe7\x31\xae\x9d\xad\x86\xa3\x94\xea\xbe\xf6\x34\x73\x3a\xeb\x1c\x92\x80\xad\xf4\xe8\x78\x35\x70\x5e\xa3\x8f\x73\xed\xd2\x9c\xac\x69\x47\xa3\xd1\xed\x9f\x6f\xcd\xaa\xd0\x61\xb1\x52\xe7\x99\x45\x6a\xa5\xd6\xc6\xd5\x02\xf2\xad\x84\xdd\xe1\xd8\x92\xec\x16\x3a\x30\xae\xa2\x3d\x39\x26\x93\xc9\x36\x3a\x2f\x2d\x0a\x30\x2c\xad\x51\x4f\x72\x2a\x86\x6b\x69\xe3\x5c\x1d\x28\x5f\xe4\x7b\xd7\x5b\x41\xff\x85\x20\x07\x5b\xbf\xcd\x41\x54\xa4\x66\x01\x1e\x81\x66\x1c\x1f\xa0\xae\x71\xb0\x41\x95\x33\x66\x72\x7b\x92\x54\x55\xf5\xd7\x5a\x59\xdc\x1c\xc9\xbc\x24\xab\x53\xa2\x42\xac\xb7\x3a\x25\x9b\xb5\x56\x2a\x9c\x92\xd5\x71\x21\xe1\xf2\x02\xbe\x63\xf9\xd9\x70\x1f\xde\x59\xe3\xee\xfb\x70\xa7\x6b\x84\x8b\xcb\x9d\x5c\x95\x4a\xa3\x2c\xd2\xab\x74\xe0\xfe\x85\x1e\x8c\xb5\x12\x3e\x18\x8f\x15\x2d\xe0\x0a\x98\xa0\xb8\x39\x42\x01\x40\xad\x54\x86\x97\x02\xa0\x48\x8c\x27\x52\x16\x6f\x5e\x9f\x44\x17\x9f\xb9\x67\x8a\xfc\x14\xe7\xd4\x21\xc3\xdd\xa2\xb5\xe4\xd1\x43\x91\x67\x45\x71\xb4\xd0\x5f\x01\x00\x00\xff\xff\xb9\x1d\xa4\x50\x9f\x06\x00\x00")

func staticMainCssBytes() ([]byte, error) {
	return bindataRead(
		_staticMainCss,
		"static/main.css",
	)
}

func staticMainCss() (*asset, error) {
	bytes, err := staticMainCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/main.css", size: 1695, mode: os.FileMode(420), modTime: time.Unix(1468801300, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_footHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x50\xb0\xd1\x4f\xca\x4f\xa9\xb4\xe3\xb2\xd1\xcf\x28\xc9\xcd\xb1\xe3\x02\x04\x00\x00\xff\xff\xd2\x42\x65\xbd\x12\x00\x00\x00")

func templates_footHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templates_footHtml,
		"templates/_foot.html",
	)
}

func templates_footHtml() (*asset, error) {
	bytes, err := templates_footHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/_foot.html", size: 18, mode: os.FileMode(420), modTime: time.Unix(1468764419, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_headHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x92\xc1\x6e\x13\x31\x10\x86\xef\x7d\x0a\xe3\x6b\xb3\xeb\x45\x44\x14\x55\xeb\xa0\x2a\xed\x01\x09\xa9\x88\x52\x09\x8e\x8e\x3d\x9b\x1d\xd5\xf6\x18\xcf\x6c\xd2\x20\x1e\x1e\xa5\xa1\x28\x15\x97\xf4\xe4\xf9\xc7\xfe\xfc\xd9\xd2\xf4\x6f\xae\x6f\x97\xdf\x7e\x7c\xb9\x51\xa3\xa4\xb8\x38\xeb\x0f\x8b\x52\xfd\x08\x2e\xec\x0b\xa5\xfa\x04\xe2\x94\x1f\x5d\x65\x10\xab\x27\x19\x9a\x0f\xfa\x78\x6b\x14\x29\x0d\xfc\x9c\x70\x63\xf5\xf7\xe6\xfe\xaa\x59\x52\x2a\x4e\x70\x15\x41\x2b\x4f\x59\x20\x8b\xd5\x9f\x6e\x2c\x84\x35\xbc\x20\xb3\x4b\x60\xf5\x06\x61\x5b\xa8\xca\xd1\xe1\x2d\x06\x19\x6d\x80\x0d\x7a\x68\x9e\xc2\x4c\x61\x46\x41\x17\x1b\xf6\x2e\x82\x7d\xfb\x7c\x91\xa0\x44\x58\xd0\xaf\xa9\x45\xea\xcd\x21\xfd\xa7\x08\xc0\xbe\x62\x11\xa4\x7c\x64\x39\x40\xaa\x51\x57\xea\xba\x4e\xf9\x01\xb2\xba\xff\xfa\x59\xdd\x8d\x54\x05\x32\xd4\x67\x45\xc4\xfc\xa0\x2a\x44\xab\x59\x76\x11\x78\x04\x10\xad\xc6\x0a\x83\xd5\xc6\x24\xf7\xe8\x43\x6e\x57\x44\xc2\x52\x5d\xd9\x07\x4f\xc9\x0c\x94\xa5\x71\x5b\x60\x4a\x60\xe6\xed\xfb\xf6\x9d\xf1\xcc\x2f\xda\x6d\xc2\xdc\x7a\xe6\xd3\x44\x7b\x92\xdb\x35\xd1\x3a\x82\x2b\xc8\x4f\x16\xcf\xfc\x71\x70\x09\xe3\xce\xde\x16\xc8\xe7\x77\x2e\xf3\xf9\x92\x72\x80\xcc\x10\x2e\x2f\xba\xee\xf7\xbf\xfe\xe5\xbc\xeb\x50\x5c\x44\x3f\x9b\x77\xdd\xec\xa2\xeb\xb4\x92\x5d\x01\xab\x05\x1e\xc5\x9c\xfa\x10\x16\x27\xe8\xcd\x30\x55\xc1\xcd\xeb\xfe\xf0\x17\x4d\xee\x88\xe9\xcd\x61\xd6\xf6\xe5\x8a\xc2\x6e\x71\xf6\x27\x00\x00\xff\xff\xf7\x58\x69\x35\x96\x02\x00\x00")

func templates_headHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templates_headHtml,
		"templates/_head.html",
	)
}

func templates_headHtml() (*asset, error) {
	bytes, err := templates_headHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/_head.html", size: 662, mode: os.FileMode(420), modTime: time.Unix(1468801417, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x53\x41\x93\x9b\x3c\x0c\xbd\xef\xaf\xd0\xf8\xb2\xc9\x01\xb8\x7f\x1f\xe4\xd0\x99\x3d\x74\x26\xa7\xb4\x39\x77\x1c\x2c\xc0\xad\xb1\x18\x59\x4e\x9a\xd2\xfc\xf7\x0e\x81\x00\x3b\xd9\xf6\x84\xd0\xb3\xde\x93\x9f\xe4\xbe\x17\x6c\x3b\xa7\x05\x41\x7d\x6b\x50\x9b\xb4\x91\xd6\xa9\xdb\xed\x25\x37\xf6\x0c\xa5\xd3\x21\x14\xea\x42\xec\x8c\xda\xbd\x00\xac\xb3\x25\x79\x41\x2f\xf7\xfc\x7b\xc4\x51\x4d\xc9\x85\x75\xd7\x21\x4f\xf0\xf3\x01\xb5\xa3\x5f\x31\xb5\x94\x67\xc6\x9e\x3f\x3a\x14\x1c\xd5\xda\xab\x9d\x06\xc3\xd1\xff\x40\x0f\xc7\xc3\x1e\x42\x43\x2c\xe8\x91\x57\x75\xab\xb0\xef\x6d\x05\xe9\xf1\xb0\xbf\xdd\x46\x28\x94\x6c\x3b\x01\xb9\x76\x58\x28\xc1\x9f\x92\x7d\xd7\x67\x3d\x66\xe7\xde\x2e\xd6\x1b\xba\xa4\xe4\x1d\x69\x03\x05\x54\xd1\x97\x62\xc9\x6f\xb6\xfd\x74\x02\xc0\x50\x19\x5b\xf4\x92\xd6\x28\x6f\x0e\x87\xf0\xd3\xf5\xb3\xd9\xbc\xde\x3b\x3a\x1e\xf6\xaf\xdb\xb4\xa2\x32\x86\xcd\xf6\xff\xa9\x68\x6a\x21\x1b\xd5\x9e\x7d\x6a\x31\x04\x5d\xe3\x62\x91\xf5\x5d\x14\xb0\xa6\x50\x0f\x52\x05\xe4\xef\xac\x85\x92\xc6\x86\x34\xa0\xc3\x52\x36\x5b\x05\x8c\xda\x90\x77\xd7\x42\x3d\x22\x05\x67\xed\x22\x16\xaa\xef\x47\x07\x3e\xf4\xde\xfa\x8a\x66\x00\xc0\x86\xd1\x52\xa8\x88\x21\xd7\xd0\x30\x56\x77\x86\xaf\x9a\x6b\x14\xf8\x0d\xc3\x46\x0c\x5c\xeb\x9c\x70\xf4\xe5\xb0\x35\x0f\x38\xcf\xf4\x2c\xf6\x97\xc1\xa0\x37\xd3\x50\xc6\x21\xbd\x31\x13\x3f\xc6\xf4\xec\x0a\xe0\x80\xcf\x9d\xf6\xfd\x58\x30\x2b\xfe\x4b\x60\x4d\x57\x11\xb7\x8b\x11\xc3\xdf\x72\xf9\xc9\xf0\x65\x39\x14\x78\xdd\x62\xa1\x22\x3b\x05\x9d\xd3\x25\x36\xe4\x0c\x72\xa1\x1a\x91\xee\xbf\x2c\x5b\x39\x97\x9f\xa2\x08\xf9\xa9\x3a\xc4\x53\x6b\x45\x3d\x54\x4f\xe2\x93\xe4\xe4\x22\xae\x0a\x00\xbe\x8c\xcb\x0b\xb9\x9d\xbb\xd3\x50\xe9\x44\x33\xd3\x25\x61\x5b\x37\xa2\x40\xb3\xd5\x49\x63\x8d\x41\x5f\x28\xe1\x81\x22\xcf\xec\x4a\x37\x1b\x85\x17\xbf\x97\x4b\xcd\x7e\x4c\xc1\xf4\x79\xf7\xcc\x2b\x22\x99\x9f\xf9\x9f\x00\x00\x00\xff\xff\x4e\x97\x9b\xac\x01\x04\x00\x00")

func templatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesIndexHtml,
		"templates/index.html",
	)
}

func templatesIndexHtml() (*asset, error) {
	bytes, err := templatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.html", size: 1025, mode: os.FileMode(420), modTime: time.Unix(1468801046, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRedirectHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x8f\x41\x0b\x82\x40\x10\x85\xef\xfe\x8a\x61\xbc\xbb\x79\x1f\x17\x34\xd7\x36\x50\x0c\x99\x43\x1d\x25\x87\x3c\xa4\x82\x59\x10\xe2\x7f\x0f\xd7\xba\xf5\x4e\x03\xef\xe3\xf1\x0d\x59\x2e\x72\xed\x01\x90\x35\x71\xba\x1e\x00\xc4\x47\xce\x8d\x2e\x86\x97\x34\x70\x92\xb1\xab\x7b\xe9\xa7\xfb\x9b\xd4\x56\xac\xb4\xfa\xe1\x94\x94\xe9\x05\x92\xc3\xbe\xcc\xcb\x2a\x42\x3f\x73\x41\x60\x73\xe6\x08\xfd\x9d\x0b\x7e\x87\x6d\xf8\x6f\xd5\x86\x5b\xcd\xad\x40\x33\x5c\x9f\x9d\xf4\x13\xb4\xf5\x03\x3a\xc7\x52\x0c\xb6\x32\x59\x84\xf3\x1c\x70\x3d\xde\x64\x5a\x16\xd4\xad\x8c\x42\x2a\xd6\x81\xb3\x59\x25\xb4\x47\x6a\xfb\xe6\x13\x00\x00\xff\xff\xf9\x4e\x03\x95\xd5\x00\x00\x00")

func templatesRedirectHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesRedirectHtml,
		"templates/redirect.html",
	)
}

func templatesRedirectHtml() (*asset, error) {
	bytes, err := templatesRedirectHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/redirect.html", size: 213, mode: os.FileMode(420), modTime: time.Unix(1468764854, 0)}
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
	"static/furtive.min.css": staticFurtiveMinCss,
	"static/main.css": staticMainCss,
	"templates/_foot.html": templates_footHtml,
	"templates/_head.html": templates_headHtml,
	"templates/index.html": templatesIndexHtml,
	"templates/redirect.html": templatesRedirectHtml,
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
	"static": &bintree{nil, map[string]*bintree{
		"furtive.min.css": &bintree{staticFurtiveMinCss, map[string]*bintree{}},
		"main.css": &bintree{staticMainCss, map[string]*bintree{}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"_foot.html": &bintree{templates_footHtml, map[string]*bintree{}},
		"_head.html": &bintree{templates_headHtml, map[string]*bintree{}},
		"index.html": &bintree{templatesIndexHtml, map[string]*bintree{}},
		"redirect.html": &bintree{templatesRedirectHtml, map[string]*bintree{}},
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

