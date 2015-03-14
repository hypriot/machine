package hypriot

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/docker/docker/api"
	"github.com/docker/machine/drivers"
	"github.com/docker/machine/provider"
	"github.com/docker/machine/state"
)

// The `hypriot` driver is used to control any Hypriot related
// Docker Host, like a Raspberry Pi connected through ssh.
type Driver struct {
	URL string
}

func init() {
	drivers.Register("hypriot", &drivers.RegisteredDriver{
		New:            NewDriver,
		GetCreateFlags: GetCreateFlags,
	})
}

func GetCreateFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "hypriot-docker-url",
			Usage: "URL of host when no driver is selected",
			Value: "",
		},
	}
}

func NewDriver(machineName string, storePath string, caCert string, privateKey string) (drivers.Driver, error) {
	return &Driver{}, nil
}

func (d *Driver) AuthorizePort(ports []*drivers.Port) error {
	return nil
}

func (d *Driver) Create() error {
	return nil
}

func (d *Driver) DeauthorizePort(ports []*drivers.Port) error {
	return nil
}

func (d *Driver) DriverName() string {
	return "hypriot"
}

func (d *Driver) GetIP() (string, error) {
	return "", nil
}

func (d *Driver) GetMachineName() string {
	return ""
}

func (d *Driver) GetSSHHostname() (string, error) {
	return "", nil
}

func (d *Driver) GetSSHKeyPath() string {
	return ""
}

func (d *Driver) GetSSHPort() (int, error) {
	return 0, nil
}

func (d *Driver) GetSSHUsername() string {
	return ""
}

func (d *Driver) GetURL() (string, error) {
	return d.URL, nil
}

func (d *Driver) GetState() (state.State, error) {
	return state.None, nil
}

func (d *Driver) GetProviderType() provider.ProviderType {
	return provider.None
}

func (d *Driver) Kill() error {
	return fmt.Errorf("hosts without a driver cannot be killed")
}

func (d *Driver) PreCreateCheck() error {
	return nil
}

func (d *Driver) Remove() error {
	return nil
}

func (d *Driver) Restart() error {
	return fmt.Errorf("hosts without a driver cannot be restarted")
}

func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	url := flags.String("hypriot-docker-url")

	if url == "" {
		return fmt.Errorf("--hypriot-docker-url option is required when no driver is selected")
	}
	validatedUrl, err := api.ValidateHost(url)
	if err != nil {
		return err
	}

	d.URL = validatedUrl
	return nil
}

func (d *Driver) Start() error {
	return fmt.Errorf("hosts without a driver cannot be started")
}

func (d *Driver) Stop() error {
	return fmt.Errorf("hosts without a driver cannot be stopped")
}
