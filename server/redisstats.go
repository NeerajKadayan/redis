package server

import "github.com/company-tests/tazapay/test/types"

func processKeys(input []string) types.Stats {
	// process each key concurrently using go routines.
	listener := make(chan types.Stats, 5)

	// worker wait groups.
	// run go routines to fetch each key.
	// NOTE : close the channel after use.

	go GetCPU(listener)
	/*CPU := */ <-listener
	// similarly repeat for all REDIS keys

	return types.Stats{}
}

func GetCPU(l chan types.Stats) types.Stats {
	v := types.Stats{}
	// iterate and store the value in types.Stats{}

	l <- v
	return types.Stats{}
}
