// Code generated by go-bindata.
// sources:
// templates/cf_lb.tf
// templates/concourse_lb.tf
// templates/network.tf
// templates/network_security_group.tf
// templates/output.tf
// templates/resource_group.tf
// templates/storage.tf
// templates/tls.tf
// templates/vars.tf
// DO NOT EDIT!

package azure

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

var _templatesCf_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x57\x3d\x8f\xe3\x36\x10\xed\xfd\x2b\x08\x21\x55\x00\x79\x3f\x71\x48\xa3\x22\x41\x8a\xa4\x0b\x70\xe9\x09\x8a\x1c\xd9\xc4\xd2\x24\x8f\x1c\x79\xcf\x39\xf8\xbf\x07\x24\x25\xdb\xa2\x29\xaf\x7c\xa9\x82\xdb\x6e\xcd\x37\x1f\x7a\xf3\x66\x38\xdc\x33\x27\x59\xab\x80\x54\xfe\xe0\x11\x76\x54\x98\x1d\x93\xba\x22\xdf\x8e\xab\xd5\xf9\xd0\x76\x5f\x29\x07\x87\xb4\x65\x1e\x3e\xbd\x96\x8e\x2d\xf3\xfe\xdd\x38\x91\xce\x1c\x78\xd3\x3b\x0e\xa4\x62\xff\xf4\x0e\xdc\x8e\xfa\xbe\xd5\x80\x15\xa9\x78\x57\xfb\x10\x60\x45\x88\x66\x3b\x20\xf9\x5f\x43\xaa\x9f\xbe\xed\x99\x5b\x83\xde\x53\x29\x8e\x75\x32\x58\x11\xc2\x84\x70\xe0\x3d\xb5\x0e\x3a\xf9\x35\x87\x6b\xc0\x77\xe3\xde\x28\x97\xc2\x1d\x03\x7c\xcc\x81\x6e\x9c\xe9\x2d\x4d\xc1\x22\x7c\xcc\x69\x8a\x58\xb7\xc6\x6f\xd7\x01\x16\xcd\xf7\xd2\x61\xcf\x14\x1d\xfd\x46\xfb\x89\x79\x86\x98\xd8\x17\x39\x18\x5d\x79\xe0\xbd\x93\x78\x48\x71\x23\x27\xf3\x84\x14\xf8\x08\xe9\x29\xc3\x19\x4a\xa3\x8b\x50\x07\x1b\x69\xf4\x2c\x0b\x4b\x49\x58\x11\x82\x6c\xe3\x63\x6a\x84\x80\xde\x4b\x67\xf4\x0e\x34\x5e\x25\x15\x22\x1d\x17\x7e\xb4\xeb\x15\x24\x1d\x6c\x11\xad\xbf\x21\x85\x39\x06\x84\xf6\x21\xa0\x75\xd2\x04\x8f\x65\x9b\xe7\xc7\xa7\x15\x21\x42\x3a\xe0\x39\x4f\x67\xbf\x7f\xea\xd6\xf4\x5a\x44\x75\x71\x0e\xde\xcf\x66\xf0\xab\x52\xe6\x3d\x45\x35\x68\xb8\x51\x33\xb8\xbf\xb9\x0d\xa8\x81\x53\x6b\x1c\x52\xc7\xf4\x06\xa6\xa8\x9f\x03\x46\x80\x47\xa9\x63\x15\xaf\x80\x0d\xa9\x5e\x5f\x5f\x2e\x3c\xcd\x89\xff\xca\x53\x0e\x1c\x31\xc5\x76\xb8\x64\x78\x51\x57\x94\x25\x5c\x90\x55\x19\xb8\xe6\xdd\x7d\x1d\x72\x16\x8b\x32\x9b\xef\x90\xca\x60\xb8\x40\x2d\xcf\xff\x77\xb5\xfc\x80\x72\xb1\x7d\xab\x24\xa7\x32\xcc\x50\xd5\x7e\x2c\x8f\xb2\x3e\xda\x5a\xda\xb9\x91\x7a\x6d\xf9\xc1\x6c\xfd\x0e\x92\x4e\x5f\x71\x2a\x06\x53\xa7\x5c\x1a\x52\x89\x83\x66\x3b\xc9\x67\x38\x60\xd6\x2a\x99\xc0\x74\xc3\x10\xde\xd9\xa1\x22\xd5\xc0\xe7\x1d\xd7\x0a\xb3\xb6\x1e\xed\xff\xe3\xbd\xb1\xf4\x76\x0a\x72\x7d\xeb\x87\xeb\xe5\x94\x64\x43\xaa\xcf\xc8\xb4\x60\x4e\xd0\xcf\x3b\xa6\x54\x15\xcf\x51\x82\xcb\xcf\xd3\x09\x67\x96\xf1\xd0\xd9\x0d\x79\x8e\xf7\x50\xea\xbb\x16\x72\xcf\xd3\x64\xfe\x0a\x90\xc7\xa7\xe4\xa3\xd4\xa7\x0d\xa9\xfe\x40\xb4\x03\x80\xe1\xb6\xe0\xe4\x41\x99\x8d\xd4\x09\xb2\x35\x1e\x0b\x90\x88\x58\xa7\x4f\x9f\x6c\x59\xc7\x64\x26\x35\x82\xdb\xb3\x2c\xf4\xa7\xc7\xe1\xab\x77\x60\x7a\x24\xc5\xc3\x5e\x6f\x81\x29\xdc\x1e\x28\x6e\x1d\xf8\xad\x51\x82\x34\xe4\x65\xe4\x60\xa8\x66\x10\x16\x37\xba\x93\x9b\xde\xa5\xa2\xe4\xb4\x94\xba\x62\x30\xae\xa5\xad\x27\xc6\x29\xe7\xb4\xcd\x51\x29\x16\x6c\x44\x52\x1c\x1f\x12\xde\x3f\x9c\xa1\xe9\x97\x75\x5c\xee\xce\xba\x89\x79\x77\xce\x68\x04\x2d\xe2\x98\xbb\x4c\xb6\x21\xd5\x78\x16\x8e\xd2\xfa\x90\xaa\x13\x90\x0d\x79\x7d\x7d\xb9\xd7\x89\x32\x9b\xdc\x47\xc1\xc9\x87\x14\xde\xea\x2c\xde\xd5\xa3\xa3\x19\x3a\xaf\x27\x40\xce\xec\x09\xb1\x56\xed\xfa\xb4\x70\xad\x08\x69\x19\x7f\x0b\x19\x9e\xe6\xb8\x31\x2a\xfb\xdc\xab\x6c\x06\x9b\x7a\xb0\xa9\x83\xcd\x95\xc3\xc0\x2e\xf5\x80\x28\xf5\x69\x05\x2c\xcf\xd6\x85\x4b\x71\xdd\x42\xbd\x45\x8f\x43\xd3\x1a\xf3\x26\x21\x3e\x28\x04\x65\x5d\x27\x75\xea\xe0\xea\x77\xe9\xc3\xab\x42\x5c\x14\xa5\x10\xf1\x97\xc7\xd9\xb6\xcd\x1a\xd7\xc1\x97\x1e\x3c\xd2\x69\x27\x35\xe4\x69\x74\xd0\x02\xcd\xbf\x6b\x3a\x1d\x22\x2d\xde\xab\xf8\x04\x92\x5d\x18\xb6\x57\xa3\xa5\x21\x95\xf7\xaa\x0e\x88\x14\x56\x30\x64\x53\x39\x64\x8f\xa8\xe3\x38\x57\xd2\xbb\x69\x8a\x1b\x7f\x3d\xd7\x39\x96\x43\x49\x8f\xa0\xc1\xdd\x2c\xc7\xdd\x75\x09\xae\x95\xc7\x41\x8b\xb3\x9a\xa7\xb3\x7a\xfa\x40\xdd\x93\x56\xbc\xe2\xfa\x56\x57\xdf\x58\x9e\xce\x65\x1e\xc0\x59\x81\xf2\x38\x59\x81\x22\xa7\xa3\x34\x9c\xe9\x83\xca\xe3\xb2\xf9\x31\xb5\x4b\x69\x75\x5f\x46\x2d\x04\xbf\x14\x0f\xb6\xdc\x3b\xbf\x31\x1f\x2e\xf7\xf0\xdf\xa4\xc8\x37\x76\xae\xc5\xc5\x2c\x4d\x87\xcb\x97\xf0\xc2\xc1\x30\x33\x15\xee\x7a\x13\x5f\xb6\x7f\x7c\x2a\x9a\x1e\x6d\x8f\x61\xb7\x0f\x3b\xcc\xb8\xbb\x44\x9f\x69\x6d\xd9\x33\xd5\x67\xee\x0b\xcb\xce\xf8\xf0\x3f\x2f\x8a\xff\x06\x00\x00\xff\xff\xd1\x5a\x82\x47\xd5\x10\x00\x00")

