package resdata

import (
	"fmt"
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

			// Add timezone to a list if it doesn't already have it
			if !slices.Contains(time_zone_list, timezone_delimited[0]){
				time_zone_list = append(time_zone_list, timezone_delimited[0])
			}
			continue
		} else { // Add those entries that doesn't have a delimier "/"
			// Append it on the map
			map_timezone[i] = append(map_timezone[i], "")
			// Add timezone to a list if it doesn't already have it
			if !slices.Contains(time_zone_list, i){
				time_zone_list = append(time_zone_list, i)
			}
		}
	}
	fmt.Println(time_zone_list)
	fmt.Println(map_timezone)
}