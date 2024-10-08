# vscode-server

[Microsoft Visual Studio Code server](https://code.visualstudio.com/Download) is supported from firmware 27. For older firmware on 32-bit devices use packages provided in this repo. There are currently two methods to connect remote from the VS-Code IDE; with the Remote SSH extension or with the Remote Tunnels extension.

#### See also [issues](https://github.com/WAGO/vscode-server/issues) and [discussion ](https://github.com/WAGO/vscode-server/discussions)section.

## Note!

VS-Code will download the server bits and need enough space. Also we need some space for apps and extensions. A good way to start is the default 'admin' user-space (not root).

_PFC300 or Edge Controller is prefered because these devices provides enough CPU power and RAM for the connection. Other devices will also work, but are slower to connect etc._

## Prerequisite

Install the ipk's in Web Based Management (WBM) software upload tab (select 'force install' if notified).

> Can be skipped for FW >= 27.

If we want to use the the remote SSH extension (instead of tunneling) we must change the settings for SSH-server Dropbear:

```sh
nano /etc/dropbear/dropbear.conf
```

Then enable for 'local port forwarding':

```sh
LOCAL_PORT_FORWARDING=true
```

Save the file and restart controller (or restart dropbear).

## Connect with SSH

Open VS-Code and install 'Remote - SSH' extension. Guideline is found in the extension.

[https://code.visualstudio.com/docs/remote/ssh](https://code.visualstudio.com/docs/remote/ssh)

## Connect with tunnels

Open VS-Code and install 'Remote - Tunnels' extension. Guideline is found in the extension.

Download and unpack the cli-server from microsoft.com, then start the tunnel as admin:

```sh
./code tunnel
```

[https://code.visualstudio.com/docs/remote/tunnels](https://code.visualstudio.com/docs/remote/tunnels)

## Nice to know about

#### Exension usage:

* [Docker](extensions/docker.md)
* [Live Preview](extensions/live-preview.md)
* [Git](extensions/git.md)
* [Python](extensions/python.md)
* [Go](extensions/go.md)

#### Various settings:

* [Terminal](settings/terminal.md)
* [User](settings/user.md)
* [Extensions](settings/extensions.md)
* [Git](settings/git.md)
* [Go](settings/go.md)

Exampels:

* [Go](examples/go/opc-ua.md)

## Tunnel as a service

Create a tunnel as described in this repo then make a startup script of your own.

```sh
touch /etc/init.d/vscode && chmod +x /etc/init.d/vscode
```

Add symlink:

```sh
ln -s /etc/init.d/vscode /etc/rc.d/S99_z1_vscode
```

Edit 'vscode' script with this content:

```sh
cd /home/admin
su -c './code tunnel --accept-server-license-terms' admin
```

Reboot the controller and test the connection.