func templatesCf_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_lbTf,
		"templates/cf_lb.tf",
	)
}

func templatesCf_lbTf() (*asset, error) {
	bytes, err := templatesCf_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_lb.tf", size: 4309, mode: os.FileMode(480), modTime: time.Unix(1516141378, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConcourse_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\x4d\x6e\xc3\x20\x10\x85\xf7\x9c\x62\x84\xba\xb5\x6f\x90\xb3\xa0\xe1\x27\x29\x2a\x61\xac\x01\xb2\x68\xe4\xbb\x57\x90\x18\xd5\x75\x9c\x46\x6a\xb3\x9e\xf9\x98\xef\xbd\x98\x5d\xa2\xc2\xc6\x81\xc4\xcf\xc2\x8e\xcf\x6a\x2a\x3a\x78\xa3\xfc\x24\x41\x1a\x8a\x86\x0a\x27\x27\xe1\x2a\x00\x22\x9e\x1d\xec\xfd\x0e\x20\xdf\xae\x17\xe4\xd1\xc5\x8b\xf2\x76\x1e\xfa\xf2\x10\xb4\x14\x00\x81\x0c\x66\x4f\xf1\xf9\x36\xbb\x93\xa7\x38\xd7\x85\xe5\x36\x75\x62\x2a\x93\x5a\xbf\xde\x16\x96\x9b\xd7\x93\xa3\xa6\xf4\x3e\xd6\xf1\x86\xe9\x42\x0a\xad\x65\x97\x92\xc2\xd0\x6f\x39\x80\x4c\x19\xb3\x37\x52\xcc\x42\x6c\xd3\x08\xfa\xb5\x18\x7e\xb5\x7f\x24\xf3\xba\xc3\xa3\xec\xb6\x91\x09\x80\x23\x53\xcc\x2e\xda\xaa\x6b\x28\x1e\xfd\xa9\xf0\x6d\xb3\x5e\xbe\x53\xe1\x93\xe3\x17\xde\xe0\xa7\x61\xc5\x93\x0d\xb7\xcd\xd6\xdb\xb5\x56\x9f\x18\x3b\x74\xf4\xb6\x49\xcd\x7b\x91\x2b\x8d\xe6\xa3\x4a\x2c\xd0\x89\x28\xfc\xb9\x87\x3b\x74\x68\xb0\x7f\x68\x04\xad\xc6\x80\xd1\x38\xae\xd2\x9b\xff\x64\xd0\x3f\x8d\x67\x21\xa8\xe4\xa9\xe4\x6f\x26\xd5\xb6\x32\x6f\x46\x17\x0c\xc5\x3d\xc1\xdc\x5f\xdf\x03\xd5\x6f\x76\x07\xf3\xb0\x86\xde\x5a\x63\x7e\x05\x00\x00\xff\xff\xbf\x6f\xe7\x2a\x0b\x04\x00\x00")

