package envconf

import (
	"os"

	"github.com/pkg/errors"

	"gopkg.in/yaml.v2"
)

// From creates a new conf.
func From(defaultConf, filename string, p interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			return errors.Wrap(err, "opening config file")
		}
		nf, err := os.Create(filename)
		if err != nil {
			return errors.Wrap(err, "creating config file")
		}
		defer func() { _ = nf.Close() }() // Best effort.

		if _, err := nf.WriteString(defaultConf); err != nil {
			return errors.Wrap(err, "writing default config")
		}
		f = nf
	}
	return errors.Wrap(yaml.NewDecoder(f).Decode(p), "decoding config file")
}
