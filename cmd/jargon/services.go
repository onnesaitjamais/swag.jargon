/*
#######
##                                     _
##        ____    _____ ____ _        (_)__ ________ ____  ___
##       (_-< |/|/ / _ `/ _ `/ _     / / _ `/ __/ _ `/ _ \/ _ \
##      /___/__,__/\_,_/\_, / (_) __/ /\_,_/_/  \_, /\___/_//_/
##                     /___/     |___/         /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package jargon

import (
	"github.com/spf13/cobra"

	"github.com/arnumina/swag.jargon/internal/services"
)

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Print the list of services (+start, +stop, +restart)",
	ValidArgs: []string{
		"basic",
		"dark",
		"double",
		"light",
		"simple",
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			return err
		}

		return cobra.OnlyValidArgs(cmd, args)
	},
	RunE: func(_ *cobra.Command, args []string) error {
		var style string

		if len(args) != 0 {
			style = args[0]
		}

		return services.List(style)
	},
}

var servicesStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start service(s)",
	RunE: func(_ *cobra.Command, args []string) error {
		return services.Start(args)
	},
}

var servicesStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop service(s)",
	Args:  cobra.MaximumNArgs(2),
	RunE: func(_ *cobra.Command, args []string) error {
		var (
			serviceName  string
			sdInstanceID string
		)

		switch len(args) {
		case 2:
			sdInstanceID = args[1]
			fallthrough
		case 1:
			serviceName = args[0]
		}

		return services.Stop(serviceName, sdInstanceID)
	},
}

var servicesRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart service(s)",
	Args:  cobra.MaximumNArgs(2),
	RunE: func(_ *cobra.Command, args []string) error {
		var (
			serviceName  string
			sdInstanceID string
		)

		switch len(args) {
		case 2:
			sdInstanceID = args[1]
			fallthrough
		case 1:
			serviceName = args[0]
		}

		return services.Restart(serviceName, sdInstanceID)
	},
}

func init() {
	rootCmd.AddCommand(servicesCmd)

	servicesCmd.AddCommand(servicesStartCmd)
	servicesCmd.AddCommand(servicesStopCmd)
	servicesCmd.AddCommand(servicesRestartCmd)
}

/*
######################################################################################################## @(°_°)@ #######
*/
