package screenshot

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/chromedp/chromedp"
)

var chromeContext context.Context
var Cancel context.CancelFunc

func StartChrome() {
	chromeContext, Cancel = chromedp.NewContext(context.Background())
}

func screenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.CaptureScreenshot(res),
		chromedp.Stop(),
	}
}

// TODO take in *ServiceDescriptor instead of urlstr -- returns error only
func GetScreenshot(urlstr string, savePath string) (string, error) {
	// log.Println("taking shot of ", urlstr)
	var buf []byte
	err := chromedp.Run(chromeContext, screenshot(urlstr, `#main`, &buf))
	if err != nil {
		// log.Fatal(err)
		return "", err
	}
	// save the screenshot to disk
	//file_path := "static/imgs/"+b64.StdEncoding.EncodeToString([]byte(in_url)[:30])+".png"
	//file_path := "static/imgs/"+url.QueryEscape(domain)+time.Now().Unix()+".png"
	file_path := fmt.Sprintf("%s/%s%d.png", savePath, "screenurls", time.Now().Unix())
	if err = ioutil.WriteFile(file_path, buf, 0644); err != nil {
		return "", err
	}
	//file_path = url.QueryEscape(file_path)
	return file_path, nil
}
