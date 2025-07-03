
# Learn Go with Tests

This repository contains Go packages and examples following the "Learn Go with Tests" approach, covering fundamental Go concepts through test-driven development.

## Project Structure

### Core Packages

#### `integers/`
- **Purpose**: Introduction to basic Go functions and testing
- **Key Files**:
  - `adder.go`: Simple addition function
  - `adder_test.go`: Tests and examples for the Add function
- **Concepts**: Basic functions, testing, documentation examples

#### `iteration/`
- **Purpose**: Learning Go's iteration constructs
- **Key Files**:
  - `repeat.go`: String repetition using loops and strings.Builder
  - `repeat_test.go`: Tests and benchmarks for string repetition
- **Concepts**: For loops, string manipulation, benchmarking

#### `arrays/`
- **Purpose**: Working with arrays and slices
- **Key Files**:
  - `sum.go`: Functions for summing arrays and slices
  - `sum_test.go`: Comprehensive tests for sum operations
- **Concepts**: Arrays, slices, variadic functions, helper functions

#### `structs/`
- **Purpose**: Custom types and methods
- **Key Files**:
  - `shapes.go`: Geometric shapes with area calculations
  - `shapes_test.go`: Tests for shape area calculations
- **Concepts**: Structs, methods, interfaces

#### `pointers/`
- **Purpose**: Memory management and pointer usage
- **Key Files**:
  - `wallet.go`: Bitcoin wallet with deposit/withdrawal operations
  - `wallet_test.go`: Tests for wallet operations and error handling
- **Concepts**: Pointers, methods on pointers, error handling, custom types

#### `maps/`
- **Purpose**: Key-value data structures
- **Key Files**:
  - `dictionary.go`: Dictionary implementation using maps
  - `dictionary_test.go`: Tests for dictionary operations
- **Concepts**: Maps, error handling, custom error types

#### `di/` (Dependency Injection)
- **Purpose**: Dependency injection patterns
- **Key Files**:
  - `di.go`: Greeting function with configurable output
  - `di_test.go`: Tests demonstrating dependency injection
- **Concepts**: Interfaces, dependency injection, io.Writer

#### `mocking/`
- **Purpose**: Test doubles and mocking
- **Key Files**:
  - `main.go`: Countdown application with configurable sleep
  - `countdown_test.go`: Tests using spies and mocks
  - `mocks.go`: Test double implementations
- **Concepts**: Interfaces, mocking, spies, test doubles, time handling

#### `concurrency/`
- **Purpose**: Concurrent programming with goroutines
- **Key Files**:
  - `check_websites.go`: Concurrent website checking
  - `check_websites_test.go`: Tests for website checker
  - `check_websites_bench_test.go`: Benchmarks comparing concurrent vs sequential execution
- **Concepts**: Goroutines, channels, concurrency, benchmarking

#### `select/`
- **Purpose**: Advanced concurrency with select statements
- **Status**: Directory exists but appears to be in development
- **Concepts**: Select statements, channel operations, timeouts

## Running the Code

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./concurrency

# Run benchmarks
go test -bench=. ./concurrency
```

### Running Individual Programs
```bash
# Run the dependency injection example
go run ./di

# Run the mocking countdown example  
go run ./mocking
```

## Key Learning Objectives

1. **Test-Driven Development**: Each package demonstrates writing tests first
2. **Go Fundamentals**: Variables, functions, control structures
3. **Data Structures**: Arrays, slices, maps, structs
4. **Object-Oriented Concepts**: Methods, interfaces, composition
5. **Error Handling**: Go's explicit error handling patterns
6. **Concurrency**: Goroutines, channels, and concurrent patterns
7. **Testing Patterns**: Unit tests, benchmarks, mocks, and test helpers

## Go Version

This project uses Go 1.22 and follows modern Go practices and idioms.
