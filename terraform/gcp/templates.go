// Code generated by go-bindata.
// sources:
// templates/bosh_director.tf
// templates/cf_dns.tf
// templates/cf_lb.tf
// templates/concourse_lb.tf
// templates/jumpbox.tf
// templates/vars.tf
// DO NOT EDIT!

package gcp

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

var _templatesBosh_directorTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x96\xdf\x4e\xe3\x3a\x10\xc6\xef\xfd\x14\x23\xeb\x5c\x80\x54\x72\x4a\x69\x4a\xf7\x82\x27\x41\x55\xe4\xa6\x6e\xc8\xe2\xc6\x91\xe3\xb4\xac\x50\xde\x7d\x65\x3b\x89\xe3\xfc\xab\x61\x41\xc0\x45\x20\x99\xf9\xe6\x9b\x9f\xc7\x89\xcf\x44\xa4\x64\xcf\x28\xe0\xa2\xdc\x67\x54\x46\x71\x7a\x10\x18\xde\x11\x80\xfc\x93\x53\x00\x80\x27\xc0\x85\x14\x69\x96\x60\x04\x70\xa0\x47\x52\x32\xa9\x6e\xde\x2f\x03\xfd\xfb\xff\xfd\x06\xa3\x0a\x21\x41\x0b\x5e\x8a\x98\x02\x4e\x38\x4f\x18\x8d\x62\x7e\xca\x4b\x49\xa3\x8c\xca\x0b\x17\xaf\x18\xf0\x7e\xcf\xee\xda\xff\x54\x8d\x8c\x9c\x74\x8d\xfe\xcf\x13\xe0\xff\xde\xcf\x44\x04\x34\x3b\x47\xe9\xa1\x6a\xb3\x10\x00\x29\x25\x8f\x62\x41\x89\xa4\x91\x31\xad\x9e\x14\xf0\x04\x47\xc2\x0a\x3a\x6b\xc5\xc6\xd7\x6e\xcc\x8d\x31\x33\x03\x0b\x75\x28\x02\x48\x73\x8d\x29\x12\x24\x4b\xa8\x0d\xec\x10\xac\x54\x58\x5d\xa9\xab\x37\x4e\x26\xe8\x70\x09\x0a\xca\x8e\x11\x4b\xb3\xd7\x6a\x9e\xea\x31\x15\xf4\x42\x18\xc3\x80\xe9\x9b\xa4\x22\x23\xcc\x6d\x63\xd0\x40\x1b\xd6\xf1\xe6\xeb\x4a\x89\x56\x18\x21\x00\x63\xc7\xf4\xae\xa0\x3f\xe3\x66\x0e\x96\x78\xa7\x02\x08\x63\xfc\xa2\x9d\x00\xe4\x5c\xc8\xc2\x98\x79\xc6\xab\x15\x5e\x00\xde\x6c\x37\x5b\x75\x5d\x85\x61\x18\xe2\x9d\x09\x13\x5c\xf2\x98\x33\x65\x47\xc6\xb9\x32\x58\x29\x29\x49\x44\x42\x65\x24\x49\x62\x2a\xb9\xfd\xec\x79\xf1\x72\xc7\x73\x9a\xe1\x9d\x2f\x29\x9b\x32\x8f\xca\xc6\x7d\x05\x2b\x0f\xff\xfe\xdc\xb6\xeb\xf5\x83\xbe\x6e\xd7\xeb\x2f\xe4\x78\x48\x05\x8d\x25\x17\x1f\x64\xd9\xa6\x79\xf0\x6c\x63\xbf\x9b\x69\xa7\x97\x3e\xd7\x4f\x01\x4a\xb3\x7a\xe3\x78\xb3\x69\x32\xee\x24\xf7\x45\x34\x9a\xf2\x8d\xa4\x3a\x4d\xcd\x0d\xdf\x7a\x65\xc6\x6f\x15\xae\xc2\xa5\xf9\xe3\xf1\xf1\xf1\x27\xe6\xed\x77\x79\xca\xf7\xfc\x4d\xf1\xd1\x37\x66\x69\xf6\x82\xbf\x91\x63\x5d\xc9\x6b\x0f\x3f\x3c\x6c\x7f\xfd\x13\xba\x76\xd1\x16\xf3\x3b\xec\xc3\x83\xea\x39\x9c\x3f\x34\x90\x1d\x56\x69\x7c\xb2\xb0\x7c\xb6\xf6\x54\x4c\x79\xf8\xdc\xf6\x67\x3c\x26\xac\x30\xb8\x9c\x0f\xbc\x7b\x72\xf0\x07\x03\x60\x4f\x25\x57\x75\x6c\x68\x60\x8f\x2f\x56\xa9\x19\x80\x68\x70\x96\xd1\xf3\xe1\xf4\xa5\x4f\x32\x75\x6b\xfa\xd0\xd2\xab\xac\x6e\x19\xfd\x9b\xde\xe9\x66\x01\xdb\x05\x2c\x6f\x5d\x85\xe4\x32\xf0\xae\x62\x5f\x78\x21\x6f\x34\xb3\xc0\x29\xb6\x80\x7b\x23\x50\xef\x9f\xa8\x7d\x9a\xe6\x9e\x02\xe1\xad\xdb\xb3\xab\xe0\x21\xb0\x31\x02\x0a\x4d\xa4\xbe\xc3\x6a\x00\x2c\xb9\x31\xfc\xcd\xce\x09\xda\x6f\xb7\x65\xaf\x65\x5a\x33\xad\xd4\x75\x99\x26\xc7\x4a\x35\x4c\x1c\x3f\xd3\x6f\x39\x67\x1d\xdc\x9c\xd9\xea\x4d\x4e\x53\xb8\x42\x88\x97\x32\x2f\x25\x60\xe7\x94\x7e\x26\xac\xac\x1b\x31\x1c\xeb\xa7\x6e\x4a\xf7\x6c\x3d\x9a\x65\x03\xdc\x44\x67\x68\x27\x72\x9d\x18\x37\xdd\x59\xd4\x89\x74\x27\x66\x22\x3d\xb9\x5c\x4b\x4e\x2e\x6e\x6a\xb3\x4c\xdd\xc9\x9b\xd0\x18\x99\xf2\x09\x08\x1e\x62\x63\x13\x3f\xe1\x4c\xbd\xd2\x5c\x95\x67\xfd\x26\x6c\xb5\x86\xc3\x5f\xe1\x85\x1b\xd2\x9f\x46\x13\xb0\x1b\x77\xef\x57\x70\xb0\x4d\x46\x34\x07\x13\x7d\x6d\x79\xac\x16\xaa\xd0\xdf\x00\x00\x00\xff\xff\x7d\x6d\xbe\xdd\xd9\x0e\x00\x00")

