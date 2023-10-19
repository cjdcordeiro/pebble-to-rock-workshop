# Exercise: Challenge B

You can find the Golang application "assistant" in this folder. The objective
is to create an Ubuntu ROCK called "challenge-b", that runs the `assistant` binary every time a container is deployed.

## Requirements

- [ ] The ROCK must be named "challenge-b"
- [ ] It must be built on Ubuntu 22.04, but have a `bare` base
- [ ] The ROCK's default user must be `_daemon_`
- [ ] The image must have an environment variable called `AUTHOR`, with your name
- [ ] The `assistant` binary must be built with Go>=1.21
- [ ] The ROCK must fetch the `README.md` file from <https://github.com/cjdcordeiro/pebble-to-rock-workshop.git>, rename it to `NOTES` and place it
and place inside the ROCK, at `/`
- [ ] The ROCK must also have a file `/build/date` with the build's timestamp
- [ ] When the container starts, it must run the `assistant` and then exit
the container (both "on-success" and "on-failure")

## Expected Outputs

The `assistant` program is executed with the image's entrypoint and can accept
arguments:

```bash
docker run challenge-b --args assistant -help
```

```log
2023-10-18T12:37:04.381Z [pebble] Started daemon.
2023-10-18T12:37:04.385Z [pebble] POST /v1/services 4.091224ms 202
2023-10-18T12:37:04.385Z [pebble] Started default services with change 1.
2023-10-18T12:37:04.389Z [pebble] Service "assistant" starting: assistant [ -help ]
2023-10-18T12:37:05.390Z [assistant] Usage: assistant [options]
2023-10-18T12:37:05.390Z [assistant] Options:
2023-10-18T12:37:05.390Z [assistant]   -hello
2023-10-18T12:37:05.390Z [assistant]            Print a greeting message
2023-10-18T12:37:05.390Z [assistant]   -help
2023-10-18T12:37:05.390Z [assistant]            Show help message
2023-10-18T12:37:05.390Z [assistant]   -inspire-me
2023-10-18T12:37:05.390Z [assistant]            Get inspired by a quote
2023-10-18T12:37:05.390Z [assistant]   -time
2023-10-18T12:37:05.390Z [assistant]            Print current time
2023-10-18T12:37:05.390Z [pebble] Service "assistant" stopped unexpectedly with code 0
2023-10-18T12:37:05.390Z [pebble] Service "assistant" on-success action is "shutdown", triggering server exit
2023-10-18T12:37:05.390Z [pebble] Server exiting!
```

---

The `NOTE` (previously named `README.md`) file from the upstream GitHub
repository is under `/`:

```bash
docker run challenge-b ls /
```

```log
2023-10-19T12:26:43.591Z [pebble] Started daemon.
2023-10-19T12:26:43.592Z [pebble] GET /v1/files?action=list&path=%2F 342.027µs 200
.dockerenv
.rock
NOTES
bin
build
dev
etc
proc
sys
var
```

---

The ROCK's build date is written in a file, at `/build/date`:

```bash
docker run challenge-b --args assistant -time -read-file /build/date
```

```log
2023-10-19T12:26:14.433Z [pebble] Started daemon.
2023-10-19T12:26:14.438Z [pebble] POST /v1/services 4.130046ms 202
2023-10-19T12:26:14.438Z [pebble] Started default services with change 1.
2023-10-19T12:26:14.442Z [pebble] Service "assistant" starting: assistant [ -time -read-file /build/date ]
2023-10-19T12:26:15.443Z [assistant] Current time: 2023-10-19 12:26:15
2023-10-19T12:26:15.443Z [assistant] Thu Oct 19 14:23:23 CEST 2023
2023-10-19T12:26:15.443Z [assistant] 
2023-10-19T12:26:15.443Z [pebble] Service "assistant" stopped unexpectedly with code 0
2023-10-19T12:26:15.443Z [pebble] Service "assistant" on-success action is "shutdown", triggering server exit
2023-10-19T12:26:15.443Z [pebble] Server exiting!
```

---

The ROCK's default user is not `root`, but `_daemon_`:

```bash
docker inspect challenge-b -f '{{.Config.User}}'
```

```log
_daemon_
```

---

The `AUTHOR` environment variable is part of the ROCK:

```bash
docker inspect challenge-b -f '{{.Config.Env}}'
```

```log
[AUTHOR=Cristovao Cordeiro]
```

---

Finally, for fun, the image has no `ubuntu:22.04` base:

```bash
dive challenge-b:latest
```

## Useful Resources

- rockcraft.yaml reference: <https://canonical-rockcraft.readthedocs-hosted.com/en/latest/reference/rockcraft.yaml/>
- plugins: <https://canonical-rockcraft.readthedocs-hosted.com/en/latest/reference/part_properties/#plugin>
- "organize" property: <https://canonical-rockcraft.readthedocs-hosted.com/en/latest/reference/part_properties/#organize>
- "stage" property: <https://canonical-rockcraft.readthedocs-hosted.com/en/latest/reference/part_properties/#stage>
- Step execution environment: <https://craft-parts.readthedocs.io/en/latest/reference/parts_steps.html#step-execution-environment>
- copying the OCI archive to Docker and running it: <https://canonical-rockcraft.readthedocs-hosted.com/en/latest/tutorials/hello-world/#run-the-rock-in-docker>
