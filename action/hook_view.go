// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/hook"

	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
)

// HookView defines the command for inspecting a hook.
var HookView = &cli.Command{
	Name:        "hook",
	Description: "Use this command to view a hook.",
	Usage:       "View details of the provided hook",
	Action:      hookView,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "HOOK_ORG"},
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "provide the organization for the hook",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "HOOK_REPO"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "provide the repository for the hook",
		},

		// Hook Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_HOOK", "HOOK_NUMBER"},
			Name:    "hook",
			Aliases: []string{"h"},
			Usage:   "provide the number for the hook",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "HOOK_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "print the output in default, yaml or json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View hook details for a repository.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --hook 1
  2. View hook details for a repository with json output.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --hook 1 --output json
  3. View hook details for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --hook 1

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/hook/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// inspect a hook.
func hookView(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the hook configuration
	h := &hook.Config{
		Action: viewAction,
		Org:    c.String("org"),
		Repo:   c.String("repo"),
		Number: c.Int("hook"),
		Output: c.String("output"),
	}

	// validate hook configuration
	err = h.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the hook configuration
	return h.View(client)
}