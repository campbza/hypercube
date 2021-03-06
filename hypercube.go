
package main

import "fmt"
import "math/rand"
import "time"
import "strconv"

//packet that gives source and destination information
type packet struct {
	srce int
	dest int
}

func make_bundle(width int) (bs []chan packet) {
	//initialize a slice of channels
	bs = make([]chan packet, width)
	for i := 0; i < width; i++ {
		bs[i] = make(chan packet)
	}
	return
}

func make_injector(chs []chan packet) {
	//takes a slice of channels, generates random source and destination nodes in the 
	//hypercube and sends corresponding packets to channels
	n := len(chs)
	//goroutine executing this process
	go func () {
		for {
			s := rand.Intn(n)
			t := rand.Intn(n)
			// send packet{s,t} to channel s
			fmt.Printf("Injection of packet {%v --> %v}.\n", s, t)

			chs[s] <- packet{s, t}
			time.Sleep(time.Duration(2)*time.Second)
		}
	} ()
}

func node(index int, in chan packet, outs [](chan packet), report chan packet) {
	for {
		// receive from in channel, assign its value to p
		p := <- in
		//if we've reached the destination node in the hypercube
		if p.dest == index {
			fmt.Printf("Node %v received packet from %v.\n", index, p.srce)
			report <- p
		//here we route the packet along the shortest 'legal' path to destination node
		} else {
			var b uint = 0
			var l int = len(outs)
			var r uint = uint(l)
			for b < r {
				if (index & (1 << b)) != (p.dest & (1 << b)) {
					fmt.Printf("Node %v forwarding packet {%v --> %v}.\n", index, p.srce, p.dest)
					outs[index ^ (1 << b)] <- p
					break
				} else {
					b = b + 1
				}
			}
		}
	}
}

func make_hypercube(chs []chan packet, rep chan packet) {
	//hypercube with n nodes
	n := len(chs)
	i := 0
	for i < n {
		//concurrently run our node function
		go node(i, chs[i], chs, rep)
		i = i + 1
	}
}

func main() {
	var s string = ""
	var n int = 0
	fmt.Println("Enter a power of two: ")
	fmt.Scanln(&s)
	n,_ = strconv.Atoi(s)

	chs := make_bundle(n)
	rep := make(chan packet)
	make_injector(chs)
	make_hypercube(chs, rep)
	for x := range rep {
		fmt.Printf("Packet {%v --> %v} routed.\n",x.srce,x.dest)
	}
}
