package wfuzz

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"wfuzz-ttoc/log"
)

func Wfuzz(targetUrl string, payloadList []string, concurrency int, showSuccess string) ([]string, error) {
	var results []string
	var wg sync.WaitGroup

	sem := make(chan struct{}, concurrency)

	for _, payload := range payloadList {

		sem <- struct{}{}
		wg.Add(1)

		go func(p string) {
			defer func() {
				<-sem
				wg.Done()
			}()

			urlParts := []string{targetUrl, p}
			url := strings.Join(urlParts, "")
			resp, err := http.Head(url)
			if err != nil {
				log.LogError(err)
				return
			}

			if showSuccess == "200" {
				if resp.Status == "200" {
					fmt.Printf("%s (%s)\n", url, resp.Status)
				}
			} else if showSuccess == "404" {
				if resp.Status == "404" {
					fmt.Printf("%s (%s)\n", url, resp.Status)
				}
			} else if showSuccess == "303" {
				if resp.Status == "303" {
					fmt.Printf("%s (%s)\n", url, resp.Status)
				}
			} else {
				fmt.Printf("%s (%s)\n", url, resp.Status)
			}

			if resp.StatusCode != http.StatusNotFound {
				results = append(results, url)
			}
		}(payload)
	}

	wg.Wait()

	return results, nil
}
