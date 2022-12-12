package ini

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"crg.eti.br/go/config"
	ini "gopkg.in/ini.v1"
)

func init() {
	f := config.Fileformat{
		Extension:   ".ini",
		Load:        LoadINI,
		PrepareHelp: PrepareHelp,
	}
	config.Formats = append(config.Formats, f)
}

// LoadINI config file.
func LoadINI(config interface{}) (err error) {
	configFile := filepath.Join(config.Path, config.File)
	file, err := os.Open(configFile)
	if os.IsNotExist(err) && !config.FileRequired {
		err = nil
		return
	}

	if err != nil {
		return
	}

	err = ini.MapTo(config, file)
	return
}

// PrepareHelp return help string for this file format.
func PrepareHelp(config interface{}) (help string, err error) {
	mAux, err := json.Marshal(config)
	if err != nil {
		return
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(mAux, &m)
	if err != nil {
		return
	}

	// TODO: implement multiple levels
	help = ""
	for k, v := range m {
		help += fmt.Sprintf("%s = %v\n", k, v)
	}
	return
}
