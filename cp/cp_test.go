package cp

import (
	"log"
	"math/rand"
	"os"
	"testing"
)

func createTestFile(name string, size uint) *os.File {
	var e [size]
	file, err := os.Create(name)
	rand.Read(e)
	file.Write(e)
	return e
}

func assertFilesEqual(t *testing.T, lhs_path string, rhs string) {
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestCopyFile(t *testing.T) {
	file, err := os.Create("tmpfile")
	if err != nil {
		log.Fatal(err)
	}
	rand.Read()
	CopyFile("tmpfile", "tmpfile.copy", -1, -1)
}
