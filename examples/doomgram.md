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
```

## Execution

When executed, it gives output (to the standard output stream) along the lines of

```
$ go run examples/doomgram.go
start: ____________ ____________ ____________
tick (11:21:36.537733): ___aa_______ ____a_______ ____________
tick (11:21:37.538065): _bdee_______ _abdec______ ___acda_____
tick (11:21:38.538103): _bdfe_______ _abded______ ___bcda_____
tick (11:21:39.538122): _bdffa______ _abded______ __abcdb_____
tick (11:21:40.538149): _bdffb______ _acded______ __abcdb_____
tick (11:21:41.538177): _bdffba_____ _acdeda_____ __abcdb_____
tick (11:21:42.538201): _bdffcbaa___ _acdedbaa___ __abdebaa___
tick (11:21:43.538250): _bdffcbaa___ _acdfdbaa___ __abdebaa___
tick (11:21:44.538274): _beffcbaa___ _acdfdbaa___ __abdebaa___
tick (11:21:45.538293): _beffcbaa___ _acefdbaa___ __acdebaa___
tick (11:21:46.538341): _ceffcbaa___ _acefdbaa___ __acdebaa___
tick (11:21:47.538370): _ceffcbaa___ _acefdbaa___ __acdebaa___
tick (11:21:48.538401): _ceffcbaa___ _acefdbaa___ __acdecaa___
tick (11:21:49.538432): _ceffcbaa___ _acefdbaa___ __acdecaa___
tick (11:21:50.538454): _ceffcbaa___ _acefdbaa___ __acdecaa___
tick (11:21:51.538479): _cegfcbaa___ _acefdbaa___ __acdecaa___
tick (11:21:52.538501): _cegfcbaa___ _acefdbaa___ __acdecaa___
tick (11:21:53.538532): _cegfcbaa___ _acefebaa___ __acdecaa___
tick (11:21:54.538561): _cegfcbaa___ _acefebaa___ __acdecaa___
tick (11:21:55.538598): _cegfcbaa___ _acefebaa___ __acdecaa___
tick (11:21:56.538652): _cegfccba___ _acefecba___ __acdecba___
tick (11:21:57.538678): _cegfccba___ _acefecba___ __acdecba___
tick (11:21:58.538967): _cegfccba___ _acefecba___ __acdecba___
tick (11:21:59.538998): _cegfdcba___ _acefecba___ __acdecba___
tick (11:22:00.539017): _cegfdcba___ _acefecba___ __acdecba___
tick (11:22:01.539043): _cegfdcba___ _acefecba___ __acdecba___
tick (11:22:02.539065): _cegfdcba___ _acefecba___ __acdecba___
tick (11:22:03.539081): _ceggdcba___ _acefecba___ __bcdecba___
tick (11:22:04.539112): _ceggdcba___ _acefecba___ __bcdedba___
tick (11:22:05.539168): _ceggdcba___ _acefecba___ __bcdedba___
tick (11:22:06.539207): _ceggdcba___ _acefecba___ __bcdedba___
tick (11:22:07.539231): _ceggdcba___ _acefecba___ __bcdedba___
tick (11:22:08.539250): _ceggdcba___ _acefecba___ __bcdedba___
tick (11:22:09.539266): _ceggdcba___ _acefecba___ __bcdedba___
tick (11:22:10.539283): _ceggdcba___ _acefecba___ __bcdedba___
tick (11:22:11.539304): _ceggdcba___ _acefecba___ __bcdedba___
tick (11:22:12.539348): _ceggdcba___ _acefecba___ __bcdedba___
tick (11:22:13.539365): _ceggdcba___ _acefecba___ __bcdedba___
tick (11:22:14.539383): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:15.539398): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:16.539419): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:17.539439): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:18.539470): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:19.539495): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:20.539512): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:21.539531): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:22.539552): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:23.539571): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:24.539594): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:25.539611): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:26.539633): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:27.539656): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:28.539672): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:29.539694): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:30.539716): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:31.539739): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:32.539753): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:33.539779): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:34.539801): _ceggdcba___ _adefecba___ __bcdedba___
tick (11:22:35.539823): _ceggdcba___ _adefecba___ __bcdedba___
done:  _ceggdcba___ _bdefecba___ __bcdfdba___
```


<!-- ########################### end of file ########################### -->

