package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenReadJsonFile(t *testing.T) {
	assert := assert.New(t)

	readFileFunc := OpenReadJsonFile()

	file, err := readFileFunc("./../../config/app.json")
	assert.Nil(err)
	assert.Equal("apptest", file.Name)
	assert.Equal("local", file.Enviroment)
	assert.Equal(false, file.Debug)

	file, err = readFileFunc("./../../config/db.json")
	assert.NotNil(err)
	assert.Equal("", file.Name)
	assert.Equal("", file.Enviroment)
	assert.Equal(false, file.Debug)

	file, err = readFileFunc("./../../config/notfoundfiletest.json")
	assert.NotNil(err)
	assert.Equal("", file.Name)
	assert.Equal("", file.Enviroment)
	assert.Equal(false, file.Debug)
}

func TestOpenFileInCycleByCount(t *testing.T) {
	assert := assert.New(t)

	files, err := OpenFileInCycleByCount("./../../config/app.json")
	assert.Equal(5, len(files))
	assert.Nil(err)

	files, err = OpenFileInCycleByCount("./../../config/db.json")
	assert.Equal(0, len(files))
	assert.NotNil(err)

	files, err = OpenFileInCycleByCount("./../../config/notfoundfiletest.json")
	assert.Equal(0, len(files))
	assert.NotNil(err)
}
