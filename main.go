package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	timer := time.NewTimer(time.Second * 3)
	defer timer.Stop()

	freq := time.Millisecond * 100 / 8
	last := time.Now()
	ticker := time.NewTicker(freq)
	defer ticker.Stop()
	for {
		select {
		case <-timer.C:
			{
				fmt.Println("exit")
				os.Exit(0)
			}
		case t := <-ticker.C:
			{
				d := t.UnixNano()/int64(time.Millisecond) - last.UnixNano()/int64(time.Millisecond)
				if time.Duration(d)*time.Millisecond > 2*freq {
					panic(d)
				}
				last = t
			}
		}
	}

}
