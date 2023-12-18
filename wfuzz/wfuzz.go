package wfuzz

import (
	"errors"
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

			urlParts := []string{targetUrl, "/" + p}
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
			// 将访问结果写入日志
			err = errors.New(url + " (" + resp.Status + ")")
			log.LogError(err)

		}(payload)
	}

	wg.Wait()

	// end
	fmt.Printf("            __                  _____ _             \n __      __/ _|_   _ ________  |_   _| |_ ___   ___ \n \\ \\ /\\ / / |_| | | |_  /_  /____| | | __/ _ \\ / __|\n  \\ V  V /|  _| |_| |/ / / /_____| | | || (_) | (__ \n   \\_/\\_/ |_|  \\__,_/___/___|    |_|  \\__\\___/ \\___|\n                                                    ")
	fmt.Printf("\n  Wfuzz-ttoc is finished.\n  THANKS FOR USING!\n")
	err := errors.New("Wfuzz-ttoc is finished.")
	log.LogError(err)
	return results, nil
}
