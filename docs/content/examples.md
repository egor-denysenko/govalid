---
title: "Examples"
description: "Real-world examples of using govalid"
weight: 40
---

# Examples

This page provides practical examples of using govalid, starting from simple use cases and progressing to more complex scenarios.

## Getting Started: Simple Validation

Let's start with the simplest possible example to understand the basics.

### Step 1: Create a Basic Struct

Create a file `main.go`:

```go
package main

import (
	"fmt"
	"log"
)

//go:generate govalid .

type User struct {
	// +govalid:required
	Name string `json:"name"`
	
	// +govalid:required
	// +govalid:email
	Email string `json:"email"`
	
	// +govalid:gte=0
	// +govalid:lte=120
	Age int `json:"age"`
}

func main() {
	// Valid user
	validUser := &User{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}
	
	if err := ValidateUser(validUser); err != nil {
		log.Printf("Valid user failed validation: %v", err)
	} else {
		fmt.Println("✓ Valid user passed validation")
	}
	
	// Invalid user - missing required fields
	invalidUser := &User{
		Name: "",        // Required field is empty
		Email: "invalid", // Invalid email format
		Age: 150,        // Age exceeds maximum
	}
	
	if err := ValidateUser(invalidUser); err != nil {
		fmt.Printf("✗ Invalid user failed validation: %v\n", err)
	}
}
```

### Step 2: Generate Validation Code

Run these commands in your terminal:

```bash
# Generate the validation code
go generate .

# Run the program
go run .
```

### Step 3: Expected Output

```
✓ Valid user passed validation
✗ Invalid user failed validation: field Name is required
```

The generated validation code creates a `ValidateUser` function that checks all the validation rules you specified with markers.

### Step 4: View Generated Code

After running `go generate .`, you'll see a new file (e.g., `main_govalid.go`) containing:

```go
// Code generated by govalid; DO NOT EDIT.

package main

import (
	"errors"
	"github.com/sivchari/govalid/validation/validationhelper"
)

var (
	ErrNilUser                     = errors.New("input User is nil")
	ErrUserNameRequiredValidation  = errors.New("field Name is required")
	ErrUserEmailRequiredValidation = errors.New("field Email is required")
	ErrUserEmailEmailValidation    = errors.New("field Email must be a valid email address")
	ErrUserAgeGTEValidation        = errors.New("field Age must be greater than or equal to 0")
	ErrUserAgeLTEValidation        = errors.New("field Age must be less than or equal to 120")
)

func ValidateUser(t *User) error {
	if t == nil {
		return ErrNilUser
	}

	if len(t.Name) == 0 {
		return ErrUserNameRequiredValidation
	}

	if len(t.Email) == 0 {
		return ErrUserEmailRequiredValidation
	}

	if !validationhelper.IsValidEmail(t.Email) {
		return ErrUserEmailEmailValidation
	}

	if !(t.Age >= 0) {
		return ErrUserAgeGTEValidation
	}

	if !(t.Age <= 120) {
		return ErrUserAgeLTEValidation
	}

	return nil
}
```

## Web API with Validation

This example shows how to use govalid in a web API for user registration with comprehensive validation.

### Project Structure

```
web-api/
├── main.go
├── go.mod
└── main_govalid.go (generated)
```

### Complete Example

**main.go**:
```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)

//go:generate govalid .

type CreateUserRequest struct {
    // +govalid:required
    // +govalid:minlength=2
    // +govalid:maxlength=50
    Name string `json:"name"`
    
    // +govalid:required
    // +govalid:email
    Email string `json:"email"`
    
    // +govalid:required
    // +govalid:minlength=8
    // +govalid:maxlength=100
    Password string `json:"password"`
    
    // +govalid:gte=13
    // +govalid:lte=120
    Age int `json:"age"`
    
    // +govalid:enum=admin,user,guest
    Role string `json:"role"`
    
    // +govalid:maxitems=10
    Tags []string `json:"tags,omitempty"`
}

type User struct {
    ID    string    `json:"id"`
    Name  string    `json:"name"`
    Email string    `json:"email"`
    Age   int       `json:"age"`
    Role  string    `json:"role"`
    Tags  []string  `json:"tags,omitempty"`
    CreatedAt time.Time `json:"created_at"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    
    // Parse JSON request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    // Validate the request using generated function
    if err := ValidateCreateUserRequest(&req); err != nil {
        http.Error(w, fmt.Sprintf("Validation failed: %v", err), http.StatusBadRequest)
        return
    }
    
    // Process the valid request
    user := createUser(req)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func createUser(req CreateUserRequest) *User {
    // Simulate user creation
    return &User{
        ID:        "user-123",
        Name:      req.Name,
        Email:     req.Email,
        Age:       req.Age,
        Role:      req.Role,
        Tags:      req.Tags,
        CreatedAt: time.Now(),
    }
}

