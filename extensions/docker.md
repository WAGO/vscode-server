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

## Connect with Docker

Connect with SSH or Tunnels then install 'Docker' extension. Guidelines is found in the extension.
