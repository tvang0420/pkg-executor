// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package linux

import (
	"fmt"

	"github.com/go-vela/pkg-runtime/runtime"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/library"
	"github.com/go-vela/types/pipeline"

	"github.com/sirupsen/logrus"
)

// Opt represents a configuration option to initialize the client.
type Opt func(*client) error

// WithBuild sets the library build in the client.
func WithBuild(b *library.Build) Opt {
	logrus.Trace("configuring build in linux client")

	return func(c *client) error {
		// check if the build provided is empty
		if b == nil {
			return fmt.Errorf("empty build provided")
		}

		// set the build in the client
		c.build = b

		return nil
	}
}

// WithPipeline sets the pipeline build in the client.
func WithPipeline(p *pipeline.Build) Opt {
	logrus.Trace("configuring pipeline in linux client")

	return func(c *client) error {
		// check if the pipeline provided is empty
		if p == nil {
			return fmt.Errorf("empty pipeline provided")
		}

		// set the pipeline in the client
		c.pipeline = p

		return nil
	}
}

// WithRepo sets the library repo in the client.
func WithRepo(r *library.Repo) Opt {
	logrus.Trace("configuring repo in linux client")

	return func(c *client) error {
		// check if the repo provided is empty
		if r == nil {
			return fmt.Errorf("empty repo provided")
		}

		// set the repo in the client
		c.repo = r

		return nil
	}
}

// WithRuntime sets the runtime engine in the client.
func WithRuntime(r runtime.Engine) Opt {
	logrus.Trace("configuring runtime in linux client")

	return func(c *client) error {
		// check if the runtime provided is empty
		if r == nil {
			return fmt.Errorf("empty runtime provided")
		}

		// set the runtime in the client
		c.Runtime = r

		return nil
	}
}

// WithUser sets the library user in the client.
func WithUser(u *library.User) Opt {
	logrus.Trace("configuring user in linux client")

	return func(c *client) error {
		// check if the user provided is empty
		if u == nil {
			return fmt.Errorf("empty user provided")
		}

		// set the user in the client
		c.user = u

		return nil
	}
}

// WithVelaClient sets the Vela client in the client.
func WithVelaClient(cli *vela.Client) Opt {
	logrus.Trace("configuring Vela client in linux client")

	return func(c *client) error {
		// check if the Vela client provided is empty
		if cli == nil {
			return fmt.Errorf("empty Vela client provided")
		}

		// set the Vela client in the client
		c.Vela = cli

		return nil
	}
}