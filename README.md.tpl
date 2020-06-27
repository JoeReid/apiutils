APIUtils
========

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

### But wait, theres more

Why not let the api consumer decide the format they want

```go
!!EXEC cat ./example/selector/main.go
```

Pagination
----------

Write your handler to fetch paginated data, and let the lib wory about
getting the values from the consumer.

```go
!!EXEC cat ./example/pagination/main.go
```
