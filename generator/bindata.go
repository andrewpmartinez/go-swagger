package generator

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _templates_model_gotmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x51\xcd\x4e\xf4\x30\x0c\xbc\xe7\x29\xac\x4a\x9f\xf4\x71\xa0\x7b\xe7\x0a\x1c\xb8\x20\x24\xf6\x01\xd6\xb4\x6e\x1a\x68\x7e\x48\x5c\x56\x55\x94\x77\x27\x49\x5b\x56\xcb\x85\x5b\x6c\xcf\xd8\x33\x93\x18\x7b\x1a\x94\x21\x68\xb4\xed\x69\x72\xde\x3a\xf2\xbc\x34\x29\x89\x18\xd5\x00\xed\x83\xed\x5e\xd9\x2b\x23\x53\x8a\xf1\xba\x22\xd3\x57\x58\xfb\xb2\xb1\x9e\x51\x53\x4a\x50\x70\xc8\x78\x5c\x5c\xa9\x4e\xef\xc1\x9a\xbb\xa6\xc0\xd0\xa3\x5e\x31\xcd\x49\xec\x7c\xe1\xb0\xfb\x40\x49\x50\x11\xf5\x59\xba\x87\x03\x1c\x47\x15\x60\x50\x13\xc1\x19\x03\x48\x32\xe4\x91\xa9\x87\xb7\x05\x78\x24\x08\x67\x94\x92\x3c\xb0\xb5\x53\x5b\xf0\x8f\xbd\xe2\x2c\x2d\x0f\x77\x9e\x56\x72\x64\xc8\xa6\xbe\x08\x86\x99\xeb\xaa\x91\x0c\x2c\x76\x06\x4f\xb7\x7e\x36\x57\x9b\xf6\x13\xd0\x59\xad\xd1\xf4\x62\xcb\xe0\x49\x3b\xeb\x39\xa4\xa4\xd6\x07\xfc\x17\x90\xe5\x7a\x34\x59\xf6\x65\x5a\x9b\xf9\x9a\x32\x3c\x40\xf3\xef\xb3\x81\x76\x6b\xae\x4e\x6f\x2e\x9e\xff\xc8\x76\x07\x72\x8e\xb0\xe4\x72\x3f\x61\x08\x5b\xba\x81\xfd\xdc\x31\x44\xf1\x23\x60\x8b\x5f\x51\xa8\x4c\x26\xed\xa6\xe2\xe2\xd7\x97\x56\x35\xbb\x82\x24\xbe\x03\x00\x00\xff\xff\xfc\x8d\x3d\xfc\xfa\x01\x00\x00")

func templates_model_gotmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_model_gotmpl,
		"templates/model.gotmpl",
	)
}

