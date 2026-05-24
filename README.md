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


---

## Example How to Use

```
package main

import (
	"fmt"

	"github.com/LiaVoitenko/golang_basic/cache"
)

func main() {
	c := cache.New()

	// Set value
	c.Set("userId", 42)

	// Get value
	value, ok := c.Get("userId")
	fmt.Println(value, ok)

	// Delete value
	c.Delete("userId")

	// Try to get deleted value
	value, ok = c.Get("userId")
	fmt.Println(value, ok)
}
```