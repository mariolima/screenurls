package screenshot

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var chromeContext context.Context
var Cancel context.CancelFunc

type ChromeClient struct {
	chromeContext context.Context
	Cancel        context.CancelFunc
}

type UrlToScreen struct {
	Url   string
	Vhost string
}

func NewClient() (cc *ChromeClient) {
	cc = &ChromeClient{}
	cc.setup()
	return
}

func (cc *ChromeClient) setup() {
	cc.chromeContext, cc.Cancel = chromedp.NewContext(context.Background())
}

func screenshot(ts *UrlToScreen, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		//TODO add ts.Vhost
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers{
			"Host": ts.Vhost,
		}),
		// network.SetExtraHTTPHeaders(network.Headers(network.Headers{"Host": ts.Vhost})),
		chromedp.Navigate(ts.Url),
		chromedp.Sleep(600 * time.Millisecond),
		chromedp.CaptureScreenshot(res),
		chromedp.Stop(),
	}
}

// TODO take in *ServiceDescriptor instead of urlstr -- returns error only
func (cc *ChromeClient) GetScreenshot(ts *UrlToScreen, savePath string) (string, error) {
	// logis.Debug("taking shot of ", ts.Url)
	var buf []byte
	err := chromedp.Run(cc.chromeContext, screenshot(ts, `#main`, &buf))

	if err != nil {
		// log.Fatal(err)
		return "", err
	}
	if err != nil {
		log.Fatal(err)
	}
	// save the screenshot to disk
	file_path := fmt.Sprintf("%s/%s%d.png", savePath, url.QueryEscape(ts.Url), time.Now().Unix())
	if err = ioutil.WriteFile(file_path, buf, 0644); err != nil {
		return "", err
	}
	//file_path = url.QueryEscape(file_path)
	return file_path, nil
}
