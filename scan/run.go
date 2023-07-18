package scan

import (
	"XSSikeng/utils"
	"fmt"
)

func Run(urlList []string, words []string, payloads []string, proxyURL string, timeOut int) {

	//check if url can be parsed, if not delete this url
	validUrlList, errList := utils.CheckParse(urlList)
	for _, err := range errList {
		fmt.Println("\nError:", err)
	}
	//add injected point to Path
	temp := append([]string{}, validUrlList...)
	modifiedURL := utils.AddInjectedPointToURL(temp)
	for _, cval := range modifiedURL {
		fmt.Println(cval)
	}
	modifiedURL = utils.RemoveDuplicates(modifiedURL)

	// check XSS in path
	for _, url := range modifiedURL {
		InPath(url, words, payloads, proxyURL, timeOut)

	}

	// check XSS in parameter
	temp1 := append([]string{}, validUrlList...)
	temp1 = utils.RemoveDuplicates(temp1)
	for _, url := range temp1 {

		InParameter(url, words, payloads, proxyURL, timeOut)

	}

}
