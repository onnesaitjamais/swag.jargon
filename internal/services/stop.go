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

// Stop AFAIRE
func (c *Cmd) Stop(service, sdInstanceID string) error {
	return c.restartOrStop("stop", service, sdInstanceID)
}

/*
######################################################################################################## @(°_°)@ #######
*/
