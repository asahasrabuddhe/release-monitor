#!/bin/bash
protoc -I proto/ proto/release-monitor.proto --go_out=plugins=grpc:../../../
