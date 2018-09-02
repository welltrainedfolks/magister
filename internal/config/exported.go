// Copyright (c) 2018, Well Trained Folks and contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject
// to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package config

import (
	// stdlib
	"flag"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"

	// other
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

var (
	configPath string
	// Configuration struct.
	Config *Configuration
)

// Initialize initializes package.
func Initialize() {
	log.Info().Msg("Initializing configuration module...")
	flag.StringVar(&configPath, "config", "", "Path to configuration file. Can be relative or absolute.")
}

// LoadConfiguration loads configuration file.
func LoadConfiguration() {
	log.Info().Msg("Trying to load configuration...")

	if configPath == "" {
		log.Fatal().Msg("Configuration file path wasn't specified. Please, specify it with '-config' parameter. See '-help'.")
	}

	// Prepare configuration path.
	if strings.Contains(configPath, "~") {
		curUser, err := user.Current()
		if err != nil {
			log.Fatal().Msg("Failed to get current user's data! This is neccessary for proper configuration file path expanding. If your configuration file is really present - try to specify ABSOLUTE path to configuration file.")
		}
		configPath = strings.Replace(configPath, "~", curUser.HomeDir, -1)
	}

	var err1 error
	configPath, err1 = filepath.Abs(configPath)
	if err1 != nil {
		log.Fatal().Msgf("Failed to obtain absolute path to configuration file: %s", err1.Error())
	}

	log.Info().Msgf("Configuration file absolute path: %s", configPath)

	// Load and parse configuration file.
	Config = &Configuration{}
	fileData, err2 := ioutil.ReadFile(configPath)
	if err2 != nil {
		log.Fatal().Msgf("Failed to read configuration file contents: %s", err2.Error())
	}

	err3 := yaml.Unmarshal(fileData, Config)
	if err3 != nil {
		log.Fatal().Msgf("Failed to parse configuration file: %s", err3.Error())
	}

	log.Debug().Msgf("Parsed configuration: %+v", Config)
}
