#!/bin/bash

# Directory containing proto files
PROTO_DIR="./grpc/proto"

# Directory where generated Go files will be placed
GEN_DIR="./grpc/gen"

# Ensure the output directory exists
mkdir -p "${GEN_DIR}"

# Generate Go code for all proto files in PROTO_DIR
for PROTO_FILE in "${PROTO_DIR}"/*.proto; do
    echo "Generating Go files for ${PROTO_FILE}..."
    protoc -I "${PROTO_DIR}" --go_out="${GEN_DIR}" --go_opt=paths=source_relative \
           --go-grpc_out="${GEN_DIR}" --go-grpc_opt=paths=source_relative \
           "${PROTO_FILE}"
done

echo "All proto files have been compiled."
