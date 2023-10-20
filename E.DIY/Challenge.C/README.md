# Exercise: Challenge C

Remember the sample? The one with the Flask application...well, let's take it,
and rebase it on Ubuntu 22.04 instead, whilst keeping it chiselled and baseless.

**HINT:** at first glance, this seems easy. Here are a few tips to explain why
it's not, while shedding some light about what needs to be done:

- there are no Python slices for the Jammy (22.04) release
- installing slices from a custom Chisel release implies the use of "*-script"s

---

**ROCKSTART BADGE:** if you succeed, you can become an official contributor by
submitting your Python slice definitions to <https://github.com/canonical/chisel-releases>.

## Requirements

- [ ] The application must be the one from the sample
- [ ] This should be a 22.04 ROCK, with a `bare` base
- [ ] It must be chiselled
- [ ] You must upgrade from Python3.8 to Python3.11

## Expected Output

The behaviour should be the same as in the sample, but the Ubuntu and Python versions should both be newer.

## Useful Resources

- How to create slices: <https://canonical-rockcraft.readthedocs-hosted.com/en/latest/how-to/create-slice/>
- How to use Chisel: <https://canonical-rockcraft.readthedocs-hosted.com/en/latest/how-to/use-chisel/>
- How to install custom slices: <https://canonical-rockcraft.readthedocs-hosted.com/en/latest/explanation/overlay-step/>
- How to contribute to a Chisel release: <https://canonical-rockcraft.readthedocs-hosted.com/en/latest/how-to/publish-slice/>
- Part properties: <https://canonical-rockcraft.readthedocs-hosted.com/en/latest/reference/part_properties/>
- copying the OCI archive to Docker and running it: <https://canonical-rockcraft.readthedocs-hosted.com/en/latest/tutorials/hello-world/#run-the-rock-in-docker>
