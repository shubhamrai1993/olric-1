// Copyright 2018-2020 Burak Sezer
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

package olric

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (db *Olric) systemStatsHTTPHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data, err := db.serializer.Marshal(db.stats())
	if err != nil {
		db.httpErrorResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		db.log.V(6).Printf("[ERROR] Failed to write to ResponseWriter: %v", err)
	}
}

func (db *Olric) systemPingHTTPHandler(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	addr := ps.ByName("addr")
	data, err := db.serializer.Marshal(db.Ping(addr))
	if err != nil {
		db.httpErrorResponse(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		db.log.V(6).Printf("[ERROR] Failed to write to ResponseWriter: %v", err)
	}
}