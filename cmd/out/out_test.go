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
	"fmt"
	"os"
	"os/exec"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Out", func() {
	var tmpdir string
	var source string

	var session *gexec.Session
	var outCmd *exec.Cmd

	BeforeEach(func() {
		var err error

		tmpdir, err = os.MkdirTemp("", "out-source")
		Expect(err).NotTo(HaveOccurred())

		source = path.Join(tmpdir, "out-dir")
		err = os.MkdirAll(source, 0755)
		Expect(err).NotTo(HaveOccurred())
		outCmd = exec.Command(outPath, source)
		fmt.Printf("%s", tmpdir)
	})

	AfterEach(func() {
		os.RemoveAll(tmpdir)
	})

	Context("when executed", func() {

		JustBeforeEach(func() {
			var err error

			session, err = gexec.Start(outCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("reports error", func() {
			<-session.Exited
			Expect(session.Err).To(gbytes.Say("out should not be used"))
			Expect(session.ExitCode()).To(Equal(1))
		})

	})

})
