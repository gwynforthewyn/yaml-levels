package levels

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Level(path string) (int, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return 0, err
	}

	var result map[string]any
	err = yaml.Unmarshal(data, &result)
	if err != nil {
		return 0, err
	}

	fmt.Printf("%s\n", result)

	maxDepth := 0

	var measureKeyDepth func(data any, depth int)
	measureKeyDepth = func(data any, depth int) {
		if maxDepth < depth {
			maxDepth = depth
		}

		// you aren't ranging over data, which is the map that has to be iterated over
		//type check for if there's a map as the value

		if datamap, ok := data.(map[string]any); ok {
			for _, child := range datamap {
				if childIsMap, ok := child.(map[string]any); ok {
					fmt.Printf("ChildIsMap %s", childIsMap)
					measureKeyDepth(childIsMap, depth+1)
				}
			}
		}

	}
	measureKeyDepth(result, 1)
	return maxDepth, nil

}
