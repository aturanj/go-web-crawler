package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	Crawl("https://www.google.com")
}

//Crawl func fetches an URL and lists URLs on the loaded pages
func Crawl(url string) {

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)

	pageContent := string(dataBytes)

	regex := regexp.MustCompile(`(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)

	URLs := regex.FindAllStringSubmatch(pageContent, -1)

	list := map[string]string{}

	for i := range URLs {

		urlFounded := URLs[i][0]

		_, ok := list[urlFounded]

		if !ok {

			list[urlFounded] = ""
			fmt.Printf("URL: %s\n", urlFounded)
		}
	}
}
