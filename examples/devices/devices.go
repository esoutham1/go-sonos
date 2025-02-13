// go-sonos
// ========
//
// Copyright (c) 2012-2015, Ian T. Richards <ianr@panix.com>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//
//   - Redistributions of source code must retain the above copyright notice,
//     this list of conditions and the following disclaimer.
//   - Redistributions in binary form must reproduce the above copyright
//     notice, this list of conditions and the following disclaimer in the
//     documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED
// TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
// LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
// NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package main

import (
	"github.com/esoutham1/go-sonos/ssdp"
	"log"
	"runtime"
)

// This code identifies lists all UPnP devices discovered
func main() {
	log.Print("go-sonos example discovery\n")
	ethernetInterfaceName := "eth0"
	goos := runtime.GOOS
	if goos == "darwin" {
		ethernetInterfaceName = "en0" //first ethernet adapter on MacOS
	}

	mgr := ssdp.MakeManager()
	mgr.Discover(ethernetInterfaceName, "0", false)
	i := 0
	dev_map := mgr.Devices()
	for _, dev := range dev_map {
		log.Printf("[%02d] %s %s %s %s %s\n", i, dev.Product(), dev.ProductVersion(), dev.Name(), dev.Location(), dev.UUID())
		i++
	}

	mgr.Close()
}