func templatesConcourse_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesConcourse_lbTf,
		"templates/concourse_lb.tf",
	)
}

func templatesConcourse_lbTf() (*asset, error) {
	bytes, err := templatesConcourse_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/concourse_lb.tf", size: 1035, mode: os.FileMode(480), modTime: time.Unix(1516149865, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetworkTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x5d\x4e\x86\x30\x10\x45\xdf\xbb\x8a\xc9\xc4\x57\xd8\x81\x2b\x31\xa6\x29\xed\x88\x8d\xd0\x92\xe9\x8f\x46\xd2\xbd\x9b\x12\x30\xa1\x42\xfc\xfa\x7c\x66\x7a\xcf\x1d\xa6\xe0\x13\x6b\x02\x54\xdf\x89\x89\x67\x99\x2d\xc7\xa4\x26\xe9\x28\x7e\x7a\xfe\x40\xc0\xc1\x87\x77\x84\x55\x00\x38\x35\x13\x34\xef\x19\xf0\x69\xcd\x8a\x7b\x72\x59\x5a\x53\xba\x8a\x77\xd9\xa1\x00\x50\xc6\x30\x85\x20\xc3\xa2\x34\xfd\xf2\x2f\xfb\xc0\xfe\x83\xd4\xd6\x70\xc1\x57\x01\x30\x79\xad\xa2\xf5\xee\x72\x3f\xd3\x68\xbd\x2b\x75\xef\x91\x5a\x8e\xec\xd3\x22\xb7\x58\x1b\x77\x48\x9c\x81\xbe\x46\xea\x2b\x55\x50\x14\x21\xfe\x4a\x87\x34\x38\x8a\xff\xba\xde\xc8\x86\x93\xec\xc2\xf4\x66\xbf\xda\x81\xb3\xec\x8d\xc3\xc3\x12\x00\xcd\x99\x2e\x3a\x68\x88\xa6\x84\x9f\x00\x00\x00\xff\xff\x2b\x4f\xd2\x88\xf9\x01\x00\x00")

func templatesNetworkTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetworkTf,
		"templates/network.tf",
	)
}

func templatesNetworkTf() (*asset, error) {
	bytes, err := templatesNetworkTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network.tf", size: 505, mode: os.FileMode(480), modTime: time.Unix(1516143936, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetwork_security_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x94\xcf\x8e\x9b\x30\x10\x87\xef\x7e\x8a\x91\xd5\xd3\x4a\x1b\x6d\xd9\xb0\xe2\x92\x43\x8f\xbd\xf7\x8e\x1c\x7b\x4a\xac\x12\x0f\x1a\x9b\xa4\x6d\xc4\xbb\x57\x86\x92\x86\x84\x44\x14\x55\xaa\xb2\x0a\xe7\x6f\x7e\xf3\xc7\x9f\x60\xf4\x54\xb3\x46\x90\xea\x67\xcd\xc8\xdb\xdc\x61\xd8\x13\x7f\xcb\x3d\xea\x9a\x6d\xf8\x91\x17\x4c\x75\x25\x41\xae\xc9\x6f\x24\x1c\x04\x80\x53\x5b\x84\xb3\x6f\x05\xf2\xc3\x61\xa7\x78\x81\x6e\x97\x5b\xd3\x3c\xb7\xb8\x00\x28\x49\xab\x60\xc9\x8d\xc2\x8c\x85\x25\xd7\x44\xae\x9f\xa4\xeb\x97\xb7\x3d\x5a\xae\x1f\x6c\x08\x2c\x62\xfe\x22\x52\x8d\x14\x02\x20\xa8\xc2\xb7\xc3\x01\xa0\xdb\x59\x26\xb7\x45\x17\x2e\xc6\x8a\x9d\x1a\xd1\x08\x31\x61\x71\xae\x4b\x94\x20\xfd\xad\xb5\xaf\xae\xef\xbb\xed\x2b\xb6\x14\xc3\xc6\x6b\x92\x97\x17\x01\x60\x2c\xa3\x3e\x3f\xd1\x9f\xdc\xcf\x6e\x4d\xb5\x33\x31\x4d\x69\x8d\xde\x5f\x9d\xe0\x53\x59\xd2\xbe\xeb\x4a\x81\x34\x95\x57\xb8\x2f\xba\x8a\xd4\xef\x73\x56\xc4\x21\x67\xe5\x0a\x1c\x52\x4f\x91\x31\xe8\x83\x75\xed\x03\x5e\x80\x2b\x90\x49\x72\x12\xa4\x8c\x61\xf4\x3e\xaf\x18\xbf\xda\xef\x37\x82\xce\xc1\x9e\x19\x53\x60\x70\xe0\x09\x2a\x00\x8c\x0b\x3c\x22\xd4\x38\x38\x48\xfb\x2b\x51\x62\xe1\xb3\x2a\xd0\x85\x19\xbe\x9c\x14\x4f\xd0\xe6\xe3\x7d\x6b\xf3\x96\xbd\x65\x0f\x71\x86\xe2\x74\xcf\x49\x3c\xd7\x9d\x63\xfd\x04\x7d\x92\xfb\xd6\x27\x49\xd3\x34\x7d\xf8\x73\xf4\xc7\x38\x3f\xc3\x9a\x58\x35\xc1\x95\xd7\xff\xe1\xca\xd3\x3f\x32\x25\x7d\x7d\x68\x72\xd4\x44\x33\x9a\x4d\xbd\x9e\xa1\x4a\x5f\x39\x41\x97\xe5\x7d\xff\x5a\xb2\x6c\xb9\x7c\xef\xca\xfc\x0a\x00\x00\xff\xff\xea\xda\x59\x3f\xf4\x0b\x00\x00")

