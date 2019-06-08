package test

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var res string
	err = c.Run(ctxt, text(&res))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("抓取澳维超-比赛球队:", res)

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("overview: %s", res)
}

func text(res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`http://live.titan007.com/`),
		chromedp.Text(`#team1_1502762`, res, chromedp.NodeVisible, chromedp.ByID),
	}
}
