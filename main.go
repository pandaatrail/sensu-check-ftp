package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	Host string
	Port int
	User string
	Password string
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "sensu-check-ftp",
			Short:    "FTP/SFTP Checks for sensu go",
			Keyspace: "sensu.io/plugins/sensu-check-ftp/config",
		},
	}

	options = []*sensu.PluginConfigOption{
		{
			Path:      "host",
			Env:       "CHECK_EXAMPLE",
			Argument:  "host",
			Shorthand: "z",
			Default:   "localhost",
			Usage:     "The host to connect to",
			Value:     &plugin.Host,
		},
		{
			Path:      "port",
			Env:       "CHECK_EXAMPLE",
			Argument:  "port",
			Shorthand: "p",
			Default:   21,
			Usage:     "The port to use",
			Value:     &plugin.Port,
		},
		{
			Path:      "user",
			Env:       "CHECK_EXAMPLE",
			Argument:  "user",
			Shorthand: "u",
			Default:   "michel",
			Usage:     "The user to use",
			Value:     &plugin.User,
		},
		{
			Path:      "password",
			Env:       "CHECK_EXAMPLE",
			Argument:  "password",
			Shorthand: "P",
			Default:   "password",
			Usage:     "The password for this user",
			Value:     &plugin.Password,
		},
	}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
	return sensu.CheckStateOK, nil
}



func executeCheck(event *types.Event) (int, error) {

	
	params := fmt.Sprintf("%s:%d", plugin.Host, plugin.Port)

	c, err := ftp.Dial(params, ftp.DialWithTimeout(5*time.Second))
	// Do something with the FTP conn
	if err != nil {
		fmt.Printf("Conn: %s", err.Error())
		os.Exit(1)
	}

	err = c.Login(plugin.User, plugin.Password)
	if err != nil {
		fmt.Printf("Login: %s", err.Error())
		os.Exit(1)
	}

	err = c.Quit()
	if err != nil {
		fmt.Printf("Quit: %s", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
	return 0, nil

	/*
	err = c.Login("test", "test")
	if err != nil {
		log.Fatal(err)
	}
	// Do something with the FTP conn
	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
	
	if err != nil {
		log.Fatal(err)
	}


	return sensu.CheckStateOK, nil
	*/
}