func main() {
    http.HandleFunc("/users", createUserHandler)
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### How to Test

1. **Generate validation code**:
```bash
go generate .
```

2. **Run the server**:
```bash
go run .
```

3. **Test with valid data**:
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "secretpassword123",
    "age": 25,
    "role": "user",
    "tags": ["developer", "golang"]
  }'
```

**Expected response**:
```json
{
  "id": "user-123",
  "name": "John Doe",
  "email": "john@example.com",
  "age": 25,
  "role": "user",
  "tags": ["developer", "golang"],
  "created_at": "2024-01-01T12:00:00Z"
}
```

4. **Test with invalid data**:
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "A",
    "email": "invalid-email",
    "password": "123",
    "age": 200,
    "role": "superuser",
    "tags": []
  }'
```

**Expected response**:
```
Validation failed: field Name must have at least 2 characters
```

### Key Features Demonstrated

- **Multiple validation rules** per field
- **Automatic error handling** with clear messages
- **Zero-allocation validation** for high-performance APIs
- **Type-safe validation** functions
- **Easy integration** with existing HTTP handlers

## Collection and Enum Validation

This example demonstrates advanced validation with collections (slices, maps) and enums.

### Project Structure

```
collections/
├── main.go
├── go.mod
└── main_govalid.go (generated)
```

### Complete Example

