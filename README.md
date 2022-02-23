# cli-charts

Generate barchart for cli app.        
Works with percentage.      

```go
package main

import (
	"fmt"

	"github.com/1-irdA/chart"
)

func main() {
	label := []string{"go", "rb", "rs", "cpp", "js", "java"}
	values := []float64{17, 13, 3.75, 16.25, 14.95, 35.05}

	barChart := chart.New("Languages", 5.0, label, values, "=")
	result := barChart.Generate(chart.Horizontal)

	fmt.Println(result)
}
```