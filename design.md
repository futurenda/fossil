# Fossil

An exercise for Casey.

- Goroutines (Section 8.1)
- Channels (Section 8.4)
- Multiplexing with select (Section 8.7)
- Interfaces (Chapter 7)
- Packages and the Go Tool (Chapter 10)
- Methods (Chapter 6)
- BDD

## CLI

### Flags

- `--verbose`

Show verbose logs and update current status (eg. progress bar)

## Concurrent IO

- `max_read_io_goroutines` (eg. 16)
- `max_write_io_goroutines` (eg. 16)

## Cancellation

`Control-C` should trigger cancellation and all IO goroutines should respond to the cancellation.

## Tests

- Dependency Injection
- Test using io.ReadCloser interface
- Test using io.WriteCloser interface

## Tools

- golang/dep as dependency tool
- onsi/ginkgo as BDD testing framework

## References

- The Go Programming Language
