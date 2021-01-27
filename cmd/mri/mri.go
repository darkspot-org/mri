package main

import (
	"context"
	"encoding/json"
	"github.com/Ullaakut/nmap/v2"
	"github.com/darkspot-org/mri/client"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/", handle)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error while listening to :8080: %s", err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	var request client.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("error while decoding request body: %s", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	scan, err := scanHost(request)
	if err != nil {
		log.Printf("error while scanning host: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(scan); err != nil {
		log.Printf("error while encoding response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("successfully scanned host: %v", request)
}

func scanHost(req client.Request) (client.Scan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	var ports []string
	for _, port := range req.Ports {
		ports = append(ports, strconv.Itoa(port))
	}

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(req.Host),
		nmap.WithPorts(ports...),
		nmap.WithServiceInfo(),
		nmap.WithContext(ctx),
	)

	if err != nil {
		return client.Scan{}, err
	}

	res, _, err := scanner.Run()
	if err != nil {
		return client.Scan{}, err
	}

	scan := client.Scan{
		Params:  req,
		Results: map[int]client.Result{},
	}

	for _, host := range res.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		for _, port := range host.Ports {
			scan.Results[int(port.ID)] = client.Result{
				Proto:   port.Protocol,
				State:   port.State.State,
				Service: port.Service.Name,
				Version: port.Service.Version,
			}
		}
	}

	return scan, nil
}