func templates_model_gotmpl() (*asset, error) {
	bytes, err := templates_model_gotmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/model.gotmpl", size: 506, mode: os.FileMode(420), modTime: time.Unix(1424115031, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_modelvalidator_gotmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x57\xdd\x6f\xdb\x36\x10\x7f\xf7\x5f\x71\x13\x36\xc0\x4e\x1b\x65\x03\x86\x3d\xac\xeb\x80\x22\xf3\xd0\x00\x5d\x1a\x24\xdd\x9e\xcb\xc8\x67\x99\x1d\x45\x3a\x24\x95\x8f\x19\xfa\xdf\x77\x24\x45\xcb\xfa\xb0\x6b\x25\x18\xb0\x27\x8b\x77\xe4\xdd\xef\x7e\xf7\x41\x7a\xb3\x59\xe0\x92\x4b\x84\x64\xad\x79\xc1\x2d\xbf\xc7\x7b\x26\xf8\x82\x59\xa5\x93\xaa\x9a\x6c\x36\x7c\x09\xe9\x35\xde\x95\x5c\xe3\x82\x04\xb4\x44\xad\xe1\xe7\xb7\x50\xef\xc3\xad\x76\xba\xd9\xa4\x57\xcc\xae\xaa\xea\x35\x24\xf4\xfd\x41\x65\xcc\x72\x25\xab\x2a\x79\x0d\xb4\xfe\x8b\x89\x12\xe7\x8f\x6b\x8d\xc6\x78\xf1\xec\x8d\xb7\xf5\xcd\x5b\x90\x5c\xc0\x66\x02\xa0\xd1\x96\x5a\x3a\xe9\xc4\xf9\x46\xb9\xd8\x62\xf8\x83\xcb\x0f\x28\x73\x67\x7e\x08\xc4\x56\x3d\x1a\x85\x97\xee\x58\x1f\x87\x8a\x3d\x1e\x44\x15\xd5\xcf\x44\xd5\x58\x1f\x85\x8a\x3c\x59\xd4\x72\x18\x53\xad\x7c\x06\xa2\xcf\xe1\x48\x30\xfd\x79\x6c\xf6\x78\x51\x16\x7b\x73\xe7\x94\x07\x11\x2d\x85\x62\xf6\xa7\x1f\xa7\x83\x75\x14\x53\x18\x5c\xf8\xd5\xfc\x31\x13\xa5\xa1\x72\xde\x8a\xc7\xe6\xf5\x00\xde\xa0\x7c\x29\xde\xe8\xa2\x83\x37\x8a\xc7\xe1\x2d\x85\xe5\x6b\x81\x1f\x97\x7b\x20\x6f\xf5\x2f\x45\xbd\xe3\x68\x14\xc2\xb9\xdc\x47\xa7\xd3\x3c\xaf\x3f\x82\xcd\xa3\x61\xc4\xdf\x38\xf2\xb2\xd2\x58\x55\x2c\x95\x2e\x98\x6d\x4d\xbd\x01\x90\xbf\xfb\x5d\x5f\xa1\xcf\x09\xc2\x46\xbf\x34\x56\x73\x99\x0f\x93\x79\x34\xe8\x08\xd6\x08\x9e\x0d\xcd\xe6\x4b\xc4\x85\xb9\xe1\xff\xa0\x97\x10\x36\xcd\x8a\x4b\x56\xd0\xd2\x09\x5d\x0c\x5c\xba\x94\x0a\x94\x7b\x90\xf4\x5b\xf5\xc2\x62\x61\xf6\xf6\xaa\xd7\x7e\x2d\x61\x1d\x1c\xb1\x43\x6b\xcb\x63\x7b\xf1\x10\xa0\x5a\xfb\x2c\x40\x5b\xcb\xa3\x00\xfd\x29\xf9\x5d\x89\x07\x30\xed\x6c\xf8\x6f\x2f\xc5\xff\x41\x53\x39\x18\x37\x54\xe6\x02\x6f\xb2\x15\x16\xec\xc6\xd5\x29\x90\xea\xec\x0c\x8c\x97\x83\xf1\x8a\x41\x8f\x13\xea\x3e\xe0\x0e\xf9\xf7\x6f\xe8\xf7\x17\xd8\x5b\xa6\xa4\x7e\xf5\x8a\x80\x6c\x36\x9a\xc9\x1c\x21\x8d\xfc\x03\x19\xa6\xcf\xb5\xa0\xb0\xdd\x33\x46\xad\x51\xdb\xa7\xa6\x53\x20\xdd\x69\x7e\xff\x25\x0c\x06\x7c\x52\xd9\x3e\xc6\xab\xda\x42\xa8\x95\x17\xfa\x0b\xfc\xbc\x5b\x2c\xb8\x23\x9e\x89\xc6\xc8\x36\x70\x72\xe9\xa5\x74\xd3\x57\x95\x23\x81\x58\xf0\xdd\x3a\x83\xd3\xb6\xd2\x09\x7e\x70\x3b\x3c\x11\x00\x47\x21\x01\xd8\x89\x99\xc0\xec\x25\x18\x7e\x6d\x7b\xeb\x24\x5d\x69\xd3\x8d\xe3\x52\xd9\x77\x42\xa8\x07\x7a\xfa\x25\x43\x26\x93\x5e\xd9\xcd\x06\xe7\x71\x77\xd4\xa9\xdb\x2f\x98\xb5\x27\x32\x25\x2b\x4c\x6b\x08\x4a\x07\xf5\x37\x66\xd9\xa7\xa7\x75\x7f\xf0\x5d\xa3\xa9\x89\xed\x41\x72\x12\xdf\x20\xd3\x99\x23\xa3\x77\x2e\x7d\xcf\xcc\xdc\x07\x4b\xf4\xf7\x09\x38\x57\xc5\x5a\x19\x6e\xb1\x36\x43\x36\xfd\xee\x69\xdf\x50\xb0\x92\xa6\xe9\x6c\x70\x9c\xf7\xb3\xb5\xad\x96\x0b\x73\x15\xdf\xe2\x55\xd5\xce\x70\xef\x89\xee\x52\xdc\xae\xb5\x0b\x73\xee\x89\x0a\xf7\x10\xbd\xd5\xda\x36\x86\xef\xbc\x41\x33\x4a\x5a\x46\x58\x3b\x06\x3a\xf7\xd0\xf0\x49\xda\x8c\x8f\x1f\x7d\xa6\xda\xa7\xbb\xa9\x6d\x1f\x0f\xbf\x6b\x96\xfd\xcd\xa8\xe3\x3c\xa7\xfe\x93\x84\xae\x02\x3e\xad\xb8\x81\x25\xa7\x6e\x7d\x60\x06\x72\x24\x6c\x64\x74\x01\xb7\x4f\x60\x57\xd4\xc2\x0f\x2c\xcf\x51\x83\x55\x4a\xa4\x6e\xff\xdc\x55\xab\xcc\x49\x19\xcf\x15\x3c\x5f\x59\x20\xee\xef\x11\x96\xa5\xf5\xa6\x56\x28\xe1\x49\x95\x94\xe9\x53\x5d\xca\x96\xa5\xe8\x02\x32\x55\x14\x4c\x2e\x26\x13\x4e\xf9\xd7\x16\xa6\x54\x19\x49\xce\xed\xaa\xbc\x4d\x49\x77\x96\x31\x53\x32\xf1\x85\x17\x67\xb9\x3a\xad\x4f\x9f\xc5\x59\x9c\x1c\xb3\x39\x14\x58\x32\xf1\xe3\x25\x4e\x1c\xef\x2c\xce\x1c\x82\x4d\x37\xf9\x12\x92\xef\xee\x62\x67\x47\xc6\x66\x9e\x9e\x58\xda\xdb\x4b\xc0\x84\xc8\x89\xc7\xf7\x25\xc1\x3f\x17\xcc\x98\x7a\xaa\x2d\x4b\x99\x81\xab\xda\x6b\xcc\x90\x2a\x4a\x07\x39\x9c\x90\x68\x67\xdf\x0c\x9a\x7e\x81\x93\x9d\x7f\x7c\x86\x9e\x81\x75\x7f\xf8\x5e\x93\xf8\x30\xed\xa8\x67\xad\x58\xea\xa9\xca\x31\x86\xe3\x4a\x85\x9a\xad\x69\xa4\xa0\x68\xae\xb2\x1e\xb8\x34\x3a\xe8\x0d\xe9\x69\xff\xb6\xf2\xc8\xdc\xc0\xaa\xbb\x99\xd4\x33\x3f\x0b\x1b\xd6\x9a\xaf\xa6\xcf\xe9\x90\x6f\x57\xf8\x36\x8b\x2c\x38\x2c\x0d\x27\xb0\x7b\x25\xb4\x82\xda\x13\xd2\x21\xaa\x1b\x27\x8e\xea\x03\xe1\xc1\x49\x3d\x80\x1a\xdb\x3e\xca\x23\xe6\x7f\x13\x1b\x51\x33\x30\x7e\xff\x0d\x00\x00\xff\xff\x65\xe5\xfe\x33\x02\x10\x00\x00")

