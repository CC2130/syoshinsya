package main

import "fmt"

type IPAddr [4]byte

// I don't know,,, but I'd like to use pointer
func (i *IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
	fmt.Println("An exercise for Stringer")
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"dnserver": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, &ip)
	}
}
