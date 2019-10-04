package cli

import (
	"github.com/KeisukeYamashita/biko/browser"
	dd "github.com/KeisukeYamashita/biko/providers/datadog"
	"github.com/KeisukeYamashita/biko/providers/gcp"
	"github.com/urfave/cli"
)

func newDatadaogCmd() cli.Command {
	return cli.Command{
		Name:    "datadog",
		Aliases: []string{"dd"},
		Usage:   "Open Datadog resource",
		Flags:   []cli.Flag{},
		Subcommands: []cli.Command{
			newDDWatchDogCmd(),
			newDDEventCmd(),
			newDDDashboardCmd(),
			newDDInfrastructureCmd(),
			newDDMonitorsCmd(),
			newDDIntegrationsCmd(),
			newDDApmCmd(),
			newDDNotebookCmd(),
			newDDLogsCmd(),
			newDDSyntheticsCmd(),
		},
	}
}

func newDDWatchDogCmd() cli.Command {
	return cli.Command{
		Name:    "watchdog",
		Aliases: []string{"wd"},
		Usage:   "Open Watchdog page",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			dd := &dd.Provider{}
			return browser.Open(c, dd)
		},
	}
}

func newDDEventCmd() cli.Command {
	return cli.Command{
		Name:  "events",
		Usage: "Open Events page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			dd := &dd.Provider{}
			return browser.Open(c, dd)
		},
	}
}

func newDDDashboardCmd() cli.Command {
	return cli.Command{
		Name:  "dashboard",
		Usage: "Open Dashboard page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newDDInfrastructureCmd() cli.Command {
	return cli.Command{
		Name:  "infrastructure",
		Usage: "Open Infrastructure page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newDDMonitorsCmd() cli.Command {
	return cli.Command{
		Name:  "monitors",
		Usage: "Open Monitors page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newDDIntegrationsCmd() cli.Command {
	return cli.Command{
		Name:  "integrations",
		Usage: "Open Integrations page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newDDApmCmd() cli.Command {
	return cli.Command{
		Name:  "apm",
		Usage: "Open APM page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newDDNotebookCmd() cli.Command {
	return cli.Command{
		Name:  "notebook",
		Usage: "Open Notebook page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newDDLogsCmd() cli.Command {
	return cli.Command{
		Name:  "logs",
		Usage: "Open Logs page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}

func newDDSyntheticsCmd() cli.Command {
	return cli.Command{
		Name:  "synthetics",
		Usage: "Open Synthetics page",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			gcp := &gcp.Provider{}
			return browser.Open(c, gcp)
		},
	}
}
