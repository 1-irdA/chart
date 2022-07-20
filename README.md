# chart

Generate barchart for cli app.        
Works with percentage.      

```go
package main

import (
	"fmt"

	"github.com/garrou/chart"
)

func main() {
	label := []string{"go", "rb", "rs", "cpp", "js", "java", "carbon", "lua"}
	values := []float64{17, 13, 3.75, 16.25, 14.95, 35.05, 54.65, 99.99}

	barChart := chart.New("Languages", 5, label, values, "=")
	result := barChart.Generate(chart.Horizontal)

	fmt.Println(result)
}
```

## Horizontal

```text
                Languages

go     | ===                        17.00 %
rb     | ==                         13.00 %
rs     |                            3.75 %
cpp    | ===                        16.25 %
js     | ==                         14.95 %
java   | =======                    35.05 %
carbon | ==========                 54.65 %
lua    | ===================        99.99 %

```

## Vertical

```text
                Languages
100.00|
 95.00|                      =
 90.00|                      =
 85.00|                      =
 80.00|                      =
 75.00|                      =
 70.00|                      =
 65.00|                      =
 60.00|                      =
 55.00|                      =
 50.00|                   =  =
 45.00|                   =  =
 40.00|                   =  =
 35.00|                =  =  =
 30.00|                =  =  =
 25.00|                =  =  =
 20.00|                =  =  =
 15.00| =        =     =  =  =
 10.00| =  =     =  =  =  =  =
  5.00| =  =     =  =  =  =  =
      -------------------------
        g  r  r  c  j  j  c  l
        o  b  s  p  s  a  a  u
                 p     v  r  a
                       a  b
                          o
                          n
```