func templatesNetwork_security_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetwork_security_groupTf,
		"templates/network_security_group.tf",
	)
}

func templatesNetwork_security_groupTf() (*asset, error) {
	bytes, err := templatesNetwork_security_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network_security_group.tf", size: 3060, mode: os.FileMode(480), modTime: time.Unix(1516127018, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOutputTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\x51\x6f\xdb\x20\x14\x85\xdf\xfd\x2b\x90\xb5\x87\x55\xda\x52\x2d\x92\xf7\x10\x69\xbf\x05\x11\xfb\xd6\x61\xc5\x80\xee\xbd\x38\xed\xaa\xfc\xf7\x09\x13\x67\x26\x35\x4b\x5a\x3f\xc2\x39\xdf\x3d\x1c\xb0\x0b\xec\x03\x8b\x7a\xb4\xc0\xd2\xaa\x01\x6a\xf1\x56\x09\x31\x2a\x13\x40\xfc\x12\xf5\x97\x37\xf5\x27\x20\xe0\x20\x47\x8d\x1c\x94\x91\x16\xf8\xe8\xf0\x79\xb3\x77\x74\xd8\x44\xc7\xa9\xae\x4e\x55\x35\x83\x28\xec\x6f\xa2\x92\xa6\x44\x40\x20\x17\xb0\x05\xd9\xa3\x0b\xfe\xff\xa4\x5c\x5b\xcc\xc4\x0e\x55\x0f\x52\xb5\xad\x0b\xf6\x56\xb8\x5c\x5c\x62\x76\xf0\xa4\x82\x61\x49\xd0\x06\xd4\xfc\x9a\x12\x14\xa9\xe7\xd6\xae\xe4\x25\x38\xbc\x30\xa0\x55\x46\xea\x32\xd1\x87\xbd\xd1\xad\xd4\x67\x88\xf6\x52\x75\x1d\x02\xd1\x55\x4e\x8d\xd0\xb2\xc3\x79\xf7\x8a\x77\x60\xf6\xb4\x7b\x7c\xbc\x87\xbb\xdb\x36\x4d\xd3\x64\x74\x8f\x7a\x54\x0c\xf2\x19\x5e\x97\xe0\xf8\x4d\x61\xd9\x90\x5c\x68\x26\xa4\x1c\x07\xda\x2c\x16\xa5\x87\xe1\x54\x57\x42\x10\x58\xd2\xac\xc7\x18\x8c\x31\x40\x36\x28\xa5\xfa\xf8\x9c\x8b\x4f\x3a\x0f\x96\xe8\xf0\x6e\xd4\x93\x32\x94\xcd\xfa\x1d\x06\xbf\x77\x2f\x32\xa0\xf9\x44\xfb\xbb\xed\x36\xab\x68\xbe\xf9\x56\x77\xf8\x0e\x37\x2a\xdc\x2c\x05\xe9\xee\x8c\x6b\x95\xa1\x49\x7b\xb9\xbe\xf8\x48\xa2\x27\x8e\xfb\x9e\x8c\x60\x47\xa9\xbb\xe9\x3c\xda\x9e\x1f\x4c\x84\xfc\x43\x67\xcb\xb9\xb0\x3f\x26\x59\xdc\x39\x38\xe2\xaf\xd3\xd0\xdc\xf1\x4d\xfc\x78\x98\x5c\x73\x23\x97\x5d\xed\xef\x71\x37\xc9\x7d\x39\xc3\x07\xed\x3f\x1f\x0a\x4f\x79\xf5\xff\x4d\x88\x4c\x93\xdb\x33\x7a\xc1\x7e\x5d\xd8\x9a\xbd\x3f\xde\x32\xf7\xc7\xdc\x3a\xd7\xb7\x2c\xa0\xc0\x58\x69\xba\x50\xc2\x1d\xb0\xb5\xe2\x13\xed\x6f\x00\x00\x00\xff\xff\x2b\xd8\x75\x99\xf7\x05\x00\x00")

func templatesOutputTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesOutputTf,
		"templates/output.tf",
	)
}

