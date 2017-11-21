package katsubushi

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

// Plugin mackerel plugin for katsubushi
type Plugin struct {
	Target   string
	Socket   string
	Tempfile string
	Prefix   string
}

// MetricKeyPrefix interface for PluginWithPrefix
func (m Plugin) MetricKeyPrefix() string {
	if m.Prefix == "" {
		m.Prefix = "katsubushi"
	}
	return m.Prefix
}

// FetchMetrics interface for mackerelplugin
func (m Plugin) FetchMetrics() (map[string]float64, error) {
	network := "tcp"
	target := m.Target
	if m.Socket != "" {
		network = "unix"
		target = m.Socket
	}
	conn, err := net.Dial(network, target)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	fmt.Fprint(conn, "STATS\r\n")

	ret, err := m.parseStats(conn)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (m Plugin) parseStats(conn io.Reader) (map[string]float64, error) {
	scanner := bufio.NewScanner(conn)
	stat := make(map[string]float64)

	for scanner.Scan() {
		line := scanner.Text()
		s := string(line)
		if s == "END" {
			stat["new_items"] = stat["total_items"]
			return stat, nil
		}

		res := strings.Split(s, " ")
		if res[0] == "STAT" {
			f, err := strconv.ParseFloat(res[2], 64)
			if err == nil {
				stat[res[1]] = f
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return stat, err
	}
	return nil, nil
}

// GraphDefinition interface for mackerelplugin
func (m Plugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(m.Prefix)

	// https://github.com/kayac/go-katsubushi#stats
	var graphdef = map[string]mp.Graphs{
		"connections": {
			Label: (labelPrefix + " Connections"),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "curr_connections", Label: "Connections"},
				{Name: "total_connections", Label: "New Connections", Diff: true},
			},
		},
		"cmd": {
			Label: (labelPrefix + " Command"),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "cmd_get", Label: "Get", Diff: true},
			},
		},
		"hitmiss": {
			Label: (labelPrefix + " Hits/Misses"),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "get_hits", Label: "Get Hits", Diff: true},
				{Name: "get_misses", Label: "Get Misses", Diff: true},
			},
		},
		"uptime": {
			Label: (labelPrefix + " Uptime"),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "uptime", Label: "Uptime"},
			},
		},
	}
	return graphdef
}

// Do the plugin
func Do() {
	optHost := flag.String("host", "localhost", "Hostname")
	optPort := flag.String("port", "11212", "Port")
	optSocket := flag.String("socket", "", "Server socket (overrides hosts and port)")
	optPrefix := flag.String("metric-key-prefix", "katsubushi", "Metric key prefix")
	optTempfile := flag.String("tempfile", "", "Temp file name")
	flag.Parse()

	var katsubushi Plugin

	katsubushi.Prefix = *optPrefix

	if *optSocket != "" {
		katsubushi.Socket = *optSocket
	} else {
		katsubushi.Target = fmt.Sprintf("%s:%s", *optHost, *optPort)
	}
	helper := mp.NewMackerelPlugin(katsubushi)
	helper.Tempfile = *optTempfile
	helper.Run()
}
