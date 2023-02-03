## template

```
	github.com/Masterminds/sprig/v3 v3.1.0
	github.com/aymerick/raymond v2.0.2+incompatible
```

## usage

- `only once template.RegisterSettings`

```go
package main

import (
	"github.com/sinlov/drone-info-tools/template"
)

func init() {
  template.RegisterSettings(template.DefaultFunctions)
}
```