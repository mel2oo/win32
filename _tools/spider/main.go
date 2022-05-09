package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

const evalJS = `
(function() { 
	var res = [];
	var tbody = document.querySelector("#main > div.content > div:nth-child(8) > table > tbody");
	for (const tr of tbody.children) {
		res.push(tr.children[0].firstChild.innerHTML);
	}
	return res;
})(); 
`

func main() {
	url := flag.String("url", "https://docs.microsoft.com/en-us/windows/win32/api/shellapi/", "")
	flag.Parse()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.0 Safari/537.36`),
	)

	var (
		ctx              context.Context
		cancelA, cancelC context.CancelFunc
	)

	ctx, cancelA = chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancelA()
	ctx, cancelC = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancelC()

	defer chromedp.Cancel(ctx)

	var res []string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(*url),
		chromedp.Evaluate(evalJS, &res),
	); err != nil {
		panic(err)
	}

	for _, api := range res {
		if strings.HasSuffix(api, "A") {
			continue
		}

		fmt.Printf("proc%s = dll.NewProc(\"%s\")\n", strings.TrimSuffix(api, "W"), api)
	}
}
