// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"

	"github.com/buildkite/yaml"
)

// create filesystem based on the operating system
var appFS = afero.NewOsFs()

// Generate produces a pipeline based off the provided configuration.
func (c *Config) Generate() error {
	// create the pipeline file content
	pipeline := steps(c.Type)

	// check if stages were enabled for the provided configuration
	if c.Stages {
		pipeline = stages(c.Type)
	}

	// create output for pipeline file
	out, err := yaml.Marshal(pipeline)
	if err != nil {
		return err
	}

	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// send Filesystem call to capture base directory path
	base, err := os.Getwd()
	if err != nil {
		return err
	}

	// create full path for pipeline file
	path := filepath.Join(base, c.File)

	// check if custom path was provided for pipeline file
	if len(c.Path) > 0 {
		// create custom full path for pipeline file
		path = filepath.Join(c.Path, c.File)
	}

	// send Filesystem call to create directory path for pipeline file
	err = a.Fs.MkdirAll(filepath.Dir(path), 0777)
	if err != nil {
		return err
	}

	// send Filesystem call to create pipeline file
	return a.WriteFile(path, []byte(out), 0644)
}