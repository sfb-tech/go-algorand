// Copyright (C) 2019-2020 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/algorand/websocket"
	"github.com/gorilla/mux"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  10240,
	WriteBufferSize: 10240,
	CheckOrigin: func(r *http.Request) bool {
		if len(r.Header.Get("Origin")) == 0 {
			return true
		}
		if strings.HasPrefix(r.Header.Get("Origin"), "devtools://") {
			return true
		}
		if strings.HasPrefix(r.Header.Get("Origin"), "http://localhost") {
			return true
		}
		return false
	},
}

func main() {
	router := mux.NewRouter()

	appAddress := "localhost:9392"

	debugger := MakeDebugger()

	remote := RemoteHookAdapter{debugger}
	remote.Setup(router)

	wa := WebPageAdapter{}
	wa.Setup(router)
	debugger.AddAdapter(&wa)

	cdta := CDTAdapter{}
	cdta.Setup(&CDTSetupParams{router, appAddress})
	debugger.AddAdapter(&cdta)

	server := http.Server{
		Handler:      router,
		Addr:         appAddress,
		WriteTimeout: time.Duration(0),
		ReadTimeout:  time.Duration(0),
	}

	log.Printf("starting server on %s", appAddress)
	server.ListenAndServe()
}
