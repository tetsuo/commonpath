# commonpath

Returns the longest common sub-path given a sequence of path names.

Ported from Python `commonpath`:

* [github.com/python/cpython/Lib/posixpath.py#L550](https://github.com/python/cpython/blob/e6ef47ac229b5c4a62b9c907e4232e350db77ce3/Lib/posixpath.py#L550)
* [github.com/python/cpython/Lib/ntpath.py#L827](https://github.com/python/cpython/blob/e6ef47ac229b5c4a62b9c907e4232e350db77ce3/Lib/ntpath.py#L827)


## Example

```go
paths := []string{
  "/a/b/c/src",
  "/a/b/c/docs",
  "/a/b/tests",
}

common, _ := commonpath.CommonPath(paths)

fmt.Println("Common path:", common)
// Output: Common path: /a/b
```
