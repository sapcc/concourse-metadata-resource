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
	"os"
	"path/filepath"

	"bufio"
	"fmt"

	"github.com/sapcc/concourse-metadata-resource/pkg/models"
)

func main() {
	if len(os.Args) < 2 {
		fatalNoErr("usage: " + os.Args[0] + " <destination>")
	}

	destination := os.Args[1]

	log("creating destination dir " + destination)
	err := os.MkdirAll(destination, 0755)
	if err != nil {
		fatal("creating destination", err)
	}

	meta := make(models.Metadata, 6)

	var request models.InRequest

	err = json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		fatal("reading request", err)
	}

	var inVersion = request.Version

	handleProp(destination, "build-id", "BUILD_ID", meta, 0)
	handleProp(destination, "build-name", "BUILD_NAME", meta, 1)
	handleProp(destination, "build-job-name", "BUILD_JOB_NAME", meta, 2)
	handleProp(destination, "build-pipeline-name", "BUILD_PIPELINE_NAME", meta, 3)
	handleProp(destination, "build-team-name", "BUILD_TEAM_NAME", meta, 4)
	handleProp(destination, "atc-external-url", "ATC_EXTERNAL_URL", meta, 5)

	err = json.NewEncoder(os.Stdout).Encode(models.InResponse{
		Version:  inVersion,
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

func fatalNoErr(doing string) {
	log(doing)
	os.Exit(1)
}

func handleProp(destination, filename, prop string, meta models.Metadata, index int) {
	output := filepath.Join(destination, filename)
	log("creating output file " + output)
	file, err := os.Create(output)
	if err != nil {
		fatal("creating output file "+output, err)
	}
	defer file.Close()

	val := os.Getenv(prop)
	meta[index] = models.MetadataField{
		Name:  prop,
		Value: val,
	}
	w := bufio.NewWriter(file)
	fmt.Fprintf(w, "%s", val)

	err = w.Flush()

	if err != nil {
		fatal("writing output file"+output, err)
	}
}
