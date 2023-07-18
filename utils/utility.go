package utils

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func RemoveDuplicates(list []string) []string {
	// Create a map to track element occurrences

	seen := make(map[string]bool)
	result := []string{}

	// Iterate over the list
	for _, element := range list {
		// Check if the element has been seen before
		if !seen[element] {
			// Add the element to the result slice
			result = append(result, element)
			// Mark the element as seen
			seen[element] = true
		}
	}

	return result
}

func InputViaPipeLine() []string {

	urls := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		urls = append(urls, url)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return urls
}

func InputViaFlag() (string, string, string, int) {
	proxyFlag := flag.String("proxy", "", "ex: http://127.0.0.1")
	wordsFlag := flag.String("words", "", "ex: yourpath/yourwordsfile.txt")
	payloadsFlag := flag.String("payloads", "", "ex: yourpath/yourpayloadsfile.txt")
	timeOutFlag := flag.Int("timeOut", 10, "ex: type \"10\" for waiting 10 seconds then stop the request")
	flag.Parse()
	// Print the proxy server address if provided
	if *wordsFlag == "" {
		fmt.Println("\nwords is missing")
	}
	// Print the proxy server address if provided
	if *payloadsFlag == "" {
		fmt.Println("payloads is missing")
	}
	return *wordsFlag, *payloadsFlag, *proxyFlag, *timeOutFlag
}

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	var uriList []string
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return nil
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read and print each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		uriList = append(uriList, line)
	}

	// Check for any errors encountered during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read file:", err)
		return nil
	}
	return uriList
}

func ModifyQueryParam(u *url.URL, paramName, paramValue string) {
	query := u.Query()
	query.Set(paramName, paramValue)
	u.RawQuery = query.Encode()
}

func AddInjectedPointToURL(urlList []string) []string {

	for idx, urlStr := range urlList {
		u0, _ := url.Parse(urlStr)
		u, _ := url.Parse(urlStr)

		u1, _ := url.Parse(urlStr)

		u2, _ := url.Parse(urlStr)

		if u0.RawPath == "" {
			u0.Host = strings.TrimSuffix(u0.Host, "/") + "/FUZZ"
			urlList = append(urlList, u0.String())
		}

		// add FUZZ
		if u.RawQuery != "" {
			u.Path = strings.TrimSuffix(u.Path, "/") + "FUZZ"
		} else {
			u.Path = u.Path + "FUZZ"
		}
		urlList[idx] = u.String()
		// add /FUZZ
		if u1.RawQuery != "" {
			u1.Path = strings.TrimSuffix(u1.Path, "/") + "/FUZZ"
		} else {
			u1.Path = strings.TrimSuffix(u1.Path, "/") + "/FUZZ"
		}
		//add /FUZZ/
		if u2.RawQuery != "" {
			u2.Path = strings.TrimSuffix(u2.Path, "/") + "/FUZZ/"
		} else {
			u2.Path = strings.TrimSuffix(u2.Path, "/") + "/FUZZ/"
		}

		urlList = append(urlList, u1.String())
		urlList = append(urlList, u2.String())

	}

	return urlList
}

func CheckParse(urlList []string) ([]string, []error) {
	var errList []error
	for idx, val := range urlList {
		_, err := url.Parse(val)
		if err != nil {
			urlList = append(urlList[:idx], urlList[idx+1:]...)
			errList = append([]error{}, err)
		}
	}
	return urlList, errList
}
