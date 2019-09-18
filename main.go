package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"gortc.io/stun"
)

func main() {
	var err error

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintln(os.Stderr, os.Args[0], "stun1.l.google.com:19302")
	}

	flag.Parse()
	addr := flag.Arg(0)
	if addr == "" {
		addr = "stun1.l.google.com:19302"
	}

	var sleepStr string = flag.Arg(1)
	if sleepStr == "" {
		sleepStr = "0"
	}
	var sleepDuration int
	sleepDuration, err = strconv.Atoi(sleepStr)
	if err != nil {
		fmt.Println("Invalid sleep duration:", sleepStr)
		panic(err)
	}

	var rtoStr string = flag.Arg(2)
	if rtoStr == "" {
		rtoStr = "300"
	}
	var rto int
	rto, err = strconv.Atoi(rtoStr)
	// ====================
	fmt.Println("target stun server                  : ", addr)
	fmt.Println("sleep interval for each iteration is: ", sleepStr)
	show_net_interfaces()
	fmt.Println("==================================================")

	// Creating a "connection" to STUN server.
	c, err := stun.Dial("udp", addr)
	if err != nil {
		panic(err)
	}

	c.SetRTO(time.Duration(rto) * time.Millisecond)

	for index := 0; index < 10; index++ {
		// Building binding request with random transaction id.
		message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

		// Sending request to STUN server, waiting for response message.
		if err := c.Do(message, func(res stun.Event) {
			if res.Error != nil {
				panic(res.Error)
			}
			// Decoding XOR-MAPPED-ADDRESS attribute from message.
			var xorAddr stun.XORMappedAddress
			if err := xorAddr.GetFrom(res.Message); err != nil {
				panic(err)
			}
			fmt.Println("your address on Gateway is, IP:", xorAddr.IP, "Port:", xorAddr.Port)
		}); err != nil {
			panic(err)
		}

		// fmt.Println("sleep ", sleepDuration, "ms")
		time.Sleep(time.Duration(sleepDuration) * time.Millisecond)
	}
}

func show_net_interfaces() {
	fmt.Println("=== interfaces ===")

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
			fmt.Println("       ", addrStr0, "is interface-local multicast :", ip.IsInterfaceLocalMulticast())
			fmt.Println("       ", addrStr0, "is link-local multicast      :", ip.IsLinkLocalMulticast())
			fmt.Println("       ", addrStr0, "is link-local unicast        :", ip.IsLinkLocalUnicast())
			fmt.Println("       ", addrStr0, "is global unicast            :", ip.IsGlobalUnicast())
			fmt.Println("       ", addrStr0, "is multicast                 :", ip.IsMulticast())
			fmt.Println("       ", addrStr0, "is loopback                  :", ip.IsLoopback())
		}
	}
}
