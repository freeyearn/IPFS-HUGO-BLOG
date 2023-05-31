package hugo

import (
	"testing"
)

func TestHugo(t *testing.T) {
	h := GetHugo(".", "test")
	h.Version()
	err := h.Init()
	if err != nil {
		return
	}
	//h.Create("test")

	h.Build()

}
