package hash

import (
	"crypto/md5"
	"crypto/sha1"

	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	sha1 := sha1.New()
	str := []byte("qazwsxeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbnRpdHlJRCI6IjEiLCJJRCI6ImJ2MmdpaGJkMGN2Z3VmN3VjdWcwIiwiT3JnSUQiOiIvMS80LzkiLCJQZXJtaXQiOlsiZmFjdWx0eSIsInN0dWRlbnQiLCJhc210OmNoZW5kdSIsImFzbXQ6Y2xhc3MiLCJsZWF2ZSIsImxlYXZlX2FwcHJvdmUiLCJcdTAwM2NuaWxcdTAwM2UiLCJ1c2VyOmluZm8iXSwiUm9sZVJhbmsiOjAsIlJvbGVTaWduIjoiQURNSU4iLCJUeXBlIjoiMiIsImV4cCI6MTYwNjc5NTM3MCwiaWF0IjoxNjA2Nzk0MTcwfQ.BzqPWjLPLCsHDIenXmtOQizCa78eVOra0JEVgb-usCU")
	sha1.Write(str)
	sum := sha1.Sum(nil)
	fmt.Printf("%x\n", sum)
	md5 := md5.New()
	md5.Write(str)
	sum = md5.Sum(nil)
	fmt.Printf("%x\n", sum)
	// // b170cb28bfc8a4b3fe4b809fb5a428692c54c6c6f8fbe72b97c690956bcaf8b4
	// b170cb28bfc8a4b3fe4b809fb5a428692c54c6c6f8fbe72b97c690956bcaf8b4
}
