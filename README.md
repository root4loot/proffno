# proffNO Go Library

A Go library used to retrieve subsidiaries for a Norwegian company based on ownership percentage and depth level. Parses data from [proff.no](https://beta.proff.no)

## Installation

```
go get github.com/root4loot/proffNO
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/root4loot/proffNO"
)

func main() {
    // Get subsidiaries of a company and its sub-subsidiaries owned by more than 50%
    subs, _ := proffno.GetSubsidiaries("DnB Bank ASA", 2, 50)

    for _, sub := range subs {
        fmt.Println(sub.Name)
    }
}
```

### Output

```
DnB Livsforsikring AS
DnB Private Equity VI (is) AS
DnB Private Equity IV (is) AS
DnB Private Equity II (is) AS
DnB Private Equity V (is) AS
DnB Næringseiendom AS
DnB Eiendomsinvest 2 AS
DnB Kjøpesenter og Hotell AS
DnB Liv Eiendom Sverige AS
DnB Kontor AS
DnB Eiendomsholding AS
Sbanken ASA
DnB Eiendomsutvikling AS
Godgata AS
Godbitene AS
Mosetertoppen Hafjell AS
Skandinaviske Handelsparker AS
B&R Holding AS
DnB Asset Management Holding AS
DnB Asset Management AS
S Fra Dnb AS
DnB Eiendom AS
DnB Næringsmegling AS
Godfjellet AS
Yellow Holding AS
DnB Ventures AS
Ocean Holding AS
Rq Invest AS
Olympic Zeus AS
Godskipet 8 AS
Godskipet 2 AS
Godskipet 4 AS
DnB Gjenstandsadministrasjon AS
DnB Boligkreditt AS
```


