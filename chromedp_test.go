package test

import (
	"testing"

	"github.com/chromedp/chromedp"
)

func TestChromdp(t *testing.T) {

}

func text(res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(`http://live.titan007.com/`),
		chromedp.Text(`#team1_1502762`, res, chromedp.NodeVisible, chromedp.ByID),
	}
}
