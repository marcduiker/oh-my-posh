package segments

import (
	"github.com/jandedobbeleer/oh-my-posh/src/properties"
	"github.com/jandedobbeleer/oh-my-posh/src/runtime"
)

type Dapr struct {
	language
}

func (d *Dapr) Template() string {
	return languageTemplate
}

func (d *Dapr) Init(props properties.Properties, env runtime.Environment) {
	d.language = language{
		env:   env,
		props: props,
		commands: []*cmd{
			{
				executable: "dapr",
				args:       []string{"--version"},
				regex:      `(?:(?P<version>((?P<major>[0-9]+).(?P<minor>[0-9]+).(?P<patch>[0-9]+)(-(?P<rc>rc\.[0-9]+))?)))`,
			},
		},
		versionURLTemplate: "https://github.com/dapr/dapr/releases/tag/v{{.Full}}",
	}
}

func (d *Dapr) Enabled() bool {
	return d.language.Enabled()
}
