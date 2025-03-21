package looplevel

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := os.Args[1]
	files, _ := os.ReadDir(dir)
	sha := sha256.New()
	var prev, next chan int

	for _, file := range files {
		if !file.IsDir() {
			next = make(chan int)

			go func(filename string, prev, next chan int) {
				fpath := filepath.Join(dir, filename)
				hashOnFile := FHash(fpath)
				if prev != nil {
					<-prev
				}
				sha.Write(hashOnFile)
				next <- 0
			}(file.Name(), prev, next)

			prev = next
		}
	}

	<-next

	fmt.Printf("%x\n", sha.Sum(nil))
}
