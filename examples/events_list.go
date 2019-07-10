// Package main provides ...
package main

import (
	"fmt"
	"github.com/staskobzar/goami2"
	"log"
	"net"
	"sort"
	"strings"
)

var list map[string]int
var listLen int

func printList() {
	fmt.Printf("\u001b[1000D")
	fmt.Printf("\u001b[%dA", int(listLen))
	ll := lineLen()
	keys := getKeys()
	sort.Strings(keys)
	for _, event := range keys {
		fmt.Printf("\u001b[2K%s:", event)
		fmt.Printf("%s%d\n", strings.Repeat(".", ll-len(event)+1), list[event])
	}
	listLen = len(list)
}

func lineLen() int {
	ll := 0
	for event := range list {
		if len(event) > ll {
			ll = len(event)
		}
	}
	return ll
}

func getKeys() []string {
	keys := make([]string, 0, len(list))
	for k := range list {
		keys = append(keys, k)
	}
	return keys
}

func consumer(ch <-chan *goami2.AMIMsg) {
	for {
		select {
		case msg := <-ch:
			val, ok := msg.Event()
			if ok {
				list[val]++
				printList()
			}
		}
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5038")
	if err != nil {
		log.Fatalln(err)
	}
	client, err := goami2.NewClient(conn)
	if err != nil {
		log.Fatalln(err)
	}
	err = client.Login("username", "secret")
	if err != nil {
		log.Fatalln(err)
	}
	ch, err := client.AnyEvent()
	if err != nil {
		log.Fatalln(err)
	}
	list = make(map[string]int)
	fmt.Printf("\n\n")
	printList()
	consumer(ch)
}
