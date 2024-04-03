// Copyright 2024 SAP SE
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
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sapcc/concourse-metadata-resource/pkg/models"
)

func main() {
	var request models.CheckRequest
	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		fmt.Fprintln(os.Stderr, "parse error:", err.Error())
		os.Exit(1)
	}
	t := strconv.FormatInt(time.Now().UnixNano(), 10)
	versions := models.CheckResponse{
		models.TimestampVersion{
			Version: t,
		},
	}
	err = json.NewEncoder(os.Stdout).Encode(versions)
	if err != nil {
		fmt.Fprintln(os.Stderr, "encode error:", err.Error())
		os.Exit(1)
	}
}
