package test

import (
	//"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type fileReader struct {
	file   string
	offset int64
}

func (f *fileReader) Read(p []byte) (n int, err error) {
	reader, err := os.Open(f.file)
	defer reader.Close()
	if err != nil {
		return 0, err
	}
	reader.Seek(f.offset, 0)

	n, err = reader.Read(p)

	if err == io.EOF {
		time.Sleep(1 * time.Second)
	}
	f.offset += int64(n)

	return n, err
}
func GetFileName(path string) string {
	s := strings.Split(path, `\`)
	return s[len(s)-1]
}
func main() {
	//	file := &fileReader{os.Args[1], 0}
	//	br := bufio.NewReader(file)
	//
	//	for {
	//		log, _, err := br.ReadLine()
	//
	//		if err == io.EOF {
	//			continue
	//		}
	//
	//		if err != nil {
	//			fmt.Printf("err: %v", err)
	//			return
	//		}
	//
	//		fmt.Printf("%s\n", string(log))
	//	}

	//	fmt.Println(strings.Replace("df  dfd       sdsfd", " ", "", -1))
	//	var n []byte
	//	b := []byte("hello")
	//	d := bytes.Buffer{}
	//	fmt.Println(BytesJoin(n, b))
	//	d.Write(b)
	//	fmt.Println(d.String())
	//	d.Write([]byte("world"))
	//	fmt.Println(d.String())

	//	s := []string{`1`, `2`, `3`, `4`}
	//	fmt.Println(s[1:])
	fmt.Println(time.Now().ISOWeek())
	fmt.Println(time.Now().Weekday().String())
	fmt.Println(strings.Contains("dfdfd", ""))
	m := map[string]string{
		"a": "sdsd",
	}
	k, v := m["a"]
	fmt.Println(k, v)

	fmt.Println(len(strings.Split("dfdfd", "|")))
	ps := string(os.PathSeparator)
	fmt.Println(strings.TrimRight(`d:\wwere\dfdfd\\\`, ps))

	s := "dddsxx"
	fmt.Println(strings.LastIndex(s, "/"))

	fmt.Println(Substr(`1234`, -1))

	fmt.Println(strings.Count("1,2,3,4", ","))

}
func BytesJoin(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	return bytes.Join(s, []byte(""))
}
func Substr(s string, n int) string {
	rs := []rune(s)
	if n > 0 {
		return string(rs[0:n])
	}
	return string(rs[-n:len(rs)])
}
