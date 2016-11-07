# LLR Package
Contains structs, consts, "enums," and methods specific to LulaRoe items.

## Usage
### Import LLR Package
```
import (
    "github.com/brianshef/roetisserie/llr"
)
```

### Status "enum" example
```
import (
    "fmt"
    "github.com/brianshef/roetisserie/llr"
)

func main() {
    status := llr.ORDERED
    fmt.Println(status)     //  ORDERED
    status++
    fmt.Println(status)     //  RECEIVED
}
```

## Testing
To test the `llr` package, execute the following:
```
go test github.com/brianshef/roetisserie/llr
```
