# envman

[![License: GPL v3](https://img.shields.io/badge/License-GPL--3.0-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.25.5-blue)](https://go.dev/)
[![Go Report Card](https://goreportcard.com/badge/github.com/hisbaan/envman)](https://goreportcard.com/report/github.com/hisbaan/envman)
[![Release](https://img.shields.io/github/v/release/hisbaan/envman)](https://github.com/hisbaan/envman/releases)

envman is a CLI tool for managing `.env` files when working with multiple environments, particularly useful in a monorepo. It allows you to hotswap between different environment configurations using symlinks.

## Features

- Manage multiple projects with different environment configurations
- Define multiple directories within a project for `.env` file synchronization
- Quickly switch between environments using symlinks
- View current project status and active environment

## Installation

### Go

Ensure you have Go installed (version 1.25.5 or later). Then run:

```bash
go install github.com/hisbaan/envman@latest
```

### Build from Source

Clone the repository and build:

```bash
git clone https://github.com/hisbaan/envman.git
cd envman
go build -o envman
```

## Usage

### Add a Project

Add your current directory or a specific path as a project:

```bash
envman proj add [project-dir]
```

### Remove a Project

Remove a project from the configuration:

```bash
envman proj rm [project-dir]
```

### Add a Directory

Add a subdirectory to the current project:

```bash
envman dir add [dir]
```

### Remove a Directory

Remove a directory from the current project:

```bash
envman dir rm [dir]
```

### Link an Environment

Link an environment configuration to `.env` files across all project directories:

```bash
envman link [environment]
```

### Unlink

Remove all `.env` symlinks:

```bash
envman unlink
```

### Status

View the current project status, directories, and active environment:

```bash
envman status
```

## Configuration

envman stores its configuration in `~/.config/envman/config.toml`. The configuration is managed automatically through the CLI commands.
