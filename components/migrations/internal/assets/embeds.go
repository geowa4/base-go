// Code generated by go-bindata. DO NOT EDIT.
// sources:
// sql/0001_first.down.sql
// sql/0001_first.up.sql
// sql/0002_second.down.sql
// sql/0002_second.up.sql
package assets

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

var __0001_firstDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\x48\xcb\xcf\x2f\x56\x70\x76\x0c\x76\x76\x74\x71\xb5\xe6\x42\x92\x48\x4a\x2c\x2a\xb6\xe6\x02\x04\x00\x00\xff\xff\xfb\xf0\xbe\xe3\x2a\x00\x00\x00")

func _0001_firstDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_firstDownSql,
		"0001_first.down.sql",
	)
}

func _0001_firstDownSql() (*asset, error) {
	bytes, err := _0001_firstDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_first.down.sql", size: 42, mode: os.FileMode(420), modTime: time.Unix(1543521617, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0001_firstUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xb1\xca\x83\x30\x14\x46\xf7\xfb\x14\xdf\xa8\xe0\x1b\xfc\x53\xfe\x78\x2d\xa1\x36\x96\x98\xa1\x4e\x25\xc5\x58\x02\xd6\x80\xda\xd2\xc7\x2f\xc1\xa1\x8b\x74\x3d\x1c\xce\xbd\x9f\x34\x2c\x2c\xc3\x8a\xff\x9a\xa1\x2a\xe8\xc6\x82\x2f\xaa\xb5\x2d\x86\x18\x17\x64\x04\x84\x1e\x8b\x9f\x83\x1b\x0b\x02\x26\xf7\xf0\x58\xfd\x7b\x2d\x88\x80\xb3\x51\x27\x61\x3a\x1c\xb9\x43\x16\xfa\x9c\xf2\x3f\xa2\x1f\xcd\x9b\x9b\x77\x9a\x43\x8c\xd7\xd0\x23\x4c\xab\xbf\xfb\x39\x91\x97\x1b\x9f\xfe\x0b\x76\x2e\x25\xab\x6a\x0c\xab\x83\xde\xd8\x16\xc9\x61\xb8\x62\xc3\x5a\xf2\xb6\x20\xb9\x68\x34\x4a\xae\xd9\x32\xa4\x68\xa5\x28\x39\xfd\xf9\x09\x00\x00\xff\xff\x16\xf7\xb2\x1e\xfb\x00\x00\x00")

func _0001_firstUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_firstUpSql,
		"0001_first.up.sql",
	)
}

func _0001_firstUpSql() (*asset, error) {
	bytes, err := _0001_firstUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_first.up.sql", size: 251, mode: os.FileMode(420), modTime: time.Unix(1543521617, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0002_secondDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _0002_secondDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0002_secondDownSql,
		"0002_second.down.sql",
	)
}

func _0002_secondDownSql() (*asset, error) {
	bytes, err := _0002_secondDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0002_second.down.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1543521617, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0002_secondUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\x41\x6f\xab\x30\x10\x84\xef\xfe\x15\x73\x88\x94\xf0\xf4\x90\x5e\xce\x9c\x1c\xb2\x41\x48\x3c\x53\x19\x88\xd4\x13\xa2\xf1\x92\x5a\xa2\x26\xb5\x69\xfb\xf7\x2b\x48\xd2\x43\xd5\x53\xaf\xb3\x3b\xe3\x6f\xbc\xa9\x26\x59\x13\x4a\x0d\x4d\x0f\x85\x4c\x09\x87\x46\xa5\x75\x5e\x2a\x04\x66\xd3\xf6\xe3\x18\xda\xce\x99\xf6\xa9\xf3\x61\x13\x09\x40\x53\xdd\x68\x55\xc1\xba\x89\xcf\xec\x21\x2b\xb1\xda\x95\xfb\xc7\x95\xd8\x53\x5a\x48\x4d\x02\x40\xdf\x5a\x73\xdf\x48\xc4\x8e\xb2\x5c\xcd\x72\xae\x2a\xd2\x35\x72\x55\x97\x98\x93\x37\xae\x7b\xe1\x39\x14\x47\x59\x34\x54\x61\xb3\xee\xad\x0f\xd3\x7a\xd1\xae\x2f\xe5\x2a\x83\x35\x37\x4f\x6b\x4d\x22\xe6\x59\x1c\x23\x77\x81\xfd\x14\x30\x3d\x33\x1c\x7f\xe0\x34\x58\x76\x13\x3a\x67\xe0\xb9\x67\xcf\xee\xc4\xd7\xa9\x5d\x36\xd9\xe0\xc2\x3e\x8c\xee\x3b\xc9\x52\xad\x1f\xc7\xd6\x9a\xbf\x78\xef\x86\x37\x8e\xbe\x78\xfa\x45\xdc\xfe\x8b\x92\x5f\xb8\xb6\xd1\x15\xf6\xd6\xe4\x86\x4f\x6a\x9f\xdc\xbf\x0c\x28\xa4\xca\x1a\x99\x11\x2e\xc3\xe5\x1c\x5e\x07\x1c\xcb\x42\xd6\x79\x41\x89\x10\x71\x8c\x8a\x0a\x4a\x6b\xfc\xc1\x41\x97\xff\x7f\x3c\x49\x22\x3e\x03\x00\x00\xff\xff\x8c\x18\xf3\xb6\xc4\x01\x00\x00")

func _0002_secondUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0002_secondUpSql,
		"0002_second.up.sql",
	)
}

func _0002_secondUpSql() (*asset, error) {
	bytes, err := _0002_secondUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0002_second.up.sql", size: 452, mode: os.FileMode(420), modTime: time.Unix(1544231603, 0)}
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
	"0001_first.down.sql":  _0001_firstDownSql,
	"0001_first.up.sql":    _0001_firstUpSql,
	"0002_second.down.sql": _0002_secondDownSql,
	"0002_second.up.sql":   _0002_secondUpSql,
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
	"0001_first.down.sql":  &bintree{_0001_firstDownSql, map[string]*bintree{}},
	"0001_first.up.sql":    &bintree{_0001_firstUpSql, map[string]*bintree{}},
	"0002_second.down.sql": &bintree{_0002_secondDownSql, map[string]*bintree{}},
	"0002_second.up.sql":   &bintree{_0002_secondUpSql, map[string]*bintree{}},
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
