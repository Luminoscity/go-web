/**
 * Author: Tim Ambrose
 * 2025
 */

package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	number  int
	done    bool
	leftCS  *ChopS
	rightCS *ChopS
}

func (p *Philo) eat(mayEat chan bool, done chan bool, wait *sync.WaitGroup) {
	defer wait.Done()
	for i := range 3 {
		<-mayEat
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("Starting to eat %v (%v)\n", p.number, i+1)
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Finishing eating %v (%v)\n", p.number, i+1)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		done <- true
	}
	p.done = true
}

const nPhilo int = 5

func incrementWhoEating(eating *[]int, philos []*Philo, idx int) error {
	other := (*eating)[1-idx]
	start := ((*eating)[idx] + 1) % nPhilo
	current := start

	for range nPhilo {
		if current != other && !philos[current].done {
			(*eating)[idx] = current
			return nil
		}
		current = (current + 1) % nPhilo
		if current == start {
			break
		}
	}

	return errors.New("All philosphers are done eating")
}

func host(philos []*Philo, mayEat []chan bool, done []chan bool, hostWait *sync.WaitGroup) {
	eating := []int{0, 2}

	for who := range nPhilo {
		go philos[who].eat(mayEat[who], done[who], hostWait)
	}

	for _, who := range eating {
		mayEat[who] <- true
	}
	// fmt.Println("Eating:", eating[0]+1, eating[1]+1)

	var allDone error = nil
	for allDone == nil {
		select {
		case <-done[eating[0]]:
			allDone = incrementWhoEating(&eating, philos, 0)
			if allDone == nil {
				mayEat[eating[0]] <- true
				// fmt.Println("Eating:", eating[0]+1, eating[1]+1)
			}
		case <-done[eating[1]]:
			allDone = incrementWhoEating(&eating, philos, 1)
			if allDone == nil {
				mayEat[eating[1]] <- true
				// fmt.Println("Eating:", eating[0]+1, eating[1]+1)
			}
		}
	}
}

func main() {

	mayEat := make([]chan bool, nPhilo)
	done := make([]chan bool, nPhilo)

	for i := range nPhilo {
		mayEat[i] = make(chan bool, 1)
		done[i] = make(chan bool, 1)
	}

	CSticks := make([]*ChopS, nPhilo)
	for i := range nPhilo {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, nPhilo)
	for i := range nPhilo {
		philos[i] = &Philo{i + 1, false, CSticks[i], CSticks[(i+1)%nPhilo]}
	}

	var hostWait sync.WaitGroup
	hostWait.Add(nPhilo)
	go host(philos, mayEat, done, &hostWait)
	hostWait.Wait()
}
