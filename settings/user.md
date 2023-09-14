---
description: User spesific settings
---

# User

Default VS-Code will download the server to the directory belonging to user launching it. If we use a memory card we can extend the space creating a symlink from default VS-Code directory to the new directory on the memory card:

```
// Log in as root
cd /home
useradd <new user>
passwd <new user>
mv <new user> /media/sd
ln -s /media/sd/<new user> <new user>
chown -h <new user>:<new user> <new user>
```

For remote SSH configure the .ssh-config parameters to the newly created user. For remote tunnel start the tunnel as the newly created user and VS-Code will download the server bits to SD card:

```
su <user>
./code tunnel
```



