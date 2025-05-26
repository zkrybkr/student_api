package main

import (
	"context"
	"fmt"
	c "studentApi/config"
)
func main() {
	db := c.InitDB()
    defer db.Close()

    // Test sorgusu
    var now string
    err := db.QueryRow(context.Background(), "SELECT NOW()").Scan(&now)
    if err != nil {
        panic(err)
    }

    fmt.Println("🕒 Current time from DB:", now)
}