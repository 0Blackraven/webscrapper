package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	var url string = "https://github.com"
	robotResolver(url);

}

func robotResolver(url string) {
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
	fmt.Printf("%v",result);
}

