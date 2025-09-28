package main

import (
	d "github.com/synesissoftware/Diagnosticism.Go"

	"fmt"
	"math/rand"
	"sync"
	"time"
)

func show_strips(dg1 *d.DOOMGram, dg2 *d.DOOMGram, dg3 *d.DOOMGram) {

	fmt.Printf("tick (%v): %s %s %s\n", time.Now().Format("15:04:05.000000"), dg1.ToStrip(), dg2.ToStrip(), dg3.ToStrip())

	for {
		select {
		case <-time.After(time.Duration(1) * time.Second):

			fmt.Printf("tick (%v): %s %s %s\n", time.Now().Format("15:04:05.000000"), dg1.ToStrip(), dg2.ToStrip(), dg3.ToStrip())
		}
	}
}

func run_waits_doomgram(dg *d.DOOMGram) {
	r := rand.New(rand.NewSource(12345678))

	for {

		waitTime := r.Int63n(1_0000)

		t_before := time.Now()

		time.Sleep(time.Duration(waitTime) * time.Nanosecond)

		t_after := time.Now()

		d_duration := t_after.Sub(t_before)

		dg.PushEventDuration(d_duration)
	}
}

func run_waits_doomscope(dg *d.DOOMGram) {
	r := rand.New(rand.NewSource(12345678))

	for {

		waitTime := r.Int63n(10_0000)

		d.DOOMScope(dg, nil, func() error {

			time.Sleep(time.Duration(waitTime) * time.Nanosecond)

			return nil
		})
	}
}

func run_waits_doomscope_locked(dg *d.DOOMGram, rwmu *sync.RWMutex) {
	r := rand.New(rand.NewSource(12345678))

	for {

		waitTime := r.Int63n(100_0000)

		d.DOOMScope(dg, rwmu, func() error {

			time.Sleep(time.Duration(waitTime) * time.Nanosecond)

			return nil
		})
	}
}

func main() {

	var dg1 d.DOOMGram
	var dg2 d.DOOMGram
	var dg3 d.DOOMGram
	var rwmu sync.RWMutex

	fmt.Println("start:", dg1.ToStrip(), dg2.ToStrip(), dg3.ToStrip())

	go show_strips(&dg1, &dg2, &dg3)
	go run_waits_doomgram(&dg1)
	go run_waits_doomscope(&dg2)
	go run_waits_doomscope_locked(&dg3, &rwmu)

	time.Sleep(time.Duration(1) * time.Minute)

	fmt.Println("done: ", dg1.ToStrip(), dg2.ToStrip(), dg3.ToStrip())
}
