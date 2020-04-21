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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/arnumina/swag/component/registry"
	"github.com/arnumina/swag/util/failure"
	"github.com/jedib0t/go-pretty/table"
)

// Cmd AFAIRE
type Cmd struct {
	BSName string
	BSPort int
}

func (c *Cmd) url() string {
	return fmt.Sprintf("http://:%d/api/v1/services", c.BSPort)
}

func (c *Cmd) listServices() (registry.Services, error) {
	res, err := http.Get(c.url())
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil,
			failure.New(nil).
				Set("status", res.StatusCode).
				Msg(string(data)) //////////////////////////////////////////////////////////////////////////////////////
	}

	var services registry.Services

	if err := json.Unmarshal(data, &services); err != nil {
		return nil, err
	}

	return services, nil
}

func (c *Cmd) render(style string, services registry.Services) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(
		table.Row{
			"#", "ID", "NAME", "VERSION", "BUILT", "FQDN", "PORT", "INSTANCE", "STATUS", "UPTIME", "HEARTBEAT",
		},
	)

	for index, s := range services {
		t.AppendRow([]interface{}{
			index + 1,
			s.ID,
			s.Name,
			s.Version,
			s.BuiltAt.Local().String(),
			s.FQDN,
			s.Port,
			s.SdInstance,
			s.Status,
			time.Since(s.StartedAt).Round(time.Second).String(),
			time.Since(s.Heartbeat).Round(time.Second).String(),
		})
	}

	switch style {
	case "basic":
		t.SetStyle(table.StyleDefault)
	case "dark":
		t.SetStyle(table.StyleColoredDark)
	case "double":
		t.SetStyle(table.StyleDouble)
	case "light":
		t.SetStyle(table.StyleColoredBright)
	case "simple":
		t.SetStyle(table.StyleLight)
	default:
		t.SetStyle(table.StyleColoredDark)
	}

	t.SetCaption("The list of services.\n")
	t.Render()
}

// List AFAIRE
func (c *Cmd) List(style string) error {
	services, err := c.listServices()
	if err != nil {
		return err
	}

	c.render(style, services)

	return nil
}

func (c *Cmd) doOne(action, service, sdInstance string) error {
	var instance string

	if sdInstance != "" {
		instance = "/" + sdInstance
	}

	res, err := http.Get(fmt.Sprintf("%s/%s/%s%s", c.url(), action, service, instance))
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent {
		return failure.New(nil).
			Set("status", res.StatusCode).
			Msg(string(data)) //////////////////////////////////////////////////////////////////////////////////////////
	}

	return nil
}

func (c *Cmd) restartOrStop(action, service, sdInstance string) error {
	services, err := c.listServices()
	if err != nil {
		return err
	}

	var bsSdInstance string

	for _, s := range services {
		if service == "" || s.Name == service {
			if sdInstance == "" || s.SdInstance == sdInstance {
				if s.Name == c.BSName {
					bsSdInstance = s.SdInstance
					continue
				}

				if err = c.doOne(action, s.Name, s.SdInstance); err != nil {
					return err
				}
			}
		}
	}

	if bsSdInstance != "" {
		return c.doOne(action, c.BSName, bsSdInstance)
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
