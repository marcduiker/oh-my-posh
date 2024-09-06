package segments

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDapr(t *testing.T) {
	cases := []struct {
		Case           string
		ExpectedString string
		Version        string
	}{
		{Case: "CLI version regular", ExpectedString: "1.14.1", Version: "CLI version: 1.14.1"},
		{Case: "CLI version RC", ExpectedString: "1.14.1-rc.8", Version: "CLI version: 1.14.1\nRuntime version: 1.14.0-rc.8"},
	}
	for _, tc := range cases {
		params := &mockedLanguageParams{
			cmd:           "dapr",
			versionParam:  "--version",
			versionOutput: tc.Version,
		}
		env, props := getMockedLanguageEnv(params)
		d := &Dapr{}
		d.Init(props, env)
		assert.True(t, d.Enabled(), fmt.Sprintf("Failed in case: %s", tc.Case))
		assert.Equal(t, tc.ExpectedString, renderTemplate(env, d.Template(), d), fmt.Sprintf("Failed in case: %s", tc.Case))
	}
}
