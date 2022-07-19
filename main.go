package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Println("Doing something!")
}

func main() {
	//ctx := context.TODO()
	ctx := context.Background()
	doSomething(ctx)
}
