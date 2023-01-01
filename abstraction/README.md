# Abstraction in Go Lang

Abstraction in Go is the process of hiding the implementation details of a type or function and exposing only the relevant information and behaviors to the users of that type or function. The goal of abstraction is to simplify the use of complex types or functions by providing a clear and easy-to-use interface.

## How to achieve abstaction ?
There are several ways to achieve abstraction in Go:

### 1. Functions
By writing functions that take arguments and return values, you can hide the implementation details of a particular task and expose only the necessary inputs and outputs to the users of the function
### 2. Interfaces
Interfaces allow you to specify a set of methods that a concrete type must implement in order to satisfy the interface. This allows you to define a common behavior that can be implemented by multiple types
### 3. Embedding
Go does not have inheritance, but it does have a mechanism called embedding that allows you to include the fields and methods of one type in another type. This can be used to achieve abstraction by creating a type with a common set of fields and methods that can be embedded in other types
### 4. Packages
By organizing your code into packages, you can hide the implementation details of certain types and functions and expose only the relevant information to other parts of the program
### 5. Pointers
By using pointers to types, you can pass around the address of a value rather than the value itself. This can be used to abstract away the details of a particular value and allow users to manipulate the value indirectly 
### 6. Constructors
In Go, you can define functions that return a pointer to a new instance of a type. These functions are known as constructors and can be used to abstract away the details of creating a new instance of a type