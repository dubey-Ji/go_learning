package main

import (
	"context"
	"fmt"
)

func processData(ctx context.Context) {
	// ** Context value recieve kar rha h
	if usrName, ok := ctx.Value("username").(string); ok {
		fmt.Println("Process data for user", usrName)
	} else {
		fmt.Println("User not found in context")
	}
}

func main() {
	ctx := context.Background()
	// fmt.Println(ctx)

	// ** Context value send karege

	contextWithUsername := context.WithValue(ctx, "username", "ashutosh")
	processData(contextWithUsername)

}
