# Golang Grpc Project Structure

- [Description](#description)
- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Usage](#usage)
- [Architecture Structure](#architecture-structure)

## Description

This project implements a simple gRPC service using Protocol Buffers for efficient data serialization between clients and servers. It's designed to demonstrate basic gRPC communication in [your programming language/environment].


## Getting Started

These instructions will guide you through the setup needed to get the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before starting, ensure you have the following software installed on your machine:
- Linux or macOS operating system
- Go, any one of the three latest major releases of Go.
- Go plugins for the protocol compiler
- Protocol Buffers Compiler
- Air for live reloading in development
- GNU Make


### Install Go
For Go installation instructions, see [Go’s Getting Started guide](https://go.dev/doc/install).

### Install Go plugins for the protocol compiler

you need to install the necessary plugins that allow `protoc` to generate Go-specific code. Run the following commands to install these plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
### Install Wire by running:
```bash
go go install github.com/google/wire/cmd/wire@latest
```
### Install Air for Live Reloading
Air is a live reloading tool for Go applications, which is useful during development. Install Air by running:
```bash
go install github.com/cosmtrek/air@latest
```

#### Update your PATH
After installing the plugins, you need to update your PATH environment variable to ensure that the `protoc` compiler can find the installed plugins. You can do this by adding the following line to your shell configuration file (such as `.bashrc` or `.zshrc`) or by running it directly in your terminal:
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Install the Protocol Buffers Compiler
Depending on your operating system, follow one of the instructions below to install `protoc`, the Protocol Buffers compiler.
Protocol Buffers Installation Guide.

**Linux:**

Using `apt` (Debian-based systems like Ubuntu):

```bash
sudo apt install -y protobuf-compiler
```

**MacOS:**

Using Homebrew:

```bash
brew install protobuf
```
##### Additional Documentation

For more detailed installation instructions, including building from source and additional operating systems, refer to the official Protocol Buffers Installation Guide:

[Protocol Buffers Installation Guide](https://grpc.io/docs/protoc-installation/).


### Install GNU Make
GNU Make is a tool which controls the generation of executables and other non-source files of a program from the program's source files.

**Linux:**

Using `apt` (Debian-based systems like Ubuntu):
```bash
sudo apt install -y make
```

**MacOS:**

GNU Make is usually pre-installed on macOS. To check if it is installed and find its version, use:

```bash
make --version
```
If it is not installed, you can install it using Homebrew:

```bash
brew install make
```

## Setup
After cloning the repository, you need to give execute permissions to the **generate_protos.sh** script and **update_module_name.sh** script: 
```bash
   chmod +x scripts/generate_protos.sh
```
```bash
   chmod +x ./scripts/update_module_name.sh
```
These scripts are essential for setting up and managing your project. Additionally, you can use the update_module_name.sh script to conveniently change the module name used in your Go project. Simply run the script with the new module name as an argument:
```bash
   ./scripts/update_module_name.sh new_module_name
```
Replace **new_module_name** with the desired new module name. This script will automatically update the module name in your **go.mod** file and replace the old module name with the new one in all relevant files in your project.

<a name="usage" id="usage"></a>
## ⚡️ Usage
  To use any of the above targets, run the make command followed by the target name. For example:

- install Downloads Go module dependencies.
```bash
   make install
```
-  Runs the Wire dependency injection tool from the **cmd/server/grpc** directory.
```bash
   make wire
```
-  [Once you've set up the necessary permissions](#setup), Generates Go code from **.proto** files using the **scripts/generate_protos.sh** script.
```bash
   make gen-protos
```
- Starts the development server with live reloading using the air configuration specified in **air.server.toml**.
```bash
  make make dev-server
```
- Starts the development client with live reloading using the air configuration specified in **air.client.toml**.
```bash
  make make dev-client
```
- Builds the application with specified linker flags to reduce binary size.
```bash
  make build
```
- Runs the application.
```bash
  make start
```
### Architecture Structure
```
project-root/
├── cmd/
│   ├── client        # Entry point of the application
│   │   └── main.go               
│   └── server               
│       └── main.go              
├── env/
│   ├── development.env       # Development environment settings
│   └── production.env        # Production environment settings
├── internal/
│   ├── entities/             # Domain types
│   ├── model/
│   │   ├── request/          # Request models
│   │   └── response/         # Response models
│   ├── repository/           # Database access layer
│   │   └── search/           # Specific implementation, e.g., search repository
│   ├── services/             # Business logic
│   │   └── cores/            # Core service implementations
│   └── rpci/                 # protocols that describe how RPC calls should be made and handled.
├── pkg/
│   ├── configs/              # Configuration file templates or default configs
│   └── utils/                # Utility functions and common helpers
├── go.mod                    # Go module dependencies
├── go.sum                    # Sum file for module verification
└── README.md                 # Project overview and documentation
