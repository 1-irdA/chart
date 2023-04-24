# chart

Generate barchart for cli app with no dependency and colors.       

## Bar chart

```go
package main

import (
	"fmt"

	"github.com/garrou/chart"
)

func main() {
	series := []chart.Series{
		chart.NewColoredSeries("go", "=", 17, chart.Yellow),
		chart.NewColoredSeries("cpp", "*", 14.95, chart.Red),
		chart.NewColoredSeries("java", "+", 54.65, chart.Blue),
	}
	ch := chart.NewBarChart("Languages", 5, series)
	result := ch.Generate(chart.Vertical)

	fmt.Println(result)
}
```

```text
      Languages

 55.00|
 50.00|       +
 45.00|       +
 40.00|       +
 35.00|       +
 30.00|       +
 25.00|       +
 20.00|       +
 15.00| =     +
 10.00| =  *  +
  5.00| =  *  +
      ----------
        g  c  j
        o  p  a
           p  v
              a
```

```text
      Languages

go   | ===             17.00
cpp  | **              14.95
java | ++++++++++      54.65

```

## Stack chart

```go
package main

import (
	"fmt"

	"github.com/garrou/chart"
)

func main() {
	series := []chart.Series{
		chart.NewColoredSeries("go", "=", 17, chart.Yellow),
		chart.NewColoredSeries("cpp", "*", 14.95, chart.Red),
		chart.NewColoredSeries("java", "+", 54.65, chart.Blue),
	}
	ch := chart.NewStackChart("Languages", 2, series)
	result := ch.Generate(chart.Horizontal)

	fmt.Println(result)
}
```

```text
      Languages

========*******+++++++++++++++++++++++++++
========*******+++++++++++++++++++++++++++
========*******+++++++++++++++++++++++++++

= go : 17.0   * cpp : 14.9   + java : 54.6
```