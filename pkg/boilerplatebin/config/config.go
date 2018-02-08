package config

import (
	"os"

	"github.com/jessevdk/go-flags"
)

//Configuration, can be go-flags too
type Config struct {
	LogLevel string `short:"v" long:"log-level" default:"0" description:"Logging level"`
}

// Parse the args and environment to fill the ServerConfig
func (c *Config) Parse() error {
	parser := flags.NewParser(c, flags.HelpFlag|flags.PrintErrors|flags.PassDoubleDash)
	_, err := parser.Parse()

	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			return err
		}
	}

	return nil
}
