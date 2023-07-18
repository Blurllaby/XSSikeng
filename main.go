package main

import (
	"XSSikeng/scan"
	"XSSikeng/utils"
	"fmt"
)

func main() {

	fmt.Println(`
	###          ##           #######          #######              /
   /####       ####  /      /       ###      /       ###    #     #/
  /   ###      /####/      /         ##     /         ##   ###    ##
       ###    /   ##       ##        #      ##        #     #     ##
        ###  /              ###              ###                  ##
         ###/              ## ###           ## ###        ###     ##  /##       /##    ###  /###        /###
          ###               ### ###          ### ###       ###    ## / ###     / ###    ###/ #### /    /  ###  /
          /###                ### ###          ### ###      ##    ##/   /     /   ###    ##   ###/    /    ###/
         /  ###                 ### /##          ### /##    ##    ##   /     ##    ###   ##    ##    ##     ##
        /    ###                  #/ /##           #/ /##   ##    ##  /      ########    ##    ##    ##     ##
       /      ###                  #/ ##            #/ ##   ##    ## ##      #######     ##    ##    ##     ##
      /        ###                  # /              # /    ##    ######     ##          ##    ##    ##     ##
     /          ###   /   /##        /     /##        /     ##    ##  ###    ####    /   ##    ##    ##     ##
    /            ####/   /  ########/     /  ########/      ### / ##   ### /  ######/    ###   ###    ########
   /              ###   /     #####      /     #####         ##/   ##   ##/    #####      ###   ###     ### ###
                        |                |                                                                   ###
                         \)               \)                                                           ####   ###
                                                                                                     /######  /#
                                                                                                    /     ###/
                                                   /| __________________
                                               O|===|* >____ZuoqTr______>
                                                   \|`)

	fmt.Println("ex: cat uriList.txt | ./XSSikeng -payloads payloads.txt -words words.txt -timeOut 10 -proxy http://127.0.0.1")
	var (
		uris, words, payloads          []string
		wordsAddr, payloadsAddr, proxy string
		timeOut                        int
	)

	uris = utils.InputViaPipeLine()
	if uris == nil {
		fmt.Printf("There is no uri list")
	}
	wordsAddr, payloadsAddr, proxy, timeOut = utils.InputViaFlag()
	words = utils.ReadFile(wordsAddr)
	payloads = utils.ReadFile(payloadsAddr)
	scan.Run(uris, words, payloads, proxy, timeOut)
	//fmt.Printf("words: %s\npayloads: %s\nproxy: %s\n, ", wordsAddr, payloadsAddr, proxy)
}
