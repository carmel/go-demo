package test

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/agent"
)

func TestHeadlessbrowser(t *testing.T) {
	// Create a new browser and open reddit.
	bow := surf.NewBrowser()
	bow.AddRequestHeader("Accept", "text/html")
	bow.AddRequestHeader("Accept-Charset", "utf8")
	//bow.SetCookieJar(http.CookieJar)
	bow.SetUserAgent(agent.CreateVersion("chrome", "67"))
	//bow.SetProxy("http://192.168.1.120:82/")
	bow.SetTimeout(time.Second * 1500)
	err := bow.Open("http://live.titan007.com/")
	fout, _ := os.Create("d:/1.html")
	bow.Download(fout)
	if err != nil {
		fmt.Println("open err", err)
		panic(err)
	}
	bow.Find("td.td_status").Each(func(_ int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
	ctx := bow.ResponseHeaders().Get("Content-Type")
	h, _ := bow.Dom().Html()
	time.Sleep(time.Second * 10)
	fmt.Println(GBConvertor(ctx, h))
	// Outputs: "reddit: the front page of the internet"

	// "Go 世界！123 Hello."

	// Click the link for the newest submissions.
	//	bow.Click("a.new")
	//	// Outputs: "newest submissions: reddit.com"
	//	fmt.Println(bow.Title())
	//
	//	// Log in to the site.
	//	fm, _ := bow.Form("form.login-form")
	//	fm.Input("user", "JoeRedditor")
	//	fm.Input("passwd", "d234rlkasd")
	//	if fm.Submit() != nil {
	//		panic(err)
	//	}
	//
	//	// Go back to the "newest submissions" page, bookmark it, and
	//	// print the title of every link on the page.
	//	bow.Back()
	//	bow.Bookmark("reddit-new")
	//	bow.Find("a.title").Each(func(_ int, s *goquery.Selection) {
	//		fmt.Println(s.Text())
	//	})
}

func GBConvertor(ctx string, text string) string {
	if strings.Contains(ctx, "gbk") || strings.Contains(ctx, "GBK") {
		enc := mahonia.NewDecoder("gbk")
		return enc.ConvertString(text)
	} else {
		return text
	}
}
