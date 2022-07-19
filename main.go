package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))

	anotherCtx := context.WithValue(ctx, "myKey", "anotherValue")
	doAnother(anotherCtx)
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))

}

func doAnother(ctx context.Context) {
	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value("myKey"))

}

func doSomething2(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan int)
	go doAnother2(ctx, ch)

	for num := 1; num <= 3; num++ {
		ch <- num
	}

	cancel()

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("doSomething: finished\n")

}

func doAnother2(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Printf("doAnother: finished\n")
			return

		case num := <-ch:
			fmt.Printf("doAnother: %d\n", num)
		}

	}
}

func ()  {
	
}

func main() {
	//ctx := context.TODO()
	ctx := context.Background()

	ctx = context.WithValue(ctx, "myKey", "myValue")
	doSomething2(ctx)
}
