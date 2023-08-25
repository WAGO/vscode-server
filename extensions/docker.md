# Docker

## Prerequisite

Enable for Docker in Web Based Management.

Root owns docker.sock and the docker-group can connect but not "others" like admin-user from VS-Code.

```
srw-rw----    1 root     docker           0 Jul  7 10:04 /var/run/docker.sock
```

Add others:

```
chmod o+rwx /var/run/docker.sock
```

## Connect with tunnel (workaround)

Connect to controller with remote tunnel then install 'Docker' extension. Guidelines is found in the extension. Please read issuetracker [here](https://github.com/WAGO/vscode-server/issues/4).

Pull a 'Debian' container, then right click the image and select 'run interactive'. Now the container shows up in the 'containers view' and the terminal connects to the container. Prepare the container with the fallowing:

```
apt update
apt Install libatomic1
apt install --reinstall ca-certificates
```

Download and run the code-server CLI in the container. Afterwards you should find the tunnel as normal in the remote explorer.

## Connect with SSH

Connect to controller with SSH then install 'Docker' extension. Guidelines is found in the extension. Please read issuetracker [here](https://github.com/WAGO/vscode-server/issues/6).

Pull a 'Debian' container, then right click the image and select 'run interactive'. Now the container shows up in the 'containers view' and the terminal connects to the container. Prepare the container with the fallowing:

```
apt update
apt Install libatomic1
```

Right click the container in the containers-view and select 'Attach Visual Studio Code'.
