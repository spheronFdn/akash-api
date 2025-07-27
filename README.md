# Akash API

[![lint](https://github.com/akash-network/akash-api/actions/workflows/lint.yaml/badge.svg)](https://github.com/akash-network/akash-api/actions/workflows/lint.yaml)
[![tests](https://github.com/akash-network/akash-api/actions/workflows/tests.yaml/badge.svg)](https://github.com/akash-network/akash-api/actions/workflows/tests.yaml)

## Brief

The purpose of this repo is to combine gRPC API definitions for [Akash node](https://github.com/akash-network/node)
and [Akash provider](https://github.com/akash-network/provider) in a single place as well as code-generation related to it.

There are currently defined two buf packages with further intentions to be published into BSR if there is demand for it:

- `buf.build/akash-network/node` - Akash protobuf definitions previously located in [proto](https://github.com/akash-network/node/tree/master/proto/akash)
  have all been moved [under](./proto/node/akash). All generated code can be found [here](./go/node)
- `buf.build/akash-network/provider` - Akash manifest definitions previously defined as plan Go structs have been converted into Protobuf [definitions](./proto/provider/akash)

Proto docs are available:

- for [node](docs/proto/node.md)
- for [provider](docs/proto/provider.md)

## Contributing

Issues should be submitted via [support repo](https://github.com/akash-network/support/issues) and tagged with `repo/akash-api`.
All PRs must have an open issue in the support repo.

## Releases

Actual releases indicating changes to the repo itself. API version are defined withing each module.

## Pre-generated packages

We provide generated code allowing developers to just import and focus on features, not how to generate stubs.

- [Go](./go)

# Archived Repository

This repository was initially used for educational and testing purposes during our early exploration of compute infrastructure, including a period of collaboration and experimentation with the Akash Network under the Apache 2.0 license.

We have now migrated to a new, independently developed production codebase under a new GitHub organization:
👉 [https://github.com/spheron-core/](https://github.com/spheron-core/)

This migration aligns with our roadmap toward TGE, Foundation-based governance, and long-term code maintainability.

Note:

- This repository is now deprecated and archived.
- It is no longer maintained and is not used in any part of the production infrastructure.
- All original attributions have been preserved in compliance with the Apache 2.0 license.
- No active development will occur on this repository moving forward.

For the latest updates and active development, please refer to the new organization above.
