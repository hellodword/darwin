package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	freq := flag.Int64("freq", 13, "")
	flag.Parse()

	timer := time.NewTimer(time.Second * 3)
	defer timer.Stop()

	last := time.Now()
	ticker := time.NewTicker(time.Duration(*freq) * time.Millisecond)
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
				if d-*freq > 20 {
					panic(d)
				}
				last = t
			}
		}
	}

}
