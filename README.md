# goami2: Simple Asterisk Manager Interface (AMI) library

Usage example:

```go
import "github.com/staskobzar/goami2"
//.....more code

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
  defer client.Close()
	for {
    select {
    case msg := <-ch:
      log.Printf("Got event: %s\n", msg.Event())
    }
  }
}
```

# API
TODO...
