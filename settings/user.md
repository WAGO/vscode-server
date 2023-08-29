# User

Default VS-Code will download the server to the user-directory belonging to this user launching it. If we use a memory card we can extend the disk space creating a symlink from default directory to the new directory on the memory card like this:

```
// Log in as root and change "myusername" to e.g. vscode..
cd /home
useradd username "myusername"
usermod -a -G admin "myusername"
ln -s "/media/myusername" "myusername"
```

Start the tunnel as you created user and VS-Code will download the server bits to SD card:

```
su "myusername"
./code tunnel
```
