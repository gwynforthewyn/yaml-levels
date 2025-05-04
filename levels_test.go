package levels_test

import (
	"levels"
	"testing"
)

func TestReadingOneYamlReturnsOne(t *testing.T) {

	num, err := levels.Level("examples/one.yaml")

	if err != nil {
		t.Error(err)
	}

	if num != 1 {
		t.Errorf("Expected 1 level; actually %d\n", num)
	}

}

func TestReadingTwoYamlReturnsTwo(t *testing.T) {

	num, err := levels.Level("examples/two.yaml")

	if err != nil {
		t.Error(err)
	}

	if num != 2 {
		t.Errorf("Expected 2 levels; actually %d\n", num)
	}

}

func TestReadingTwoInOneSameYamlReturnsTwo(t *testing.T) {

	num, err := levels.Level("examples/two-in-one-same.yaml")

	if err != nil {
		t.Error(err)
	}

	if num != 2 {
		t.Errorf("Expected 2 levels; actually %d\n", num)
	}

}
