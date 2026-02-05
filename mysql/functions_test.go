package mysql

import (
	"testing"
)

func TestUUIDToBin(t *testing.T) {
	assertSerialize(t, UUID_TO_BIN(String("294b7e43-a21f-4088-8f04-fb8de7770587")),
		`uuid_to_bin(?)`, "294b7e43-a21f-4088-8f04-fb8de7770587")
}
