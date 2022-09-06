<img src="https://raw.githubusercontent.com/finacore/.github/main/horizontal.svg" width="30%">
<br><br>

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
	Message: "the error message goes here",
}
```

2. __Use the CreateDefaultError function__, the _CreateDefaultError__ function provide a way to create a new __DefaultError__ object just passing a desired error message.

```go
err := commonserrors.CreateDefaultError("the error message goes here")
```

3. __Use the MakeDefaultError function__, This way is used when you already have somme error and wanto to oconvert it directly to __JSON__ using this model.

```go
err := commonserrors.MakeDefaultError(sommeError)
```

In order to facilitate the processing of errors created, this kind of error provides a method to identificate the error
through the code. the method is __Status__ and return the error object, that facilitate the usability as shows below.

```go
err := commonserrors.CreateDefaultError("internal server error").Status(500)
```

### ValidationError

The __ValidationError__ data structure has as objetive to map a __field__ with some __validation error__ received from the validation process. 

Is important to know that this process has no intention to make validation. It's just a data structure to standardize an validation output in __JSON__ format.

There is two ways to create the __ValidationError__ instance, these are arranged below:

1. __Create object from struct__, this is simple but verbose. I really think is no make sense use this way when there is disponible a very simple method that can be used instead. See the next one.

```go
err := &commonserrors.ValidationError{
	Field: "field_name",
	Message: "the error message goes here",
}
```
2. __Use the CreateValidationError function__, the _CreateValidationError__ function provide a way to create a new __ValidationError__ object just passing the name of field and the desired error message.

```go
err := commonserrors.CreateValidationError("field_name", "the error message goes here")
```

### Error()

Both errors objects provide the method __Error()__ that enable the user to obtains the error in a string format. It is simple and follow the patter of the default __error__ object provided nativelly in __Go__ programing language.

This method usage is simple and can be seen in the example bellow.

```go
formatedStringError := err.Error()
```
