package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Go program can handle *Unix signal*
// Unix signal - controlling terminal of a running process
// Ex : Ctrl-C send a SIGINT - cause the process to terminate
//      shell command "kill" send a SIGTERM - cause the process to terminate

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Register signal channel to received any notification from *Unix signal* - SIGNIT, SIGTERM
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // Similar to this sigs <- SIGINT or sigs <- SIGTERM

	go func() {
		sig := <-sigs
		fmt.Println("Got a signal", sig)

		done <- true
	}()

	fmt.Println("Running the program ...")
	<-done // Wait until got a *Unix signal*
	fmt.Println("Exiting the program ...")
}
