# Prerequisites

- Rockcraft is installed

# Instructions

- There is already a sample `rockcraft.yaml` file in this folder.
  - If for some reason you don't have it, then simply use the templated
    one from `rockcraft init`.

- Pack the ROCK:

    ```bash
    rockcraft -v
    ```

- Choose a parts lifecycle `<stage>` and try to `--shell-after`:

    ```bash
    rockcraft prime --shell-after
    # You'll then be inside the LXC builder instance
    # and you'll be able to inspect your project contents
    # right after the Rockcraft <stage> was executed.
    ```

- If you didn't run with enough verbosity, you can still find the logs under
`/home/cris/.local/state/rockcraft/log/`.

## Troubleshooting

**Got an LXC-related timeout while building the ROCK?**

This could be related with a known networking conflict between Docker and LXD.
This this quick fix:

```bash
sudo iptables -P FORWARD ACCEPT
```

**Are you running Rockcraft in an LXC instance? Is it failing with an LXD error?**

If you see something like

> Error: Failed to run: /snap/lxd/current/bin/lxd forkstart

or

> Permission denied - Failed to mount "proc" on "/var/snap/lxd/common/lxc//proc" with flags 14

then you should try configuring your parent LXC instance to support nesting:

1. exit the LXC instance where you are running Rockcraft from,
2. run: `lxc config set <instance-name> security.nesting=true`,
3. go back inside the instance and retry packing your ROCK.
