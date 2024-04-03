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
	"encoding/json"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/sapcc/concourse-metadata-resource/pkg/models"
)

var _ = Describe("Out", func() {
	var destination string
	var outCmd *exec.Cmd

	BeforeEach(func() {
		outCmd = exec.Command(outPath, destination)

		outCmd.Env = append(os.Environ(),
			"BUILD_ID=1",
			"BUILD_NAME=2",
			"BUILD_JOB_NAME=3",
			"BUILD_PIPELINE_NAME=4",
			"BUILD_TEAM_NAME=5",
			"ATC_EXTERNAL_URL=6",
			"BUILD_CREATED_BY=7",
		)
	})

	AfterEach(func() {
	})

	Context("when executed", func() {
		var request models.OutRequest
		var response models.OutResponse

		BeforeEach(func() {

			request = models.InRequest{
				Version: models.TimestampVersion{
					Version: "1",
				},
				Source: models.Source{},
			}

			response = models.OutResponse{}
		})

		JustBeforeEach(func() {
			stdin, err := outCmd.StdinPipe()
			Expect(err).NotTo(HaveOccurred())

			session, err := gexec.Start(outCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			err = json.NewEncoder(stdin).Encode(request)
			Expect(err).NotTo(HaveOccurred())

			<-session.Exited
			Expect(session.ExitCode()).To(Equal(0))

			err = json.Unmarshal(session.Out.Contents(), &response)
			Expect(err).NotTo(HaveOccurred())
		})

		It("reports the version to be the input version", func() {
			Expect(response.Version.Version).To(Equal("1+7"))
		})
	})
})
