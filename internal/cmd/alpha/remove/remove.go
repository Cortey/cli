package remove

import (
	"github.com/kyma-project/cli.v3/internal/clierror"
	"github.com/kyma-project/cli.v3/internal/cmd/alpha/add/managed"
	"github.com/kyma-project/cli.v3/internal/cmdcommon"
	"github.com/kyma-project/cli.v3/internal/communitymodules/cluster"
	"github.com/spf13/cobra"
)

type removeConfig struct {
	*cmdcommon.KymaConfig
	cmdcommon.KubeClientConfig

	modules []string
}

func NewRemoveCMD(kymaConfig *cmdcommon.KymaConfig) *cobra.Command {
	cfg := removeConfig{
		KymaConfig:       kymaConfig,
		KubeClientConfig: cmdcommon.KubeClientConfig{},
	}

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Removes Kyma modules.",
		Long:  `Use this command to remove Kyma modules`,
		PreRun: func(_ *cobra.Command, _ []string) {
			clierror.Check(cfg.KubeClientConfig.Complete())
		},
		Run: func(_ *cobra.Command, _ []string) {
			clierror.Check(runRemove(&cfg))
		},
	}

	cmd.AddCommand(managed.NewManagedCMD(kymaConfig))

	cfg.KubeClientConfig.AddFlag(cmd)
	cmd.Flags().StringSliceVar(&cfg.modules, "module", []string{}, "Name and version of the modules to remove. Example: --module serverless,keda:1.1.1,etc...")
	return cmd
}

func runRemove(cfg *removeConfig) clierror.Error {
	modules := cluster.ParseModules(cfg.modules)
	return cluster.RemoveSpecifiedModules(cfg.Ctx, cfg.KubeClient.RootlessDynamic(), modules)
}
