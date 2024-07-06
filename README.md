## Error

```bash
goa gen github.com/bkeane/gists/design

exit status 1
0x1400035b0c0[1/1]0x1400019ac40[services/imports_shared/design/design.go:13] unknown attribute "id" in attribute
```

## Behavioral Observances

### Remove `View(...)` & Kepp `CollectionOf(...)` => Works!

```go
// `services/imports_shared/design/design.go`
var _ = Service("AsJson", func() {
	Method("index", func() {
		Description("Returns all rows")
		Result(CollectionOf(shared.Row))
		HTTP(func() {
			GET("/")
		})
	})
})
```

### Remove `CollectionOf(...)` & Keep `View(...)` => Works!

```go
// `services/imports_shared/design/design.go`
var _ = Service("AsJson", func() {
	Method("index", func() {
		Description("Returns all rows")
		Result(shared.Row, func() {
			View("benign")
		})
		HTTP(func() {
			GET("/")
		})
	})
})
```
