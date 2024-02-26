# Billion Row Challenge

This is small project inspired by another similarly name Java project.

## Generator

A Go binary to generate a billion row text file (csv with 2 columns) of keys
and fixed-point decimals.
Use `go build && ./billionrows <file> <number_of_rows>` to generate the data.

## Aggregate

### Rust
A Rust binary to report statistics about the csv file. This is very slow!

### duckdb
Use duckdb with `cat duck.sql | duckdb`

### Go
Todo!
