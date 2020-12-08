package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {

	foundedUrllist := map[string]string{}

	Crawl("http://www.google.com", 1, foundedUrllist)

	fmt.Println("Founded URL Count: ", len(foundedUrllist))
}

//Crawl func fetches an URL and lists URLs on the loaded pages
func Crawl(url string, depth int, foundedUrllist map[string]string) {

	if depth < 0 {
		return
	}

	//get the page
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	//read the page as binary
	dataBytes, err := ioutil.ReadAll(response.Body)

	//convert the page to string
	pageContent := string(dataBytes)

	//regex for URL detection
	regex := regexp.MustCompile(`(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)

	URLs := regex.FindAllStringSubmatch(pageContent, -1)

	depth--

	for i := range URLs {

		urlFounded := URLs[i][0]

		_, ok := foundedUrllist[urlFounded]

		if !ok {

			foundedUrllist[urlFounded] = ""

			fmt.Printf("URL: %s\n", urlFounded)

			Crawl(urlFounded, depth, foundedUrllist)
		}
	}
}
