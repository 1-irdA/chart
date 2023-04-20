# chart

Generate barchart for cli app with no dependency and colors.       

## Barchart

```go
package main

import (
	"fmt"

	"github.com/garrou/chart"
)

func main() {
	label := []string{"go", "cpp", "js", "java"}
	values := []float64{17, 14.95, 35.05, 54.65}
	colors := []chart.AnsiColor{chart.Yellow, chart.Red, chart.Green, chart.Blue}

	barChart := chart.NewBarChart("Languages", 5, label, values, colors, "=")
	result := barChart.Generate(chart.Horizontal)

	fmt.Println(result)
}
```

```text
                Languages

go     | ===                        17.00
rb     | ==                         13.00
rs     |                            3.75
cpp    | ===                        16.25
js     | ==                         14.95
java   | =======                    35.05
carbon | ==========                 54.65
lua    | ===================        99.99

```

## Vertical barchart

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

## Stackchart

```go
package main

import (
	"fmt"

	"github.com/garrou/chart"
)

func main() {
	label := []string{"go", "cpp", "js", "java"}
	values := []float64{17, 13, 16.25, 54.65}
	fills := []string{"*", "+", "=", "-"}
	colors := []chart.AnsiColor{chart.Yellow, chart.Red, chart.Green, chart.Blue}

	stackChart := chart.NewStackChart("Languages", 2, label, values, fills, colors)
	result := stackChart.Generate(chart.Horizontal)

	fmt.Println(result)
}
```

```text
        Languages

********++++++========---------------------------
********++++++========---------------------------
********++++++========---------------------------

* go : 17.0   + cpp : 13.0   = js : 16.2   - java : 54.6
```