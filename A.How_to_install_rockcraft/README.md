# Prerequisites

- snap enabled system (<https://snapcraft.io>)
- LXD installed (<https://linuxcontainers.org/lxd/getting-started-cli/>)
- skopeo installed (<https://github.com/containers/skopeo>)
- Docker installed (<https://snapcraft.io/docker>)
- a text editor

# Instructions

- Install the Rockcraft snap:

    ```bash
    # Use --stable if available
    sudo snap install rockcraft --classic --edge
    ```

- Make sure it's installed:

    ```bash
    rockcraft --version
    ```

- Check what commands are available:

    ```bash
    rockcraft help --all
    ```

- Start a new Rockcraft project:

    ```bash
    rockcraft init
    ```

## Troubleshooting

**Failed to install the Rockcraft snap?**

If it is a networking or store availability issue, then you'll find a backup
snap file in this directory that you can use to install Rockcraft:

```bash
sudo snap install --classic --dangerous rockcraft_amd64.snap
```
