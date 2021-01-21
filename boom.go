package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func boom() {
	n := 1440
	if s := os.Getenv("LIVE_FOR"); s != "" {
		v, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			log.Printf("unable to parse LIVE_FOR=%s: %s", s, err)
		} else {
			n = int(v)
		}
	}

	t := time.NewTicker(time.Minute * time.Duration(n))
	<-t.C

	log.Printf("\n")
	log.Printf("########   #######   #######  ##     ##\n")
	log.Printf("##     ## ##     ## ##     ## ###   ###\n")
	log.Printf("##     ## ##     ## ##     ## #### ####\n")
	log.Printf("########  ##     ## ##     ## ## ### ##\n")
	log.Printf("##     ## ##     ## ##     ## ##     ##\n")
	log.Printf("##     ## ##     ## ##     ## ##     ##\n")
	log.Printf("########   #######   #######  ##     ##\n")
	log.Printf(".......................................\n")
	log.Printf("time expired; shutting down!\n")
	os.Exit(0)
}