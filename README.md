# nuri
Nuri is web library on golang

## Version
v1.0.0

## Example

print "Hello, World!" with status code, 200
```go
nuri.GET("/", func(c nuri.Context) (int, string) {
  return 200, "Hello, World!"
})
```

---

print "<Hello, World!>"
```go
nuri.GET("/", func(c nuri.Context) (int, string) {
  return c.ToText(200, "<Hello, World!>")
})
```

---

print HTML

```go
nuri.GET("/", func(c nuri.Context) (int, string) {
  return 200, "<h1>Hello, World!</h1>"
})
```

---

How to run

```go
nuri.Run(":5000")
```
It runs on localhost:5000
