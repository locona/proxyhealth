package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/locona/proxyhealth/pkg/fileutil"
	"github.com/locona/proxyhealth/pkg/proxy"
)

const successProxyFile = "success_proxy_list"

func main() {
	proxyFQDNList, err := proxyFQDNList()
	if err != nil {
		panic(err)
	}

	fp, _, err := fileutil.Recreate(successProxyFile)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	var wg sync.WaitGroup
	for idx := range proxyFQDNList {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := proxy.Health(proxyFQDNList[i])
			if err == nil {
				fmt.Fprintln(fp, proxyFQDNList[i])
			}

			if err != nil {
				log.Printf("Failure Proxy: %v \n", proxyFQDNList[i])
			}
		}(idx)
	}

	wg.Wait()

}

func proxyFQDNList() ([]string, error) {
	fp, err := os.Open("proxy_list")
	defer fp.Close()
	if err != nil {
		return []string{}, err
	}

	scanner := bufio.NewScanner(fp)
	res := make([]string, 0)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res, nil
}

func createFile() {
}
