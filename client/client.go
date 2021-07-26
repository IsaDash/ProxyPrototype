package main

import (
	"bufio"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Request: ")
	request, _ := reader.ReadString('\n')
	request = request[:len(request) - 1]
	fmt.Println("URI: ")
	URI, _ := reader.ReadString('\n')
	URI = URI[:len(URI) - 1]

	// load client cert and key pair
	//clientCert, err := tls.LoadX509KeyPair("client1.crt", "client1.key")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// load ca cert
	caCert, err := ioutil.ReadFile("../nginx/certs/ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	//append ca cert to cert pool
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// tls config for client
	//tlsConf := &tls.Config{
	//	Certificates: []tls.Certificate{clientCert},
	//	//RootCAs: caCertPool, // root certificate authorities used when verifying server certs
	//	InsecureSkipVerify: true, // when set to false, tls will verify server's certificate chain and host name
	//}

	// http Transport for supporting HTTP proxies and manage underlying tcp connection
	//transport := &http.Transport{TLSClientConfig: tlsConf}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:7000", nil)
	req.Header.Set("request", request)
	req.Header.Set("URI", URI)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response status:", res.Status)

	scanner := bufio.NewScanner(res.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

