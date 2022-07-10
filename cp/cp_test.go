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

func assertFilesEqualP(lhs_path string, rhs_path string, start int, end int) bool {
	lhs_bytes, err := os.ReadFile(lhs_path)
	if err != nil {
		log.Fatal(err)
		return false
	}
	rhs_bytes, err := os.ReadFile(rhs_path)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if start <= 0 {
		start = 0
	}
	if end <= 0 {
		end = len(lhs_bytes)
	}
	return bytes.Equal(lhs_bytes[start:end], rhs_bytes)
}

func assertFilesEqual(lhs_path string, rhs_path string) bool {
	return assertFilesEqualP(lhs_path, rhs_path, 0, 0)
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
func TestAssertFilesEqual(t *testing.T) { // comment here
	f1, f2 := "f1.txt", "f2.txt"
	defer cleanUp(f1, f2)
	createTestFile("f1.txt", 100)
	createTestFile("f2.txt", 100)
	if assertFilesEqual("f1.txt", "f2.txt") != false {
		t.Errorf("Files f1 and f2 are equal")
	}
}

func TestCopyFile(t *testing.T) {
	fn := "f1.txt"
	fd := "f2.txt"
	defer cleanUp(fn, fd)
	createTestFile(fn, 100)
	CopyFile(fn, fd, 0, 0)
	if assertFilesEqual(fn, fd) != true {
		t.Errorf("Files f1 and f2 are not equal")
	}
}

func TestCopyOffset(t *testing.T) {
	fn := "f1.txt"
	fd := "f2.txt"
	defer cleanUp(fn, fd)
	createTestFile(fn, 1000)
	CopyFile(fn, fd, 10, 0)
	if assertFilesEqualP(fn, fd, 10, 0) != true {
		t.Errorf("Files f1 and f2 are not equal")
	}
}

func TestCopyLimit(t *testing.T) {
	fn := "f1.txt"
	fd := "f2.txt"
	defer cleanUp(fn, fd)
	createTestFile(fn, 1000)
	CopyFile(fn, fd, 0, 10)
	if assertFilesEqualP(fn, fd, 0, 10) != true {
		t.Errorf("Files f1 and f2 are not equal")
	}
}

func TestCopyOffsetAndLimit(t *testing.T) {
	fn := "f1.txt"
	fd := "f2.txt"
	defer cleanUp(fn, fd)
	createTestFile(fn, 1000)
	CopyFile(fn, fd, 10, 200)
	if assertFilesEqualP(fn, fd, 10, 210) != true {
		t.Errorf("Files f1 and f2 are not equal")
	}
}