func templates_modelvalidator_gotmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_modelvalidator_gotmpl,
		"templates/modelvalidator.gotmpl",
	)
}

func templates_modelvalidator_gotmpl() (*asset, error) {
	bytes, err := templates_modelvalidator_gotmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/modelvalidator.gotmpl", size: 4098, mode: os.FileMode(420), modTime: time.Unix(1424115031, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_server_operation_gotmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x56\x4d\x6f\xe3\x46\x0c\xbd\xeb\x57\xb0\x06\x5a\x48\x81\x56\xbe\xa7\xd8\x43\x9b\x16\xcd\x1e\xba\x0d\x92\xa0\x3d\x16\x93\x11\x2d\xcd\x56\x9a\x51\xe6\x23\x8e\x6b\xf8\xbf\x97\xf3\x21\xc7\x4a\x14\xd7\xc5\x62\x17\x41\x4f\x96\x34\xe4\x23\xf9\xf8\xc8\xf1\xc0\xf8\x5f\xac\x41\xd8\x6e\xab\xab\xf8\xb8\xdb\x65\xd9\x72\x09\xb7\xad\x30\xb0\x12\x1d\xc2\x9a\x19\x68\x50\xa2\x66\x16\x6b\xb8\xdb\x80\x6d\x11\xcc\x9a\x35\x0d\x6a\xb0\x4a\x75\x95\xb7\xff\xb9\x16\x56\xc8\x86\x0e\x47\xbf\x5e\x34\xad\x85\x41\xab\x07\x84\x95\xb3\x01\xaa\x45\x09\x1b\xe5\x40\xe3\x3b\xed\xe4\x04\x69\x0c\x01\x5c\xf5\x3d\x93\x75\x96\x89\x7e\x50\xda\x42\x9e\x01\x2c\x24\xda\x65\x6b\xed\xb0\xc8\xfc\x5b\x23\x6c\xeb\xee\x2a\xb2\x5c\x72\x66\x1c\xeb\x3e\x89\x7e\xd9\xa8\x77\x09\x6b\x71\x82\xcd\x12\xb5\x56\xda\x9c\x64\x6a\x06\xe4\x27\x19\x6a\xe5\xac\x0f\x4f\xb6\xdb\xad\x66\x92\x88\xad\x3e\x84\x2a\x0c\xd1\xea\x3f\x12\x1f\x42\xda\x15\x2c\xbe\xbd\x5f\x40\x95\x3e\xa2\xac\xe9\xa9\x08\xbc\x53\x23\x2e\x3a\x66\xcc\x47\xd6\x53\x2b\x2e\x89\x88\x8e\xc8\x21\x1f\xd4\x2b\xc6\x89\x49\x45\xa4\xb7\xcc\x02\x67\x12\xda\x70\x0c\x0f\xac\x13\xb5\xf7\xbc\x74\xc4\xdc\x81\x3b\x0c\x4c\xb3\xde\x64\x76\x33\xe0\xbf\x22\x6f\x43\x2e\x62\x05\xd5\x0f\xce\xb6\x4a\x8b\xbf\xb1\x0e\x09\x52\x56\xd1\xda\x37\xb3\x76\x1c\x0d\x30\xea\xa0\x71\x9d\x05\xca\x86\xf2\x08\x5c\x56\xf0\x0b\x5a\x13\x7a\xaa\xf1\xde\xa1\xb1\x29\x3c\x90\x80\xd6\xd8\x75\xfe\x37\x74\x1c\xb9\xd3\xc2\x6e\x02\x17\x5c\x0c\xac\x8b\x61\x6f\x1c\x27\x6c\xf3\xab\xaa\xb1\x0b\x81\x63\xd4\x7c\x9a\xf8\x55\xc0\x2c\xe1\xcc\x4b\x76\x04\xd8\xed\x0a\xc8\x23\xca\x35\x5a\xa7\xa5\xb9\x50\xfd\xd0\xe1\xe3\x6f\x77\x9f\x90\xdb\xdd\xee\x2c\xb1\x4c\x4e\xd3\x30\x65\x4c\xbe\x88\x8d\xe8\xcc\x9e\x9a\x53\xc3\x06\xf7\x83\x36\x3e\xe1\x7c\x2e\x6f\xff\x9d\x95\xaf\x43\xc2\x2b\x35\x87\xa7\xb8\x3c\x3e\xe2\x7a\xea\x08\x5c\x23\xcd\xb6\x27\x40\xe2\x1a\xfc\x24\x57\xa3\x02\xa3\xa2\x71\x56\xbf\x6a\xf0\x3b\x41\x28\x99\xad\x9c\xe4\x2f\x71\x73\x6e\x1f\xe1\xac\x17\x35\x21\xad\x99\xc6\xea\x42\x91\x9e\x1f\x6d\x99\x46\x43\xcf\xab\xbe\x08\x6d\x3c\x0c\xe4\xc5\xaf\x03\x69\xf0\xdd\xf4\x68\x9b\x20\xcf\x81\x62\x95\x89\x7d\x7d\x3e\x06\x08\x25\x47\xd2\x7f\x52\xfc\xc6\x92\x34\x9a\xc0\xf0\xe4\x2d\x92\x33\x33\x86\x60\xac\x76\xdc\x86\xf8\x29\xd0\x5c\x3d\x74\x1a\xb9\x07\x98\xeb\xc8\x5e\x15\xaf\xd4\xeb\x73\x0c\x04\xfa\x7e\x5e\x23\x47\xf1\x80\x3a\x25\xf0\x8c\x89\x02\x6e\x50\x3f\xe0\xe5\xed\xed\x55\xae\x53\xa7\xae\xd1\x0c\x4a\x1a\xfc\x83\xa6\x16\x75\x09\x1a\xce\xd2\xf7\x20\xd7\x22\xb2\xe7\x77\x5f\x09\x7f\xc2\xf9\x7b\x78\x11\x65\xac\xa3\xba\xf6\x56\x1f\xe4\x4a\xe5\xa4\xb6\xd9\x85\xe3\xc2\x56\x08\x8a\x3c\x0e\xb5\x77\xca\x7d\x4a\x1e\xd7\xeb\x97\xe0\xbc\xe7\x37\xef\x41\x8a\x2e\x24\x06\xc7\xd2\x09\x95\xd5\x54\x29\x41\x24\x14\x1a\xef\x38\xad\xe5\x58\x13\x01\x16\x01\x28\x2a\x84\x1e\xbd\xe4\xf7\xdb\x2b\xa6\xe9\x03\xdf\xc3\xd3\x6e\x80\xc5\x7e\xb5\x6e\x77\x8b\xb1\xae\x71\xbc\xe2\x5b\x95\x3f\x5f\x27\xa3\x54\x26\xc3\x95\x6a\x3a\xca\xc6\x8f\x42\xd6\xbf\xfb\x7b\x20\x35\x65\x4f\x4a\x19\x04\xfd\xcc\x2b\x0d\xf2\xf7\x13\xae\xfc\xa2\xbf\x23\x98\xf1\xd6\xf8\x52\xd4\xbd\x76\xcf\xcc\xed\x3b\xed\xb1\x5e\x2d\x3e\xc9\xbb\x7a\x5a\x56\xb3\x75\x96\x4f\xbd\x2a\x7c\x91\x8c\x5b\xba\xbb\xbb\xcd\x78\x7f\x1e\xac\xde\xaf\xa8\xa0\xec\xf3\x61\x89\x9d\xc3\x9d\x7d\x4c\x27\x89\xaa\x93\x38\x7a\x29\x8b\xe3\x8c\xbd\x59\x7e\x28\xff\x62\xf6\x6a\xfe\x92\x6a\x7b\x43\x12\x7b\x53\x0a\xfb\x7f\xc9\x6a\xff\x57\x6a\xfc\xdf\xf3\x4f\x00\x00\x00\xff\xff\x3d\xa5\x5c\x34\x4a\x0d\x00\x00")

