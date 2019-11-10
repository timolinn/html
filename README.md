# HTML Parser

This is a simple html parser, it really is nothing super serious. However, it will power [nginB](https://github.com/timolinn/nginB), a mini browser engine.

## Usage

```go
    package main

    import (
            "fmt"
            "github.com/timolinn/html-parser"
        )

    func main() {
        tmpl := []byte(`<html><body><p id="hw">Hello, world</p></body></html>`)
        parsed := html.Parse(tmpl)
        fmt.Printf("%+v\n", parsed)
    }
```

## Contributing

While I'm not currently investing heavily because of time. I'm open to welcoming new pull requests for new features or general improvements.
