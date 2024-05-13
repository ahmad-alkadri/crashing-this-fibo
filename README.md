
# Crashing Fibonacci API with Memory Watcher

## Overview

This example, playground repository contains a setup of a Go application that
provides a REST API for calculating Fibonacci numbers. It includes a memory
monitoring component that ensures the application does not exceed a predefined
memory limit and can crash itself if necessary. Additionally, this project
includes a script for load testing the API using Vegeta, a versatile HTTP load
testing tool.

## Installation and Setup

### Prerequisites
- **Go**: Ensure you have Go installed on your system. You can download it from
  [the official site](https://golang.org/dl/).
- **Vegeta**: Vegeta needs to be installed for load testing. Install it using
  Homebrew on macOS:

  ```
  brew install vegeta
  ```
  Or download and install it from [GitHub](https://github.com/tsenart/vegeta)
  for other platforms.

### Launching the API
Run the application using the following command from the root of the repository:
```
go run cmd/crashing-this-fibo/main.go
```

You can also opt to build them and run the compiled binary.

### Load Testing with Vegeta

Navigate to the `scripts` directory where the `load_test.sh` script is located.
Make the script executable (if not already done):
```
chmod +x load_test.sh
```
Launch the load test by running:
```
./load_test.sh
```

## Configuration

### Adjusting the Load Test Rate

You can modify the rate of requests in the `load_test.sh` script. Find the line
with `-rate=10` and change the rate to your desired number of requests per
second.

### Setting the Memory Limit

The memory limit is controlled by an environment variable `APP_MEMORY_LIMIT`.
This variable is read as megabytes. To set the memory limit to 5 MB, for
example, use the following command before starting the application:
```
export APP_MEMORY_LIMIT=5
```
Adjust this value based on your system's capacity and testing requirements.

## Conclusion

This repo serves as a playground and example on setting up a Go REST API app
able to crashing itself. Blog post will follow.