package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	var url string = "https://github.com/0Blackraven/rateLimiter"
	baseUrl, path := baseUrl(url);
	fmt.Printf("%v",baseUrl);
	var restrictedPaths []string = robotResolver(baseUrl);
	for index := range restrictedPaths {
		if strings.HasPrefix(path, restrictedPaths[index]) {
			fmt.Print("cannot read is a restricted link");
		}
		htmlParser(url);
	}
	fmt.Printf("%v",restrictedPaths);
}

func robotResolver(url string) []string {
	var link string = url + "/robots.txt"
	response, err := http.Get(link)
	if err != nil {
		fmt.Printf("err occured: %v", err)
	}
	defer response.Body.Close();
	var result []string;
	scanner:= bufio.NewScanner(response.Body);
	for scanner.Scan() {
		line := scanner.Text();
		if strings.HasPrefix(line, "Disallow:"){
			path:=strings.Split(line, ":")[1];
			result = append(result, path);
		}
	}
	return result;
}

func baseUrl (input string) (string,string) {
	response, err := url.Parse(input);
	if err != nil {
		fmt.Printf("%v", err);
	}
	return 	response.Scheme + "://" + response.Host, response.Path;

}

func htmlParser (url string) {
	response,geterr := http.Get(url);
	if geterr != nil {
		fmt.Print("error getting file");
	}
	defer response.Body.Close();
	doc,err := goquery.NewDocumentFromReader(response.Body);
	if err != nil {
		fmt.Print("error reading the body of the file");
	}
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			fmt.Println(link)
		}
	})
}