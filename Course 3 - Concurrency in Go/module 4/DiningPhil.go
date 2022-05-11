package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	num             int
	leftCS, rightCS *ChopS
}

func (p Philo) eat(perm <-chan bool, fin chan<- bool) {
	defer wg.Done()
	for j := 0; j < 3; j++ {
		if ran() {
			p.leftCS.Lock()
			p.rightCS.Lock()
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock()
		}
		<-perm
		fmt.Printf("starting to eat %d\n", p.num+1)

		time.Sleep(10)

		fmt.Printf("finishing  to eat %d\n", p.num+1)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		fin <- true
	}

}

func host(perm chan<- bool, fin <-chan bool) {
	count := 0
	for {
		if count < 2 {
			select {
			case <-fin:
				count--
			case perm <- true:
				count++
			}
		} else {
			<-fin
			count--
		}
	}

}

var wg sync.WaitGroup

func main() {
	CSticks := make([]*ChopS, 5)
	perm := make(chan bool)
	fin := make(chan bool)
	go host(perm, fin)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5]}
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(perm, fin)
	}
	wg.Wait()

}

func ran() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}
