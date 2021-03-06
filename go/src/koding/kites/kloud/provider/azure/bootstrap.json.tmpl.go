// Code generated by go-bindata.
// sources:
// bootstrap.json.tmpl
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

var _bootstrapJsonTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x55\x4d\x6b\xdc\x30\x10\x3d\xdb\xbf\xc2\x88\x9c\x4a\xb2\x5d\x7a\xcc\x2d\xf4\xd0\x96\x42\x28\xdd\xd0\x4b\x29\x42\x6b\x4f\x1c\x11\x59\x12\xfa\xd8\x74\x6b\xf4\xdf\x8b\x24\xcb\xb6\xbc\x0e\xbb\xcb\x9e\x2c\xbd\x79\xf3\xe6\xcd\x68\xb6\x2f\x0b\x24\x95\x38\xd0\x06\x14\xba\xaf\xfa\xb2\x28\x10\xf9\x67\x15\x0c\x1f\x05\x92\x76\xcf\xa8\x7e\xc1\x1a\x8c\xa1\xbc\xd5\xe8\xbe\x42\x37\xfd\x81\xa8\x4d\xc0\xe1\xe5\xbd\x43\xb7\x21\x4e\xdb\xbd\xae\x15\x95\x86\x0a\x8e\x69\xb3\x0c\x5b\x5c\x3b\x54\x16\x85\x2b\x0b\x77\x5b\x16\x48\x81\x16\x56\xd5\x90\x09\xc2\x2f\x42\x1b\x68\xb0\x06\x75\xa0\xf5\xa4\xef\x55\x34\x94\xb7\xe9\xb3\x40\x9c\x74\xfe\x12\xf5\xfd\xe6\x6b\x88\xd8\xc5\x80\x47\xd2\x81\x1b\xc4\x15\x88\x89\x9a\xf8\xd4\x4b\x59\xe9\x7c\x04\x82\x7c\x81\x0e\x14\x61\xb8\x16\xdc\x00\x37\xde\x80\x67\xc2\x34\x0c\x80\x06\xc6\x3a\x3c\x59\xcc\x59\x0d\x2a\xab\x5a\x01\xf1\xdf\xfb\x63\xf5\x3d\x0a\x4d\x02\xc8\x1e\xd8\xa0\xf3\x09\x48\xb7\x63\xb6\x75\x6e\x13\xab\xd9\xd4\xa2\xf3\x86\x78\x47\x0a\x77\xdb\xf7\xf4\xb9\xda\xec\x8c\x50\xa4\x85\xac\x9c\xc9\x1d\x1d\x6f\xaf\xb1\x67\x8d\xf0\x6a\x7f\x16\xe5\x0f\x9c\xe7\xeb\x27\x75\x2d\x2c\x37\xd8\x1c\xe5\x42\xcf\xd3\x51\x7a\x21\xf3\xf2\x81\x37\x59\xad\x50\x5b\x45\xcd\x11\xb7\x4a\x58\x79\x51\xa9\x43\xc4\x17\x1f\x70\x6d\xa5\x93\x92\x49\xc2\x81\x2a\x63\x09\xc3\x1c\xcc\x9b\x50\xaf\x97\x68\xf8\x15\x43\x1e\x63\x44\x2e\x82\x34\x8d\x02\xad\xb1\x96\x24\xf4\xee\xb7\x0f\x78\x88\x87\x3b\x7f\xe6\x1c\xfa\x73\x6d\x6b\xb4\xdd\x73\x30\xa3\x96\xdc\x90\x70\x97\x89\x98\x54\x48\x05\xcf\xf4\xef\x80\x5c\xa8\x48\xd8\x93\x1e\xa0\x9b\x7e\xad\x3b\x69\xa2\x7d\xee\x68\x65\xf0\x32\x1a\x1a\x07\xfb\xa7\x65\xe0\xe6\xe6\xe6\x0c\x58\x59\x76\x7e\x9e\x75\xab\xee\xe2\xdd\x9d\x01\xd2\x8d\x1e\xe4\x54\x1e\xad\x93\xbf\x6b\x43\x91\x4c\x4e\x63\xf9\x8d\xef\x85\xe5\xcd\x34\xb5\xc9\xf9\x07\xc6\xc4\x5b\x3a\x96\x8a\x0a\xcf\x84\xee\xab\x4f\xdb\x6d\x4a\x1d\x76\x18\x3e\x35\x75\xbb\x09\xbf\x8f\x5b\x94\x23\xa5\x50\x06\x2b\xc2\xdb\x90\xf9\xc3\xec\x85\x19\xca\x43\x6b\x2f\x21\x9b\xc3\xdf\x61\x94\x4a\x18\x51\x8b\xb0\x7e\x9e\x3e\xff\x40\xb3\x7e\x0c\x2f\x2d\xac\x61\x61\x8d\xb4\x69\x7e\x50\xbe\x7e\xf1\xe0\x7b\x6c\xca\x81\x30\x0b\xf3\x19\xc8\xc1\x69\x06\xd2\x9a\x3f\xb3\xd0\x16\xab\xec\x4c\xaa\x05\x7a\x25\xd7\xb8\x3d\xde\xd9\x1b\xa7\x94\xab\x13\x3c\x32\x7a\xaa\xf7\xde\xff\x09\xd7\x02\xb8\x4a\x96\x3f\xd4\x89\x63\xf9\x4e\x87\xbf\xc8\xd2\x95\xe5\xff\x00\x00\x00\xff\xff\x3c\x96\xf0\xec\xb8\x07\x00\x00")

func bootstrapJsonTmplBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapJsonTmpl,
		"bootstrap.json.tmpl",
	)
}

func bootstrapJsonTmpl() (*asset, error) {
	bytes, err := bootstrapJsonTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.json.tmpl", size: 1976, mode: os.FileMode(420), modTime: time.Unix(1470666525, 0)}
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
	"bootstrap.json.tmpl": bootstrapJsonTmpl,
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
	"bootstrap.json.tmpl": {bootstrapJsonTmpl, map[string]*bintree{}},
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
