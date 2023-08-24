# vscode-server

How to support MS VS-Code server on Wago Devices

## About

These packages enables Microsoft Visual Studio Code server to be installed on Wago 32-bit arm based devices. There are currently two methods to connect remote from the VS-Code IDE; with the Remote SSH extension or with the Remote Tunnels extension.

## Note!

VS-Code will download the server bits and need enough space. Also we need some space for apps and extensions. A good way to start is the default 'admin' user-space (not root).

_Currently Edge Controller 750-8303 is prefered beacause this device provides enough CPU power and RAM for the connection. Other devices will work, but are slower to connect etc._

## Building the IPK's

Using Wago SDK we can build the two IPK's with selecting _'core -> gcc libraries -> libatomic_' and 'Shell & Console Tools -> Busibox -> Coreutils -> readlink' and '_Enable canonicalization by following all symlinks (-f)_'.

Pay attention to that the ipk's may be rebuild with every firmware to match the allready installed version.

## Prerequisite

Install the two ipk's found in this repo with the Web Based Management (WBM) 'Configuration -> Software Upload' or with the command line interface:

```
opkg install --force-reinstall <IPK>
```

> Packages with the same name could allready excist (but without all features). With the new ipk's provided here only the content has changed therefor we get a warning. Using WBM we must select 'force install'. NB! If libraries need to be downgraded use 'opkg --downgrade' (not possible with WBM).

If we want to use the the remote SSH extension (instead of tunneling) we must change the settings for SSH-server Dropbear:

```
nano /etc/dropbear/dropbear.conf
```

Then enable for 'local port forwarding':

```
LOCAL_PORT_FORWARDING=true
```

Save the file and restart controller (or restart dropbear).

## Connect with SSH

Open VS-Code and install 'Remote - SSH' extension. Guideline is found in the extension.

[https://code.visualstudio.com/docs/remote/ssh](https://code.visualstudio.com/docs/remote/ssh)

## Connect with tunnels

Open VS-Code and install 'Remote - Tunnels' extension. Guideline is found in the extension.

Download and unpack the cli-server from microsoft.com, then start the tunnel as admin:

```
./code tunnel
```

[https://code.visualstudio.com/docs/remote/tunnels](https://code.visualstudio.com/docs/remote/tunnels)

## Extensions on remote host

It's possible to use several Vs-Code extensions on the Wago remote device, e.g. Docker.

## Tunnel as a service (workaround)

Please read ussuetracker [here](https://github.com/WAGO/vscode-server/issues/2). Create a tunnel as described in this repo then make a startup script of your own.

```bash
touch /etc/init.d/vscode && chmod +x /etc/init.d/vscode
```

Add symlink:

```bash
ln -s /etc/init.d/vscode /etc/rc.d/S99_z1_vscode
```

Edit 'vscode' script with this content:

```bash
cd /home/admin
su -c './code tunnel --accept-server-license-terms' admin
```

Reboot the controller and test the connection.
