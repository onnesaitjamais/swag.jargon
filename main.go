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

package main

import (
	"os"

	"github.com/arnumina/swag.jargon/cmd/jargon"
)

var version, builtAt string

func main() {
	if jargon.Run(version, builtAt) != nil {
		os.Exit(-1)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
