// Copyright 2018, Special Brands BV
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/goburrow/modbus"
)

func main() {
	serial := "/dev/cu.usbserial-FT1TEP2G"
	if len(os.Args) > 1 {
		serial = os.Args[1]
	}
	fmt.Println(serial)

	handler := modbus.NewRTUClientHandler(serial)

	err := handler.Connect()
	defer handler.Close()
	client := modbus.NewClient(handler)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	for true {
		results, err := client.ReadHoldingRegisters(0, 3)
		if err != nil {
			fmt.Printf("%v\n", err)
			result := fmt.Sprintf("%v", err)
			if result != "serial: timeout" {
				break
			}
		}
		fmt.Printf("results %v\n", results)
	}
}
