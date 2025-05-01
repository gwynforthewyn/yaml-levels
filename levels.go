package levels

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Stack []any

func (s *Stack) Push(item any) {
	*s = append(*s, item)
}

func (s *Stack) Pop() any {
	// Pop
	n := len(*s)
	// Parens around *s is to enforce precedence of the dereference
	top := (*s)[n-1]
	*s = (*s)[:n-1]

	return top
}

func (s *Stack) Len() int {
	return len(*s)
}

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

	stack := Stack{}

	var iterateOverKeys func(data any)

	iterateOverKeys = func(data any) {
		if res, ok := data.(map[string]any); ok {
			for key, val := range res {
				stack = append(stack, key)

				//type check if the value is itself a map with a value.
				iterateOverKeys(val)
			}
		}
	}

	iterateOverKeys(result)

	return stack.Len(), nil

}
