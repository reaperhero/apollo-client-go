# Agollo - Go Client for Apollo

[![Build Status](https://travis-ci.org/shima-park/agollo.svg?branch=master)](https://travis-ci.org/reaperhero/apollo-client-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/reaperhero/apollo-client-go)](https://goreportcard.com/report/github.com/reaperhero/apollo-client-go)
[![codebeat badge](https://codebeat.co/badges/bc2009d6-84f1-4f11-803e-fc571a12a1c0)](https://codebeat.co/projects/github-com-shima-park-agollo-master)
[![golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org)
[![GoDoc](http://godoc.org/github.com/reaperhero/apollo-client-go?status.svg)](http://godoc.org/github.com/reaperhero/apollo-client-go/agollo)
[![GitHub release](https://img.shields.io/github/release/shima-park/agollo.svg)](https://github.com/reaperhero/apollo-client-go/releases)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

携程Apollo Golang版客户端

## 快速开始
### 获取安装
```
go get -u github.com/reaperhero/apollo-client-go
```

## Features
* 实时同步配置,配置改动监听
* 配置文件容灾
* 支持多namespace, cluster
* 客户端SLB
* 支持通过 ACCESSKEY_SECRET 来实现 client 安全访问
* 支持自定义签名认证

## 示例

### 读取配置
此示例场景适用于程序启动时读取一次。不会额外启动goroutine同步配置
```
package main

import (
	"fmt"

	agollo "github.com/reaperhero/apollo-client-go"
)

var (
	configServerURL  = []string{"http://192.168.50.24:8080"}
	configAppid      = "testId"
	configNameSpaces = []string{"application", "mysql"}
)

func main() {
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
```
