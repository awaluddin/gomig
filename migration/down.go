package main

import (
	"fmt"
	// "time"
)

type Payload struct {
	data int
}

func (p *Payload) UploadToS3() error {
	fmt.Println("payload data:", p.data)
	return nil
}

func main() {
	payloads := []Payload{Payload{1}, Payload{2}, Payload{3}, Payload{4}}

	// payloads.
	for _, payload := range payloads {
		// fmt.Println("payload data:", payload.data)
		payload.UploadToS3()
	}

	// time.Sleep(1 * time.Second)
	/*
	   PRINTS:

	   payload data: 4
	   payload data: 4
	   payload data: 4
	   payload data: 4
	*/
}
