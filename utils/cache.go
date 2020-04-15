package utils

import (
	"bytes"
	"encoding/gob"
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var cc  *cache.Cache

func InitCache() {

	cc = cache.New(5*time.Minute, 10*time.Minute)
}

// SetCache
func SetCache(key string, value interface{}, timeout int) error {
	data, err := Encode(value)
	if err != nil {
		return err
	}
	if cc == nil {
		return errors.New("cc is nil")
	}
	defer func() {
		if r := recover(); r != nil {
			LogError(r)
			cc = nil
		}
	}()
	timeouts := time.Duration(timeout) * time.Second
	cc.Set(key, data, timeouts)
	return nil

}

func GetCache(key string, to interface{}) error {
	if cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			LogError(r)
			cc = nil
		}
	}()
	data, _ := cc.Get(key)
	if data == nil {
		return errors.New("cache is not exist")
	}
	err := Decode(data.([]byte), to)
	if err != nil {
		LogError(err)
		LogError("GetCache失败，key:" + key)
	}
	return err
}
// Encode
// 用gob进行数据编码
//
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode
// 用gob进行数据解码
//
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
