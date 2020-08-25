package passkit_golang_sdk

import (
	"fmt"
	"log"
	"google.golang.org/grpc"
	"io/ioutil"
	"github.com/PassKit/passkit-golang-sdk/helpers/router"
)

const (
	// The address & port of the PassKit gRPC service.
	gRPCHost = "grpc-dev.passkit.io"
	gRPCPort = "443"

	// The location of your client certificates.
	clientCertFile = "certs/certificate.pem" // [Required] Please store your certificate.pem at ./certs directory. Your client certificate you receive by email or on Settings > Developer Credential page (https://dev-app.passkit.io/login).
	clientKeyFile  = "certs/key.pem"         // [Required] Please store your key.pem at ./certs directory. Your private key you receive by email or on Settings > Developer Credential page (https://dev-app.passkit.io/login). You need to decrypt the key before use. Check README.md for details.
	clientCAFile   = "certs/ca-chain.pem"    // [Required] Please store your ca-chain.pem at ./certs directory. The certificate chain you receive by email or on Settings > Developer Credential page (https://dev-app.passkit.io/login)
	 )

var conn *grpc.ClientConn

// ConnectPasskitSdk takes your credentials and establish connection with PassKit SDK.
func main(){
	var err error

	cert, err := ioutil.ReadFile(clientCertFile)

	if err != nil {
		log.Fatalf("could not load certificate file: %v", err)
	}

	key, err := ioutil.ReadFile(clientKeyFile)

	if err != nil {
		log.Fatalf("could not load key file: %v", err)
	}

	ca, err := ioutil.ReadFile(clientCAFile)

	if err != nil {
		log.Fatalf("could not load ca file: %v", err)
	}

	// Generate context object to connect to the server.
	if conn, err = router.NewCertAuthTLSGRPCClient(fmt.Sprintf("%s:%s", gRPCHost, gRPCPort), string(cert), string(key), string(ca)); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connect SDK Success: Established connection to the server successfully.")
}
