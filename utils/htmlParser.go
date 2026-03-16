package utils

import (
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
)

func HtmlParser (url string) []string {
	var links []string
	response,geterr := http.Get(url);
	if geterr != nil {
		fmt.Printf("Error in the get response for [%v] \n Error: [%v]",url,geterr);
	}
	defer response.Body.Close();
	doc,err := goquery.NewDocumentFromReader(response.Body);
	if err != nil {
		fmt.Printf("Could not read the body of the url: [%v]\nError: [%v]",url,err);
	}
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			links = append(links, link)
			fmt.Println(link)
		}
	})
	return links
}

func BaseUrl (input string) (string,string) {
	response, err := url.Parse(input);
	if err != nil {
		fmt.Printf("could not parse the provided url [%v]\n Error : [%v]\n", input, err);
	}
	return 	response.Scheme + "://" + response.Host, response.Path;

}
