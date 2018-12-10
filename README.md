# Base Go v0.0.2

## Purpose

Serve as a foundation upon which I can build future Go projects.

## Contributing

### Getting Started

Since this project uses Go modules, ensure this repo is cloned _outside_ your `$GOPATH`.

Install these tools in addition to those for quick edits.

1. Go installation matching [.tool-versions](./.tool-versions)
1. Make
1. Docker
1. [ag](https://geoff.greer.fm/ag/)
1. [entr](http://entrproject.org/)

To build and run the application, run these commands.

1. `make deps`
1. `make clean dev`
1. `make clean`

Changing any static file (e.g. HTML, SQL), run `make embeds`.

### API Testing

A [Postman](https://www.getpostman.com/) collection is provided in [build/postman](./build/postman).
This can either be imported into Postman or run on the command line with [Newman](https://learning.getpostman.com/docs/postman/collection_runs/command_line_integration_with_newman/).

### Pull Requests

If you find any bugs or have any suggestions, please open an issue on GitHub.
I will take pull requests for many things, but I will be more hesitant with new features.

For each PR, a test suite will run to ensure correctness.
To see what commands it runs, see [build/ci](./build/ci).

### Building a Release

```bash
make release
```

This will generate static binaries for Mac and Linux and a Docker image.

## Operating

Before running this, ensure you have a running Postgres server.
If one is not provided, the process will crash.
The main application runs on port 8000, and the metrics server is available on port 8001.
Configuring those ports, log level, and database settings can be done via environment variables.

_All environment variables are prefixed with `BASE_GO_`.\_

| Variable     | Default   |
| ------------ | --------- |
| APP_PORT     | 8000      |
| METRICS_PORT | 8001      |
| LOG_LEVEL    | info      |
| DB_HOST      | 127.0.0.2 |
| DB_PORT      | 5432      |
| DB_USER      | postgres  |
| DB_SSLMODE   | disable   |

Valid values for `LOG_LEVEL` include `debug`, `info`, `warn`, `error`, `fatal`, and `panic`.
Valid values for `DB_SSLMODE` include the following.

- `disable` - No SSL
- `require` - Always SSL (skip verification)
- `verify-ca` - Always SSL (verify that the certificate presented by the
  server was signed by a trusted CA)
- `verify-full` - Always SSL (verify that the certification presented by
  the server was signed by a trusted CA and the server host name
  matches the one in the certificate)
