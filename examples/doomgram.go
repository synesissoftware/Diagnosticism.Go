package main

import (
	d "github.com/synesissoftware/Diagnosticism.Go"

	"fmt"
	"math/rand"
	"time"
)

func show_strip(dg *d.DOOMGram) {

	fmt.Printf("tick (%v): %s\n", time.Now().Format("15:04:05.000000"), dg.ToStrip())

	for {
		select {
		case <-time.After(time.Duration(1) * time.Second):

			fmt.Printf("tick (%v): %s\n", time.Now().Format("15:04:05.000000"), dg.ToStrip())
		}
	}
}

func run_waits(dg *d.DOOMGram) {
	r := rand.New(rand.NewSource(12345678))

	for {

		waitTime := r.Int63n(10_0000)

		t_before := time.Now()

		time.Sleep(time.Duration(waitTime) * time.Nanosecond)

		t_after := time.Now()

		d_duration := t_after.Sub(t_before)

		dg.PushEventDuration(d_duration)
	}
}

func main() {

	var dg d.DOOMGram

	fmt.Println("start: ", dg.ToStrip())

	go show_strip(&dg)
	go run_waits(&dg)

	time.Sleep(time.Duration(1) * time.Minute)

	fmt.Println("done: ", dg.ToStrip())
}
