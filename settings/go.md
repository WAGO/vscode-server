---
description: Golang spesific settings
---

# Go

Here we can set the environment variables for GOOS (local OS) and GOARCH (host OS):

```json
"go.toolsEnvVars": {
        // For ARM 32-bit (Wago)
        "GOOS": "linux",  
        "GOARCH": "arm", 

        // For WSL:
        //"GOOS": "linux",  
        //"GOARCH": "amd64", 

        // For WINDOWS 
        //"GOOS": "windows",  
        //"GOARCH": "amd64", 
    },
```

