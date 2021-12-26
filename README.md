# compute-average

An API that uses Client Streaming gRPC framework to capture all integer numbers on Requests that the client sent and returns just the average of all in one Response.

  

#

### To execute

First initialize the Server

```

go run average/average_server/server.go

```

After, execute the client passing two arguments as below

```

go run average/average_client/client.go 720 120 10 21 31

```

#

### Requirements

[![Stable release](https://img.shields.io/badge/libprotoc-3.17.3-green.svg)](https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3)

[![Stable release](https://img.shields.io/badge/golang-v1.17.5-blue.svg)](https://go.dev/doc/go1.17)