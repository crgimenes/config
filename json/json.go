package json

import (
	"encoding/json"
	"os"
	"path/filepath"

	"crg.eti.br/go/config"
	"crg.eti.br/go/config/helper"
)

func init() {
	f := config.Fileformat{
		Extension:   ".json",
		Load:        LoadJSON,
		PrepareHelp: PrepareHelp,
	}
	config.Formats = append(config.Formats, f)
}

// LoadJSON config file.
func LoadJSON(config interface{}) (err error) {
	configFile := filepath.Join(config.Path, config.File)
	file, err := os.Open(configFile)
	if err != nil {
		if os.IsNotExist(err) && !config.FileRequired {
			err = nil
		}
		return
	}
	defer helper.Closer(file)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return
	}

	return
}

// PrepareHelp return help string for this file format.
func PrepareHelp(config interface{}) (help string, err error) {
	var helpAux []byte
	helpAux, err = json.MarshalIndent(&config, "", "    ")
	if err != nil {
		return
	}
	help = string(helpAux)
	return
}
