<div align="center">

# 🔥 Go Bytes

**A curated collection of Go patterns, idioms, and techniques**

![Go Version](https://img.shields.io/badge/Go-1.25.4-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

*Learn by example. Build with confidence.*

[Features](#-features) • [Quick Start](#-quick-start) • [Examples](#-code-examples) • [Testing](#-testing)

</div>

---

## 📖 Overview

**Go Bytes** is a comprehensive playground demonstrating modern Go programming patterns, from basic data structures to advanced concurrency techniques. Each example is production-ready and battle-tested, perfect for learning or reference.

## ✨ Features

### 🏗️ Data Structures
- **Generic Set** - Type-safe set implementation with full CRUD operations
- **TTL Cache** - Thread-safe in-memory cache with automatic cleanup and expiration

### 🎨 Design Patterns
- **Functional Options Pattern** - Flexible configuration with optional parameters
- **Fluent Builder Pattern** - Chainable API for object construction
- **Branded Types** - Compile-time type safety for domain models

### ⚡ Concurrency
- **Goroutines & Channels** - Parallel HTTP requests with WaitGroups
- **Mutex Synchronization** - Thread-safe cache operations

### 🔧 Advanced Go Features
- **Generics** - Type parameters for reusable data structures
- **Iterators (Go 1.23+)** - Lazy sequence evaluation and transformation
- **Reflection** - Dynamic struct manipulation and introspection
- **Build Tags** - Conditional compilation for feature flags
- **Bitmask Operations** - Efficient permission management

### 🛠️ Utilities
- **Custom JSON Marshaling** - Deterministic JSON output with sorted keys
- **Struct Field Copying** - Reflection-based field mapping
- **Pointer Helpers** - Convenient value-to-pointer conversions

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/ryands17/go-bytes.git
cd go-bytes

# Run with basic features
go run main.go

# Run with premium features enabled
go run -tags=premium main.go

# Run tests
go test ./...

# Build with premium features
go build -tags=premium -o go-bytes
```

## 📂 Project Structure

```
go-bytes/
├── cmd/
│   ├── brands/         → Branded types with type-safe validation
│   ├── builders/       → Builder patterns (functional & fluent)
│   ├── features/       → Build tags and feature flags
│   ├── iterators/      → Iterator implementations (primes, sequences)
│   ├── routines/       → Concurrency patterns with goroutines
│   ├── structures/     → Generic data structures (Set, etc.)
│   └── utils/          → Helper functions and utilities
│       └── bitmasks/   → Bitwise operations for permissions
└── main.go             → Comprehensive examples and usage
```

## 💡 Code Examples

### Generic Set
```go
set := structures.NewSet[int]()
set.Add(1)
set.Add(2)
set.Add(1) // Duplicate ignored
fmt.Println(set.Size()) // Output: 2
fmt.Println(set.Contains(1)) // Output: true
```

### TTL Cache
```go
cache := utils.NewCache(5*time.Second, nil)
cache.Set("user:1", "John Doe", nil)

if value, found := cache.Get("user:1"); found {
    fmt.Println(value) // Output: John Doe
}
// Automatically expires after 5 seconds
```

### Functional Options Pattern
```go
db := builders.NewDBClient(
    builders.WithUrl("localhost:5432"),
    builders.WithConnections(10),
    builders.WithTimeout(10*time.Second),
)
db.Connect()
```

### Fluent Builder Pattern
```go
db := builders.NewDBClientFluent().
    WithURL("localhost:5432").
    WithConnections(10).
    WithTimeout(10 * time.Second).
    Build()
db.Connect()
```

### Branded Types
```go
user := brands.GeneralUser{
    ID:       "1",
    Name:     "Alice",
    UserType: brands.Admin,
}

adminUser, err := brands.IsAdmin(user)
if err != nil {
    log.Fatal("Not an admin!")
}
// adminUser is now type AdminUser, not GeneralUser
```

### Bitmask Permissions
```go
perms := bitmasks.READ
perms.Set(bitmasks.WRITE)
perms.Toggle(bitmasks.EXECUTE)

fmt.Println(perms.Has(bitmasks.READ))    // true
fmt.Println(perms.Has(bitmasks.WRITE))   // true
fmt.Println(perms.Has(bitmasks.EXECUTE)) // true
```

### Iterator Chains (Go 1.23+)
```go
// Generate first 100 prime numbers
primes := iterators.Primes(iterators.NonZeroIntegers(100))
result := slices.Collect(primes)
fmt.Println(result) // [2, 3, 5, 7, 11, 13, ...]
```

### Concurrent HTTP Requests
```go
// Fetch 5 users concurrently with goroutines and channels
routines.FetchAllUsers(5)
// Uses WaitGroups and channels for synchronization
```

### Custom JSON Marshaling
```go
// Produces deterministic JSON with sorted keys
type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}

user := User{Name: "Bob", Age: 30}
jsonBytes, _ := utils.MarshalStruct(user)
fmt.Println(string(jsonBytes)) // {"age":30,"name":"Bob"}
```

## 🧪 Testing

Run the test suite:

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific package tests
go test ./cmd/structures
go test ./cmd/utils
```

Tests include:
- Set operations (add, remove, contains)
- Custom JSON marshaling with field ordering
- Cache expiration and cleanup

## 🔧 Dependencies

- **[resty/v2](https://github.com/go-resty/resty)** - Elegant HTTP client for Go
- **[testy](https://github.com/peterldowns/testy)** - Modern testing assertions

## 📚 What You'll Learn

| Topic | Location | Description |
|-------|----------|-------------|
| **Generics** | `cmd/structures/` | Type-safe data structures |
| **Concurrency** | `cmd/routines/` | Goroutines, channels, WaitGroups |
| **Reflection** | `cmd/utils/helpers.go` | Dynamic struct operations |
| **Build Tags** | `cmd/features/` | Conditional compilation |
| **Iterators** | `cmd/iterators/` | Lazy sequences (Go 1.23+) |
| **Patterns** | `cmd/builders/` | Builder and Options patterns |
| **Bitmasks** | `cmd/utils/bitmasks/` | Bitwise operations |
| **Branded Types** | `cmd/brands/` | Type-safe domain modeling |

## 🎯 Use Cases

- **Learning Resource** - Study modern Go patterns and idioms
- **Reference Implementation** - Copy-paste production-ready code
- **Interview Prep** - Review common Go interview topics
- **Code Review** - Examples of idiomatic Go code
- **Teaching Tool** - Educational resource for Go workshops

## 🤝 Contributing

Contributions are welcome! Feel free to:
- Add new patterns and examples
- Improve documentation
- Fix bugs or add tests
- Suggest new features

## 📝 License

MIT License - feel free to use this code in your projects!

---

<div align="center">

**Built with ❤️ using Go 1.25.4**

[⬆ Back to Top](#-go-bytes)

</div>
