# webhook-demultiplexer

**webhook-demultiplexer** is a REST API, developed in Go, to provide a translation layer between StatusCake webhooks and a given application's API. It only exposes only one endpoint &#8212; configured as a webhook in StatusCake &#8212; and all its data processing is focused in HTTP payloads received from StatusCake requests. 

At the moment, it only supports the following applications' API:

- [Cachet](https://github.com/CachetHQ/Cachet): StatusCake creates or updates a given incident in Cachet

## Getting Started

By default, this application runs at the top of a Docker container and its configured using variables read as environment variables or from Vault.

However, it's still possible to run it without a container &#8212; let me call this **standalone** mode.

### Standalone

### Docker