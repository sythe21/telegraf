package config

import (
	"os"
	"testing"
	"time"

	"github.com/sythe21/telegraf/internal/models"
	"github.com/sythe21/telegraf/plugins/inputs"
	"github.com/sythe21/telegraf/plugins/inputs/exec"
	"github.com/sythe21/telegraf/plugins/inputs/memcached"
	"github.com/sythe21/telegraf/plugins/inputs/procstat"
	"github.com/sythe21/telegraf/plugins/parsers"

	"github.com/stretchr/testify/assert"
)

func TestConfig_LoadSingleInputWithEnvVars(t *testing.T) {
	c := NewConfig()
	err := os.Setenv("MY_TEST_SERVER", "192.168.1.1")
	assert.NoError(t, err)
	err = os.Setenv("TEST_INTERVAL", "10s")
	assert.NoError(t, err)
	c.LoadConfig("./testdata/single_plugin_env_vars.toml")

	memcached := inputs.Inputs["memcached"]().(*memcached.Memcached)
	memcached.Servers = []string{"192.168.1.1"}

	filter := internal_models.Filter{
		NameDrop:  []string{"metricname2"},
		NamePass:  []string{"metricname1"},
		FieldDrop: []string{"other", "stuff"},
		FieldPass: []string{"some", "strings"},
		TagDrop: []internal_models.TagFilter{
			internal_models.TagFilter{
				Name:   "badtag",
				Filter: []string{"othertag"},
			},
		},
		TagPass: []internal_models.TagFilter{
			internal_models.TagFilter{
				Name:   "goodtag",
				Filter: []string{"mytag"},
			},
		},
		IsActive: true,
	}
	assert.NoError(t, filter.CompileFilter())
	mConfig := &internal_models.InputConfig{
		Name:     "memcached",
		Filter:   filter,
		Interval: 10 * time.Second,
	}
	mConfig.Tags = make(map[string]string)

	assert.Equal(t, memcached, c.Inputs[0].Input,
		"Testdata did not produce a correct memcached struct.")
	assert.Equal(t, mConfig, c.Inputs[0].Config,
		"Testdata did not produce correct memcached metadata.")
}

func TestConfig_LoadSingleInput(t *testing.T) {
	c := NewConfig()
	c.LoadConfig("./testdata/single_plugin.toml")

	memcached := inputs.Inputs["memcached"]().(*memcached.Memcached)
	memcached.Servers = []string{"localhost"}

	filter := internal_models.Filter{
		NameDrop:  []string{"metricname2"},
		NamePass:  []string{"metricname1"},
		FieldDrop: []string{"other", "stuff"},
		FieldPass: []string{"some", "strings"},
		TagDrop: []internal_models.TagFilter{
			internal_models.TagFilter{
				Name:   "badtag",
				Filter: []string{"othertag"},
			},
		},
		TagPass: []internal_models.TagFilter{
			internal_models.TagFilter{
				Name:   "goodtag",
				Filter: []string{"mytag"},
			},
		},
		IsActive: true,
	}
	assert.NoError(t, filter.CompileFilter())
	mConfig := &internal_models.InputConfig{
		Name:     "memcached",
		Filter:   filter,
		Interval: 5 * time.Second,
	}
	mConfig.Tags = make(map[string]string)

	assert.Equal(t, memcached, c.Inputs[0].Input,
		"Testdata did not produce a correct memcached struct.")
	assert.Equal(t, mConfig, c.Inputs[0].Config,
		"Testdata did not produce correct memcached metadata.")
}

func TestConfig_LoadDirectory(t *testing.T) {
	c := NewConfig()
	err := c.LoadConfig("./testdata/single_plugin.toml")
	if err != nil {
		t.Error(err)
	}
	err = c.LoadDirectory("./testdata/subconfig")
	if err != nil {
		t.Error(err)
	}

	memcached := inputs.Inputs["memcached"]().(*memcached.Memcached)
	memcached.Servers = []string{"localhost"}

	filter := internal_models.Filter{
		NameDrop:  []string{"metricname2"},
		NamePass:  []string{"metricname1"},
		FieldDrop: []string{"other", "stuff"},
		FieldPass: []string{"some", "strings"},
		TagDrop: []internal_models.TagFilter{
			internal_models.TagFilter{
				Name:   "badtag",
				Filter: []string{"othertag"},
			},
		},
		TagPass: []internal_models.TagFilter{
			internal_models.TagFilter{
				Name:   "goodtag",
				Filter: []string{"mytag"},
			},
		},
		IsActive: true,
	}
	assert.NoError(t, filter.CompileFilter())
	mConfig := &internal_models.InputConfig{
		Name:     "memcached",
		Filter:   filter,
		Interval: 5 * time.Second,
	}
	mConfig.Tags = make(map[string]string)

	assert.Equal(t, memcached, c.Inputs[0].Input,
		"Testdata did not produce a correct memcached struct.")
	assert.Equal(t, mConfig, c.Inputs[0].Config,
		"Testdata did not produce correct memcached metadata.")

	ex := inputs.Inputs["exec"]().(*exec.Exec)
	p, err := parsers.NewJSONParser("exec", nil, nil)
	assert.NoError(t, err)
	ex.SetParser(p)
	ex.Command = "/usr/bin/myothercollector --foo=bar"
	eConfig := &internal_models.InputConfig{
		Name:              "exec",
		MeasurementSuffix: "_myothercollector",
	}
	eConfig.Tags = make(map[string]string)
	assert.Equal(t, ex, c.Inputs[1].Input,
		"Merged Testdata did not produce a correct exec struct.")
	assert.Equal(t, eConfig, c.Inputs[1].Config,
		"Merged Testdata did not produce correct exec metadata.")

	memcached.Servers = []string{"192.168.1.1"}
	assert.Equal(t, memcached, c.Inputs[2].Input,
		"Testdata did not produce a correct memcached struct.")
	assert.Equal(t, mConfig, c.Inputs[2].Config,
		"Testdata did not produce correct memcached metadata.")

	pstat := inputs.Inputs["procstat"]().(*procstat.Procstat)
	pstat.PidFile = "/var/run/grafana-server.pid"

	pConfig := &internal_models.InputConfig{Name: "procstat"}
	pConfig.Tags = make(map[string]string)

	assert.Equal(t, pstat, c.Inputs[3].Input,
		"Merged Testdata did not produce a correct procstat struct.")
	assert.Equal(t, pConfig, c.Inputs[3].Config,
		"Merged Testdata did not produce correct procstat metadata.")
}
