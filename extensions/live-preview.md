---
description: Local web server extension from Microsoft
---

# Live Preview

## Prerequisite

1. Enable HTTP in WBM.
2. SSH connection.&#x20;

## Demo

Download or make a sample webpage to admin-user space. Open this folder.

CTRL+SHIFT+p and 'Live Preview: 'Start server' or pres the icon 'Show Preview' at top right corner.

<figure><img src="../.gitbook/assets/image (1).png" alt=""><figcaption></figcaption></figure>

## Wago Lighttpd server

To start making web pages with the lighttpd server we could use this extension for testing the result as we allready have a connection to the device.

```bash
// Quick and dirty
chown admin /var/www
```

Open /var/www folder in VS-Code and create a test page. Then start the live preview server as described above to see the result.
