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
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func header(w io.Writer) {
	fmt.Fprintln(w, "\tswag.jargon")
	fmt.Fprintln(w, "=================================================================================================")
}

func footer(w io.Writer) {
	fmt.Fprintln(w, "-------------------------------------------------------------------------------------------------")
}

func helpFunc(hf func(*cobra.Command, []string)) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		header(cmd.OutOrStdout())
		hf(cmd, args)
		footer(cmd.OutOrStdout())
	}
}

// Run AFAIRE
func Run(version, builtAt string) error {
	ts, err := strconv.ParseInt(builtAt, 0, 64)
	if err != nil {
		return err
	}

	root := &cobra.Command{
		Use:   "swag.jargon",
		Short: "The swag command line client",
	}

	root.Version = fmt.Sprintf(
		"%s built at %s by Archivage Numérique © INA %d\n",
		version,
		time.Unix(ts, 0).Local().String(),
		time.Now().Year(),
	)

	root.SetHelpFunc(helpFunc(root.HelpFunc()))

	addLog(root)
	addServices(root)

	return root.Execute()
}

/*
######################################################################################################## @(°_°)@ #######
*/