**main.go**:
```go
package main

import (
    "fmt"
    "log"
)

//go:generate govalid .

type Product struct {
    // +govalid:required
    // +govalid:uuid
    ID string `json:"id"`
    
    // +govalid:required
    // +govalid:minlength=3
    // +govalid:maxlength=200
    Name string `json:"name"`
    
    // +govalid:required
    // +govalid:gt=0
    Price float64 `json:"price"`
    
    // +govalid:required
    // +govalid:enum=electronics,clothing,books,home,sports
    Category string `json:"category"`
    
    // +govalid:required
    // +govalid:minitems=1
    // +govalid:maxitems=10
    Images []string `json:"images"`
    
    // +govalid:maxitems=20
    Tags []string `json:"tags,omitempty"`
    
    // +govalid:gte=0
    Stock int `json:"stock"`
    
    // +govalid:required
    // +govalid:enum=draft,active,inactive,discontinued
    Status string `json:"status"`
    
    // +govalid:maxitems=50
    Attributes map[string]string `json:"attributes,omitempty"`
}

type ProductFilter struct {
    // +govalid:enum=electronics,clothing,books,home,sports
    Category string `json:"category,omitempty"`
    
    // +govalid:gte=0
    MinPrice float64 `json:"min_price,omitempty"`
    
    // +govalid:gte=0
    MaxPrice float64 `json:"max_price,omitempty"`
    
    // +govalid:maxitems=10
    Tags []string `json:"tags,omitempty"`
    
    // +govalid:gte=1
    // +govalid:lte=100
    Limit int `json:"limit,omitempty"`
    
    // +govalid:gte=0
    Offset int `json:"offset,omitempty"`
}

func main() {
    // Example 1: Valid product
    validProduct := &Product{
        ID:       "550e8400-e29b-41d4-a716-446655440000",
        Name:     "Wireless Headphones",
        Price:    199.99,
        Category: "electronics",
        Images:   []string{"img1.jpg", "img2.jpg"},
        Tags:     []string{"wireless", "audio", "bluetooth"},
        Stock:    50,
        Status:   "active",
        Attributes: map[string]string{
            "brand":    "TechCorp",
            "warranty": "2 years",
        },
    }
    
    if err := ValidateProduct(validProduct); err != nil {
        log.Printf("Valid product failed: %v", err)
    } else {
        fmt.Println("✓ Valid product passed validation")
    }
    
    // Example 2: Invalid product - multiple validation errors
    invalidProduct := &Product{
        ID:       "invalid-uuid",           // Invalid UUID format
        Name:     "AB",                     // Too short (minlength=3)
        Price:    -10.0,                    // Must be > 0
        Category: "invalid-category",       // Not in enum
        Images:   []string{},               // Empty slice (minitems=1)
        Tags:     make([]string, 25),       // Too many items (maxitems=20)
        Stock:    -5,                       // Must be >= 0
        Status:   "unknown",                // Not in enum
        Attributes: make(map[string]string, 60), // Too many items (maxitems=50)
    }
    
    if err := ValidateProduct(invalidProduct); err != nil {
        fmt.Printf("✗ Invalid product failed: %v\n", err)
    }
    
    // Example 3: Product filter validation
    filter := &ProductFilter{
        Category: "books",
        MinPrice: 0,
        MaxPrice: 50,
        Tags:     []string{"fiction", "bestseller"},
        Limit:    20,
        Offset:   0,
    }
    
    if err := ValidateProductFilter(filter); err != nil {
        log.Printf("Filter validation failed: %v", err)
    } else {
        fmt.Println("✓ Product filter is valid")
    }
    
    // Example 4: Invalid filter
    invalidFilter := &ProductFilter{
        Category: "invalid",     // Not in enum
        MinPrice: -10,          // Must be >= 0
        MaxPrice: -20,          // Must be >= 0
        Tags:     make([]string, 15), // Too many items (maxitems=10)
        Limit:    200,          // Exceeds maximum (lte=100)
        Offset:   -1,           // Must be >= 0
    }
    
    if err := ValidateProductFilter(invalidFilter); err != nil {
        fmt.Printf("✗ Invalid filter failed: %v\n", err)
    }
}
```

### How to Test

1. **Generate validation code**:
```bash
go generate .
```

2. **Run the example**:
```bash
go run .
```

### Expected Output

```
✓ Valid product passed validation
✗ Invalid product failed: field ID must be a valid UUID
✓ Product filter is valid
✗ Invalid filter failed: field Category must be one of: electronics, clothing, books, home, sports
```

### Key Features Demonstrated

- **UUID validation** for unique identifiers
- **Collection validation** with `minitems` and `maxitems`
- **Enum validation** for restricted value sets
- **Map validation** with size constraints
- **Numeric range validation** with `gt`, `gte`, `lt`, `lte`
- **String length validation** with `minlength` and `maxlength`
- **Optional field validation** with `omitempty` JSON tags

## Custom Types and Advanced Patterns

This example shows how to work with custom types and combine multiple validation patterns.

### Project Structure

```
advanced/
├── main.go
├── go.mod
└── main_govalid.go (generated)
```

### Complete Example

