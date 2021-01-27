package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	Host  string `json:"host"`
	Ports []int  `json:"ports"`
}

type Scan struct {
	Params  Request        `json:"request"`
	Results map[int]Result `json:"results"`
}

type Result struct {
	Proto   string `json:"proto"`
	State   string `json:"state"`
	Service string `json:"service"`
	Version string `json:"version"`
}

type Client struct {
	address string
}

func NewClient(address string) *Client {
	return &Client{address: address}
}

func (c *Client) Scan(host string, ports []int) (Scan, error) {
	req := Request{
		Host:  host,
		Ports: ports,
	}

	b, err := json.Marshal(req)
	if err != nil {
		return Scan{}, err
	}

	res, err := http.Post(c.address, "application/json", bytes.NewReader(b))
	if err != nil {
		return Scan{}, err
	}

	var body Scan
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return Scan{}, err
	}

	return body, nil
}
