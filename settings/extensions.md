---
description: Extension spesific settings
---

# Extensions

We can install extensions automatically when connecting. In the example we install a github add-on:

```json
{
"git.enableSmartCommit": true,
"remote.SSH.defaultExtensions": [
    "GitHub.vscode-pull-request-github"
],
}

```

To get the id of extensions:

```bash
code --list-extensions
```