**main.go**:
```go
package main

import (
    "fmt"
    "log"
)

//go:generate govalid .

type UserRole string
type Priority int
type Status string

const (
    RoleAdmin UserRole = "admin"
    RoleUser  UserRole = "user"
    RoleGuest UserRole = "guest"
)

const (
    PriorityLow    Priority = 1
    PriorityMedium Priority = 2
    PriorityHigh   Priority = 3
)

type Task struct {
    // +govalid:required
    // +govalid:uuid
    ID string `json:"id"`
    
    // +govalid:required
    // +govalid:minlength=3
    // +govalid:maxlength=200
    Title string `json:"title"`
    
    // +govalid:maxlength=2000
    Description string `json:"description,omitempty"`
    
    // +govalid:required
    // +govalid:enum=admin,user,guest
    AssignedToRole UserRole `json:"assigned_to_role"`
    
    // +govalid:required
    // +govalid:enum=1,2,3
    Priority Priority `json:"priority"`
    
    // +govalid:required
    // +govalid:enum=pending,in_progress,completed,cancelled
    Status Status `json:"status"`
    
    // +govalid:minitems=1
    // +govalid:maxitems=10
    Tags []string `json:"tags"`
    
    // +govalid:maxitems=20
    Metadata map[string]string `json:"metadata,omitempty"`
    
    // +govalid:gte=1
    // +govalid:lte=100
    EstimatedHours int `json:"estimated_hours,omitempty"`
    
    // +govalid:required
    // +govalid:email
    CreatedBy string `json:"created_by"`
}

func main() {
    // Example 1: Valid task
    validTask := &Task{
        ID:             "550e8400-e29b-41d4-a716-446655440000",
        Title:          "Implement user authentication",
        Description:    "Add OAuth2 integration for user login",
        AssignedToRole: RoleUser,
        Priority:       PriorityMedium,
        Status:         "pending",
        Tags:           []string{"auth", "security", "backend"},
        Metadata: map[string]string{
            "sprint": "2024-Q1",
            "team":   "backend",
        },
        EstimatedHours: 8,
        CreatedBy:      "project.manager@company.com",
    }
    
    if err := ValidateTask(validTask); err != nil {
        log.Printf("Valid task failed: %v", err)
    } else {
        fmt.Println("✓ Valid task passed validation")
    }
    
    // Example 2: Invalid task - business logic validation
    invalidTask := &Task{
        ID:             "550e8400-e29b-41d4-a716-446655440000",
        Title:          "Critical security fix",
        Description:    "Fix SQL injection vulnerability",
        AssignedToRole: RoleGuest,      // High priority task assigned to guest
        Priority:       PriorityHigh,   // This combination should fail business logic
        Status:         "pending",
        Tags:           []string{"security", "urgent"},
        Metadata:       map[string]string{},
        EstimatedHours: 4,
        CreatedBy:      "security@company.com",
    }
    
    if err := ValidateTask(invalidTask); err != nil {
        log.Printf("Task validation failed: %v", err)
    } else {
        // Additional business logic validation
        if err := CreateTask(invalidTask); err != nil {
            fmt.Printf("✗ Business logic validation failed: %v\n", err)
        }
    }
    
    // Example 3: Multiple validation errors
    brokenTask := &Task{
        ID:             "invalid-uuid",        // Invalid UUID
        Title:          "AB",                  // Too short
        Description:    "",                    // Valid (optional)
        AssignedToRole: "invalid-role",        // Not in enum
        Priority:       0,                     // Not in enum
        Status:         "unknown",             // Not in enum
        Tags:           []string{},            // Too few items (minitems=1)
        Metadata:       make(map[string]string, 25), // Too many items (maxitems=20)
        EstimatedHours: 200,                   // Too high (lte=100)
        CreatedBy:      "invalid-email",       // Invalid email
    }
    
    if err := ValidateTask(brokenTask); err != nil {
        fmt.Printf("✗ Multiple validation errors: %v\n", err)
    }
}

func CreateTask(task *Task) error {
    // First run generated validation
    if err := ValidateTask(task); err != nil {
        return fmt.Errorf("invalid task: %w", err)
    }
    
    // Additional business logic validation
    if task.Priority == PriorityHigh && task.AssignedToRole == RoleGuest {
        return fmt.Errorf("high priority tasks cannot be assigned to guests")
    }
    
    // Simulate saving task
    fmt.Printf("✓ Task created successfully: %s\n", task.Title)
    return nil
}
```

### How to Test

1. **Generate validation code**:
```bash
go generate .
```

2. **Run the example**:
```bash
go run .
```

### Expected Output

```
✓ Valid task passed validation
✗ Business logic validation failed: high priority tasks cannot be assigned to guests
✗ Multiple validation errors: field ID must be a valid UUID
```

### Key Features Demonstrated

- **Custom type validation** with UserRole, Priority, Status
- **Combined validation rules** (required + enum, minlength + maxlength)
- **Business logic validation** on top of generated validation
- **Multiple validation errors** showing first failure
- **Optional field handling** with `omitempty` tags
- **Complex data structures** with nested validation

