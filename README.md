# Gedis distributed Key/Value store

## Overview

### Components

- `cmd/server` - grpc interface serving a database node
- ` cmd/migrate` - simple migration script to test functionality
- `pkg/db` - key/value store operations of a single node
- `pkg/grpc` - generated go libraries from `gedis.proto`
- `pkg/node` - the implementation of the grpc server and core of the distributed coordination efforts
### Prerequisites
- ko, https://github.com/google/ko
- go (supported version 1.18)

### Building
To containerize the server simply run `make containerize`.

### Deployment to k8s
To deploy the server to kubernetes simply run `make deploy-k8s`.
Additionally, there is an optional migration job to seed the database. 
One can build and deploy it using `make deploy-migration-job`.