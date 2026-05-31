# In-Memory Cache (Go)

A simple in-memory key-value cache library written in Go.  
Supports storing, retrieving, and deleting values of any type with **TTL (time-to-live)** and **concurrent-safe access**.

---

## Features

- Set key-value pairs in memory  
- Get values with existence check  
- Delete keys from cache  
- TTL support (auto-expiration of values)  
- Thread-safe (safe for concurrent goroutines)

---

## Installation

Install the library using `go get`:

```bash
go get github.com/LiaVoitenko/golang_basic

---

## Example Usage

```
package main

import (
	"fmt"
	"time"

	"github.com/LiaVoitenko/golang_basic/cache"
)

func main() {
	c := cache.New()

	// set value with TTL
	c.Set("userId", 42, time.Second*5)

	// get value
	userId, ok := c.Get("userId")
	fmt.Println(userId, ok)

	// wait for expiration
	time.Sleep(time.Second * 6)

	// now expired
	userId2, ok2 := c.Get("userId")
	fmt.Println(userId2, ok2)
}

```

---
