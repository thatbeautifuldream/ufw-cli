# UFW CLI Tool

A command-line interface tool to simplify the management of UFW (Uncomplicated Firewall) on Linux systems. This tool provides easy-to-use commands for installing, configuring, and managing UFW firewall rules.

## Features

- UFW installation check and automated installation
- Basic firewall rules setup (SSH, HTTP, HTTPS)
- Custom port configuration
- Firewall status checking
- Interactive command-line interface

## Prerequisites

- Go 1.x or higher
- Linux-based system (Debian/Ubuntu recommended)
- Sudo privileges

## Installation

To install the tool globally, run the following commands:

```bash
go install github.com/thatbeautifuldream/ufw-cli
```

### Usage

- Install UFW:

```bash
ufw-cli install
```

- Set up basic UFW rules:

```bash
ufw-cli setup
```

- Configure additional ports interactively:

```bash
ufw-cli configure
```

- Check the current UFW status:

```bash
ufw-cli status
```

## Development

Clone the repository:

```bash
git clone --depth 1 https://github.com/thatbeautifuldream/ufw-cli.git
```

Navigate to the project directory:

```bash
cd ufw-cli
```

## Build and Run Locally

The project includes several Make commands for building and running the application:

| Command        | Description                                            |
| -------------- | ------------------------------------------------------ |
| `make build`   | Builds the binary in the `build` directory             |
| `make run`     | Runs the application directly                          |
| `make clean`   | Removes build artifacts and cleans the build directory |
| `make install` | Builds and installs the binary to `/usr/local/bin`     |
| `make all`     | Runs clean and then builds the application             |
| `make help`    | Displays all available make commands                   |

## Contributing

Contributions are welcome! Please feel free to submit a pull request.
