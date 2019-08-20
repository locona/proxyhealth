package proxy

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

const healthURL = "https://api.ipify.org?format=json"

func Health(proxyFQDN string) error {
	httpProxy := fmt.Sprintf("http://%v", proxyFQDN)
	httpProxyURL, _ := url.Parse(httpProxy)

	timeout := time.Duration(60 * time.Second)
	httpClient := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			Proxy:             http.ProxyURL(httpProxyURL),
		},
	}
	resp, err := httpClient.Get(healthURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	realIP, err := ClientIP(healthURL)
	if err != nil {
		return err
	}
	proxyClientIP := extractIP(body)

	if realIP == proxyClientIP {
		return fmt.Errorf("Use real ip: %v", realIP)
	}
	return nil
}

func ClientIP(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	clientIP := extractIP(body)
	return clientIP, nil
}

func extractIP(body []byte) string {
	re, _ := regexp.Compile(`^.......|..$`)
	return re.ReplaceAllString(string(body), "")
}
