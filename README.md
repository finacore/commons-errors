# Commons Library

Projet to organize some code that are used in a set of this projects repository.

## Validation

The validation model receives a model (struct) that contains the validation tag, in case of some validation fail the return will be a _ValidationError_ array, other else _nil_.

```go
import "github.com/finacore/commons/validation"

//define struct model
type User struct {
	Name    string `validate:"required,min=3,max=256"`
	Surname string `validate:"required,min=3,max=256"`
	Email   string `validate:"omitempty,email"`
}

//create a model based on struct
userModel := User{
    Surname: "da Silva"
    Email: "dasilva@gmail.com"
}

err := validation.ValidateModel(userModel)
```

## Errors

in this package you can use the follow two errors:

* __Default__
* __Validation__

The default error generate an interface containing just one field __Message__.

```go
import "github.com/finacore/commons/baseerror"

err := baseerror.Default("some error message")
```

The validationError is formed by the fields  __Field:__ and __Message:__ and can be created as showed below.

```go
import "github.com/finacore/commons/baseerror"

err := baseerror.Validation("field-name", "some error message")
```