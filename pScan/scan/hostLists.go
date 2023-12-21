package scan

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

var (
	ErrExists = errors.New("host exists in the list")
	ErrNotExists = errors.New("host not in list")
)

//Hostlists represents a list of hosts to run port scan
type HostLists struct{
	Hosts []string
}


//search searches for hosts in the list
func (hl *HostLists) search(host string) (bool, int) {
	sort.Strings(hl.Hosts) //sorts all hosts alphabetically

	i := sort.SearchStrings(hl.Hosts, host)
	if i < len(hl.Hosts) && hl.Hosts[i] == host {
		return true, i
	}
	return false, -1
}

//Add adds a host to the list
func (hl *HostLists) Add(host string) error {
	if found, _ := hl.search(host); found {
		return fmt.Errorf("%w: %s", ErrExists, host)
	}
	hl.Hosts = append(hl.Hosts, host)
	return nil
}

//Remove deletes a host from the list
func (hl *HostLists) Remove(host string) error {
	if found, i := hl.search(host); found {
		hl.Hosts = append(hl.Hosts[:i],hl.Hosts[i+1:]... )
		return nil
	}
	return fmt.Errorf("%w:%s", ErrNotExists, host)
}

//Load obtains hosts from a hosts file
func (hl *HostLists) Load(hostsFile string) error {
	f, err := os.Open(hostsFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist){
			return nil
		}
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		hl.Hosts = append(hl.Hosts, scanner.Text())
	}
	return nil
}

//Save saves hosts to a host file
func (hl *HostLists) Save(hostsFile string) error {
	output := ""
	for _, h := range hl.Hosts{
		output += fmt.Sprintln(h)
	}
	return os.WriteFile(hostsFile, []byte(output), 0644)
}
