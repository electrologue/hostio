# Host.io API Go Client

A simple API Client written in Go for https://Host.io/

API Documentation: https://host.io/docs

## Examples

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/hostio"
)

func main() {
	client := hostio.NewClient("token")

	data, err := client.Web(context.Background(), "example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/hostio"
)

func main() {
	client := hostio.NewClient("token")

	data, err := client.DNS(context.Background(), "example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/hostio"
)

func main() {
	client := hostio.NewClient("token")

	data, err := client.Related(context.Background(), "example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/hostio"
)

func main() {
	client := hostio.NewClient("token")

	data, err := client.Full(context.Background(), "example.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/electrologue/hostio"
)

func main() {
	client := hostio.NewClient("token")

	pager := &hostio.Pager{
		Limit: 5,
		Page:  5,
	}

	data, err := client.Domains(context.Background(), hostio.NS, "google.com", pager)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
```