func templatesBosh_directorTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesBosh_directorTf,
		"templates/bosh_director.tf",
	)
}

func templatesBosh_directorTf() (*asset, error) {
	bytes, err := templatesBosh_directorTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/bosh_director.tf", size: 3801, mode: os.FileMode(480), modTime: time.Unix(1515635419, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_dnsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x95\x3f\xcf\x9b\x30\x10\xc6\x77\x3e\xc5\xc9\xea\x54\x09\xf4\x4a\x9d\x33\x54\xea\xdc\xa5\x63\xf5\x0a\x39\xf6\x85\x20\x19\xdb\xba\x3b\x48\xd3\x57\x7c\xf7\xca\x04\xf2\xa7\x0d\x6a\x18\x22\x31\x24\x4b\x00\x3f\xbe\xe7\xb9\x9f\x11\xd7\x69\xaa\xf5\xd6\x21\x28\x3e\xb2\x60\x53\xda\xd0\xe8\xda\x2b\xf8\xc8\x00\xe4\x18\x11\x36\xa0\x58\xa8\xf6\x95\xca\xfa\x2c\x23\xe4\xd0\x92\x41\x50\x55\x08\x95\xc3\xd2\x7a\x2e\x1b\xed\x75\x85\xb6\xfc\x1d\x3c\x2a\x50\xe8\xbb\xe1\xf1\xe9\x36\x15\xf2\xba\x41\x18\x7f\x1b\x50\x9f\x3e\x3a\x4d\x45\x92\xd5\xb6\xcf\x07\x59\x06\x90\xb6\x4c\xc2\xb3\xe8\x26\x55\x5f\x0c\x3a\x64\x43\x75\x94\x3a\xf8\xa4\xfb\xf6\xfd\x07\xa4\x12\xb0\x0b\x04\xb2\x47\xb8\xa9\x0e\xe8\xbb\x9a\x82\x6f\xd0\xcb\xd0\x40\x68\x25\xb6\xf2\x57\xbb\x43\x5c\x46\xea\x90\xf8\x94\xb8\xd3\xae\xc5\x53\x8c\x99\x46\x8b\xeb\x36\x8b\x14\x7c\xaa\xd0\xcf\x93\x22\x34\x81\x6c\xc9\x28\x0a\xd4\xa1\x76\xd6\x68\xb2\xb9\xf5\xfc\x0f\xa7\x0d\xa8\xcf\xc5\x83\xe6\x13\xb9\xfe\x84\x27\xa2\xb7\x5c\x0e\x74\x7e\x4e\xe6\x26\x34\xb1\x15\x2c\x2b\x17\xb6\xda\x95\xda\x5a\x42\xe6\xc2\xec\xf2\xf1\x52\xbd\x4f\x07\x7e\xf6\xff\x9a\xca\x89\xb8\xcb\xc9\x7d\x79\x7b\xcb\x32\x80\xeb\x24\x0b\x19\xf5\x2a\x15\x20\xb2\x5a\x34\x0f\x01\xcf\x9b\xff\x1b\xb1\x18\xff\x7b\xf5\xfe\x18\x60\xb3\xcb\x99\xf7\x79\xa4\xf0\xeb\x78\x0f\x30\xf3\xfe\x09\x88\xaf\x82\x5f\xdc\x57\x43\xf7\x5e\xba\xc5\x60\xc5\xc4\xb9\x97\x56\x4c\x7c\x2e\xd3\xe4\x4d\xa1\x15\xa4\x55\x42\xbd\xc4\x5b\x4c\xd5\x86\x18\x1d\xd2\x1c\xd9\x71\xf9\xb9\x74\x0f\x2b\xfa\x10\xdc\xc4\x5a\x4c\xd3\x85\xaa\x22\xac\xb4\x84\x59\xa2\x57\x92\x17\xd5\x85\x33\xeb\xc0\xf3\x63\xeb\xc0\x2f\x9c\x0f\x4e\x28\x42\xbb\x6f\xb7\xf7\x30\x8e\x4b\xcf\x24\x39\x9a\xaf\x8f\xe5\xd8\xfa\x0d\xcd\x3f\x01\x00\x00\xff\xff\xc5\xca\x78\x60\xaa\x0a\x00\x00")

