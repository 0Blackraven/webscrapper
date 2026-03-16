package utils

import (
	"net/http"
	"bufio"
	"strings"
)

func RobotResolver(url string) []string {
    link := url + "/robots.txt";
    response, err := http.Get(link);
    if err != nil || response.StatusCode != http.StatusOK {
        return nil;
    }
    defer response.Body.Close();
    var result []string;
    scanner := bufio.NewScanner(response.Body);
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text());
        if strings.HasPrefix(strings.ToLower(line), "disallow:") {
            path := strings.SplitN(line, ":", 2)[1];
            if path != "" {
                result = append(result, strings.TrimSpace(path))
            }
        }
    }
    return result
}