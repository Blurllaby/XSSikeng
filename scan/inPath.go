package scan

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func InPath(url string, words []string, payloads []string, proxyURL string, timeOut int) {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set up options to run Chrome in headless mode
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.ProxyServer(proxyURL),
	)

	// Create a new context with the modified options
	ctx, cancel = chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	// Create a new Chrome instance
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Create a new context with a timeout of any seconds
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(timeOut)*time.Second)
	defer cancel()

	var check bool

	// Alert Handler
	chromedp.ListenTarget(ctx, func(ev interface{}) {

		if ev, ok := ev.(*page.EventJavascriptDialogOpening); ok {
			if ev.Message == "14091608" {
				// fmt.Println("***	Alert method was triggered for URL: ", ev.URL)
				go func() {
					if err := chromedp.Run(ctx,
						page.HandleJavaScriptDialog(true),
					); err != nil {
						fmt.Printf("%s", err)
					}
				}()
			}
			check = true
		}
	})

	// if the value of is reflected
	for _, word := range words {
		// Replace "FUZZ" with the payload in the URL
		modifiedURL := strings.ReplaceAll(url, "FUZZ", word)

		// Navigate to the modified URL
		err := chromedp.Run(timeoutCtx, chromedp.Navigate(modifiedURL))
		if err != nil {
			fmt.Printf("Failed to navigate to URL %s: %v\n", modifiedURL, err)
			chromedp.Run(timeoutCtx, chromedp.Navigate("about:blank"))
			continue

		}

		// Get the element content
		var elementContent string
		err = chromedp.Run(ctx, chromedp.OuterHTML("html", &elementContent, chromedp.NodeReady, chromedp.ByQuery))
		if err != nil {
			fmt.Printf("Failed to get element content for URL %s: %v\n", modifiedURL, err)

		}
		// Check if the payload appears in the element content
		if strings.Contains(elementContent, word) {
			fmt.Printf("[%s] reflected with URL:\n***	%s \n", word, modifiedURL)
		}

	}

	// if the alert is executed
	for _, payload := range payloads {
		// Replace "FUZZ" with the payload in the URL
		modifiedURL := strings.ReplaceAll(url, "FUZZ", payload)

		check = false
		// Navigate to the modified URL
		err := chromedp.Run(timeoutCtx, chromedp.Navigate(modifiedURL))
		if err != nil {
			fmt.Printf("Failed to navigate to URL %s: %v\n", modifiedURL, err)
			chromedp.Run(timeoutCtx, chromedp.Navigate("about:blank"))
			continue

		}
		if check {
			fmt.Printf("[Alert] method was triggered for URL:\n***	%s \n", modifiedURL)
		}

	}
}
