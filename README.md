# GOTE (go-telnet)
[![Build Status](https://travis-ci.org/morganhein/go-telnet.svg?branch=master)](https://travis-ci.org/morganhein/go-telnet) [![Go Report Card](https://goreportcard.com/badge/github.com/morganhein/go-telnet)](https://goreportcard.com/report/github.com/morganhein/go-telnet)    

[![GoDoc](https://godoc.org/github.com/morganhein/go-telnet?status.svg)](http://godoc.org/github.com/morganhein/go-telnet) [![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/morganhein/go-telnet)


A net.Conn compatible implementation with telnet support.

This is a drop-in replacement for net.Dial that handles telnet negotiaton and other out-of-band messages transparently.

It currently refuses and/or disables all options in a sane manner, except for binary transmission. It disables the Go-Ahead and and ECHO options as well.

Further work needs to be done to implement other telnet options. This is planned, however I have little motivation to do so at the moment.

Current version requires Go1.8 to utilize the os specific writev functions.

### example

```go
package main

import ssh "github.com/wangyide/go-telnet"
import "fmt"

func main() {
    user := "admin"
    password := "yourpasswd"
    ipPort := "ip:23"

        cmds := make([]string, 0)
    cmds = append(cmds, "dis cu")
    cmds = append(cmds, "dir")
    str,_:=ssh.RunCommands(user, password, ipPort, cmds...)
    fmt.Println("===")
    fmt.Println(str)
    fmt.Println("===")
}

```