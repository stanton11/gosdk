package main

import (
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/couchbase/gocb/v2"
)

func main() {
	connection := flag.String("connection", "", "Hostname")
	username := flag.String("username", "", "Username")
	password := flag.String("password", "", "Password")
	cafile := flag.String("cafile", "", "CA filename")

	flag.Parse()

	ca, err := ioutil.ReadFile(*cafile)
	if err != nil {
		fmt.Println("failed to load CA:", err)
		return
	}

	caCerts := x509.NewCertPool()
	caCerts.AppendCertsFromPEM(ca)

	options := gocb.ClusterOptions{
		Username: *username,
		Password: *password,
		SecurityConfig: gocb.SecurityConfig{
			TLSRootCAs: caCerts,
		},
	}

	if _, err := gocb.Connect(*connection, options); err != nil {
		fmt.Println("failed to open connection:", err)
	}
}
