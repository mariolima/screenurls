package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

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
	flag.IntVar(&concurrency, "c", 20, "set the concurrency level")

	// probe flags
	var probes probeArgs
	flag.Var(&probes, "p", "add additional probe (proto:port)")

	// skip default probes flag
	var skipDefault bool
	flag.BoolVar(&skipDefault, "s", false, "skip the default probes (http:80 and https:443)")

	var web bool
	flag.BoolVar(&skipDefault, "w", false, "setup a webui that shows all the screenshots")

	// timeout flag
	var to int
	flag.IntVar(&to, "t", 10000, "timeout (milliseconds)")

	// verbose flag
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "output errors to stderr")

	flag.Parse()

	// make an actual time.Duration out of the timeout
	//timeout := time.Duration(to * 1000000)

	var savePath string = "out"

	// we send urls to check on the urls channel,
	// but only get them on the output channel if
	// they are accepting connections
	urls := make(chan string)

	// Spin up a bunch of workers
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			for url := range urls {
				_, err := screenshot.GetScreenshot(url, savePath)
				if err != nil && verbose {
					fmt.Fprintf(os.Stderr, "failed: %s\n", url)
				}
				fmt.Println(url)
			}
			wg.Done()
		}()
	}

	if web {
		ms := &webui.MatchServer{
			Port:     8080,
			Hostname: "127.0.0.1",
		}
		ms.Setup()
	}
	screenshot.StartChrome()

	// accept domains on stdin
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		url := fmt.Sprintf("https://%s", strings.ToLower(sc.Text()))
		urls <- url
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
