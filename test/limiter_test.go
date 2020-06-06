package test

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"testing"
	"time"
)

func TestLimiter(t *testing.T)  {
	limiter := rate.NewLimiter(rate.Every(time.Millisecond), 10000)
	now := time.Now()
	count := 0
	for {
		if time.Since(now) > time.Second {
			break
		}

		if err := limiter.Wait(context.Background()); err != nil {
			panic(err)
		}
		count += 1
		log.Printf("get a token:%v", time.Now())
	}
	log.Printf("total requests:%v,%vs", count, time.Since(now).Seconds())
}
