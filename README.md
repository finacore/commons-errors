# Commons Errors

This module encapsulate some methods and data structure responsibles to organize and manage the erros obtained
duruing the API execution.

Among the characteristics, the one that stand out are the capability to convert the data structure to an string, following the error interface implemented by the golang 'error'.

To fulfill their objective, this package encapsulate the following errors data structure that will be better described lower.

* __DefaultError:__ Error data structure composed by the key _message_ and their value;
* __ValidationError:__  Error data structure composed by the tuple of keys (_field_ and _message_).

## Usage

To use this module, the first thing that you want to do is get it using go module. The first thing to do is execute the following command.

```bash
go get github.com/finacore/commons-errors
```
Now you are ready to use the application in the right way.

### DefaultError

The __DefaultError__ is an implementation similar to the default __error__ type, but it is ready to be convertible to __JSON__ in a friendly format.

There are three ways to create an __DefaultError__ instance, these are arranged below:

1. __Create object from struct__, this is simple but verbose. I really think is no make sense use this way when there is disponible a very simple methods that can be used instead. See the next two ways.

```go
err := &commonserrors.DefaultError{
	Message: "the error message goes here"
}
```

2. __Use the CreateDefaultError function__, the _CreateDefaultError__ method provide a way to create a new __DefaultError__ object just passing a desired error message.

```go
err := commonserrors.CreateDefaultError("the error message goes here")
```

3. __Use the MakeDefaultError function__, This way is used when you already have somme error and wanto to oconvert it directly to __JSON__ using this model.

```go
err := commonserrors.MakeDefaultError(sommeError)
```

<!-- ## Validation

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
``` -->