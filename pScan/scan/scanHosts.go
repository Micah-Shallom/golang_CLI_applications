package scan

import (
	"fmt"
	"net"
	"time"
)

type PortState struct {
	Port int
	TCPOpen state
	UDPOpen state
}

type state bool

//String converts the boolean value of state to a human readable string

func (s state) String() string {
	if s {
		return "opened"
	}
	return "closed"
}

// scanPort performs a port scan on a single TCP port
func scanPort(host string, port int) PortState {
	p := PortState{
		Port: port,
		//we didnt have to provide the value of Open as the zero value of a boolean is false
	}

	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))//used instead of concatenating strings in order to cater for a case of an ipV6 host address

	//tcp connection to host address
	tcpConn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return p
	}
	tcpConn.Close()
	p.TCPOpen = true

	//udp connection to host address
	udpConn, err := net.DialTimeout("udp", address, 1*time.Second)
	if err != nil {
		return p
	}
	udpConn.Close()
	p.UDPOpen = true
	fmt.Println(p)
	return p
}

//Results represents the scan results for a single host

type Results struct {
	Host		string
	NotFound	bool
	PortStates	[]PortState
}

//Run performs a port scan on the hosts list
func Run(hl *HostLists, ports []int) []Results{
	// res := make([]Results, len(hl.Hosts))
	res := []Results{}
	for _, host := range hl.Hosts{
		r := Results{
			Host: host,
		}
		if _, err := net.LookupHost(host); err != nil {
			r.NotFound = true
			res = append(res, r)
			continue
		}
		for _, port := range ports{
			r.PortStates = append(r.PortStates, scanPort(host,port))
		}
		res = append(res, r)
	}
	return res
}