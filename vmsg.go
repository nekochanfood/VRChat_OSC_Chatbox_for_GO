package main

import (
	"flag"

	"github.com/hypebeast/go-osc/osc"
)

func main() {
	var (
		msg  = flag.String("msg", "null!", "the message")
		ip   = flag.String("ip", "127.0.0.1", "ip address")
		port = flag.Int("port", 9000, "port")
	)
	flag.Parse()
	client := osc.NewClient(*ip, *port)
	message := osc.NewMessage("/chatbox/input")
	message.Append(*msg, true)
	client.Send(message)
}
