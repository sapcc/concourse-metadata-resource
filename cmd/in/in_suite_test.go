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

package main_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var inPath string

var _ = BeforeSuite(func() {
	var err error

	if _, err = os.Stat("/opt/resource/check"); err == nil {
		inPath = "/opt/resource/in"
	} else {
		inPath, err = gexec.Build("github.com/sapcc/concourse-metadata-resource/cmd/in")
		Expect(err).NotTo(HaveOccurred())
	}
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func TestIn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "In Suite")
}
