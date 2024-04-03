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

	"github.com/sapcc/concourse-metadata-resource/pkg/models"
)

func main() {
	meta := make(models.Metadata, 7)

	var request models.OutRequest

	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		fatal("reading request", err)
	}

	buildCreatedBy := os.Getenv("BUILD_CREATED_BY")

	var outVersion = request.Version

	if outVersion.Version == "" {
		outVersion.Version = os.Getenv("BUILD_ID")
	}

	if buildCreatedBy != "" {
		outVersion.Version = outVersion.Version + "+" + buildCreatedBy
	}

	handleProp("BUILD_ID", meta, 0)
	handleProp("BUILD_NAME", meta, 1)
	handleProp("BUILD_JOB_NAME", meta, 2)
	handleProp("BUILD_PIPELINE_NAME", meta, 3)
	handleProp("BUILD_TEAM_NAME", meta, 4)
	handleProp("ATC_EXTERNAL_URL", meta, 5)
	handleProp("BUILD_CREATED_BY", meta, 6)

	err = json.NewEncoder(os.Stdout).Encode(models.OutResponse{
		Version:  outVersion,
		Metadata: meta,
	})

	if err != nil {
		fatal("encoding metadata", err)
	}

	log("Done")
}

func fatal(doing string, err error) {
	fmt.Fprintln(os.Stderr, "error "+doing+": "+err.Error())
	os.Exit(1)
}

func log(doing string) {
	fmt.Fprintln(os.Stderr, doing)
}

func handleProp(prop string, meta models.Metadata, index int) {
	val := os.Getenv(prop)
	meta[index] = models.MetadataField{
		Name:  prop,
		Value: val,
	}
}
