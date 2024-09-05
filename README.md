# RSX

`rsx` can package your [Risor](https://risor.io) scripts in a Go binary, so you can distribute your app without dependencies.

> [!WARNING]
> rsx is currently Alpha quality.

## Install

```
go install  github.com/rubiojr/rsx@latest
```

## Usage

Create a new Risor project:

```sh
rsx new myapp
cd myapp
```

The entry point will be `main.risor`:

```
rsx.log("hello from Risor!")
```

Build the Go binary:

```
rsx build
```

Run it:

```
./myapp
Hello, World!
```

## Adding modules

Any additional `.risor` files in the `lib` directory will be available at run time.

## Built-in rsx module

There's a built-in `rsx` module that provides some basic functionality. More modules will be added in the future.

Check [rsx.risor](lib/rsx.risor) for the available functions.

## Development tips

While in development, you can run the Risor scripts directly:

```
risor --modules lib main.risor
```
