# Gen Charging Station

## instalation

```terminal
go get -v github.com/ghuvrons/genchargingstation
```

## Example

``` go
package main

import (
    "fmt"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/ghuvrons/genchargingstation"
)

func main() {
    r := chi.NewRouter()

    r.Handle("/gen-charge", genchargingstation.Handler(func(data *genchargingstation.ChargingData) {
        fmt.Println(data.ToString())
        return
    }))

    http.ListenAndServe(":3000", r)
}
```
