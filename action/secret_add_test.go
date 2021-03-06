// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"flag"
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/urfave/cli/v2"
)

func TestAction_SecretAdd(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup flags
	authSet := flag.NewFlagSet("test", 0)
	authSet.String("api.addr", s.URL, "doc")
	authSet.String("api.token", "superSecretToken", "doc")

	fullSet := flag.NewFlagSet("test", 0)
	fullSet.String("api.addr", s.URL, "doc")
	fullSet.String("api.token", "superSecretToken", "doc")
	fullSet.String("secret.engine", "native", "doc")
	fullSet.String("secret.type", "repo", "doc")
	fullSet.String("org", "github", "doc")
	fullSet.String("repo", "octocat", "doc")
	fullSet.String("name", "foo", "doc")
	fullSet.String("value", "bar", "doc")
	fullSet.String("output", "json", "doc")

	fileSet := flag.NewFlagSet("test", 0)
	fileSet.String("api.addr", s.URL, "doc")
	fileSet.String("api.token", "superSecretToken", "doc")
	fileSet.String("file", "secret/testdata/repo.yml", "doc")
	fileSet.String("output", "json", "doc")

	// setup tests
	tests := []struct {
		failure bool
		set     *flag.FlagSet
	}{
		{
			failure: false,
			set:     fileSet,
		},
		{
			failure: false,
			set:     fullSet,
		},
		{
			failure: true,
			set:     authSet,
		},
		{
			failure: true,
			set:     flag.NewFlagSet("test", 0),
		},
	}

	// run tests
	for _, test := range tests {
		err := secretAdd(cli.NewContext(nil, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("secretAdd should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("secretAdd returned err: %v", err)
		}
	}
}
