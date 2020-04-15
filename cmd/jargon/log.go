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

	"github.com/arnumina/swag.jargon/internal/log"
)

const _defaultLogFile = "/var/log/swag/swag.log"

var file string

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Print the log file in real time",
	Args:  cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
		return log.TailFile(file)
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
	logCmd.Flags().StringVarP(&file, "file", "f", _defaultLogFile, "the log file to be printed")
}

/*
######################################################################################################## @(°_°)@ #######
*/
