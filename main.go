package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/mariolima/screenurls/pkg/screenshot"
	"github.com/mariolima/screenurls/pkg/webui"
)

type probeArgs []string

func (p *probeArgs) Set(val string) error {
	*p = append(*p, val)
	return nil
}

func (p probeArgs) String() string {
	return strings.Join(p, ",")
}

func main() {
	// concurrency flag
	var concurrency int
	flag.IntVar(&concurrency, "c", 10, "set the concurrency level")

	// probe flags
	var probes probeArgs
	flag.Var(&probes, "p", "add additional probe (proto:port)")

	var web bool
	flag.BoolVar(&web, "w", false, "setup a webui that shows all the screenshots")

	var savePath string
	flag.StringVar(&savePath, "o", "screenurls_out", "directory output for screensurls")

	// timeout flag
	var to int
	flag.IntVar(&to, "t", 10000, "timeout (milliseconds)")

	// verbose flag
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "output errors to stderr")

	flag.Parse()

	// make an actual time.Duration out of the timeout
	//timeout := time.Duration(to * 1000000)

	if verbose {
		log.SetLevel(log.TraceLevel)
	}

	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		os.Mkdir(savePath, os.ModePerm)
	}

	// we send urls to check on the urls channel,
	// but only get them on the output channel if
	// they are accepting connections
	urls := make(chan string)

	// gowitness module
	// chrm := &chrome.Chrome{
	// 	Resolution:       "800x200",
	// 	ChromeTimeout:    5,
	// 	ChromeTimeBudget: 5,
	// 	UserAgent:        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36",
	// 	ScreenshotPath:   "out",
	// }
	// chrm.Setup()

	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			cc := screenshot.NewClient()
			for u := range urls {
				path, err := cc.GetScreenshot(u, savePath)
				log.Trace("got shot at ", path)
				// u, err := url.Parse(u)
				if err != nil && verbose {
					fmt.Fprintf(os.Stderr, "failed: %v %s\n", err, u)
				}
				// chrm.ScreenshotURL(u, savePath)
				fmt.Println(u)
			}
			wg.Done()
		}()
	}

	if web {
		ms := &webui.MatchServer{
			Port:     1337,
			Hostname: "127.0.0.1",
		}
		go ms.Setup()
	}

	// accept domains on stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		// url := fmt.Sprintf("https://%s", strings.ToLower(sc.Text()))
		urls <- sc.Text()
	}

	// once we've sent all the URLs off we can close the
	// input channel. The workers will finish what they're
	// doing and then call 'Done' on the WaitGroup
	close(urls)

	// check there were no errors reading stdin (unlikely)
	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}

	// Wait until all the workers have finished
	wg.Wait()
}
