# cli-charts

Generate charts for cli app

```go
import (
	"fmt"

	"github.com/1-irdA/cli-charts/chart"
)

func main() {

	label := []string{"c", "go", "rb", "rs", "cpp", "js", "java"}
	values := []float64{10, 15, 5, 10, 26, 14, 20}

	barChart := chart.New("Languages", 5.0, label, values, "*")
	result := barChart.Generate()

	fmt.Println(result)
}
```