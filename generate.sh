#!/bin/bash

protoc average/averagepb/average.proto --go_out=plugins=grpc:.