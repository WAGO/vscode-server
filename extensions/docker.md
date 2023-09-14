---
description: Docker extension from Microsoft
---

# Docker

## Prerequisite

Enable for Docker in Web Based Management.

Root owns docker.sock and the docker-group can connect but not "others" like admin-user from VS-Code.

```bash
srw-rw----    1 root     docker           0 Jul  7 10:04 /var/run/docker.sock
```

Add others:

```bash
chmod o+rwx /var/run/docker.sock
```

## Connect with tunnel (workaround)

Connect to controller with remote tunnel then install 'Docker' extension. Guidelines is found in the extension.&#x20;

Pull a 'Debian' container, then right click the image and select 'run interactive'. Now the container shows up in the 'containers view' and the terminal connects to the container. Prepare the container with the fallowing:

```bash
apt update
apt Install libatomic1
apt install --reinstall ca-certificates
```

Download and run the code-server CLI in the container. Afterwards you should find the tunnel as normal in the remote explorer.

## Connect with SSH

Connect to controller with SSH then install 'Docker' extension. Guidelines is found in the extension.&#x20;

Pull a 'Debian' container, then right click the image and select 'run interactive'. Now the container shows up in the 'containers view' and the terminal connects to the container. Prepare the container with the fallowing:

```bash
apt update
apt Install libatomic1
```

Right click the container in the containers-view and select 'Attach Visual Studio Code'.&#x20;

## Building with Dockerfile

Create a Dockerfile in admin user space with "New File" button and create some content:

```docker
FROM debian
RUN apt update
RUN apt install libatomic1
CMD echo "Hello"
```

Right click the Dockerfile and 'Build Image..'. The image can now be found in the Docker extension.

## Building dev container

Its also passible to build a developement container. Connect to controller with SSH then install 'Dev Container' extension. Guidelines is found in the extension. Currently there is only support for 64-bit containers.

### Workaround for 32-bit

Make a Dockerfile. Hit CTRL+SHIFT+p to open command palette and 'Dev Containers: Add Dev Containers Configuration Files'.  Choose 'From Dockerfile'.

To run dev container open command palette and ''Dev containers: Reopen in container' to build and connect to the new container specified in the Dockerfile.





