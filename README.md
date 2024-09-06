# RSX: Package Risor Scripts into Go Binaries

`rsx` is a tool that packages your [Risor](https://risor.io) scripts into a standalone Go binary, allowing for easy distribution of your applications without dependencies.

> [!WARNING]
> RSX is currently in alpha.

## Features

- Package Risor scripts into a single executable
- Built-in `rsx` module for extended functionality
- Easy project initialization and management

## Pre-requisites

You need to have Go installed on your system. If you don't have it, you can download it from the [official website](https://golang.org/dl/).

## Installation

Install RSX using Go:

```bash
go install github.com/rubiojr/rsx@latest
```

Ensure that your Go bin directory is in your PATH.

## Quick Start

1. Create a new Risor project:
   ```bash
   rsx new myapp
   cd myapp
   ```

2. Edit the main script (`main.risor`) to add your code:
   ```risor
   import rsx

   rsx.log("Hello from Risor!")
   ```

3. Build the Go binary:
   ```bash
   rsx build
   ```

> [!NOTE]
> rbx runs `go build` under the hood, so you can pass any additional environment variables to it, like `GOOS`, `GOARCH`, etc.

4. Run your application:
   ```bash
   ./myapp
   # Output: Hello from Risor!
   ```

## Project Structure

- `main.risor`: The entry point of your application.
- `lib/`: Directory for additional Risor modules.
- `lib/rsx.risor`: Built-in RSX module with utility functions.

## Adding Custom Modules

Place any additional `.risor` files in the `lib` directory. They will be automatically available at runtime.

## The RSX Module

RSX comes with a built-in `rsx` module providing basic functionality. For available functions, refer to [rsx.risor](lib/rsx.risor).

## Development Workflow

During development, you can run your Risor scripts directly:

```bash
risor --modules lib main.risor
```

This allows for faster iteration without needing to rebuild the binary.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

For issues, questions, or contributions, please [open an issue](https://github.com/rubiojr/rsx/issues) on GitHub.
