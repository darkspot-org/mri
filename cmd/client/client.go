package main

import (
	"fmt"
	"github.com/darkspot-org/mri/client"
	"log"
	"os"
)

func main() {
	c := client.NewClient("http://localhost:8080")

	res, err := c.Scan(os.Args[1], []int{80, 443})
	if err != nil {
		log.Fatal(err)
	}

	for port, info := range res.Results {
		fmt.Printf("%d %v \n", port, info)
	}
}
