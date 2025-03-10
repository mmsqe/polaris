// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"github.com/magefile/mage/mg"
)

const (
	baseImage = "polard/base:v0.0.0"

	polardClientPath   = "./cosmos/testing/e2e/polard/"
	localnetRepository = "localnet"
	localnetVersion    = "latest"
)

type Localnet mg.Namespace

func (Localnet) Build() error {
	return ExecuteInDirectory(polardClientPath,
		func(...string) error {
			return dockerBuildFn(false)(
				"--build-arg", "GO_VERSION="+goVersion,
				"--build-arg", "BASE_IMAGE="+baseImage,
				"-t", localnetRepository+":"+localnetVersion,
				".",
			)
		}, false)
}

// Runs the localnet tooling sanity tests.
func (Localnet) Test() error {
	if err := (Contracts{}).Build(); err != nil {
		return err
	}
	LogGreen("Running all localnet tests")
	args := []string{
		"-timeout", "30m",
		"--focus", ".*e2e/localnet.*",
	}
	return ginkgoTest(args...)
}
