package demo

import (
	"fmt"
	"regexp"
	"testing"
)

func expandTest() {
	pat := `(((abc.)def.)ghi)`
	reg := regexp.MustCompile(pat)
	fmt.Println(reg.NumSubexp())

	src := []byte(`abc-def-ghi abc+def+ghi`)
	template := []byte(`$0   $1   $2   $3`)

	// 替换第一次匹配结果
	match := reg.FindSubmatchIndex(src)
	fmt.Printf("%v\n", match) // [0 11 0 11 0 8 0 4]
	dst := reg.Expand(nil, template, src, match)
	fmt.Printf("%s\n\n", dst)
	// abc-def-ghi   abc-def-ghi   abc-def-   abc-

	// 替换所有匹配结果
	for _, match := range reg.FindAllSubmatchIndex(src, -1) {
		fmt.Printf("%v\n", match)
		dst := reg.Expand(nil, template, src, match)
		fmt.Printf("%s\n", dst)
	}
	// [0 11 0 11 0 8 0 4]
	// abc-def-ghi   abc-def-ghi   abc-def-   abc-
	// [12 23 12 23 12 20 12 16]
	// abc+def+ghi   abc+def+ghi   abc+def+   abc+
}

func testFind(t *testing.T) {
	re := regexp.MustCompile("a*r")
	fmt.Println(string(re.Find([]byte("paranoabrmal"))))
	fmt.Println(re.NumSubexp())

	rep := regexp.MustCompilePOSIX("a*r|ara")
	fmt.Println(string(rep.Find([]byte("paranoabrmal"))))
	fmt.Println(rep.NumSubexp())

	b := []byte("abc1def1")
	pat := `abc1|abc1def1`
	reg1 := regexp.MustCompile(pat)      // 第一匹配
	reg2 := regexp.MustCompilePOSIX(pat) // 最长匹配
	fmt.Printf("%s\n", reg1.Find(b))     // abc1
	fmt.Printf("%s\n", reg2.Find(b))     // abc1def1
	fmt.Println(reg1.NumSubexp())

	b = []byte("abc1def1")
	pat = `(abc|abc1def)*1`
	reg1 = regexp.MustCompile(pat)      // 第一匹配
	reg2 = regexp.MustCompilePOSIX(pat) // 最长匹配
	fmt.Printf("%s\n", reg1.Find(b))    // abc1
	fmt.Printf("%s\n", reg2.Find(b))    // abc1def1
	fmt.Println(reg1.NumSubexp())
}

func testFindAll(t *testing.T) {
	re := regexp.MustCompile("ar")
	fmt.Printf("%q\n", (re.FindAll([]byte("paranoarmal"), -1)))

	rep := regexp.MustCompilePOSIX("ar")
	fmt.Printf("%q\n", (rep.FindAll([]byte("paranoarmal"), -1)))

	pat := `(((abc.)def.)ghi)`
	src := []byte(`abc-def-ghi abc+def+ghi`)

	reg := regexp.MustCompile(pat)
	fmt.Printf("%q\n", (reg.Find(src)))
	fmt.Printf("%q\n", (reg.FindAll(src, -1)))
	regp := regexp.MustCompilePOSIX(pat)
	fmt.Printf("%q\n", (regp.Find(src)))
	fmt.Printf("%q\n", (regp.FindAll(src, -1)))

}

func testReg(pat, srcStr string) {

	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	src := []byte(srcStr)

	reg := regexp.MustCompile(pat)
	fmt.Printf("%q\n", (reg.Find(src)))
	fmt.Printf("%q\n", (reg.FindString(srcStr)))

	fmt.Printf("%q\n", (reg.FindAll(src, -1)))
	fmt.Printf("%q\n", (reg.FindAllString(srcStr, -1)))

	fmt.Printf("%d\n", (reg.FindIndex(src)))
	fmt.Printf("%d\n", (reg.FindStringIndex(srcStr)))

	fmt.Printf("%d\n", (reg.FindAllIndex(src, -1)))
	fmt.Printf("%d\n", (reg.FindAllStringIndex(srcStr, -1)))

	fmt.Println("begin submatch")
	fmt.Printf("%q\n", (reg.FindSubmatch(src)))
	fmt.Printf("%q\n", (reg.FindStringSubmatch(srcStr)))

	fmt.Printf("%d\n", (reg.FindSubmatchIndex(src)))
	fmt.Printf("%d\n", (reg.FindStringSubmatchIndex(srcStr)))

	fmt.Printf("%q\n", (reg.FindAllSubmatch(src, -1)))
	fmt.Printf("%q\n", (reg.FindAllStringSubmatch(srcStr, -1)))

	fmt.Printf("%d\n", (reg.FindAllSubmatchIndex(src, -1)))
	fmt.Printf("%d\n", (reg.FindAllStringSubmatchIndex(srcStr, -1)))

	regp := regexp.MustCompilePOSIX(pat)
	fmt.Printf("%q\n", (regp.Find(src)))
	fmt.Printf("%q\n", (regp.FindString(srcStr)))

	fmt.Printf("%q\n", (regp.FindAll(src, -1)))
	fmt.Printf("%q\n", (regp.FindAllString(srcStr, -1)))

	fmt.Printf("%d\n", (regp.FindIndex(src)))
	fmt.Printf("%d\n", (regp.FindStringIndex(srcStr)))

	fmt.Printf("%d\n", (regp.FindAllIndex(src, -1)))
	fmt.Printf("%d\n", (regp.FindAllStringIndex(srcStr, -1)))

	fmt.Println("begin submatch")
	fmt.Printf("%q\n", (regp.FindSubmatch(src)))
	fmt.Printf("%q\n", (regp.FindStringSubmatch(srcStr)))

	fmt.Printf("%q\n", (regp.FindAllSubmatch(src, -1)))
	fmt.Printf("%q\n", (regp.FindAllStringSubmatch(srcStr, -1)))

	fmt.Printf("%d\n", (regp.FindSubmatchIndex(src)))
	fmt.Printf("%d\n", (regp.FindStringSubmatchIndex(srcStr)))

	fmt.Printf("%d\n", (regp.FindAllSubmatchIndex(src, -1)))
	fmt.Printf("%d\n", (regp.FindAllStringSubmatchIndex(srcStr, -1)))

}

func TestRegexp(t *testing.T) {

	testReg(`(((abc.)def.)ghi)x*`, `abc-def-ghixxa abc+def+ghixx`)
	testReg(`(((abc.)def.)ghi)`, `abc-def-ghixxa abc+def+ghixx`)
	testReg(`a(x*)b(y|z)c`, `-axxxbyc- -abzc-`)
	testReg(`a*r`, `paranoabrmal`)

}
