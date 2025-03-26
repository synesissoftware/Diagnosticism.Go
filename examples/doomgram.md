# Diagnosticism.Go Example - **doomgram**

## Summary

Example illustrating use of `DOOMGram` type, via a simple program that does random waits in one goroutine and simple logging in another, running for a minute.

## Source

``` Go
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
```

## Execution

When executed, it gives output (to the standard output stream) along the lines of

```
$ go run examples/doomgram.go
start:  ____________
tick (08:19:28.709291): ____a_______
tick (08:19:29.709681): _abded______
tick (08:19:30.709794): _abded______
tick (08:19:31.709866): _abded______
tick (08:19:32.709934): _abdee______
tick (08:19:33.710035): _abdee______
tick (08:19:34.710129): _abdee______
tick (08:19:35.710215): _abdee______
tick (08:19:36.710294): _acdfe______
tick (08:19:37.710346): _acdfeaa____
tick (08:19:38.710408): _acefeaa____
tick (08:19:39.710475): _acefeaa____
tick (08:19:40.710532): _acefeaa____
tick (08:19:41.710631): _acefeaa____
tick (08:19:42.710680): _acefeaa____
tick (08:19:43.710720): _acefeaa____
tick (08:19:44.710777): _acefeaa____
tick (08:19:45.710852): _acefeaa____
tick (08:19:46.710922): _acefeaa____
tick (08:19:47.711046): _acefeaa____
tick (08:19:48.711122): _acefeaa____
tick (08:19:49.711217): _acefeaa____
tick (08:19:50.711259): _acefeaa____
tick (08:19:51.711309): _acefeaa____
tick (08:19:52.711351): _acefeaa____
tick (08:19:53.711434): _acefeaa____
tick (08:19:54.711481): _acefeaa____
tick (08:19:55.711523): _acefeba____
tick (08:19:56.711569): _acefeba____
tick (08:19:57.711633): _acefeba____
tick (08:19:58.711679): _acefeba____
tick (08:19:59.711770): _acefeba____
tick (08:20:00.711811): _acefeba____
tick (08:20:01.711855): _acefeba____
tick (08:20:02.711898): _acefeba____
tick (08:20:03.711947): _acefeba____
tick (08:20:04.712011): _acefeba____
tick (08:20:05.712058): _acefeba____
tick (08:20:06.712126): _aceffba____
tick (08:20:07.712178): _aceffba____
tick (08:20:08.712221): _aceffba____
tick (08:20:09.712268): _aceffba____
tick (08:20:10.712312): _aceffba____
tick (08:20:11.712356): _aceffba____
tick (08:20:12.712398): _aceffba____
tick (08:20:13.712443): _aceffba____
tick (08:20:14.712485): _aceffba____
tick (08:20:15.712554): _aceffba____
tick (08:20:16.712650): _aceffba____
tick (08:20:17.712695): _aceffba____
tick (08:20:18.712775): _aceffba____
tick (08:20:19.712827): _aceffba____
tick (08:20:20.712872): _aceffba____
tick (08:20:21.712928): _aceffba____
tick (08:20:22.712966): _aceffba____
tick (08:20:23.713023): _aceffba____
tick (08:20:24.713066): _aceffba____
tick (08:20:25.713120): _aceffba____
tick (08:20:26.713210): _aceffba____
tick (08:20:27.713272): _aceffba____
done:  _aceffba____
```


<!-- ########################### end of file ########################### -->