func templatesCf_dnsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_dnsTf,
		"templates/cf_dns.tf",
	)
}

func templatesCf_dnsTf() (*asset, error) {
	bytes, err := templatesCf_dnsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_dns.tf", size: 2730, mode: os.FileMode(480), modTime: time.Unix(1515632091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x59\xcb\x6e\xe3\x36\x14\xdd\xfb\x2b\x08\xa1\xcb\x91\xc7\x76\xd2\xa9\xbb\x98\x55\xd1\xed\xb4\x8b\xee\x06\x01\x41\x51\x57\x36\x61\x46\x64\x49\xca\x1e\x63\x90\x7f\x2f\x48\x4a\xb6\x9e\xd4\x23\x09\x8a\xcc\x2c\xa2\x88\xbc\xe7\x92\xe7\x3e\x78\xc4\x9c\x89\x62\x24\xe1\x80\x22\xad\x39\xa6\xa0\x0c\xcb\x18\x25\x06\x22\xf4\x73\x85\x90\xb9\x4a\x40\x5f\x51\xa4\x8d\x62\xf9\x21\x5a\xbd\xac\x56\x83\x16\x58\x2a\x76\xb6\x3f\x4f\x70\x1d\xb4\x16\x85\x91\x85\x41\x91\x12\x85\x01\x85\x13\x42\x4f\x90\xa7\x58\x83\x3a\x33\x5a\x3a\x3d\x13\x5e\x38\xbb\x5f\x7e\x1e\x84\x38\x70\xc0\x54\x3c\xcb\xc2\x40\x7b\xfa\xda\xa3\xc4\x3c\x89\xcb\x91\xb8\x1a\xc9\xc9\x33\xbc\xf4\x79\xe4\x09\x66\x72\xcc\xcf\x81\x8b\x84\x70\x4c\xd2\x54\x81\xd6\x6b\x9a\xc5\xd5\x63\xf9\xb3\x09\xad\xf5\x11\x4b\x25\x7e\x5c\xa7\xa1\xd7\x60\xb5\x3e\xc6\xce\xb2\x1f\xd8\x50\x89\xe7\xac\xbb\x86\x6c\xa8\x8c\xbd\x69\x3f\xf4\x45\xcf\x86\xbc\x0c\x6c\x9f\x2a\x48\x8f\x45\x32\x13\xcf\x1b\x35\x11\x15\x68\x51\x28\x0a\x28\x6a\x59\x65\x4c\xc1\x85\x70\x1e\xa1\xa8\x7a\x8c\x69\xe6\x7d\xd9\x50\x23\xff\xcf\x39\x3c\x13\xb5\x86\xfc\x8c\x59\xfa\x12\xd3\x2c\x16\x12\xf2\x68\x85\x50\x0a\x12\xf2\x54\x63\x91\xa3\xaf\xe8\x7b\xdb\x41\x0e\xe6\x22\xd4\x69\x9d\x24\x3c\x2e\x9f\xa3\x27\x0b\xee\x9f\x6f\xe0\xe3\x66\x55\xea\xad\x10\x22\x9c\x8b\x8b\x5b\x23\x42\x52\x09\x23\xa8\xe0\x16\xc6\x50\x19\xf9\x97\x42\x19\xed\xb1\xbf\x47\xfb\x4d\xf4\x09\x45\x8f\x8f\x0f\xce\xf1\x8b\x05\xf0\x6c\x60\x45\xf2\x03\x68\x37\x69\xb3\x76\xff\x3f\x6f\xa2\x27\x3b\xc1\x10\x75\x00\x83\x0d\x39\xf8\xe1\x57\x57\xcc\x53\x30\x0c\xcd\xba\x88\x50\x74\xaf\x8c\x5a\x2c\x7a\xa2\x10\x8e\x6e\x09\x9b\x09\x75\x21\x2a\x65\xf9\x01\xab\x82\x83\x87\x3f\x1a\x23\xe3\xfb\x48\xec\x47\x26\xc4\xdd\x1a\x5a\x96\x99\xac\xd6\xbb\xb8\xd4\x2b\x9e\xd1\x50\x1a\x94\x61\xb0\x2e\x7d\x23\x58\x57\x2b\xe7\x49\x59\xdf\x1a\x78\x86\x39\xcb\x4f\x0e\xcf\x06\xde\x87\xd5\xe2\xed\x37\xaf\xe3\x47\x2f\x26\x48\xff\x0f\x0c\xe9\x26\x45\x7a\x1a\x47\xb6\x2e\x82\x24\x75\x62\x50\xcb\x9f\xca\x43\x87\x97\x2e\x31\x6e\xbe\x9f\xec\x9a\x86\xa6\x8a\x49\xc3\x5c\xd7\x88\x14\x10\xce\xaf\x88\x20\x2e\x48\x8a\x12\xc2\x49\x4e\x41\xa1\xa4\x30\x88\x33\x6d\x20\x45\x44\x23\x92\x23\x0b\x82\x6e\x20\x85\xe2\xf8\x99\xc8\x41\x6e\xca\xf1\x06\x21\x85\xe2\xb1\x7d\x57\xa7\x64\xe2\xee\x75\x7b\xfb\x3a\xb0\xff\x61\x12\x74\x3f\x0b\x95\xc1\x1c\x2a\x74\x3f\x17\xaf\x26\x04\xa1\x96\x04\x19\x68\x82\xad\x59\x16\xd7\xfe\x5a\xc7\x0a\xf7\xbd\x8e\x36\x8a\x4a\x88\x3b\xa1\x58\x2a\xc8\xd8\x8f\x0e\x97\x3d\x59\x54\x68\x50\x96\x91\x33\x4b\x21\xb5\x5b\x40\xa5\x72\x42\x27\xb8\xa2\xcf\xee\x4d\xcd\x1b\x92\x84\x29\x57\x10\x77\x7d\x75\x77\x13\x10\x61\xce\x77\x1d\x68\xc8\xc8\x9f\x56\x9c\x65\x40\xaf\x94\x43\x79\x62\x51\x05\x16\x28\x81\x4c\x28\xc0\x29\x68\xa3\x84\x75\x6c\x54\x01\xee\x80\x0a\x31\x56\x86\xb0\x95\x84\x65\x10\xc3\x67\x45\xd9\xb9\x1d\x6f\x19\x29\xb8\xa9\x0e\xaf\x57\x4a\xc2\xa9\xa5\x74\x04\xc2\xcd\x11\xd3\x23\xd0\x93\x5f\xbf\x2c\x12\xce\x68\xec\x07\xe2\x72\x20\xb8\x05\x6f\xe1\x36\xe1\x1a\x52\x1d\xb3\x12\x04\x42\x99\x5a\x15\xec\x37\xfb\x8d\x7b\xaf\xe0\xdf\x02\xb4\xc1\x92\x98\xa3\xc5\xfe\xec\x6d\xa3\x51\xca\x3b\x8e\xa6\x2c\x7e\xb0\x05\xd8\x33\x7b\x68\x91\x83\x4b\x9c\x28\xe1\x6c\x8c\x43\xcb\xe9\x4d\x8a\xba\xc1\xc7\x90\x73\x5e\xd0\xed\x37\x21\x3d\xb7\x7d\xd8\xac\x77\xdb\xad\xd3\x74\xbb\x9d\x9d\xff\xf0\xeb\x7a\xfb\xbb\x7f\xb1\xfd\xe2\x4c\xeb\x22\x0f\xbd\xa1\xcc\xeb\x7e\xbe\x94\x9e\xa4\x10\x7c\x4c\xc7\xd7\xa6\x36\x3f\x64\xee\x5f\x5e\x83\xa9\xd0\xd0\x8f\x37\xcb\x91\x92\xba\xcf\x9b\x91\x66\x7d\xe0\xc3\x39\x76\x9b\xfd\x71\x3e\x1a\x76\xbb\xdd\xee\x9e\x5f\xa3\x9f\x03\x23\x51\x0b\x9f\x82\x8d\xec\x58\x18\x3a\x5b\x04\xa0\x35\x13\x39\x26\x59\xc6\x72\x66\xdc\x59\xf6\xed\xaf\x6f\x7f\x8e\xc4\xb5\x4f\xfc\x0e\x87\x77\x6c\x1d\x0d\xc1\x3a\x2f\xc1\x07\x55\xaa\x85\x71\xf1\xf0\x9a\xba\x1e\xbc\x7f\xfe\xf8\xbb\xa5\xb4\xdf\xee\x66\x60\x79\xd1\xd6\xee\x08\x26\x54\x6d\xb3\xb2\xee\xb6\x93\x4a\xab\x36\xfd\x23\x94\xd5\x76\xb3\x7b\x8c\x1f\x76\xbf\x7d\xd9\x2f\x2f\xae\x0e\xbb\xe1\xea\x6a\x34\xc5\x5e\x76\xc7\x78\x5d\x20\x0e\x02\x51\x9c\x12\xc7\x8e\x3c\x58\x2a\x0e\x3a\xad\x65\x11\x01\xc1\xe6\x62\xa5\x58\x6d\xff\x2e\x86\x2e\xf0\xdd\x40\x76\xc8\xea\x0d\xe7\xa7\x15\x42\xe1\x90\xf6\xf6\xac\x10\xe5\xe3\x8c\xcf\xec\x5a\xb5\x45\x07\xdb\x56\x2d\xdf\xdf\xa2\x79\x4d\xb8\x7c\x5c\xde\xb5\x2e\x7a\xb6\xc6\xb8\x8c\x5c\x4f\xd9\x09\xf3\xf2\x73\x12\xe2\xec\x7c\x9c\x98\x8a\x3d\x92\x7e\x52\x8b\xe9\xcd\xc7\x8b\x2e\x6f\x82\x26\x65\xe3\x6d\xf6\xfc\x5c\xbc\xe8\x70\x0e\xba\x1b\x9e\x37\x48\xbe\xf6\x35\xf5\x22\x3a\x66\xb1\xf1\x0e\x64\xec\x37\xef\xc2\x45\xfb\xca\x7e\x69\x15\x96\x97\xf7\xdd\xbf\xb2\xb4\x80\xed\x61\x39\x06\x5c\xe9\x8a\x1b\x6a\xcd\x76\x86\x24\xf1\xc6\xcd\xb8\x75\x63\xe6\x67\xdd\xf4\x7d\x25\x29\xde\xfd\x63\x70\xff\xf8\xb8\x54\x49\x34\xb8\x9e\xac\x21\x3a\x74\x0c\x71\x31\xab\xf1\xf5\x91\x1c\xd0\x0b\x95\x87\x37\x14\xfd\x43\x2b\x08\x39\x9f\x55\x9b\x25\xdd\xe1\xea\xb4\xe1\x7c\x7d\x7d\xf6\xfc\x09\xec\xbf\x00\x00\x00\xff\xff\x83\x1c\x99\xf1\x90\x1d\x00\x00")

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

	info := bindataFileInfo{name: "templates/cf_lb.tf", size: 7568, mode: os.FileMode(480), modTime: time.Unix(1515632091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConcourse_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x93\xc1\x6e\xdb\x30\x0c\x86\xef\x7a\x0a\x82\xd8\x71\x32\x0a\xaf\x87\x5e\x76\x1a\x76\xed\x76\xd8\x6d\x28\x04\xc5\xa6\x1d\xa1\xaa\x28\x48\x72\x8c\xa1\xf0\xbb\x0f\xb2\x1d\xc7\x5d\xb3\x24\x40\x10\xa0\x27\x13\x34\xf5\x93\xfa\x7e\x8a\xbb\xe4\xbb\x04\x58\xb1\xab\xb8\x0b\x91\x54\xd2\xa1\xa5\xa4\x3c\xb3\x45\x78\x15\x00\x3b\x6d\x3b\x82\xaf\x80\x9f\x5e\x5b\xe6\xd6\x92\xaa\xf8\xc5\x77\xe9\x4d\x69\x31\xc5\x72\x8c\x9d\x7e\xa1\x01\xc5\x20\xc4\x7b\x79\xbb\x51\xc6\x9f\x13\xd6\x75\x1d\x28\xc6\x62\x39\x26\xf7\x99\xf9\x3b\xa9\x07\x8a\xdc\x85\x8a\x00\xff\x39\xdf\x98\x40\xbd\xb6\x16\x01\xf7\xa1\x5c\xb4\xa6\xe6\x79\x46\x00\x98\xda\xef\x74\x28\xc8\xed\x94\xa9\x87\x43\x9d\x64\x4f\x0e\x73\x29\xa5\x9e\xc3\xf3\xd1\x49\xe7\x7f\xc5\x66\x63\xe5\x3e\x9e\xaf\x2f\x00\xb4\xb5\xdc\x8f\xed\x00\x7c\xe0\xc4\x15\xdb\x2c\x93\x2a\x8f\x53\x92\x43\x8a\xd3\x18\xbf\xf1\xe1\x0e\x3f\x03\xde\xdf\x7f\xc9\x9f\xb2\x2c\x4b\x7c\x12\x00\x43\x16\x9a\x49\x27\xdd\xc6\xb1\xf4\x70\x99\xa7\x93\x20\x66\x5c\xb8\x72\x40\x2e\xb9\x05\xc3\xff\x19\x9c\xc6\xfc\x66\x55\x70\xb5\x01\x17\x6a\x0b\x80\x48\x31\x1a\x76\x4a\x37\x8d\x71\x26\xfd\xc9\xf5\x8f\x3f\x1e\xbf\x9f\xf1\x97\x43\xaf\x43\x6d\x5c\xab\x42\x67\x09\x01\x63\xdc\xca\x43\x56\x4e\xd9\xb5\xcf\x67\xbc\x8e\x71\x8b\x0b\xe7\x55\xf5\x85\x1b\x1f\xc9\x36\xca\x1a\xf7\x3c\x64\x95\xec\xaa\x0a\xda\xb5\x34\xaa\x8c\x56\x0a\x00\xe3\xd5\x7a\x09\x7e\x7d\xfb\x39\x67\x67\x47\x8e\xb7\xbc\xfa\x2d\xbc\x63\xb5\x4d\xc9\xc7\xab\x68\x8d\x0a\x37\xe3\x95\x5f\xc0\x07\xc3\x75\x35\xad\x9b\xc1\x7a\xb8\xbb\x35\xab\xbf\x01\x00\x00\xff\xff\x5f\xcd\x56\x41\x23\x06\x00\x00")

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

	info := bindataFileInfo{name: "templates/concourse_lb.tf", size: 1571, mode: os.FileMode(480), modTime: time.Unix(1515632091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJumpboxTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8f\x41\x0a\x83\x30\x10\x45\xf7\x39\x45\x08\xdd\xaa\x20\x64\x23\xf4\x2c\x21\x35\x83\xb5\x44\x27\x4c\x66\x44\x10\xef\x5e\x4a\x6d\xad\xd0\x4d\xbb\x1d\xe6\xbd\xff\x3f\x41\x46\xa1\x16\xb4\xe9\x10\xbb\x08\xae\xc5\x21\x09\x83\xf3\x21\x10\xe4\x6c\xb4\xb9\xc9\x90\x2e\x38\x17\x7d\x32\x7a\x51\x5a\x8f\x7e\x00\x7d\xd6\xe6\xb4\x4c\x9e\x4a\x18\x27\xd7\x87\xb5\xf8\xf8\x52\xab\x52\x28\x9c\x84\xdf\xb0\x13\x8a\x4f\x7a\xf2\x51\x36\xfc\x7b\x62\xb9\x9b\xca\xed\xb4\x36\x75\x7d\xb0\xc2\xcc\x40\xa3\x8f\xee\xd5\xe9\x2f\xeb\x41\x19\x7a\x82\x96\x91\xf6\xe1\x07\xef\x95\x39\xe5\xa6\xaa\x7e\x6b\x6d\xad\xb5\x8f\x94\x7b\x00\x00\x00\xff\xff\x10\xf1\x6b\x1b\x66\x01\x00\x00")

func templatesJumpboxTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesJumpboxTf,
		"templates/jumpbox.tf",
	)
}

func templatesJumpboxTf() (*asset, error) {
	bytes, err := templatesJumpboxTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jumpbox.tf", size: 358, mode: os.FileMode(480), modTime: time.Unix(1515632091, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x41\xaa\x03\x21\x0c\x06\xe0\xbd\xa7\x08\x61\x16\xef\x6d\x7a\x83\x9e\xa5\xd8\x31\x95\x14\x31\x92\x11\xa1\x15\xef\x5e\x9c\x19\x98\xe9\xa2\xa0\x4b\xf3\x85\x3f\xfc\xc5\x2a\xdb\x7b\x20\xc0\xa4\xf2\xa4\x39\xdf\xd8\x21\x54\x03\x90\x5f\x89\xe0\x0a\xb8\x64\xe5\xe8\xd1\x34\x63\x0e\xac\xe4\x59\xe2\x00\x7c\x4b\xa4\x01\x46\xb1\x8c\x05\xcf\x4a\x8e\x62\x66\x1b\x96\x9f\x3a\xa9\x14\x76\xa4\x80\x5e\xc4\x87\x3d\xff\xb4\xd9\xfd\x54\x1f\x1c\xe8\x0f\xa7\x5a\xac\x5e\x4e\xc3\x86\xff\x0d\x0d\xc0\xde\x07\xf4\xb7\xfa\xee\x8e\x92\x56\xb3\xd5\x00\xdf\x66\xfb\x6c\xfd\x94\x4f\x00\x00\x00\xff\xff\xeb\xf1\x4c\x79\x5e\x01\x00\x00")

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

	info := bindataFileInfo{name: "templates/vars.tf", size: 350, mode: os.FileMode(480), modTime: time.Unix(1515632091, 0)}
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
	"templates/bosh_director.tf": templatesBosh_directorTf,
	"templates/cf_dns.tf": templatesCf_dnsTf,
	"templates/cf_lb.tf": templatesCf_lbTf,
	"templates/concourse_lb.tf": templatesConcourse_lbTf,
	"templates/jumpbox.tf": templatesJumpboxTf,
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
		"bosh_director.tf": &bintree{templatesBosh_directorTf, map[string]*bintree{}},
		"cf_dns.tf": &bintree{templatesCf_dnsTf, map[string]*bintree{}},
		"cf_lb.tf": &bintree{templatesCf_lbTf, map[string]*bintree{}},
		"concourse_lb.tf": &bintree{templatesConcourse_lbTf, map[string]*bintree{}},
		"jumpbox.tf": &bintree{templatesJumpboxTf, map[string]*bintree{}},
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

