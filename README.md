# Base Home Automation

[![Build Status](https://img.shields.io/github/workflow/status/elispeigel/base/Go)](https://github.com/elispeigel/base/actions/workflows/go.yaml)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/elispeigel/base)
[![Go Report Card](https://goreportcard.com/badge/github.com/elispeigel/base)](https://goreportcard.com/report/github.com/elispeigel/base)

Base is a home automation platform written in Go, which allows you to control smart devices such as lights, speakers, and thermostats. The platform provides a simple API for sending commands between devices and managing their states.

## Features

- Control multiple devices through a single interface
- Support for smart lights, speakers, and thermostats
- Easy registration and deregistration of devices
- Use a factory pattern to create devices with default configurations
- Get the status of a device, including its on/off state, brightness, volume, or temperature

## Installation

To install the `base` package and its dependencies, you can use Go's built-in package manager:

```sh
go get github.com/elispeigel/base
```