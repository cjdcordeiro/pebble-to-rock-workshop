# Exercise: Challenge A

You ca find the Golang application "assistant" in this folder. The objective
is to create an Ubuntu ROCK called "challenge-a", that runs the `assistant` binary every time a container is deployed.

## Requirements

 - [ ] The ROCK must be named "challenge-a"
 - [ ] It must be based on Ubuntu 22.04
 - [ ] When the container starts, it must run the `assistant` and then exit
the container (both "on-success" and "on-failure")

## Expected Output

```bash
$ docker run challenge-a --args assistant -help
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


## Useful Resources

- `services` reference: <https://canonical-rockcraft.readthedocs-hosted.com/en/latest/reference/rockcraft.yaml/#services>
- Go plugin: <https://snapcraft.io/docs/go-plugin>
- copying the OCI archive to Docker and running it: <https://canonical-rockcraft.readthedocs-hosted.com/en/latest/tutorials/hello-world/#run-the-rock-in-docker>
