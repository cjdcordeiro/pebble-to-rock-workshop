# Prerequisites

- Rockcraft is installed

# Instructions

- There is already a sample `rockcraft.yaml` file and `src` code in this folder.
  - If for some reason you don't have it, then simply use the templated
    one from `rockcraft init`.

- Pack the ROCK:

    ```bash
    rockcraft pack -v
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

This could be related with a [known networking conflict between Docker and
LXD](https://documentation.ubuntu.com/lxd/en/latest/howto/network_bridge_firewalld/#prevent-connectivity-issues-with-lxd-and-docker).
The following command may fix this:

```bash
sudo iptables -P FORWARD ACCEPT
```

If you are still getting the same error (e.g. `craft-providers error: Instance
setup failed.`), check if the lxd instance is running in the rockcraft project.

```bash
lxc --project rockcraft list
```

If it is in a `RUNNING` state, deleting it and running `rockcraft pack` again
might solve the issue.

```bash
lxc --project rockcraft delete <instance-id>
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
