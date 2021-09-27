# grok
A simple library to parse grok patterns with Go. Forked from 
[vjeantet/grok](https://github.com/vjeantet/grok).

# Installation
Make sure you have a working Go environment.

```sh
go get github.com/occamshub-dev/grok
```

# Use in your project
```go
import "github.com/occamshub-dev/grok"
```

# Usage
## Available patterns and custom ones
By default this grok package contains only patterns you can see in patterns/grok-patterns file.

When you want to add a custom pattern, use the grok.AddPattern(nameOfPattern, pattern), see the example folder for an example of usage.
You also can load your custom patterns from a file (or folder) using grok.AddPatternsFromPath(path), or PatterndDir configuration.

## Parse all or only named captures
```go
g, _ := grok.New()
values, _  := g.Parse("%{COMMONAPACHELOG}", `127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`)

g, _ = grok.NewWithConfig(&grok.Config{NamedCapturesOnly: true})
values2, _ := g.Parse("%{COMMONAPACHELOG}", `127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`)
```
values is a map with all captured groups
values2 contains only named captures

# Examples
```go
package main

import (
	"fmt"

	"github.com/occamshub-dev/grok"
)

func main() {
	g, _ := grok.New()
	values, _ := g.Parse("%{COMMONAPACHELOG}", `127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`)

	for k, v := range values {
		fmt.Printf("%+15s: %s\n", k, v)
	}
}
```

output:
```
       response: 404
          bytes: 207
       HOSTNAME: 127.0.0.1
       USERNAME: -
       MONTHDAY: 23
        request: /index.php
      BASE10NUM: 207
           IPV6:
           auth: -
      timestamp: 23/Apr/2014:22:58:32 +0200
           verb: GET
    httpversion: 1.1
           TIME: 22:58:32
           HOUR: 22
COMMONAPACHELOG: 127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207
       clientip: 127.0.0.1
             IP:
          ident: -
          MONTH: Apr
           YEAR: 2014
         SECOND: 32
            INT: +0200
           IPV4:
         MINUTE: 58
     rawrequest:
```

# Example 2
```go
package main

import (
  "fmt"

  "github.com/occamshub-dev/grok"
)

func main() {
  g, _ := grok.NewWithConfig(&grok.Config{NamedCapturesOnly: true})
  values, _ := g.Parse("%{COMMONAPACHELOG}", `127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`)

  for k, v := range values {
    fmt.Printf("%+15s: %s\n", k, v)
  }
}
```

output:
```
      timestamp: 23/Apr/2014:22:58:32 +0200
           verb: GET
     rawrequest:
          bytes: 207
           auth: -
        request: /index.php
    httpversion: 1.1
       response: 404
COMMONAPACHELOG: 127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207
       clientip: 127.0.0.1
          ident: -
```
