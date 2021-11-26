package apollo_client_test

import (
	"fmt"
	"github.com/reaperhero/agollo"
	"os"
	"testing"
)

var (
	configServerURL  = []string{"http://192.168.50.24:8080"}
	configAppid      = "id"
	configNameSpaces = []string{"namespace"}
)

func TestCommonApollo(t *testing.T) {
	client,_ := agollo.NewAgolloOnce(
		configServerURL,
		configAppid,
		agollo.WithNameSpaces(configNameSpaces),
		agollo.WithLogger(agollo.NewLogger(agollo.LoggerWriter(os.Stdout))),
		agollo.AutoFetchOnCacheMiss(),
		agollo.FailTolerantOnBackupExists(),
	)
	for n, v := range client.GetAllNameSpaceValue() {
		fmt.Println(n,v)
	}
}
