package apollo_client_test

import (
	"fmt"
	"github.com/reaperhero/agollo"
	"os"
	"testing"
)

var (
	configServerURL  = []string{"http://192.168.50.24:8080"}
	configAppid      = "testId"
	configNameSpaces = []string{"application", "mysql"}
)

func TestWatchApollo(t *testing.T) {
	client, _ := agollo.NewAgolloOnce(
		configServerURL,
		configAppid,
		agollo.WithNameSpaces(configNameSpaces),
		agollo.WithLogger(agollo.NewLogger(agollo.LoggerWriter(os.Stdout))),
		agollo.AutoFetchOnCacheMiss(),
		agollo.FailTolerantOnBackupExists(),
	)
	// 一次性获取
	for n, v := range client.GetAllNameSpaceValue() {
		fmt.Println(n, v)
	}

	// 监听变化
	errCh := client.Start()
	respCh := client.Watch()
	for {
		select {
		case err := <-errCh:
			fmt.Println(err)
		case resp :=<-respCh:
			fmt.Println(resp)
		}
	}
}

