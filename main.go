package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"

	"github.com/f-secure-foundry/tamago/pi"
	_ "github.com/f-secure-foundry/tamago/pi/pi2"
)

func main() {
	sleep := 10 * time.Second

	log.Println("Hello World!")

	log.Println("-- rng -------------------------------------------------------------")

	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	log.Printf("random bytes %s", hex.EncodeToString(b))

	size := 32

	for i := 0; i < 10; i++ {
		rng := make([]byte, size)
		rand.Read(rng)
		log.Printf("%x", rng)
	}

	count := 1000
	start := time.Now()

	for i := 0; i < count; i++ {
		rng := make([]byte, size)
		rand.Read(rng)
	}

	log.Printf("retrieved %d random bytes in %s", size*count, time.Since(start))

	log.Println("-- timer -------------------------------------------------------------")

	t := time.NewTimer(sleep)
	log.Printf("waking up timer after %v", sleep)

	start = time.Now()

	for now := range t.C {
		log.Printf("woke up at %d (%v)", now.Nanosecond(), now.Sub(start))
		break
	}

	pi.InitLED()
	for {
		pi.WatchdogStart(0xFFFFF)
		for {
			ra := pi.WatchdogRemaining()
			if ra < 0xFC2F6 {
				pi.ToggleLED()
				break
			}
		}
	}
}
