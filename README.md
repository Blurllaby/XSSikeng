# XSSikeng
* A golang tool for scanning DOM, Reflected XSS and Client-side prototype pollution if it exists in url via path or parameter. 
* By using Chromedp, this tool checks if the payload would be reflected in the "Elements" tab or if the alert() method would be executed.

## Usage
| Flag |  |
| ------------- | ------------- |
| -payloads  | XSS payload or Client-side prototype pollution payload. _read the Notice below_  |
| -words  | unique words for reflection only check  |
| -timeOut  | waiting time for drop out of the request (optional)|
| -proxy  | your proxy address (optional)|
```
go build .
cat uriList.txt | ./XSSikeng -payloads payloads.txt -words words.txt [-timeOut 10 -proxy http://127.0.0.1]
```
## Notice 
* Payloads must contain the alert() function or the fuzz format of it, since the tool works by realizing if the alert() popup fired. By default, the number in alert() is 14091608 for the tool can be aware of the alert(14091608) is triggered, then automatically sets this popup to true so the tool proceeds. You are able to change it in /scan/inParam and /scan/inPath.
* For Client-side prototype pollution payload, refer to this: [client-side-prototype-pollution](https://github.com/BlackFan/client-side-prototype-pollution).
