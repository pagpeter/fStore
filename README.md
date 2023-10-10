# fStore - efficient storage for ~~browser fingerprints~~ data (PoC)

fStore is a database wrapper that compacts data before it gets saved.

Its useful when you know that you will store mostly the same data, but in a different order.
It is not production ready and should not be used in a serious enviroment, this is just a proof of concept!

## Example

```go
f := Listener()

// Setup
f.EnableDebug()
f.DontHash = []string{"DynamicValue"} // names of keys that shouldnt be minified
f.Threshhold = 5 // String length threshhold, must be over 1
f.UseKeyCompression = false // If keys should be compressed as well

res, err := f.Store(data) // Result (interface{}) and error
fmt.Println(json.Marshal(res), err)
```
