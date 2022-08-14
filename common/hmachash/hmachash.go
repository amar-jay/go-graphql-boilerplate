package hmachash

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"
)

type HMAC interface {
  Hash(input string) string
} 

type hm struct {
  hmac hash.Hash
}

func NewHMAC(key string) HMAC {
  h := hmac.New(sha256.New, []byte(key))

  return hm{
    hmac: h,
  }
}

func (h hm) Hash(input string) string {
  return "not implemented"
}
