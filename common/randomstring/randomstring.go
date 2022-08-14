package randomstring

import ()

type RandomString interface {} 

type randomstr struct {}
func CreateRandomString() RandomString {
  return &randomstr{}
}
