package ini

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/crgimenes/goconfig"
	ini "gopkg.in/ini.v1"
)

func init() {
	f := goconfig.Fileformat{
		Extension:   ".ini",
		Load:        LoadINI,
		PrepareHelp: PrepareHelp,
	}
	goconfig.Formats = append(goconfig.Formats, f)
}

// LoadINI config file.
func LoadINI(config interface{}) (err error) {
	configFile := filepath.Join(goconfig.Path, goconfig.File)
	file, err := os.Open(configFile)
	if os.IsNotExist(err) && !goconfig.FileRequired {
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
	m := make(map[string]any)
	err = ini.MapTo(config, m)
	if err != nil {
		return
	}

	for k, v := range m {
		s := ""

		switch v.(type) {
		case string:
			s = v.(string)
		case int:
			s = strconv.Itoa(v.(int))
		case bool:
			s = strconv.FormatBool(v.(bool))
		}

		help += k + " = " + s + "\n"
	}
	return
}
