package util_test

import (
	"testing"

	"github.com/alc6/aoc2023/util"
	"github.com/stretchr/testify/assert"
)

func TestReadFileLines(t *testing.T) {
	expectedLines := []string{"line1", "line2", "line3"}

	lines, err := util.ReadFileLines("testdata/read_file_lines.txt")
	assert.NoError(t, err)
	assert.Equal(t, expectedLines, lines)
}