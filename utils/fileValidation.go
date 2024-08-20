package utils

import (
	"crypto/sha256"
	"reflect"
)

func CheckFileIntegrity(fileValue []byte) bool {
	shadowKey := []byte{38, 185, 77, 11, 19, 75, 119, 233, 253, 35, 224, 54, 11, 253, 129, 116, 15, 128, 251, 127, 101, 65, 209, 216, 197, 216, 94, 115, 238, 85, 15, 115}
	standardKey := []byte{225, 148, 241, 3, 52, 66, 97, 122, 184, 167, 142, 28, 166, 58, 32, 97, 245, 204, 7, 163, 240, 90, 194, 38, 237, 50, 235, 157, 253, 34, 166, 191}
	thinkerKey := []byte{100, 40, 94, 73, 96, 209, 153, 244, 129, 147, 35, 196, 220, 99, 25, 186, 52, 241, 240, 221, 157, 161, 77, 7, 17, 19, 69, 245, 215, 108, 63, 163}

	fileHash := sha256.New()
	fileHash.Write(fileValue)
	bs := fileHash.Sum(nil)

	if reflect.DeepEqual(bs, shadowKey) || reflect.DeepEqual(bs, standardKey) || reflect.DeepEqual(bs, thinkerKey) {
		return true
	} else {
		return false
	}
}
