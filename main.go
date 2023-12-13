package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/examples/lib/dev"
)

func main() {

	d, err := dev.NewDevice("default")
	if err != nil {
		log.Fatalf("can't create new device : %s", err)
	}
	ble.SetDefaultDevice(d)

	// Scan for specified durantion, or until interrupted by user.
	fmt.Printf("Scanning for %s...\n", 5*time.Second)
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), 5*time.Second))
	chkErr(ble.Scan(ctx, false, advHandler, nil))
}

func advHandler(a ble.Advertisement) {
	if a.Connectable() {
		fmt.Printf("[%s] RSSI: %3d: %s \n", a.Addr(), a.RSSI(), a.LocalName())
	}
}

func chkErr(err error) {
	switch {
	case err == nil:
	case err == context.DeadlineExceeded:
		fmt.Printf("done\n")
	case err == context.Canceled:
		fmt.Printf("canceled\n")
	default:
		log.Fatalf(err.Error())
	}
}
