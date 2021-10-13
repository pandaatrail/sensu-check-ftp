package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types",
	"github.com/jlaffaye/ftp"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	Example string
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "sensu-check-ftp",
			Short:    "FTP/SFT Checks for sensu go",
			Keyspace: "sensu.io/plugins/sensu-check-ftp/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		&sensu.PluginConfigOption{
			Path:      "host",
			Env:       "CHECK_EXAMPLE",
			Argument:  "host",
			Shorthand: "h",
			Default:   "localhost",
			Usage:     "The host to connect to",
			Value:     &plugin.Host,
		},
		&sensu.PluginConfigOption{
			Path:      "port",
			Env:       "CHECK_EXAMPLE",
			Argument:  "port",
			Shorthand: "p",
			Default:   "21",
			Usage:     "The port to use",
			Value:     &plugin.Port,
		},
		&sensu.PluginConfigOption{
			Path:      "user",
			Env:       "CHECK_EXAMPLE",
			Argument:  "user",
			Shorthand: "u",
			Default:   "michel",
			Usage:     "The user to use",
			Value:     &plugin.User,
		},
		&sensu.PluginConfigOption{
			Path:      "password",
			Env:       "CHECK_EXAMPLE",
			Argument:  "password",
			Shorthand: "P",
			Default:   "password",
			Usage:     "The password for this user",
			Value:     &plugin.User,
		},
	}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, useStdin)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
	return sensu.CheckStateOK, nil
}

func executeCheck(event *types.Event) (int, error) {
	c, err := ftp.Dial("%s:%d", ftp.DialWithTimeout(5*time.Second), plugin.Host, plugin.Port)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("%s", "%s", plugin.User, plugin.Password)
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the FTP conn
	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
	return sensu.CheckStateOK, nil
}
