package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, errCtx := errgroup.WithContext(ctx)

	if _, err := InitApp(errCtx); err != nil {
		panic(err)
	}

	group.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		select {
		case <-quit:
			cancel()
		case <-errCtx.Done():
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("Get errors: ", err)
	} else {
		fmt.Println("end")
	}
}
