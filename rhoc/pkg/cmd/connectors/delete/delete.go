package delete

import (
	"github.com/bf2fc6cc711aee1a0c2a/cos-tools/rhoc/pkg/cmd/commands"
	"github.com/bf2fc6cc711aee1a0c2a/cos-tools/rhoc/pkg/service"
	"github.com/redhat-developer/app-services-cli/pkg/core/cmdutil/flagutil"
	"github.com/redhat-developer/app-services-cli/pkg/core/ioutil/dump"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
	"strconv"
)

type options struct {
	outputFormat string
	id           string
	force        bool
	skipConfirm  bool

	f *factory.Factory
}

func NewDeletesCommand(f *factory.Factory) *cobra.Command {
	opts := options{
		f: f,
	}

	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "delete",
		Long:    "delete",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.outputFormat != "" && !flagutil.IsValidInput(opts.outputFormat, flagutil.ValidOutputFormats...) {
				return flagutil.InvalidValueError("output", opts.outputFormat, flagutil.ValidOutputFormats...)
			}

			return run(&opts)
		},
	}

	flags := flagutil.NewFlagSet(cmd, f.Localizer)
	flags.AddOutput(&opts.outputFormat)
	flags.AddYes(&opts.skipConfirm)
	flags.StringVar(&opts.id, "id", "", "id")
	flags.BoolVarP(&opts.force, "force", "f", false, "force")

	cmd.MarkFlagRequired("id")

	return cmd
}

func run(opts *options) error {
	if !opts.skipConfirm {
		confirm, promptErr := commands.PromptConfirm("Are you sure you want to delete the connector with id '%s'?", opts.id)
		if promptErr != nil {
			return promptErr
		}
		if !confirm {
			opts.f.Logger.Debug("User has chosen to not delete connector")
			return nil
		}
	}

	c, err := service.NewAdminClient(&service.Config{
		F: opts.f,
	})
	if err != nil {
		return err
	}

	e := c.ConnectorClustersAdminApi.DeleteConnector(opts.f.Context, opts.id)
	e = e.Force(strconv.FormatBool(opts.force))

	response, httpRes, err := e.Execute()
	if httpRes != nil {
		defer func() {
			_ = httpRes.Body.Close()
		}()
	}
	if err != nil {
		return err
	}

	if httpRes != nil && httpRes.StatusCode == 204 {
		return nil
	}

	return dump.Formatted(opts.f.IOStreams.Out, opts.outputFormat, response)
}