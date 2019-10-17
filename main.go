package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
	"gortc.io/stun"
)

func runStunClient(c *cli.Context) {
	addr := c.String("server")
	sleepDuration := c.Int("sleep")
	rto := c.Int("rto")
	totalCount := c.Int("count")

	// ====================
	if c.Bool("show-interfaces") {
		showNetInterfaces(c.Bool("verbose"))
	}
	fmt.Println("target stun server                  : ", addr)
	fmt.Println("sleep interval for each iteration is: ", sleepDuration)

	// Creating a "connection" to STUN server.
	client, err := stun.Dial("udp", addr)
	if err != nil {
		panic(err)
	}

	client.SetRTO(time.Duration(rto) * time.Millisecond)

	for index := 0; index < totalCount; index++ {
		// Building binding request with random transaction id.
		message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

		// Sending request to STUN server, waiting for response message.
		if err := client.Do(message, func(res stun.Event) {
			if res.Error != nil {
				panic(res.Error)
			}
			if c.Bool("show-stun-message") {
				fmt.Println(serilizeToString(res))
			}

			// Decoding XOR-MAPPED-ADDRESS attribute from message.
			if !c.Bool("ignore-mapped-address") {
				var xorAddr stun.XORMappedAddress
				if err := xorAddr.GetFrom(res.Message); err != nil {
					panic(err)
				}
				fmt.Println("your address on Gateway is, IP:", xorAddr.IP, "Port:", xorAddr.Port)
			}
		}); err != nil {
			panic(err)
		}

		// fmt.Println("sleep ", sleepDuration, "ms")
		time.Sleep(time.Duration(sleepDuration) * time.Millisecond)
	}
}

func showNetInterfaces(verbose bool) {
	fmt.Println("=============== network interfaces ===============")

	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		fmt.Println("net.Interface:", iface)

		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			addrStr := addr.String()
			fmt.Println("    net.Addr: ", addr.Network(), addrStr)

			// Must drop the stuff after the slash in order to convert it to an IP instance
			split := strings.Split(addrStr, "/")
			addrStr0 := split[0]

			// Parse the string to an IP instance
			ip := net.ParseIP(addrStr0)
			if ip.To4() != nil {
				fmt.Println("       ", addrStr0, "is ipv4")
			} else {
				fmt.Println("       ", addrStr0, "is ipv6")
			}
			if verbose {
				fmt.Println("       ", addrStr0, "is interface-local multicast :", ip.IsInterfaceLocalMulticast())
				fmt.Println("       ", addrStr0, "is link-local multicast      :", ip.IsLinkLocalMulticast())
				fmt.Println("       ", addrStr0, "is link-local unicast        :", ip.IsLinkLocalUnicast())
				fmt.Println("       ", addrStr0, "is global unicast            :", ip.IsGlobalUnicast())
				fmt.Println("       ", addrStr0, "is multicast                 :", ip.IsMulticast())
				fmt.Println("       ", addrStr0, "is loopback                  :", ip.IsLoopback())
			}
		}
	}
	fmt.Println("==================================================")
}

func serilizeToString(e stun.Event) string {
	out, err := json.Marshal(e)
	if err != nil {
		return ""
	}

	return string(out)
}

func main() {
	app := cli.NewApp()
	app.Version = "0.0.3"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "server, s",
			Value: "stun1.l.google.com:19302",
			Usage: "STUN server address in format like IP:UDP_Port",
		},
		cli.IntFlag{
			Name:  "sleep, i",
			Value: 0,
			Usage: "Sleep interval for each iteration, in ms",
		},
		cli.IntFlag{
			Name:  "rto, r",
			Value: 300,
			Usage: "RTO timeout for each STUN Message, in ms",
		},
		cli.IntFlag{
			Name:  "count, c",
			Value: 10,
			Usage: "Loop count",
		},
		cli.BoolFlag{
			Name:  "ignore-mapped-address, e",
			Usage: "Ignore show mapped address from STUN message",
		},
		cli.BoolFlag{
			Name:  "show-interfaces, si",
			Usage: "Show network interfaces",
		},
		cli.BoolFlag{
			Name:  "show-stun-message, ss",
			Usage: "Show STUN event message details",
		},
		cli.BoolFlag{
			Name:  "verbose, V",
			Usage: "Verbose output",
		},
	}

	// sort.Sort(cli.FlagsByName(app.Flags))

	app.Action = func(c *cli.Context) error {
		runStunClient(c)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
