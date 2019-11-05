# go-utils
自己的一些公用简单包


## logex

使用方法

```go
logex.SetName("logname")
logger := logex.Wrap("request_id", "module name")
```