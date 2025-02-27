package linkctl

import (
	"flag"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/klog/v2"
	ctlutil "k8s.io/kubectl/pkg/cmd/util"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"

	"github.com/kosmos.io/kosmos/pkg/linkctl/get"
)

// DefaultConfigFlags It composes the set of values necessary for obtaining a REST client config with default values set.
var DefaultConfigFlags = genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag().WithDiscoveryBurst(300).WithDiscoveryQPS(50.0)

// NewLinkCtlCommand creates the `linkctl` command with arguments.
func NewLinkCtlCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use:   "linkctl",
		Short: i18n.T("linkctl controls a Kubernetes cluster network"),
		Long:  templates.LongDesc(`linkctl controls a Kubernetes cluster network.`),
		RunE:  runHelp,
	}

	klog.InitFlags(flag.CommandLine)

	pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	cmds.PersistentFlags().AddFlagSet(pflag.CommandLine)

	if err := flag.CommandLine.Parse(nil); err != nil {
		klog.Warning(err)
	}

	f := ctlutil.NewFactory(DefaultConfigFlags)
	ioStreams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}

	groups := templates.CommandGroups{
		{
			Message: "Basic Commands:",
			Commands: []*cobra.Command{
				get.NewCmdGet(f, ioStreams),
			},
		},
		//{
		//	Message: "Install Commands:",
		//	Commands: []*cobra.Command{
		//		install.NewCmdInstall(f, ioStreams),
		//		uninstall.NewCmdUninstall(f, ioStreams),
		//	},
		//},
		//{
		//	Message: "Cluster Member Join/UnJoin Commands:",
		//	Commands: []*cobra.Command{
		//		join.NewCmdJoin(),
		//		unjoin.NewCmdUnJoin(),
		//	},
		//},
		//{
		//	Message: "Cluster Doctor/Floater Commands:",
		//	Commands: []*cobra.Command{
		//		floater.NewCmdDoctor(),
		//	},
		//},
	}
	groups.Add(cmds)

	return cmds
}

func runHelp(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}
