package charlatan

import (
	md5Hash "crypto/md5"
	sha1Hash "crypto/sha1"
	sha256Hash "crypto/sha256"
	"fmt"
	goTime "time"
)

func sha1() interface{} {
	h := sha1Hash.New()
	now := goTime.Now().UnixNano()
	h.Write([]byte(fmt.Sprintf("%d", now)))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func sha256() interface{} {
	h := sha256Hash.New()
	now := goTime.Now().UnixNano()
	h.Write([]byte(fmt.Sprintf("%d", now)))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func md5() interface{} {
	h := md5Hash.New()
	now := goTime.Now().UnixNano()
	h.Write([]byte(fmt.Sprintf("%d", now)))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
