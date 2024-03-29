package resdata

import (
	// "fmt"
	"slices"
	"strings"
)

func TimeZoneList(timezone []string) {
	var time_zone_list []string
	delimiter := "/"
	map_timezone := make(map[string][]string)

	for _, i := range timezone{
		// Condition for elements with "/"
		if strings.Contains(i, delimiter){
			// fmt.Println(i)

			//Spliting the element
			timezone_delimited := strings.Split(i, delimiter)

			// Added timezone and area to a map(key:pair)
			map_timezone[timezone_delimited[0]] = append(map_timezone[timezone_delimited[0]], timezone_delimited[1])

			// Adding timezone to a list
			if !slices.Contains(time_zone_list, timezone_delimited[0]){
				time_zone_list = append(time_zone_list, timezone_delimited[0])
			}

		}
	}
	// fmt.Println(time_zone_list)
	// fmt.Println(map_timezone)
}