func templates_server_operation_gotmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_server_operation_gotmpl,
		"templates/server/operation.gotmpl",
	)
}

func templates_server_operation_gotmpl() (*asset, error) {
	bytes, err := templates_server_operation_gotmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/server/operation.gotmpl", size: 3402, mode: os.FileMode(420), modTime: time.Unix(1424125862, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templates_server_parameter_gotmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x57\x4b\x6f\xdb\x46\x10\xbe\xeb\x57\x4c\x85\xb6\xa0\x0c\x85\xbe\x14\x3d\xb4\xc8\xa1\x71\x5d\x44\x40\x92\xa6\xce\xe3\x52\x14\xc8\x8a\x1c\x51\xdb\x92\x4b\x7a\x77\x19\xd9\x15\xf4\xdf\x3b\x3b\xcb\x97\x48\x4a\x96\x9c\xf8\xd0\x83\x61\x71\x76\x76\xe6\x9b\xf7\xec\x76\x1b\xe3\x4a\x2a\x84\x69\xa1\x65\x26\xad\xfc\x8c\x9f\x45\x2a\x63\x61\x73\x3d\xdd\xed\x26\xdb\xad\x5c\x41\xf8\x5a\xaa\x57\xa8\x12\xbb\x26\x0a\x7d\xa3\xd6\xf0\xd3\x73\xa8\x18\xb1\x3d\x0e\xb6\xdb\xf0\xad\x70\x6c\x73\x98\xd2\xef\x57\x79\x24\xac\xcc\xd5\x6e\x37\x9d\x03\x7d\x7f\x14\x69\x89\xd7\x77\x85\x46\x63\x98\xcc\xd4\x8e\xf4\xd9\xcf\x2c\xfc\x9b\xe7\xa0\x64\x0a\xdb\x09\x80\x46\x5b\x6a\xe5\xa8\x13\x87\x06\x55\xdc\xa2\x12\x77\x47\x51\xd5\xc7\x8f\x44\xd5\x4a\x3f\x0b\x15\x69\xb2\xa8\xd5\x38\xa6\xea\xf0\x11\x88\x3e\xf9\x2b\x5e\xf4\xa7\xf3\xfc\x24\x95\xcc\xca\xec\x60\xec\xdc\xe1\x51\x44\xab\x34\x17\xf6\xc7\x1f\x82\x31\x64\xb3\x3a\x84\x5e\x05\x7f\x5d\xdf\x45\x69\x69\x28\x95\x1a\xf2\xb9\x71\x3d\x82\xd7\x1f\x7e\x29\xde\x5a\x45\x0f\x6f\x4d\x3e\x0f\x6f\x99\x5a\x59\xa4\xf8\xfb\xea\x00\xe4\xe6\xfc\x4b\x51\x77\x14\x9d\x85\xf0\x5a\x1d\x72\xa7\x3b\x79\x5c\x7d\x78\x99\x27\xc3\xe0\xff\x6d\xb7\x89\x4a\x63\xf3\x6c\x95\xeb\x4c\xd8\xbd\x86\x33\x82\xf1\x37\xe6\x7a\xc0\x7b\x8e\xe0\x19\xf9\xd3\x58\x2d\x55\x32\xee\xcb\xd3\x30\xb7\x58\x4d\x2a\xa3\xb1\xae\xf8\x06\x31\x36\xef\xe4\xbf\xc8\x14\x82\xa6\x45\xf6\x46\x64\xf4\xe9\x88\xce\x04\xa9\x5c\x40\x53\x54\x07\x80\x0c\x0b\x75\x61\x31\x33\x07\x2b\x95\x4f\x1f\x0a\x57\x0f\x47\x5d\x9f\x95\xe4\x73\x2b\xf1\x18\xa0\xea\xf4\x51\x80\x1a\xc9\x67\x01\xfa\xa0\xe4\x6d\x89\x47\x30\x75\x18\xce\x4e\xeb\xff\x79\x49\x15\x3a\x2f\x50\xdb\xfb\x91\x4c\x5d\x98\xb7\xf5\x74\x77\x37\xc8\x3b\x45\x4a\x50\x47\x87\x3e\x84\x8e\xa5\x6b\xea\xc2\x5c\x71\xb5\xfa\xf2\xa2\x09\xb4\x2f\x63\xbc\x94\x47\xc5\xe4\xca\x0a\xc2\xda\x13\xd0\xab\xaf\xfd\x9b\xfc\xbf\x10\xd1\x3f\x22\x41\x9f\x4c\xfc\x93\xce\x26\x97\x97\xf0\x7e\x2d\x0d\xac\x64\x8a\xb0\x11\x06\x12\x24\xd9\x24\x32\x86\xe5\x3d\xd8\x35\x82\xd9\x88\x24\x41\x0d\x36\xcf\xd3\xd0\xf1\x5f\xc7\x64\xad\x4a\xe8\xb0\xbe\x97\xc9\x64\x6d\x81\x7c\xf7\x19\x61\x55\x5a\x16\xb5\x46\x05\xf7\x79\x49\x7e\x7f\xa6\x4b\xb5\x27\xa9\x56\x01\x51\x9e\x65\x42\xc5\x93\x89\xcc\x8a\x5c\x5b\x08\x28\x4e\x53\x85\xf6\x72\x6d\x6d\x31\x75\x1f\x89\xb4\xeb\x72\x19\x12\xe3\x65\x24\x4c\x29\xd2\xbf\x65\x76\x99\xe4\xcf\x2a\x51\xa7\xf0\x5c\x52\xc8\x73\x6d\x4e\x62\x75\x7a\x1d\xfe\xd3\xb8\x75\x5e\xda\x13\x31\xd4\x89\x3d\x9d\x10\xf7\x76\xab\x85\xa2\x48\x84\x0b\x36\xdb\x55\xa1\x23\x92\x03\xa9\xd7\xad\x60\xfa\xdd\x2d\xc7\x8f\x89\x3e\x84\x33\x0e\x14\x45\xee\x2a\x15\xc6\xf8\x36\xc0\x1d\xc1\x90\x0f\x39\x1f\x0c\x88\x34\x65\x2f\x2f\xf3\x52\xc5\x50\xf8\x53\xca\x28\x26\xd2\xd5\x97\x25\xf9\xba\x73\x1f\x5c\xa6\x73\x41\x39\xd9\xf6\xbe\x90\x11\x89\xe0\x98\x1b\x04\xa1\x11\xf2\x25\x67\x5a\x0c\x2b\x9d\x67\x20\xc0\x79\x27\xbc\x41\xea\x0d\xc6\x4e\xe8\x02\x8e\x23\xa2\xb1\x51\x46\x96\x8b\xae\xb1\xd4\x1f\xb9\x9c\x74\x59\xfc\x2b\x9a\x48\xcb\xc2\x17\xb3\x37\x6c\x8f\x54\x9b\x0d\x9c\xac\x55\x49\x56\xa8\x89\xf2\x9e\x54\xef\xb9\xc7\xe7\xf1\x0b\xa9\xe2\x0a\x1d\x39\xc1\xae\x61\x49\x04\xf2\x0b\x79\xa3\x76\x3f\x7d\x51\x42\x32\xcb\x1c\xa4\x05\x82\x5e\x66\x44\xb5\x6b\x61\x5d\x36\xd2\x5e\x70\xe7\xf2\x5a\x25\x06\xa4\xfb\xca\x50\x11\x17\x7c\xf4\xf7\xc5\x32\xc5\x60\x06\x17\x4d\x9b\xba\x41\x43\xeb\x84\x1b\x51\xa8\x57\x22\x42\x07\xc3\xb9\xdc\xf0\x65\xa7\x96\x54\x39\x45\x1b\x49\xd1\x29\xc9\xaf\xe4\x1c\xc1\xd5\x92\xa1\x5d\xe7\x31\x38\x9f\x9b\xc9\xaa\x54\x11\xb8\x7e\x77\x83\x11\x52\x27\xd1\x95\xb1\x17\x63\x0e\x9e\x75\x2d\x0d\x34\x5c\x74\xe3\x32\x07\xce\x4a\xb8\xf0\xc9\x49\x33\xc2\x46\x6b\x8c\x6f\xdc\xd7\xdc\x25\x8b\x33\x58\xd7\xa5\x18\x5e\x55\x84\x11\xa3\x7c\xd3\x34\xae\x2d\x2b\xdc\x04\xbd\xe3\x19\xbb\xdf\x05\xf3\xa5\x30\x7f\x94\xa8\xef\xeb\x10\xdf\xf2\x95\xca\xcb\xe1\x87\x9b\x57\x21\x1f\x07\xb3\x4e\xc0\xc6\x52\xa3\x11\xd8\x6b\xb4\x9e\xd4\xaa\xd8\xed\x08\x55\xf8\x1a\x75\x82\x43\x8f\x85\x2e\xe6\x83\xa4\x09\x9a\xba\x26\x76\x11\xbf\xa3\xf8\xa6\xc8\xe3\x23\xb8\x35\xd5\x78\xe5\x37\xca\xac\x02\x99\x52\xa4\x2a\x28\x74\xf0\x04\x6a\x39\x3a\x95\xe9\xc7\x01\xbc\xa4\x9b\xa8\x9f\x02\x42\xe8\x45\x8f\xa9\x6f\xaa\xaf\x03\x64\x6f\xf0\x7c\xfd\xa0\x5c\xe5\x69\x8a\x91\xab\xff\x91\xc0\xf8\x05\xa0\x65\x69\x96\xd4\xa7\x8d\x58\x1f\xd3\x81\xa8\x9d\x83\xee\xab\x86\x73\x80\x6f\x18\xd2\x87\xb1\xf9\x05\xa8\x53\x97\x1e\xe8\x8b\x3c\xae\x43\x4b\xd4\x76\x45\xab\x5b\x48\xdd\x3a\x48\xa9\x63\x9d\xc3\xf7\x43\x1b\x06\xf8\x87\x6b\x19\xf7\x98\xf0\x97\x38\xbe\xe6\x29\x1d\xf8\x61\x4d\xcf\x82\x0d\x29\x37\xc8\xd4\x60\xba\xbf\x00\x4f\xc7\x9e\x2e\xf4\x47\x77\xd9\xa8\x1d\xb0\xc3\x5b\xf1\x87\x7c\x3c\xc0\x17\x56\x8d\x9e\xba\x3c\x0b\xea\x57\x43\xed\xa4\x6a\x99\x24\xd9\x6e\xf6\xd0\xe8\xfe\x36\xaa\xbb\xb4\xf3\x51\xdb\xb3\x81\xf7\xaf\x7e\xa7\x63\x1f\xd3\x98\xe8\xb6\xba\xd1\x15\xf1\xd8\x60\x68\x55\xb6\x83\x61\x3c\x69\xb4\xd8\x54\x6f\xb9\x47\xb4\xf9\x26\x25\xdc\x80\x91\x1a\xc9\x03\x23\xfb\x7a\x7d\xf8\xae\x79\x32\x1e\xdc\xdc\x09\xcd\x49\x69\x30\xab\xc8\x8d\xa7\x7d\x3c\x1c\xdf\x42\x45\x41\xbf\x53\x39\x88\x94\x92\xe4\x22\xf6\x1c\x8f\xde\x79\x8d\x93\x4b\xa0\x39\x73\x0e\x99\xb5\x49\xfd\x70\x36\x2e\x14\x1b\xea\xf6\x8e\x07\xdf\xcf\x7e\x39\xa9\xec\x3c\xc9\x86\xe1\x43\x06\x9e\xfb\xd5\xa1\x69\x1e\xf5\x2e\x34\xc6\x49\x7a\x46\x5c\x41\x53\xb9\xca\x65\xe2\x33\xdd\x98\x0d\xcb\xa0\x8e\xe2\x30\x73\x4e\x0e\x14\x57\x05\xe9\x60\xcb\xbc\x66\x5a\x94\xda\x0b\xb5\xb5\xfb\x15\xb5\x57\x46\x87\x86\xcd\xd7\xab\x81\x3f\xff\xfa\x92\x2a\xe8\x9a\xd1\x07\xde\x79\xa1\xf5\x1d\x7f\x2e\xfc\x23\xc1\x80\x8b\x2a\x1f\x5b\x05\xd5\x96\xdd\x7d\x8c\xf6\x1f\xb0\xfc\x94\x68\x31\x53\x20\x87\x0f\xe0\xc9\x7f\x01\x00\x00\xff\xff\xab\xcb\x6e\x98\xc3\x16\x00\x00")

func templates_server_parameter_gotmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_server_parameter_gotmpl,
		"templates/server/parameter.gotmpl",
	)
}

func templates_server_parameter_gotmpl() (*asset, error) {
	bytes, err := templates_server_parameter_gotmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/server/parameter.gotmpl", size: 5827, mode: os.FileMode(420), modTime: time.Unix(1424121596, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	"templates/model.gotmpl": templates_model_gotmpl,
	"templates/modelvalidator.gotmpl": templates_modelvalidator_gotmpl,
	"templates/server/operation.gotmpl": templates_server_operation_gotmpl,
	"templates/server/parameter.gotmpl": templates_server_parameter_gotmpl,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"model.gotmpl": &_bintree_t{templates_model_gotmpl, map[string]*_bintree_t{
		}},
		"modelvalidator.gotmpl": &_bintree_t{templates_modelvalidator_gotmpl, map[string]*_bintree_t{
		}},
		"server": &_bintree_t{nil, map[string]*_bintree_t{
			"operation.gotmpl": &_bintree_t{templates_server_operation_gotmpl, map[string]*_bintree_t{
			}},
			"parameter.gotmpl": &_bintree_t{templates_server_parameter_gotmpl, map[string]*_bintree_t{
			}},
		}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

