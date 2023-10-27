# Prerequisites

- Rockcraft is installed

# Instructions

- Grab the `rockcraft.yaml` and `src` samples (from the previous steps).

- Exercise the following:
  - change the `base` to `bare`,
  - install the `python3.8_standard` and `libgcc-s1_libs` slices instead of the
  whole Python Debian package
  - you'll also need an env var `PYTHONPATH` set to
  `/lib/python3.8/site-packages/`

- Pack the ROCK:

    ```bash
    rockcraft pack -v
    ```

- Convert it to the Docker daemon and run it (like in the previous step).

- If it runs, double check the size and compare with the previous version.
