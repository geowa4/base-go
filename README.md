# Base Go v0.0.1

## Purpose

Serve as a foundation upon which I can build future Go projects.

## Contributing

### Quick Edit

If you're looking to make a quick contribution to the Go source, install Docker and Make.
With those, run `make ci` to run all the tests.

### Bigger Changes

Install these tools in addition to those for quick edits.

1. Go installation matching [.tool-versions](./.tool-versions)
1. [ag](https://geoff.greer.fm/ag/)
1. [entr](http://entrproject.org/)

Run `make dev` to run the application and watch for changes.
When done, run `make clean`.

### Pull Requests

If you find any bugs or have any suggestions, please open an issue on GitHub.
I will take pull requests for many things, but I will be more hesitant with new features.
