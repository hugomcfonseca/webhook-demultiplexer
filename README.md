# webhook-demultiplexer

**webhook-demultiplexer** is a REST API, developed in Go, to provide a translation layer between StatusCake and a given application's API. It only exposes only one endpoint &#8212; a webhook, exposed in StatusCake &#8212; and all its data processing is focused in HTTP payloads received from StatusCake requests.

At the moment, it supports the following applications' API:

- [Cachet](https://github.com/CachetHQ/Cachet): StatusCake manages a given incident in Cachet by creating a new incident or updating status of an existing one.

## Getting Started

By default, this application runs at the top of a Docker container and it is configured either using environment variables or read them from Vault. However, you may not want to use it in a container. To do it so, you just need to clone this project and install it in your target host using provided `Makefile` -- for future reference, we are call **standalone** to this running option.

### Standalone

### Docker
