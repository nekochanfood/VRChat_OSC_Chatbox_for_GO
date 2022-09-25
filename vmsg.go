package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/hypebeast/go-osc/osc"
)

func main() {
	var (
		msg             = flag.String("msg", "message", "A message")   // -msg "A message"
		ip              = flag.String("ip", "127.0.0.1", "ip address") // -ip 127.0.0.1
		port            = flag.Int("port", 9000, "port")               // -port 0000
		continuous_mode = flag.Bool("c", false, "continuous mode")     // -c continuous mode
	)
	flag.Parse()

	client := osc.NewClient(*ip, *port)

	if *continuous_mode {
		continuous(*client)
	} else {
		message := osc.NewMessage("/chatbox/input")
		message.Append(*msg, true)
		client.Send(message)
	}
}

func continuous(client osc.Client) {
	print("Ctrl + C or type \"/exit\" to exit\n")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("> ")
		scanner.Scan()
		if scanner.Text() == "/exit" {
			print("exit")
			return
		}
		message := osc.NewMessage("/chatbox/input")
		message.Append(scanner.Text(), true)
		client.Send(message)
	}
}
