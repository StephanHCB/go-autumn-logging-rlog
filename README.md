# go-autumn-logging-rlog

[rlog](https://github.com/go-logr/logr) golang logging plugin for go-autumn-logging.

## About go-autumn

A collection of libraries for [enterprise microservices](https://github.com/StephanHCB/go-mailer-service/blob/master/README.md) in golang that
- is heavily inspired by Spring Boot / Spring Cloud
- is very opinionated
- names modules by what they do
- unlike Spring Boot avoids certain types of auto-magical behaviour
- is not a library monolith, that is every part only depends on the api parts of the other components
  at most, and the api parts do not add any dependencies.  

Fall is my favourite season, so I'm calling it go-autumn.

## About go-autumn-logging

An interface for pluggable logging support with context integration, 
see [go-autumn-logging](https://github.com/StephanHCB/go-autumn-logging).

## About go-autumn-logging-rlog

Implementation that pulls in logging using [rlog](https://github.com/go-logr/logr),
which funny enough is another (though very minimal) log abstraction.

### Features

* supports multiple actual backend loggers

### Limitations

* only has numeric log levels, so we had to map to those.
* no default logger, you will always need to place a logger on the context

## Usage

### Library Authors

Do NOT import this. You are doing it WRONG. You'll take away the application author's choice of 
logging libraries, which is the whole point of using go-autumn-logging in the first place.

Use [go-autumn-logging](https://github.com/StephanHCB/go-autumn-logging)!

### Application Authors

You're all set with this dependency.

Rlog's interface does not support logging without a context, so you must call the function provided by rlog 
to add the logger to the context.

```
ctx = logr.NewContext(ctx, logger)
```

### How To Use

Use call chaining style:

```
import "github.com/StephanHCB/go-autumn-logging"

func someFunction(ctx context.Context) {
    aulogging.Logger.Ctx(ctx).Warn().Print("something bad has happened")
}
```

or if you prefer (does the same thing):

```
import "github.com/StephanHCB/go-autumn-logging"

func someFunction(ctx context.Context) {
    aulogging.Warn(ctx, "something bad has happened")
}
```
