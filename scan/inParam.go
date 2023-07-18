package scan

import (
	"XSSikeng/utils"
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func InParameter(baseURL string, words []string, payloads []string, proxyURL string, timeOut int) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.ProxyServer(proxyURL),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Create a new context with a timeout of any seconds
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(timeOut)*time.Second)
	defer cancel()

	//Alert Handler
	var check bool

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

	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("%s", err)
	}
	queryParams := u.Query()
	if len(queryParams) > 0 {
		//check if value of parameter is reflected
		for _, word := range words {

			for paramName := range queryParams {
				originalValue := queryParams.Get(paramName)

				utils.ModifyQueryParam(u, paramName, word)
				modifiedURL := u.String()

				var htmlContent string

				// Navigate to the modified URL and retrieve the entire HTML content
				err := chromedp.Run(timeoutCtx, chromedp.Navigate(modifiedURL))
				if err != nil {
					fmt.Printf("Failed to navigate to URL %s: %v\n", modifiedURL, err)
					chromedp.Run(timeoutCtx, chromedp.Navigate("about:blank"))
					continue

				}

				// Get the element content
				err = chromedp.Run(ctx, chromedp.OuterHTML("html", &htmlContent, chromedp.NodeReady, chromedp.ByQuery))
				if err != nil {
					fmt.Printf("Failed to get element content for URL %s: %v\n", modifiedURL, err)
					chromedp.Run(timeoutCtx, chromedp.Navigate("about:blank"))

				}

				// Check if the parameter value is reflected in the HTML content

				if strings.Contains(htmlContent, word) {
					fmt.Printf("[%s] parameter reflected with uri:\n***	%s \n", paramName, modifiedURL)
				}

				utils.ModifyQueryParam(u, paramName, originalValue)

			}

		}
		//check if alert is executed
		for _, payload := range payloads {

			for paramName := range queryParams {
				originalValue := queryParams.Get(paramName)

				utils.ModifyQueryParam(u, paramName, payload)
				modifiedURL := u.String()

				check = false
				// Navigate to the modified URL and retrieve the entire HTML content
				err := chromedp.Run(timeoutCtx, chromedp.Navigate(modifiedURL))
				if err != nil {
					fmt.Printf("Failed to navigate to URL %s: %v\n", modifiedURL, err)
					chromedp.Run(timeoutCtx, chromedp.Navigate("about:blank"))
					continue

				}

				// if the alert is pop up ?
				if check {
					fmt.Printf("[Alert] method was triggered for URL:\n***	%s \n", modifiedURL)
				}

				utils.ModifyQueryParam(u, paramName, originalValue)

			}
		}

	}
}
