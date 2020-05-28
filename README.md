# GoMakefile

[![Build Status](https://travis-ci.org/AndrewDonelson/GoMakefile.svg?branch=master)](https://travis-ci.org/AndrewDonelson/GoMakefile)
![GitHub last commit](https://img.shields.io/github/last-commit/AndrewDonelson/GoMakefile)
[![Coverage Status](https://coveralls.io/repos/github/AndrewDonelson/GoMakefile/badge.svg)](https://coveralls.io/github/AndrewDonelson/GoMakefile)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/488f571baa13489494fa6002dbdf0897)](https://www.codacy.com/manual/AndrewDonelson/GoMakefile?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=AndrewDonelson/GoMakefile&amp;utm_campaign=Badge_Grade)
[![GoDoc](https://godoc.org/github.com/AndrewDonelson/GoMakefile?status.svg)](http://godoc.org/github.com/AndrewDonelson/GoMakefile)
![GitHub stars](https://img.shields.io/github/stars/AndrewDonelson/GoMakefile?style=flat)

## Summary

Standardized GoLang Makefile suitable for all projects

## Usage

```
deploy            Execute everything
all               Build binary
docker            Deploy + Docker Container
test-bench        Run benchmarks
test-short        Run only short tests
test-verbose      Run tests in verbose mode with coverage reporting
test-race         Run tests with race detector
check test tests  Run tests
test-xml          Run tests with xUnit output
test-coverage     Run coverage tests
lint              Run golint
fmt               Run gofmt on all source files
clean             Cleanup everything
```