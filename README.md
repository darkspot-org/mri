# Sonar

Little RESTful API that exposes NMAP functionalities.

## Request

```json
{
  "host": "8.8.8.8",
  "ports": [
    53
  ]
}
```

## Response

```json
{
  "request": {
    "host": "8.8.8.8",
    "ports": [
      53
    ]
  },
  "results": {
    "53": {
      "proto": "tcp",
      "service": "tcpwrapped",
      "state": "open",
      "version": ""
    }
  }
}
```

## Using dedicated client

```golang
package main

import (
	"fmt"
	"github.com/darkspot-org/sonar"
	"log"
)

func main() {
	c := sonar.NewClient("http://localhost:8080")

	res, err := c.Scan("1.1.1.1", []int{80, 443})
	if err != nil {
		log.Fatal(err)
	}

	for port, info := range res.Results {
		fmt.Printf("%d %v \n", port, info)
	}
}
```