func templatesOutputTf() (*asset, error) {
	bytes, err := templatesOutputTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/output.tf", size: 1527, mode: os.FileMode(480), modTime: time.Unix(1515632091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResource_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\x8a\xc4\x20\x10\x45\xf7\x9e\xe2\x23\xb3\x9d\xdc\x60\xce\x22\x15\x53\x64\x84\x44\x43\xa9\x59\x4c\xf0\xee\x83\x42\xa7\x3b\x24\xdd\x0d\x5d\x4b\xf9\xf5\xeb\xf9\x84\x63\xc8\x62\x19\x9a\xfe\xb2\xb0\xcc\xe6\xf6\x62\x46\x09\x79\xd1\xd0\x7d\x88\xbf\x1a\x9b\x02\x3c\xcd\x8c\x3a\x3f\xd0\x5f\xdb\x4a\xd2\xb1\x5f\x8d\x1b\xca\x77\xcb\x28\x60\x0a\x96\x92\x0b\xfe\x9e\x10\x1e\x5d\xf0\x45\x2b\x05\x24\x1a\x63\x2b\x02\xd8\xaf\x4e\x82\x9f\xd9\xa7\x53\x5b\x2d\x2a\xaa\x28\x75\x86\x5b\x72\x3f\x39\x6b\xdc\x13\xae\xab\x79\xcf\xfa\x72\x6b\xe7\x07\x8e\x66\xcc\xf1\x6a\x5b\xb8\x76\xd8\xd5\x8b\x5d\x8d\xb7\x9a\xfd\x0f\x86\x86\x41\x38\x46\x43\xd3\xa3\xb7\x98\x28\x39\xfb\x89\xb0\xff\x00\x00\x00\xff\xff\x58\xec\xbb\xb4\xcd\x01\x00\x00")

func templatesResource_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesResource_groupTf,
		"templates/resource_group.tf",
	)
}

func templatesResource_groupTf() (*asset, error) {
	bytes, err := templatesResource_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/resource_group.tf", size: 461, mode: os.FileMode(480), modTime: time.Unix(1513729670, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesStorageTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\xcd\x6a\xc3\x30\x0c\xc7\xef\x7e\x0a\x61\x76\xee\x1b\xec\xbc\xfb\xfa\x00\x46\x71\x44\x66\x70\xa4\x20\x2b\x81\xad\xe4\xdd\x47\xbc\x36\x6d\x0d\xa5\x3d\x2e\xd7\xfc\xf4\xff\xb2\x52\x91\x59\x23\x81\xc7\x9f\x59\x49\xc7\x50\x4c\x14\x07\x0a\x18\xa3\xcc\x6c\x1e\x7c\x27\xe5\xcb\xc3\xc9\x01\x30\x8e\x04\xcd\xf7\x0e\xfe\xed\xb4\xa0\x1e\x4a\x1a\xa7\x4c\x81\x78\x09\xa9\x5f\xbd\x03\xb8\x88\x87\x41\x65\x9e\x42\xbd\xae\xf8\xc5\xeb\x1e\x38\x6c\x46\x87\x8d\x5a\xbd\x73\x00\x59\x22\x5a\x12\x6e\x1d\xaf\x96\x4a\x43\x12\xae\x5e\xe7\xb8\xc1\x12\x69\x0b\x1f\x0d\xb9\x47\xed\x6f\x39\xa5\x29\xa7\x3f\xfd\x60\xdf\x53\x0d\xf6\xf1\x79\xac\xc6\x86\x43\xa9\x7d\x01\x88\x97\xa4\xc2\x23\xb1\x5d\x6d\x6f\x2a\xae\x6e\x75\xee\xf1\x88\x51\xd8\x30\x31\xe9\xd3\x19\x6b\xd0\x8a\x3c\x18\x0e\x5e\x9e\x0e\xa0\x79\xc3\xb3\xc0\xdd\x7d\x83\x34\x02\x7b\xee\xed\x3f\x95\xb2\x4f\x34\x69\x5a\xd0\xc8\xbf\x5e\xbb\x18\x8d\x91\x72\x7e\x52\x7d\xc7\xfe\x75\xfd\x2e\x4b\xb7\x75\xff\x0d\x00\x00\xff\xff\x23\xe4\x2a\x58\x37\x03\x00\x00")

func templatesStorageTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesStorageTf,
		"templates/storage.tf",
	)
}

