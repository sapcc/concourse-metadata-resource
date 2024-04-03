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
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"

	"github.com/onsi/gomega/gbytes"
	"github.com/sapcc/concourse-metadata-resource/pkg/models"
)

var _ = Describe("Check", func() {
	var (
		checkCmd *exec.Cmd
	)

	BeforeEach(func() {
		checkCmd = exec.Command(checkPath)
	})

	Context("when executed", func() {
		var source map[string]interface{}
		var version *models.TimestampVersion
		var response models.CheckResponse

		BeforeEach(func() {
			source = map[string]interface{}{}
			response = models.CheckResponse{}
			version = nil
		})

		JustBeforeEach(func() {
			stdin, err := checkCmd.StdinPipe()
			Expect(err).NotTo(HaveOccurred())

			session, err := gexec.Start(checkCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			err = json.NewEncoder(stdin).Encode(map[string]interface{}{
				"source":  source,
				"version": version,
			})
			Expect(err).NotTo(HaveOccurred())

			<-session.Exited
			Expect(session.ExitCode()).To(Equal(0))

			err = json.Unmarshal(session.Out.Contents(), &response)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when no version is given", func() {
			It("outputs a single element version array", func() {
				Expect(response).To(HaveLen(1))
				Expect(response[0].Version).ToNot(BeEmpty())
			})
		})

		Context("when version is given", func() {

			BeforeEach(func() {
				version = &models.TimestampVersion{
					Version: "1",
				}
			})

			It("outputs a new single element version array", func() {
				Expect(response).To(HaveLen(1))
				Expect(response[0].Version).ToNot(Equal("1"))
			})
		})

	})

	Context("with invalid inputs", func() {
		var session *gexec.Session

		JustBeforeEach(func() {
			stdin, err := checkCmd.StdinPipe()
			Expect(err).NotTo(HaveOccurred())

			session, err = gexec.Start(checkCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())

			stdin.Close()
		})
		Context("with a missing everything", func() {
			It("returns an error", func() {
				<-session.Exited
				Expect(session.Err).To(gbytes.Say("parse error: EOF"))
				Expect(session.ExitCode()).To(Equal(1))
			})
		})

	})
})
