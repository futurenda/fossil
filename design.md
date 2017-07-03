# Design Goals

## CLI

### Flags

- `--verbose`

Show verbose logs and update current status (eg. progress bar)

## Example

```bash
tree

.
├── design.md
├── examples
│   ├── sql
│   │   └── select.sql
│   └── tree
│       └── tree.go
├── fossil
├── LICENSE
├── main.go
├── README.md
└── tree.go

```

```bash
```

## Concurrent IO

- `max_read_io_goroutines` (eg. 16)
- `max_write_io_goroutines` (eg. 16)

## Cancellation

`Control-C` should trigger cancellation and all IO goroutines should respond to the cancellation.

## Tests

Use Ginkgo to do BDD testing.

https://github.com/onsi/ginkgo

## References

- The Go Programming
