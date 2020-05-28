# GoMakefile
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