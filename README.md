# webhook-demultiplexer

**webhook-demultiplexer** is a REST API, developed in Go, to provide a translation layer between StatusCake and a given application's API. It only exposes only one endpoint &#8212; a webhook, exposed in StatusCake &#8212; and all its data processing is focused in HTTP payloads received from StatusCake requests.

At the moment, it supports the following applications' API:

- [Cachet](https://github.com/CachetHQ/Cachet): StatusCake manages a given incident in Cachet by creating a new incident or updating status of an existing one.

## Getting Started

By default, this application runs at the top of a Docker container and it is configured either using environment variables or read them from Vault. However, you may not want to use it in a container. To do it so, you just need to clone this project and install it in your target host using provided `Makefile` -- for future reference, we are call **standalone** to this running option.

### How it works

**webhook-demultiplexer** follows the following workflow everytime it receives a request from StatusCake:

1. _POST_ request received from StatusCake. It is time to process and parse **body** to separate alert and authentication contents.
2. Does the _POST_ have authorization to use accessed API's endpoint? Time to evaluate token received in the **body** of the request and check if is authorized.
3. Extracts target application from endpoint, if any, and starts configurating client by parsing JSON configuration.
4. Start sending alerts to target APIs by every tag present in _POST_ request.

### Standalone

### Docker
