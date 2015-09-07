<!--[metadata]>
+++
title = "Hypriot"
description = "Hypriot driver for machine"
keywords = ["machine, Hypriot, driver"]
[menu.main]
parent="smn_machine_drivers"
+++
<![end-metadata]-->

#### Hypriot
Create machines using an existing Raspberry Pi with installed [Hypriot](http://blog.hypriot.com/) image (based on Raspbian).

For more details on installing the basic system (including a preinstalled Docker daemon) please
see the detailled tutorial [Getting started with Docker on your Raspberry Pi](http://blog.hypriot.com/getting-started-with-docker-on-your-arm-device/).

Before using the `docker-machine create` command you have to copy your SSH keys to the device:

    $ ssh-copy-id root@$HYPRIOT_IP_ADDRESS

Now you can easily provision your Docker engine with:

    $ docker-machine create --driver hypriot --hypriot-ip-address=$HYPRIOT_IP_ADDRESS black-pearl

In order to upgrade to the latest available Docker version, just run:

    $ docker-machine upgrade black-pearl

Options:

 - `--hypriot-ip-address`: **required** IP Address of host.
 - `--hypriot-ssh-user`: SSH username used to connect.
 - `--hypriot-ssh-key`: Path to the SSH user private key.
 - `--hypriot-ssh-port`: Port to use for SSH.

> Note: currently only a numerical IP address is supported.

Environment variables and default values:

| CLI option                 | Environment variable | Default             |
|----------------------------|----------------------|---------------------|
| **`--hypriot-ip-address`** | -                    | -                   |
| `--hypriot-ssh-user`       | -                    | `root`              |
| `--hypriot-ssh-key`        | -                    | `$HOME/.ssh/id_rsa` |
| `--hypriot-ssh-port`       | -                    | `22`                |

Once you have created your Docker machine run this command to configure your shell:

    $ eval "$(docker-machine env black-pearl)"

Now you can use the Docker machine as usual:

    $ docker version
    Client version: 1.6.2
    Client API version: 1.18
    Go version (client): go1.4.2
    Git commit (client): 7c8fca2
    OS/Arch (client): darwin/amd64
    Server version: 1.6.2
    Server API version: 1.18
    Go version (server): go1.4.2
    Git commit (server): 7c8fca2
    OS/Arch (server): linux/arm
