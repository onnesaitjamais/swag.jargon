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

package services

import (
	"fmt"
	"os"
	"os/exec"
)

// Start AFAIRE
func Start(services []string) error {
	if len(services) == 0 {
		if os.Getuid() != 0 {
			return nil
		}

		return exec.Command("systemctl", "start", fmt.Sprintf("swag.%s@0.service", _bsName)).Run() //nolint:gosec
	}

	for _, name := range services {
		if name == _bsName {
			continue
		}

		if err := doOne("start", name, ""); err != nil {
			return err
		}
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
