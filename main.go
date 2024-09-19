package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//eeae e

	go listen(ctx)

	<-ctx.Done()
}

func listen(ctx context.Context) {
	for {
		var input string
		println("> [" + time.Now().Format("2006-01-02 15:04:05") + "] Please enter something:")
		_, _ = fmt.Scanln(&input)
		println("[" + time.Now().Format("2006-01-02 15:04:05") + "] You entered: " + input)

		if input == "exit" {
			ctx.Done()
			break
		}
	}
}
