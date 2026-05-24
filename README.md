# In-Memory Cache (Go)

A simple in-memory key-value cache library written in Go.  
Supports storing, retrieving, and deleting values of any type.

---

## Features

- Set key-value pairs in memory
- Get values with existence check
- Delete keys from cache
- Supports any data type

---

## Installation

Install the library using `go get`:

```bash
go get github.com/LiaVoitenko/golang_basic
```

---


## API Reference

New() *Cache — creates a new cache instance

Set(key string, value any) — stores a value in cache by key

Get(key string) (any, bool) — retrieves value by key

Delete(key string) — deletes value from cache

---

## Example Usage

```
package main

import (
	"fmt"

	"github.com/LiaVoitenko/golang_basic/cache"
)

func main() {
	c := cache.New()

	c.Set("userId", 42)
	userId, ok := c.Get("userId")
	c.Set("userId", 43)

	fmt.Println(userId, ok)

	c.Delete("userId")
	userId2, _ := c.Get("userId")

	fmt.Println(userId2)
}

```

---
