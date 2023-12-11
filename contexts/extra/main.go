package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type Response struct {
	Payload int
	Error   error
}

func main() {
	start := time.Now()
	val, err := fetchUserData(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result: ", val)
	fmt.Println("took: ", time.Since(start))

}

func fetchUserData(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()
	
	respch := make(chan Response)
	
	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respch <- Response{val, err}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, errors.New("timeout")
		case resp := <-respch:
			return resp.Payload, resp.Error
		}
	}
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 500)

	return 666, nil
}
