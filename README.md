# test ebiten game

## references

- https://ebitengine.org/en/tour/hello_world.html
- https://ebitengine.org/en/documents/cheatsheet.html
- https://github.com/hajimehoshi/ebiten/tree/main/examples
- https://ebitengine.org/en/documents/webassembly.html

## resources from ebiten examples plus:
- https://www.fontsquirrel.com/fonts/Silkscreen


## desktop

### run

```
go run main.go
```

### build

```
go build
```


## web

### run

```
go run github.com/hajimehoshi/wasmserve@latest .
```

### build

```
env GOOS=js GOARCH=wasm go build -o game.wasm ebitengine_test
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
http-server . -a 0.0.0.0 -p 8080 -s -c-1 --cors &
```
