package utils

import "testing"

func TestCreateRandomId(t *testing.T) {
	t.Log(CreateRandomId(4))
}

func TestCreateFileUploadDir(t *testing.T) {
	t.Log(CreateFileUploadDir("resources/content/post"))
}

func TestGetCurrentDay(t *testing.T) {
	t.Log(GetCurrentDay())
}
