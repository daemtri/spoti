package flags

import (
	"strings"

	"github.com/spf13/cobra"

	"go.uber.org/multierr"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// NamedFlagSets stores named flag sets in the order of calling FlagSet.
type NamedFlagSets struct {
	// Order is an ordered list of flag set names.
	Order []string
	// FlagSets stores the flag sets by name.
	FlagSets map[string]*pflag.FlagSet
}

// FlagSet returns the flag set with the given name and adds it to the
// ordered name list if it is not in there yet.
func (nfs *NamedFlagSets) FlagSet(name string) *pflag.FlagSet {
	if nfs.FlagSets == nil {
		nfs.FlagSets = map[string]*pflag.FlagSet{}
	}
	if _, ok := nfs.FlagSets[name]; !ok {
		nfs.FlagSets[name] = pflag.NewFlagSet(name, pflag.ExitOnError)
		nfs.Order = append(nfs.Order, name)
	}
	return nfs.FlagSets[name]
}

func (nfs *NamedFlagSets) BindEnvsForViper(prefix string) error {
	viper.SetEnvPrefix(prefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	var result error
	for name, fs := range nfs.FlagSets {
		fs.VisitAll(func(flag *pflag.Flag) {
			tagName := strings.ReplaceAll(flag.Name, "-", "_")
			err := viper.BindEnv(name+"."+tagName, prefix+"_"+name+"_"+tagName)
			if err != nil {
				result = multierr.Append(result, err)
			}
		})
	}

	return result
}

func (nfs *NamedFlagSets) AddToCobraCommand(cmd *cobra.Command) {
	fs := cmd.Flags()
	for prefix, f := range nfs.FlagSets {
		f.VisitAll(func(flag *pflag.Flag) {
			if prefix != "generic" {
				flag.Name = prefix + "-" + flag.Name
			}
			fs.AddFlag(flag)
		})
	}
}