func templatesStorageTf() (*asset, error) {
	bytes, err := templatesStorageTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/storage.tf", size: 823, mode: os.FileMode(480), modTime: time.Unix(1513729670, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTlsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x41\x0a\x02\x31\x0c\x05\xd0\x7d\x4e\xf1\xc9\x09\x5c\x88\xe0\xa2\x0b\xaf\xa0\x07\x08\xad\x04\x5b\x6c\xa9\x24\xb1\x30\x0c\x73\xf7\x79\xa6\x3e\xff\xf6\x56\x70\x74\x97\x9f\xb5\x95\x43\xe5\xab\x1b\x83\xcb\xf4\x2a\x6b\x38\x63\x27\x20\xf7\xcf\xb4\x16\x75\x20\x81\x9f\xaf\x07\x13\x60\x9e\xa5\xb4\x70\x20\xe1\x7a\xb9\xdf\xe8\xa0\x33\x00\x00\xff\xff\x52\x4d\xac\xad\x51\x00\x00\x00")

func templatesTlsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesTlsTf,
		"templates/tls.tf",
	)
}

func templatesTlsTf() (*asset, error) {
	bytes, err := templatesTlsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/tls.tf", size: 81, mode: os.FileMode(480), modTime: time.Unix(1513729670, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x6a\xc4\x20\x10\x40\xef\x7e\xc5\x20\x3d\xa7\xcd\xa5\xb7\x7e\x4b\x30\x3a\x2d\x43\xcd\x18\x46\x63\xa1\x21\xff\x5e\x12\xc1\xee\x8a\x2c\x9b\xdc\xe6\xbd\x07\xce\x64\x23\x64\x66\x8f\xa0\x91\xf3\x44\x4e\xc3\x7e\x28\xf5\x3f\x15\xfc\xa2\xc0\xed\x34\xd2\xb2\x7a\x9c\xfa\x49\xdc\xe6\x68\x85\xd6\x44\x81\x3b\x38\x21\x1b\x4e\x1d\x60\x3d\xe1\x23\x10\xd1\x0a\xa6\x16\x32\xa6\x9f\x20\xdf\x93\x25\x27\x1a\x76\x05\xe0\xf0\xd3\x6c\x3e\xc1\x07\xe8\xf1\x6d\xb8\xfe\xd7\xf1\x5d\xab\xbb\x8c\x38\xa1\xb0\xf1\xcf\x75\xab\x84\x4c\x0e\x05\xb4\xf9\xdd\x04\x65\x29\x45\xb3\xe9\x59\xbe\xec\xd9\xc8\xd0\x80\x43\x2b\x80\xba\x37\x94\xaf\xca\x15\x5c\x5a\xbd\x42\xab\x55\x70\xab\x95\x9b\x74\xb4\x02\x8e\xf3\xf5\x7f\x01\x00\x00\xff\xff\x59\x69\xa5\xda\xe3\x01\x00\x00")

func templatesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVarsTf,
		"templates/vars.tf",
	)
}

func templatesVarsTf() (*asset, error) {
	bytes, err := templatesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars.tf", size: 483, mode: os.FileMode(480), modTime: time.Unix(1516130326, 0)}
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
	"templates/cf_lb.tf": templatesCf_lbTf,
	"templates/concourse_lb.tf": templatesConcourse_lbTf,
	"templates/network.tf": templatesNetworkTf,
	"templates/network_security_group.tf": templatesNetwork_security_groupTf,
	"templates/output.tf": templatesOutputTf,
	"templates/resource_group.tf": templatesResource_groupTf,
	"templates/storage.tf": templatesStorageTf,
	"templates/tls.tf": templatesTlsTf,
	"templates/vars.tf": templatesVarsTf,
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
	"templates": &bintree{nil, map[string]*bintree{
		"cf_lb.tf": &bintree{templatesCf_lbTf, map[string]*bintree{}},
		"concourse_lb.tf": &bintree{templatesConcourse_lbTf, map[string]*bintree{}},
		"network.tf": &bintree{templatesNetworkTf, map[string]*bintree{}},
		"network_security_group.tf": &bintree{templatesNetwork_security_groupTf, map[string]*bintree{}},
		"output.tf": &bintree{templatesOutputTf, map[string]*bintree{}},
		"resource_group.tf": &bintree{templatesResource_groupTf, map[string]*bintree{}},
		"storage.tf": &bintree{templatesStorageTf, map[string]*bintree{}},
		"tls.tf": &bintree{templatesTlsTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
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

