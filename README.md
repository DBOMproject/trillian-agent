# Trillian Agent

The implementation of an agent that uses trillian as the storage mechanism

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [How to Use](#how-to-use)
  - [API](#api)
  - [Configuration](#configuration)
- [Development](#development)
  - [Regenerate API](#regenerate-api)
- [Usage](#usage)
  - [Deploy Trillian](#deploy-trillian)
  - [Create Master Channel Tree](#create-master-channel-tree)
- [Helm Deployment](#helm-deployment)
- [Platform Support](#platform-support)
- [Getting Help](#getting-help)
- [Getting Involved](#getting-involved)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## How to Use

### API

Latest OpenAPI Specification for this API is available on the [api-specs repository](https://github.com/DBOMproject/api-specs/tree/master/agent)

### Configuration

| Environment Variable         | Default          | Description                                            |
|------------------------------|------------------|--------------------------------------------------------|
| LOG_LEVEL                    | `info`           | The verbosity of the logging                           |
| PORT                         | `5000`           | Port on which the agent listens                        |
| HOST                         | `0.0.0.0`        | The host address of the agent                          |
| TRILLIAN_ENDPOINT            | `localhost:8091` | The endpoint of the trillian server connect to         |
| CHANNEL_CONFIG_MAP_ID        | `0`              | The id of the trillian map to store the channel config |
| JAEGER_ENABLED               | `false`          | Is jaeger tracing enabled                              |
| JAEGER_HOST                  | ``               | The jaeger host to send traces to                      |
| JAEGER_SAMPLER_PARAM         | `1`              | The parameter to pass to the jaeger sampler            |
| JAEGER_SAMPLER_TYPE          | `const`          | The jaeger sampler type to use                         |
| JAEGER_SERVICE_NAME          | `Trillian Agent` | The name of the service passed to jaeger               |
| JAEGER_AGENT_SIDECAR_ENABLED | `false`          | Is jaeger agent sidecar injection enabled              |


## Development
### Regenerate API
[go-swagger](https://github.com/go-swagger/go-swagger) was used to generate the server api

**Please note: Make sure the trillian agent repository is included in your go path src folder**

```
swagger generate server -f {path-to-swagger} -A trillian-agent -t={path-to-trillian-repo}
```

**Please don't make updates to generated files as those changes will be overwritten when the server is regenerated**

## Usage
### Deploy Trillian 
Find the information about to deploy trillian [here](https://github.com/google/trillian/tree/master/deployment)

**Make sure to install the trillian map server as part of the deployment**

### Create Master Channel Tree
After trillian is deployed, channel config trillian map needs to be created that the trillian agent will use to store channel configurations.

Use the following command to create the tree and set the environmental variable the trillian agent will use

**Make sure to save the the map id as it will be needed for the deployment**

```
GRPC="8091"
CHANNEL_CONFIG_MAP_ID=$(\
go run github.com/google/trillian/cmd/createtree \
--admin_server=:${GRPC} \
--tree_type=MAP \
--description='Trillian Agent Channel Config' \
--display_name='TrillAgentChanConf' \
--hash_strategy=CONIKS_SHA512_256) && echo ${CHANNEL_CONFIG_MAP_ID}
```
### Helm Deployment

Instructions for deploying the trillian-agent using helm charts can be found [here](https://github.com/DBOMproject/deployments/tree/master/charts/trillian-agent)

## Platform Support

Currently, we provide pre-built container images for linux amd64 and arm64 architectures via our Github Actions Pipeline. Find the images [here](https://hub.docker.com/r/dbomproject/trillian-agent)

## Getting Help

If you have any queries on trillian-agent, feel free to reach us on any of our [communication channels](https://github.com/DBOMproject/community/blob/master/COMMUNICATION.md)

If you have questions, concerns, bug reports, etc, please file an issue in this repository's [issue tracker](https://github.com/DBOMproject/trillian-agent/issues).

## Getting Involved

Find instructions on how you can contribute in [CONTRIBUTING](CONTRIBUTING.md).
