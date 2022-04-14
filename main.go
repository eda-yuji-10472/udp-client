// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"log"
	"net"
	"os"
	"time"
)

func say() {
	conn, err := net.Dial("udp", "udp-server.eda-test.verification-gcp.colopl.jp.:8080")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second * 1)
		n, err := conn.Write([]byte("Ping"))
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}

		if len(buf) != n {
			log.Printf("data size is %d, but sent data size is %d", len(buf), n)
		}

		recvBuf := make([]byte, 1024)

		n, err = conn.Read(recvBuf)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}

		log.Printf("Received data: %s", string(recvBuf[:n]))
	}
}

func main() {
	go say()
}
