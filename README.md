# proffno

[![Test](https://github.com/root4loot/proffno/actions/workflows/test-daily.yml/badge.svg?branch=main)](https://github.com/root4loot/proffno/actions/workflows/test-daily.yml)

A simple package used to obtain a list of subsidiaries for a Norwegian company based on ownership percentage and depth level. Parses data from [proff.no](https://proff.no)

## Installation

```
go get github.com/root4loot/proffno
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/root4loot/proffno"
)

func main() {
	// Get subsidiaries with ownership percentage greater than 50%
	subs, _ := proffno.GetSubsidiaries("DnB Bank ASA", 50)

	for _, sub := range subs {
		fmt.Println(sub.Name)
	}
}
```

### Output

```
DnB Livsforsikring AS
DnB Eiendomsutvikling AS
Imove AS
B&R Holding AS
DnB Asset Management Holding AS
DnB Invest Holding AS
DnB Eiendom AS
Godskipet 9 AS
DnB Næringsmegling AS
Godfjellet AS
Ocean Holding AS
DnB Ventures AS
DnB Gjenstandsadministrasjon AS
DnB Boligkreditt AS
```


