package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" file content: %s", data[:count])

}

/*package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	createfile()
	writefile()
}

func createfile() {
	newfile, err := os.Create("textfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newfile.Close()
	//os.Remove("textfile.txt")// to remove file from current directory
	//	log.Println(newfile)
}

func writefile() {
	err := ioutil.WriteFile("textfile.txt", []byte("Dumping bytes to a file\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}*/

/*package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	fptr := flag.String("fpath", "textfile.txt", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
*/
