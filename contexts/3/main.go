package main

import (
	"context"
	"fmt"
)

type ContextKey string

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ContextKey("token"), "senha")
	BookHotel(ctx)
}

func BookHotel(ctx context.Context) {
	token := ctx.Value(ContextKey("token"))
	fmt.Println(token)
}
