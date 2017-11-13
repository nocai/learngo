package cartesian

import (
	"runtime"
	"fmt"
)

type polar struct {
	radius float64
	o      float64
}

type cartesian struct {
	x float64
	y float64
}


var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to qiut"

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl + Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl + D")
	}
}
