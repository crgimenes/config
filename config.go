// Package config uses a struct as input and populates the
// fields of this struct with parameters fom command
// line, environment variables and configuration file.
package config

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"crg.eti.br/go/config/goenv"
	"crg.eti.br/go/config/goflags"
	"crg.eti.br/go/config/structtag"
	"crg.eti.br/go/config/validate"
)

// Fileformat struct holds the functions to Load the file containing the settings.
type Fileformat struct {
	Extension   string
	Load        func(config interface{}) (err error)
	PrepareHelp func(config interface{}) (help string, err error)
}

var (

	// Tag to set main name of field.
	Tag = "cfg"

	// TagDefault to set default value.
	TagDefault = "cfgDefault"

	// TagHelper to set usage help line.
	TagHelper = "cfgHelper"

	// Path sets default config path.
	Path string

	// File name of default config file.
	File string

	// FileRequired config file required.
	FileRequired bool

	// HelpString temporarily saves help.
	HelpString string

	// PrefixFlag is a string that would be placed at the beginning of the generated Flag tags.
	PrefixFlag string

	// PrefixEnv is a string that would be placed at the beginning of the generated Event tags.
	PrefixEnv string

	// ErrFileFormatNotDefined Is the error that is returned when there is no defined configuration file format.
	ErrFileFormatNotDefined = errors.New("file format not defined")

	Usage func()

	// Formats is the list of registered formats.
	Formats []Fileformat

	// FileEnv is the enviroment variable that define the config file.
	FileEnv string

	// PathEnv is the enviroment variable that define the config file path.
	PathEnv string

	// WatchConfigFile is the flag to update the config when the config file changes.
	WatchConfigFile bool

	// DisableFlags on the command line.
	DisableFlags bool
)

func findFileFormat(extension string) (format Fileformat, err error) {
	format = Fileformat{}
	for _, f := range Formats {
		if f.Extension == extension {
			format = f
			return
		}
	}
	err = ErrFileFormatNotDefined
	return
}

func init() {
	Usage = DefaultUsage
	Path = "./"
	File = ""
	FileRequired = false

	FileEnv = "GO_CONFIG_FILE"
	PathEnv = "GO_CONFIG_PATH"

	WatchConfigFile = false
}

// Parse configuration.
func Parse(config interface{}) (err error) {
	goenv.Prefix = PrefixEnv
	goenv.Setup(Tag, TagDefault)
	err = structtag.SetBoolDefaults(config, "")
	if err != nil {
		return
	}

	lookupEnv()

	ext := path.Ext(File)
	if ext != "" {
		if err = loadConfigFromFile(ext, config); err != nil {
			return
		}
	}

	goenv.Prefix = PrefixEnv
	goenv.Setup(Tag, TagDefault)
	err = goenv.Parse(config)
	if err != nil {
		return
	}

	if !DisableFlags {
		goflags.Prefix = PrefixFlag
		goflags.Setup(Tag, TagDefault, TagHelper)
		goflags.Usage = Usage
		goflags.Preserve = true
		err = goflags.Parse(config)
		if err != nil {
			return
		}
	}

	validate.Prefix = PrefixFlag
	validate.Setup(Tag, TagDefault)
	err = validate.Parse(config)

	return
}

// PrintDefaults print the default help.
func PrintDefaults() {
	if File != "" {
		fmt.Printf("Config file %q:\n", filepath.Join(Path, File))
		fmt.Println(HelpString)
	}
}

// DefaultUsage is assigned for Usage function by default.
func DefaultUsage() {
	fmt.Println("Usage")
	goflags.PrintDefaults()
	goenv.PrintDefaults()
	PrintDefaults()
}

func lookupEnv() {
	pref := PrefixEnv
	if pref != "" {
		pref = pref + structtag.TagSeparator
	}

	if val, set := os.LookupEnv(pref + FileEnv); set {
		File = val
	}

	if val, set := os.LookupEnv(pref + PathEnv); set {
		Path = val
	}
}

func loadConfigFromFile(ext string, config interface{}) (err error) {
	var format Fileformat
	format, err = findFileFormat(ext)
	if err != nil {
		return
	}
	err = format.Load(config)
	if err != nil {
		return
	}
	HelpString, err = format.PrepareHelp(config)
	return
}
