---
description: Go extension from Google
---

# Go

The Go extension describes some initial steps like installing and using GO, but not how to configure for remote developement on a Wago controller. In this case we don't use Visual Studio Code server to connect to the controller because GO is installed on our developement PC, not remote.&#x20;

Add task.json to build, transfer and execute the program. Here we use [PuTTY Secure Copy Protocol](https://www.chiark.greenend.org.uk/\~sgtatham/putty/latest.html) instead of the native SSH because PuTTY also transfer password information. Afterwards execute "Meny->terminal->run task" and select the task "BUILD", "TRANSFER or "EXECUTE".&#x20;

### task.json

```json
{
	
	"version": "2.0.0",
	"tasks": [
        {
            "type": "go",
            "label": "GO BUILD",
            "command": "build",
            "args": [
                "${fileDirname}"
            ],
            "problemMatcher": [
                "$go"
            ],
            "group": {
                "kind": "build",
                "isDefault": true
            }
        },
        {
            "type": "shell",
            "label": "TRANSFER",
            "command": "pscp.exe",
            "args": [
                "-l",
                "<YOUR USER>",
                "-pw",
                "<YOUR PASSWORD>",
                "${workspaceFolder}/<MODULE_NAME>",
                "<USER>@<IP_ADDR>:"
            ],
            "group": "build",
            "problemMatcher": [
                "$tsc"
            ]
        },
        {
            "type": "shell",
            "label": "EXECUTE",
            "command": "plink.exe",
            "isBackground": false,
            "args": [
                "<IP_ADDR>",
                "-l",
                "root",
                "-pw",
                "wago",
                "chmod +x <PATH>/<MODULE_NAME> && <PATH>/<MODULE_NAME>"
            ],
            "group": "build",
            "problemMatcher": [
                "$tsc"
            ]
        }
    ]
}
```
