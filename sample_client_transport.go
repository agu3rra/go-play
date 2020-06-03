package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true, // test server certificate is not trusted.
		MaxVersion:         tls.VersionSSL30,
	}

	proxyUrl, err := url.Parse("http://someproxy.com:8000")
	if err != nil {
		log.Fatal(err)
	}

	tr := &http.Transport{
		MaxIdleConns:        10,
		IdleConnTimeout:     5 * time.Second,
		DisableCompression:  true,
		TLSClientConfig:     tlsConf,
		TLSHandshakeTimeout: 3 * time.Second,
		Proxy:               http.ProxyURL(proxyUrl),
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", "https://www.google.com/robots.txt", nil)
	// req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
