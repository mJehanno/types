# Types


## Introduction

Types is a library providing some usefule types like Option (found in Haskell as `Maybe` and in Rust/Gleam as `Option`).

## Usage

### Option

Option is a type made to handle case where there can be an absence of value. Most of the time, in go we use pointer to express this scenario.
But Golang does not force you to check if a pointer is nil before using it leading into the famous `nil pointer dereference error`.

Using it is pretty simple, instead of using a pointer to express a nillable value, use `Option[T]` instead. When you want to access the value, just use `UnwrapOrDefault` or `UnwrapOrElse`. 
The first let you provide a default value if the `Option` is a `None` type (no value), or will return the underlying value if it's a `Some`.

You can also switch on the `Option` type : 
```go
func GenerateValue() Option<int> {
    // ... do some stuff here
}

func main() {
    v :=GenerateValue() 

    switch v.(type){
        case Some[int]:
            // do something
        case None[int]:
            // do other thing
    }
}

```

To sum it up, `Option[T]` is used when something can either have a value or be null. This is reflected by the provided type `Some` and `None`. You have to **unwrap** if you want to use the value.
You can switch on its type (`Some` or `None`) if you want to do something in either cases.
