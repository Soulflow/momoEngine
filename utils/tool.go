package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/satori/go.uuid"
	"io"
	"net/http"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func CreateUUID() string {
	u1 := uuid.NewV4()
	return u1.String()
}

/**
 * md5生成
 */
func CreateMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * 序列化json
 */
func JsonBind(ptr interface{}, rq *http.Request) error {
	if rq.Body != nil {
		defer rq.Body.Close()
		err := json.NewDecoder(rq.Body).Decode(ptr)
		if err != nil && err != io.EOF {
			return err
		}
		return nil
	} else {
		return errors.New("empty request body")
	}
}