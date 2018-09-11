# Library to interface with traffic-light.co API

[![Documentation](https://godoc.org/github.com/aquiferinstitute/trafficlightco?status.svg)](http://godoc.org/github.com/aquiferinstitute/trafficlightco) [![Go Report Card](https://goreportcard.com/badge/github.com/aquiferinstitute/trafficlightco)](https://goreportcard.com/report/github.com/aquiferinstitute/trafficlightco) [![Coverage Status](https://coveralls.io/repos/github/aquiferinstitute/trafficlightco/badge.svg?branch=master)](https://coveralls.io/github/aquiferinstitute/trafficlightco?branch=master) [![CircleCI](https://circleci.com/gh/aquiferinstitute/trafficlightco.svg?style=svg)](https://circleci.com/gh/aquiferinstitute/trafficlightco) [![Maintainability](https://api.codeclimate.com/v1/badges/8bc5fe508b7d3891f813/maintainability)](https://codeclimate.com/github/aquiferinstitute/trafficlightco/maintainability)

## How to get an API key

Simply subscribe to the [traffic light portal](https://traffic-light.co)

## How to use

Look in the cmd folder

## Build the command line tool

```bash
GOOS=darwin GOARCH=amd64 go build -o trafficlight ./cmd/main.go
GOOS=windows GOARCH=amd64 go build -o trafficlight.exe ./cmd/main.go
```
## Notice on Traffic Light Portal data reliance

Please contact the Euler Hermes Traffic Light Portal Sandbox team directly for verification of data accuracy should you wish to make use of their service. 

Meanwhile, make comments, share feedback and best of luck! 
