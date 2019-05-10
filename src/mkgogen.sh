#!/usr/bin/env bash

# Run this if opencensus-proto is checked in the GOPATH.
# go get -d github.com/census-instrumentation/opencensus-proto
# to check in the repo to the GOAPTH.
#
# This also requires the grpc-gateway plugin.
# See: https://github.com/grpc-ecosystem/grpc-gateway#installation
#
# To generate:
#
# cd $(go env GOPATH)/census-instrumentation/opencensus-proto
# ./mkgogen.sh

OUTDIR="$(go env GOPATH)/src"
INCLUDES="$(go env GOPATH)/src/github.com/gogo/protobuf/protobuf"

ARGS="Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
goproto_registration=true"


protoc --gogofaster_out=$ARGS,plugins=grpc:$OUTDIR opencensus/proto/stats/v1/stats.proto \
    && protoc --gogofaster_out=$ARGS,plugins=grpc:$OUTDIR opencensus/proto/metrics/v1/metrics.proto \
    && protoc --gogofaster_out=$ARGS,plugins=grpc:$OUTDIR opencensus/proto/resource/v1/resource.proto \
    && protoc --gogofaster_out=$ARGS,plugins=grpc:$OUTDIR opencensus/proto/trace/v1/trace.proto \
    && protoc --gogofaster_out=$ARGS,plugins=grpc:$OUTDIR opencensus/proto/trace/v1/trace_config.proto \
    && protoc -I=. --gogofaster_out=$ARGS,plugins=grpc:$OUTDIR opencensus/proto/agent/common/v1/common.proto \
    && protoc -I=. --gogofaster_out=$ARGS,plugins=grpc:$OUTDIR opencensus/proto/agent/metrics/v1/metrics_service.proto \
    && protoc -I=. --gogofaster_out=$ARGS,plugins=grpc:$OUTDIR opencensus/proto/agent/trace/v1/trace_service.proto \
    && protoc --grpc-gateway_out=logtostderr=true,grpc_api_configuration=./opencensus/proto/agent/trace/v1/trace_service_http.yaml:$OUTDIR opencensus/proto/agent/trace/v1/trace_service.proto

# Generate OpenApi (Swagger) documentation file for grpc-gateway endpoints.
OPENAPI_OUTDIR=../gen-openapi
mkdir -p $OPENAPI_OUTDIR
protoc --swagger_out=logtostderr=true,grpc_api_configuration=./opencensus/proto/agent/trace/v1/trace_service_http.yaml:$OPENAPI_OUTDIR \
  opencensus/proto/agent/trace/v1/trace_service.proto
