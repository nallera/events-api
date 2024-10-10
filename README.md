# meli_traffic

![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

This is a basic Go application created by Fury to be used as a starting point for your project.

## First steps

### Go Runtime Version

Specify the Go runtime version tag you desire in your `Dockerfile`. If in doubt, it's completely safe to always use the
latest one given the [Go 1 compatibility guarantees](https://golang.org/doc/go1compat).

```docker
FROM hub.furycloud.io/mercadolibre/go:1.21-mini
```

> You can find all available image tags for your Dockerfile
> [here](https://github.com/melisource/fury_go-mini#supported-tags).

### Dependency Management

Run `go mod tidy` to ensure that the `go.mod` file matches the source code in the module.

For more information refer to the
[`go-mini` docs](https://github.com/melisource/fury_go-mini#dependency-management-support).

## API Documentation

### Fury Specs Hub

To simplify the management and maintainability of your API specs, we present [Fury Specs Hub](https://furydocs.io/specs-hub/latest/guide/#/). Fury Specs Hub is a new service from Fury that aims to be a one-stop solution for API definition. With Specs Hub, you will be able to:
- Define your APIs using OpenAPI or AsyncAPI.
- Automate the configuration and generation of your API specs with the help of new commands from the Fury CLI.
- Have all your specs in one place for visualization and management.
- Share them with other teams.
- Find available APIs based on the information you need.
- Usage documentation [Fury Specs Hub - Getting started](https://furydocs.io/specs-hub/latest/guide/#/tutorial/).

#### Usage guide

1. [Installing the Specs Hub plugin for Fury CLI.](https://furydocs.io/specs-hub/latest/guide/#/tutorial/install-specs-hub-furycli)
2. [Installing the OpenAPI plugin and initializing a basic configuration.](https://furydocs.io/specs-hub/latest/guide/#/tutorial/install-open-api)
3. [Generating your first API specification.](https://furydocs.io/specs-hub/latest/guide/#/tutorial/generate-open-api-spec)
4. [Validating your API specification.](https://furydocs.io/specs-hub/latest/guide/#/tutorial/validate-specs)
5. [Uploading your first specification.](https://furydocs.io/specs-hub/latest/guide/#/tutorial/upload-spec)
6. [Viewing your specification in Fury web.](https://furydocs.io/specs-hub/latest/guide/#/tutorial/view-spec)
7. [Managing your specification in Fury web.](https://furydocs.io/specs-hub/latest/guide/#/tutorial/manage-spec)

## Questions

* [Fury Issue Tracker](https://github.com/melisource/fury/issues)