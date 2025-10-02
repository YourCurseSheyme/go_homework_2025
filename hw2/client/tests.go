package client

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/YourCurseSheyme/go_homework_2025/hw2/json"
)

func TestVersion(client *Client, ctx context.Context) {
	fmt.Println("> GET /version")
	value, code, err := client.RequestVersion(ctx)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("code:", code)
		json.PrintJSON(os.Stdout, value)
	}
	fmt.Println("+-----------")
}

func TestDecode(client *Client, ctx context.Context) {
	fmt.Println("> POST /decode")
	arg := "Судьба - это не то, что вы делаете, это то, что с вами происходит"
	argB64 := base64.StdEncoding.EncodeToString([]byte(arg))
	value, code, err := client.RequestDecode(ctx, argB64)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("code:", code)
		json.PrintJSON(os.Stdout, value)
	}
	fmt.Println("+-----------")
}

func TestHardOp(client *Client, ctx context.Context) {
	fmt.Println("> GET /hard-op")
	timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	start := time.Now()
	value, _, err := client.RequestHardOp(timeout)
	elapsed := time.Since(start)
	if err != nil {
		fmt.Println("success:", false)
		fmt.Println("error:", err)
		fmt.Println("elapsed:", elapsed.Truncate(time.Millisecond))
	} else {
		fmt.Println("success:", true)
		json.PrintJSON(os.Stdout, value)
	}
	fmt.Println("+-----------")
}

func Demo() {
	fmt.Println("> Homework №2")

	hostUrl := "http://localhost:8080"
	client, err := NewClient(hostUrl)
	if err != nil {
		fmt.Printf("NewClient error: %v\n", err)
		return
	}
	root := context.Background()
	TestVersion(client, root)
	TestDecode(client, root)
	TestHardOp(client, root)

	fmt.Println("> Test have been done")
}
