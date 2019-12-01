package main

import (
	"io/ioutil"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	cacert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		panic(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cacert)
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}
	tlsConfig.BuildNameToCertificate()

	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		panic(nil)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
				RootCAs: certPool,
			},
		},
	}

	resp, err := client.Get("https://localhost:18443")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
