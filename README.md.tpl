APIUtils
========
[![codecov](https://codecov.io/gh/JoeReid/apiutils/branch/master/graph/badge.svg)](https://codecov.io/gh/JoeReid/apiutils)
[![Go Report Card](https://goreportcard.com/badge/github.com/JoeReid/apiutils)](https://goreportcard.com/report/github.com/JoeReid/apiutils)
[![GoDoc](https://godoc.org/github.com/JoeReid/apiutils?status.svg)](https://godoc.org/github.com/JoeReid/apiutils)

A handy collection of utility libs for designing apis in go

Codecs
------

Codecs should live seperate to handlers. 
The lines between business logic and data encoding should be kept sharp.
The codec package enables this to be done.

e.g. a simple helloworld application

```go
!!EXEC cat ./example/codec/main.go
```

```
$ curl 'localhost:8080/json'
{"hello":"world","timestamp":"2020-06-27T01:29:38.271839357+01:00"}

$ curl 'localhost:8080/yaml'
yamlhello: world
yamltimestamp: 2020-06-27T01:29:42.85396202+01:00
```


### But wait, theres more

Why not let the api consumer decide the format they want

```go
!!EXEC cat ./example/selector/main.go
```

```
$ curl 'localhost:8080/hello?codec=json'
{"hello":"world","timestamp":"2020-06-27T01:32:01.945250157+01:00"}

$ curl 'localhost:8080/hello?codec=yaml'
yamlhello: world
yamltimestamp: 2020-06-27T01:32:06.962818404+01:00
```

Pagination
----------

Write your handler to fetch paginated data, and let the lib wory about
getting the values from the consumer.

```go
!!EXEC cat ./example/pagination/main.go
```
