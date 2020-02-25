package screenshot

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"time"

	"github.com/chromedp/chromedp"
)

var chromeContext context.Context
var Cancel context.CancelFunc

type ChromeClient struct {
	chromeContext context.Context
	Cancel        context.CancelFunc
}

func NewClient() (cc *ChromeClient) {
	cc = &ChromeClient{}
	cc.setup()
	return
}

func (cc *ChromeClient) setup() {
	cc.chromeContext, cc.Cancel = chromedp.NewContext(context.Background())
}

func screenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.CaptureScreenshot(res),
		chromedp.Stop(),
	}
}

// TODO take in *ServiceDescriptor instead of urlstr -- returns error only
func (cc *ChromeClient) GetScreenshot(urlstr string, savePath string) (string, error) {
	// log.Println("taking shot of ", urlstr)
	var buf []byte
	err := chromedp.Run(cc.chromeContext, screenshot(urlstr, `#main`, &buf))
	if err != nil {
		// log.Fatal(err)
		return "", err
	}
	if err != nil {
		log.Fatal(err)
	}
	// save the screenshot to disk
	file_path := fmt.Sprintf("%s/%s%d.png", savePath, url.QueryEscape(urlstr), time.Now().Unix())
	if err = ioutil.WriteFile(file_path, buf, 0644); err != nil {
		return "", err
	}
	//file_path = url.QueryEscape(file_path)
	return file_path, nil
}
