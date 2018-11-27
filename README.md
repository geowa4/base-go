# Base Go v0.0.1

## Purpose

Serve as a foundation upon which I can build future Go projects.

## Contributing

### Quick Edit

If you're looking to make a quick contribution to the Go source, install Docker and Make.
With those, run `make ci` to run all the tests.

### Bigger Changes

Since this project uses Go modules, ensure this repo is cloned _outside_ your `$GOPATH`.

Install these tools in addition to those for quick edits.

1. Go installation matching [.tool-versions](./.tool-versions)
1. [ag](https://geoff.greer.fm/ag/)
1. [entr](http://entrproject.org/)

To build and run the application, run these commands.

1. `make deps`
1. `make dev`
1. `make clean`

Changing any static file (e.g. HTML, SQL), run `make embeds`.

### Pull Requests

If you find any bugs or have any suggestions, please open an issue on GitHub.
I will take pull requests for many things, but I will be more hesitant with new features.
