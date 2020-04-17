package xmc2go

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"time"
)

func getToken() string {
	//纳秒级别的
	t := time.Now().UnixNano()
	h := md5.New()
	b := make([]byte, 16)

	binary.BigEndian.PutUint64(b, uint64(t))
	h.Write([]byte(b))

	return hex.EncodeToString(h.Sum(nil))
}
