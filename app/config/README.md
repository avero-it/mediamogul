# Config package

This package helps getting the  configuration

## Default behaviour
Default loader behaviour is to put priority within environment variables, allowing to either specify a config file flag or a direct file name to read from. In case nothing is specified, it will take `.env` as default file name and `-config` as default flag name.
In case no custom behaviour is defined, if flag is specified by the application will return an error in case the file doesn't exists. In case nothing is specified by the application and the flag is not set either, it will not thrown an error, even in the case the default file is not accesible, as it is only a backup.

In case specific file name is specified, will take priority over flag and will not use backup (default config file).

Environment variables will always take priority over the files so variables can be overwritten.

## Validations
It also includes a validation part that can be defined within the structure that you want to parse against.

## Usage
```go
type config struct {
	Service1Config struct {
		Host string `env:"SERVICE_1_HOST" validate:"required"`
	}
	Service2Config struct {
        Host string `env:"SERVICE_2_HOST" validate:"required"`
        Limit int `env:"SERVICE_2_LIMIT" validate:"required"`
	}
}

func main() {
    cfg := &config{}

    if err := configx.Loader().Load(cfg); err != nil {
        panic(err)
    }
}
```

## Customising loader:

### Reading only from environment variables
```go
    cfg := &config{}
    loader := configx.Loader().OnlyEnvironment()

    if err := cloader.Load(cfg); err != nil {
        panic(err)
    }
```

### Specifying a custom configuration file only
***Notice that it will not backup into the default configuration file as it was explicitly defined and it cannot work together with flag, so you have to choose either one or the other.***

```go
    cfg := &config{}
    loader := configx.Loader().WithFileName("whatever")

    if err := cloader.Load(cfg); err != nil {
        panic(err)
    }
```

### Specifying a custom flag
***Notice that it cannot work together with flag, so you have to choose either one or the other.***

```go
    cfg := &config{}
    
    loader := configx.Loader().
        WithFileFlag("new-flag-for-file").
        WithDefaultFileName(".default-file")

    if err := cloader.Load(cfg); err != nil {
        panic(err)
    }
```

## Validations

As mentioned before, this package also includes the capability of validate the defined config structure using [go-playground/validator](https://github.com/go-playground/validator) package.

Take a look to other validations examples [here](https://github.com/go-playground/validator/tree/master/_examples).