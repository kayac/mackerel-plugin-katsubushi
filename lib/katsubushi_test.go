package katsubushi

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphDefinition(t *testing.T) {
	var katsubushi Plugin

	graphdef := katsubushi.GraphDefinition()
	if len(graphdef) != 4 {
		t.Errorf("GetTempfilename: %d should be 4", len(graphdef))
	}
}

func TestParse(t *testing.T) {
	var katsubushi Plugin
	stub := `STAT pid 8018
STAT uptime 17
STAT time 1487754986
STAT version 1.1.2
STAT curr_connections 10
STAT total_connections 123
STAT cmd_get 4
STAT get_hits 3
STAT get_misses 1
END
`

	katsubushiStats := bytes.NewBufferString(stub)

	stat, err := katsubushi.parseStats(katsubushiStats)
	fmt.Println(stat)
	assert.Nil(t, err)
	// Katsubushi Stats
	assert.EqualValues(t, stat["cmd_get"], 4)
	assert.EqualValues(t, stat["get_hits"], 3)
	assert.EqualValues(t, stat["get_misses"], 1)
	assert.EqualValues(t, stat["curr_connections"], 10)
	assert.EqualValues(t, stat["total_connections"], 123)
	assert.EqualValues(t, stat["uptime"], 17)
}
