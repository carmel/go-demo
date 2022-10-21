package demo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"testing"

	"golang.org/x/net/proxy"
)

func ProxyAwareHttpClient() *http.Client {
	// sane default
	var dialer proxy.Dialer
	// eh, I want the type to be proxy.Dialer but assigning proxy.Direct makes the type proxy.direct
	dialer = proxy.Direct
	proxyServer, isSet := os.LookupEnv("HTTP_PROXY")
	if isSet {
		proxyUrl, err := url.Parse(proxyServer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid proxy url %q\n", proxyUrl)
		}
		dialer, err = proxy.FromURL(proxyUrl, proxy.Direct)
	}

	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	httpTransport.Dial = dialer.Dial
	return httpClient
}

func TestProxy(t *testing.T) {
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	client := ProxyAwareHttpClient()
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(contents))
}
