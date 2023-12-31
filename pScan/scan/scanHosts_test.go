package scan_test

import (
	"fmt"
	"net"
	"strconv"
	"testing"

	"github.com/Micah-Shallom/modules/scan"
)

func TestStateString(t *testing.T) {
	ps := scan.PortState{}
	fmt.Println(t,ps)
	// if ps.tcpOpen.String() != "closed" {
	// 	t.Errorf("Expected %q, got %q instead\n", "closed", ps.tcpOpen.String())
	// }
	// ps.tcpOpen = true
	// if ps.tcpOpen.String() != "open" {
	// 	t.Errorf("Expected %q, got %q instead\n", "open", ps.tcpOpen.String())
	// }
}

func TestRunHostFound(t *testing.T) {
	testCases := []struct {
		name        string
		expectState string
	}{
		{"OpenPort", "open"},
		{"ClosedPort", "closed"},
	}
	host := "localhost"
	hl := &scan.HostLists{}
	hl.Add(host)

	ports := []int{}
	//Init ports, 1 open, 1 closed
	for _, tc := range testCases {
		ln, err := net.Listen("tcp", net.JoinHostPort(host, "0")) //using a zero makes use of the available ports in the host machine
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		_, portStr, err := net.SplitHostPort(ln.Addr().String()) //since the ports were fetched from the hostmachine, we use the SplitHostPort to get the ports
		if err != nil {
			t.Fatal(err)
		}
		port, err := strconv.Atoi(portStr)
		if err != nil {
			t.Fatal(err)
		}
		ports = append(ports, port)
		if tc.name == "ClosedPort" {
			ln.Close()
		}
	}
	res := scan.Run(hl, ports)

	//Verify results for HostFound test
	if len(res) != 1 {
		t.Fatalf("Expected 1 results, got %d instead\n", len(res))
	}
	if res[0].Host != host {
		t.Errorf("Expected host %q, got %q instead\n", host, res[0].Host)
	}
	if res[0].NotFound {
		t.Errorf("Expected host %q to be found\n", host)
	}
	if len(res[0].PortStates) != 2 {
		t.Fatalf("Expected 2 port states,got %d instead\n", len(res[0].PortStates))
	}
	for i, tc := range testCases {
		if res[0].PortStates[i].Port != ports[i] {
			t.Errorf("Expected port %d, got %d instead\n", ports[0], res[0].PortStates[i].Port)
		}
		_ = tc
		// if res[0].PortStates[i].tcpOpen.String() != tc.expectState {
		// 	t.Errorf("Expected port %d to be %s\n", ports[i], tc.expectState)
		// }
	}
}

func TestRunHostNotFound(t *testing.T) {
	host := "389.389.389.389"
	hl := &scan.HostLists{}
	hl.Add(host)

	res := scan.Run(hl, []int{})

	//verify results for HostNotFound test
	if len(res) != 1 {
		t.Fatalf("Expected 1 results, got %d instead\n", len(res))
	}
	if res[0].Host != host {
		t.Errorf("Expected host %q, got %q instead\n", host, res[0].Host)
	}
	if !res[0].NotFound {
		t.Errorf("Expected host %q NOT to be found\n", host)
	}
	if len(res[0].PortStates) != 0 {
		t.Fatalf("Expected 0 port states, got %d instead\n", len(res[0].PortStates))
	}
}
