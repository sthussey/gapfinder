package main

import (
	"flag"
	"math"
	"fmt"
	"os"
)

func main() {
	var xinit = flag.Float64("maxgap",3.5,"Maximum gap allowed in inches")
	var xdelta = flag.Float64("delta",.125,"Delta for gap trials in inches")
	var target = flag.Float64("width",0,"Width from post to post to the 1/8th inch")
	var bwidth = flag.Float64("bwidth",1.5,"Width of the boundary divisions, in inches")
	var y float64;

	flag.Parse()

	if(*target == 0){
		fmt.Printf("Must specify target width (--width).\n")
		os.Exit(-1)
	}

  for x := *xinit; x > 0; x -= *xdelta {
		y = (*target - x)/(*bwidth + x)
		fmt.Printf("Gap value %.4f requires %.2f segments (%.0f spindles).\n",x, y,round(y))
		var fraction = math.Abs(y - round(y))
		if fraction < 0.01 {
			fmt.Printf("Use %.4f gaps: First mark %.4f, then %.4f center-to-center.\n", x,x+.75,x+*bwidth)
			os.Exit(0)
		}
	  if y > round(y) && (((fraction * (x+*bwidth)/2)+x) <= *xinit) {
			fmt.Printf("Use %.4f gaps: First mark %.4f, then %.4f center-to-center.\n", x,x+.75+((fraction * (x+*bwidth)) / 2),x+*bwidth)
			os.Exit(0)
		}
	}
} 

func round(x float64) float64 {
	if x - math.Trunc(x) > .5 {
    return(math.Floor(x+1))
  } else {
		return(math.Floor(x))
  }
}
