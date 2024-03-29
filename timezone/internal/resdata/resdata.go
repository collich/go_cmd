package resdata

import (
	"fmt"
	"strings"
)

func TimeZoneList(timezone []string) {
	for _, i := range timezone{
		if strings.Contains(i, "/"){
			fmt.Println(i)
		}
	}
}