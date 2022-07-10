package cp

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"testing"
)

func createTestFile(name string, size uint) {
	var bytes_slice []byte = make([]byte, size)
	file, _ := os.Create(name)
	rand.Read(bytes_slice)
	file.Write(bytes_slice)
	file.Close()
}

func assertFilesEqualP(t *testing.T, lhs_path string, rhs_path string, offset int, limit int) (bool, error) {
	lhs_bytes, err := os.ReadFile(lhs_path)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	rhs_bytes, err := os.ReadFile(rhs_path)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	if limit <= 0 {
		limit = len(lhs_bytes)
	}
	if offset <= 0 {
		offset = 0
	}
	return bytes.Equal(lhs_bytes[offset:limit], rhs_bytes), nil
}

func assertFilesEqual(t *testing.T, lhs_path string, rhs_path string) (bool, error) {
	return assertFilesEqualP(t, lhs_path, rhs_path, 0, 0)
}

func cleanUp(file_names ...string) {
	for _, fn := range file_names {
		err := os.Remove(fn)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestAssertFilesEqual(t *testing.T) {
	f1, f2 := "f1.txt", "f2.txt"
	defer cleanUp(f1, f2)
	createTestFile("f1.txt", 100)
	createTestFile("f2.txt", 100)
	assertFilesEqual(t, "f1.txt", "f2.txt")
}

func TestCopyFile(t *testing.T) {
	fn := "f1.txt"
	fd := "f2.txt"
	defer cleanUp(fn, fd)
	createTestFile(fn, 100)
	CopyFile(fn, fd, 0, 0)
	assertFilesEqual(t, fn, fd)
}

func TestCopyOffset(t *testing.T) {
	fn := "f1.txt"
	fd := "f2.txt"
	defer cleanUp(fn, fd)
	createTestFile(fn, 1000)
	CopyFile(fn, fd, 10, 0)
	assertFilesEqualP(t, fn, fd, 10, 0)
}

func TestCopyLimit(t *testing.T) {
	fn := "f1.txt"
	fd := "f2.txt"
	defer cleanUp(fn, fd)
	createTestFile(fn, 1000)
	CopyFile(fn, fd, 10, 0)
	assertFilesEqualP(t, fn, fd, 0, 10)
}

func TestCopyOffsetAndLimit(t *testing.T) {
	fn := "f1.txt"
	fd := "f2.txt"
	defer cleanUp(fn, fd)
	createTestFile(fn, 1000)
	CopyFile(fn, fd, 0, 10)
	assertFilesEqualP(t, fn, fd, 10, 10)
}
