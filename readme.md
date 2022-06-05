# General
This repository contains implementations of types of API in gRPC: Unary, Server Streaming, Client Streaming & Bi-Directional Streaming

<div align="center">
 <img src="https://user-images.githubusercontent.com/27688422/164034486-34865f69-f652-49e7-868d-aa8ecba80c26.png" width="600">
</div>

1. Clone the repository
2. Enter into project folder
```shell
cd grpc-golang-examples
```
3. Update dependencies
```shell
go mod tidy
```
4. Generate proto (must be have Makefile installed):
```shell
make <name_project>
```
> Example ``make grpc_simple``

or

```shell
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/luizmoitinho/grpc-golang-examples  --go-grpc_out=. --go-grpc_opt=module=github.com/luizmoitinho/grpc-golang-examples greet/proto/dummy.proto
```
---
