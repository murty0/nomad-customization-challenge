package command

import (
	"fmt"
	"io"
	"sort"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/hashicorp/nomad/api"
	"github.com/hashicorp/nomad/api/contexts"
	"github.com/pkg/errors"
	"github.com/posener/complete"
)

type VolumeSnapshotListCommand struct {
	Meta
}

func (c *VolumeSnapshotListCommand) Help() string {
	helpText := `
Usage: nomad volume snapshot list [-plugin plugin_id]

  Display a list of CSI volume snapshots for a plugin along
  with their source volume ID as known to the external
  storage provider.

  When ACLs are enabled, this command requires a token with the
  'csi-list-volumes' capability for the plugin's namespace.

General Options:

  ` + generalOptionsUsage(usageOptsDefault) + `

List Options:

  -plugin: Display only snapshots managed by a particular plugin. This
    parameter is required.

  -secrets: A set of key/value secrets to be used when listing snapshots.
`
	return strings.TrimSpace(helpText)
}

func (c *VolumeSnapshotListCommand) Synopsis() string {
	return "Display a list of volume snapshots for plugin"
}

func (c *VolumeSnapshotListCommand) AutocompleteFlags() complete.Flags {
	return mergeAutocompleteFlags(c.Meta.AutocompleteFlags(FlagSetClient),
		complete.Flags{})
}

func (c *VolumeSnapshotListCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictFunc(func(a complete.Args) []string {
		client, err := c.Meta.Client()
		if err != nil {
			return nil
		}

		resp, _, err := client.Search().PrefixSearch(a.Last, contexts.Plugins, nil)
		if err != nil {
			return []string{}
		}
		return resp.Matches[contexts.Plugins]
	})
}

func (c *VolumeSnapshotListCommand) Name() string { return "volume snapshot list" }

func (c *VolumeSnapshotListCommand) Run(args []string) int {
	var pluginID string
	var verbose bool
	var secrets string

	flags := c.Meta.FlagSet(c.Name(), FlagSetClient)
	flags.Usage = func() { c.Ui.Output(c.Help()) }
	flags.StringVar(&pluginID, "plugin", "", "")
	flags.BoolVar(&verbose, "verbose", false, "")
	flags.StringVar(&secrets, "secrets", "", "")

	if err := flags.Parse(args); err != nil {
		c.Ui.Error(fmt.Sprintf("Error parsing arguments %s", err))
		return 1
	}

	args = flags.Args()
	if len(args) > 0 {
		c.Ui.Error("This command takes no arguments")
		c.Ui.Error(commandErrorText(c))
		return 1
	}

	// Get the HTTP client
	client, err := c.Meta.Client()
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error initializing client: %s", err))
		return 1
	}

	plugs, _, err := client.CSIPlugins().List(&api.QueryOptions{Prefix: pluginID})
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error querying CSI plugins: %s", err))
		return 1
	}
	if len(plugs) == 0 {
		c.Ui.Error(fmt.Sprintf("No plugins(s) with prefix or ID %q found", pluginID))
		return 1
	}
	if len(plugs) > 1 {
		if pluginID != plugs[0].ID {
			out, err := c.csiFormatPlugins(plugs)
			if err != nil {
				c.Ui.Error(fmt.Sprintf("Error formatting: %s", err))
				return 1
			}
			c.Ui.Error(fmt.Sprintf("Prefix matched multiple plugins\n\n%s", out))
			return 1
		}
	}
	pluginID = plugs[0].ID

	q := &api.QueryOptions{PerPage: 30} // TODO: tune page size

	for {
		resp, _, err := client.CSIVolumes().ListSnapshots(pluginID, secrets, q)
		if err != nil && !errors.Is(err, io.EOF) {
			c.Ui.Error(fmt.Sprintf(
				"Error querying CSI external snapshots for plugin %q: %s", pluginID, err))
			return 1
		}
		if resp == nil || len(resp.Snapshots) == 0 {
			// several plugins return EOF once you hit the end of the page,
			// rather than an empty list
			break
		}

		c.Ui.Output(csiFormatSnapshots(resp.Snapshots, verbose))
		q.NextToken = resp.NextToken
		if q.NextToken == "" {
			break
		}
		// we can't know the shape of arbitrarily-sized lists of snapshots,
		// so break after each page
		c.Ui.Output("...")
	}

	return 0
}

func csiFormatSnapshots(snapshots []*api.CSISnapshot, verbose bool) string {
	rows := []string{"Snapshot ID|Volume ID|Size|Create Time|Ready?"}
	length := 12
	if verbose {
		length = 30
	}
	for _, v := range snapshots {
		rows = append(rows, fmt.Sprintf("%s|%s|%s|%s|%v",
			v.ID,
			limit(v.ExternalSourceVolumeID, length),
			humanize.IBytes(uint64(v.SizeBytes)),
			formatUnixNanoTime(v.CreateTime),
			v.IsReady,
		))
	}
	return formatList(rows)
}

func (c *VolumeSnapshotListCommand) csiFormatPlugins(plugs []*api.CSIPluginListStub) (string, error) {
	// TODO: this has a lot of overlap with 'nomad plugin status', so we
	// should factor out some shared formatting helpers.
	sort.Slice(plugs, func(i, j int) bool { return plugs[i].ID < plugs[j].ID })
	length := 30
	rows := make([]string, len(plugs)+1)
	rows[0] = "ID|Provider|Controllers Healthy/Expected|Nodes Healthy/Expected"
	for i, p := range plugs {
		rows[i+1] = fmt.Sprintf("%s|%s|%d/%d|%d/%d",
			limit(p.ID, length),
			p.Provider,
			p.ControllersHealthy,
			p.ControllersExpected,
			p.NodesHealthy,
			p.NodesExpected,
		)
	}
	return formatList(rows), nil
}