## Best Practices Summary

### 1. Start Simple
Begin with basic validation like `required` and `minlength` before adding complex rules.

### 2. Layer Your Validation
1. **Generated validation** - Basic field validation
2. **Business logic** - Complex rules and relationships
3. **Integration validation** - External dependencies

### 3. Use Appropriate Markers
- **String validation**: `required` , `minlength` , `maxlength` , `email` , `url` , `uuid`

- **Numeric validation**: `gt` , `gte` , `lt` , `lte`

- **Collection validation**: `minitems` , `maxitems`

- **Enum validation**: `enum` with comma-separated values

### 4. Handle Errors Gracefully

#### Basic Error Handling
```go
if err := ValidateStruct(data); err != nil {
    // Log the error for debugging
    log.Printf("Validation failed: %v", err)
    
    // Return user-friendly error
    return fmt.Errorf("invalid data: %w", err)
}
```

#### Advanced Error Handling with errors.Is
```go
package main

import (
    "errors"
    "fmt"
    "log"
)

//go:generate govalid .

type User struct {
    // +govalid:required
    Name string `json:"name"`
    
    // +govalid:required
    // +govalid:email
    Email string `json:"email"`
    
    // +govalid:gte=18
    Age int `json:"age"`
}

func ProcessUser(user *User) error {
    if err := ValidateUser(user); err != nil {
        // Handle specific validation errors
        if errors.Is(err, ErrUserNameRequiredValidation) {
            return fmt.Errorf("user name is mandatory")
        }
        
        if errors.Is(err, ErrUserEmailRequiredValidation) {
            return fmt.Errorf("email address is required")
        }
        
        if errors.Is(err, ErrUserEmailEmailValidation) {
            return fmt.Errorf("please provide a valid email address")
        }
        
        if errors.Is(err, ErrUserAgeGTEValidation) {
            return fmt.Errorf("user must be at least 18 years old")
        }
        
        // Handle nil pointer error
        if errors.Is(err, ErrNilUser) {
            return fmt.Errorf("user data is missing")
        }
        
        // Generic fallback
        return fmt.Errorf("validation failed: %w", err)
    }
    
    // Process valid user
    fmt.Printf("Processing user: %s (%s)\n", user.Name, user.Email)
    return nil
}

func main() {
    // Example usage with specific error handling
    users := []*User{
        nil,                                    // Nil user
        {Name: "", Email: "", Age: 0},         // Missing required fields
        {Name: "John", Email: "invalid", Age: 16}, // Invalid email and age
        {Name: "Jane", Email: "jane@example.com", Age: 25}, // Valid user
    }
    
    for i, user := range users {
        fmt.Printf("User %d: ", i+1)
        if err := ProcessUser(user); err != nil {
            fmt.Printf("Error - %v\n", err)
        } else {
            fmt.Println("Success")
        }
    }
}
```

#### HTTP API Error Handling
```go
func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    if err := ValidateUser(&user); err != nil {
        // Map validation errors to HTTP status codes
        var statusCode int
        var message string
        
        switch {
        case errors.Is(err, ErrNilUser):
            statusCode = http.StatusBadRequest
            message = "User data is required"
        case errors.Is(err, ErrUserNameRequiredValidation):
            statusCode = http.StatusBadRequest
            message = "Name is required"
        case errors.Is(err, ErrUserEmailEmailValidation):
            statusCode = http.StatusBadRequest
            message = "Invalid email format"
        case errors.Is(err, ErrUserAgeGTEValidation):
            statusCode = http.StatusBadRequest
            message = "User must be at least 18 years old"
        default:
            statusCode = http.StatusBadRequest
            message = "Validation failed"
        }
        
        http.Error(w, message, statusCode)
        return
    }
    
    // Process valid user...
}
```

### 5. Performance Benefits
- **Zero allocations** for all validation operations
- **Compile-time safety** with generated functions
- **5x to 50x faster** than runtime validation libraries

These examples demonstrate the power and flexibility of govalid for creating robust, high-performance validation in Go applications.
