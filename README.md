# Commons Library

Projet to organize some code that are used in a set of this projects repository.

## Errors

Package to group the different kind of errors.

### default Error

```golang
err := baseerror.Default("Somme error message")
```

### validation error

```golang
err := baseerror.Validation("fieldName", "somme message")
```