# Prerequisites

- Rockcraft is installed
- A ROCK has been build (i.e. you have a `*.rock` file)

# Instructions

- Copy the ROCK (which is an OCI archive) to the Docker daemon:

    ```bash
    # Assuming the ROCK name is "my-flask-app" and the version is "0.1", on amd64
    /snap/rockcraft/current/bin/skopeo copy oci-archive:my-flask-app_0.1_amd64.rock docker-daemon:sample:latest
    ```

- Inspect the ROCK:

  ```bash
  # With OCI management tools
  /snap/rockcraft/current/bin/skopeo inspect oci-archive:my-flask-app_0.1_amd64.rock 

  /snap/rockcraft/current/bin/skopeo inspect --raw oci-archive:my-flask-app_0.1_amd64.rock | jq

  # Or with a container runtime like Docker
  docker inspect sample
  ```

- Run the ROCK:

    ```bash
    docker run --rm sample
    ```

## Troubleshooting
