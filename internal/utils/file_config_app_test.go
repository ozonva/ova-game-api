package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenReadJsonFileApp(t *testing.T) {
	assert := assert.New(t)

	readFileFunc := OpenReadJsonFileApp()

	file, err := readFileFunc("./testfiles/app.json")
	assert.Nil(err)
	assert.Equal("appdb", file.Name)
	assert.Equal("local", file.Environment)
	assert.Equal(false, file.Debug)

	file, err = readFileFunc("./testfiles/database.json")
	assert.NotNil(err)
	assert.Equal("", file.Name)
	assert.Equal("", file.Environment)
	assert.Equal(false, file.Debug)

	file, err = readFileFunc("./testfiles/notfoundfiletest.json")
	assert.NotNil(err)
	assert.Equal("", file.Name)
	assert.Equal("", file.Environment)
	assert.Equal(false, file.Debug)
}

func TestOpenFileInCycleByCount(t *testing.T) {
	assert := assert.New(t)

	files, err := OpenFileInCycleByCount("./testfiles/app.json")
	assert.Equal(5, len(files))
	assert.Nil(err)

	files, err = OpenFileInCycleByCount("./testfiles/database.json")
	assert.Equal(0, len(files))
	assert.NotNil(err)

	files, err = OpenFileInCycleByCount("./testfiles/notfoundfiletest.json")
	assert.Equal(0, len(files))
	assert.NotNil(err)
}
