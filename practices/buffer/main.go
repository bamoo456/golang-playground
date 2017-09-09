package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// basically buffer stores the []byte internally
	buf1 := bytes.NewBufferString("hello")
	buf2 := bytes.NewBuffer([]byte("hello"))

	fmt.Printf("buf1 [%s] should be equal to buf2 [%s]\n", buf1, buf2)
	fmt.Println("buf1 with size=", buf1.Len(), buf1.Cap())

	bufStr := buf1.String()
	fmt.Println("Get the buf1 string=", bufStr)

	bufBytes := buf1.Bytes()
	fmt.Println("Get the buf1 bytes=", bufBytes)

	//
	// Write buffer with bytes
	//
	n, err := buf1.Write([]byte(" World, "))
	fmt.Println("Write buffer bytes result...", n, err)

	//
	// Write buffer with string
	//
	n, err = buf1.WriteString("This is the buffer string")
	fmt.Println("Write buffer string result...", n, err)

	fmt.Println("Append something into current buf1=", buf1)

	var b byte = '!'
	err = buf1.WriteByte(b)
	fmt.Println("Append '!', buf1=", buf1)

	//
	// Write buffer with rune
	//
	var r rune = '是'
	// n is 3, because this chinese world use 3 bytes
	n, err = buf1.WriteRune(r)
	fmt.Println("Write buffer rune result...", n, err)

	fmt.Println("Append rune into current buf1=", buf1)

	//
	// dump the bufer content into file
	//
	file, _ := os.Create("text.txt")
	// write buf1 to the text.txt file
	// NOTE: we can also use fmt.Fprintf to do the same thing "fmt.Fprintf(file, buf1.String())"
	buf1.WriteTo(file)
	file.Close()
	fmt.Println("Check the result after buf1 was dump to text.txt file", buf1.Len(), buf1.Cap())

	//
	// Read the buffer content into anther []byte
	//
	tmp := make([]byte, 3)
	fmt.Printf("Before Read, the buf2=[%s], tmp=[%s]\n", buf2, string(tmp))
	n, err = buf2.Read(tmp)
	fmt.Printf("%d bytes has been retrieved from buf2 with error=%s\n", n, err)
	fmt.Printf("After Read, the buf2=[%s], tmp=[%s]\n", buf2, string(tmp))
	// the tmp will be overwrite after Read again
	n, err = buf2.Read(tmp)
	fmt.Printf("After Read again, the buf2=[%s], tmp=[%s]\n", buf2, string(tmp))

	//
	// Read the single byte from buffer
	//
	buf3 := bytes.NewBufferString("Hello World 3")
	b3, _ := buf3.ReadByte()

	fmt.Printf("After ReadByte, The buf3=[%s], the b3=[%s]\n", buf3, string(b3))

	r3, r3Size, _ := buf3.ReadRune()
	fmt.Printf("After ReadRune, The buf3=[%s], the r3=[%s], the r3 size=[%d]\n", buf3, string(r3), r3Size)

	//
	// Read file content
	//
	fileRead, _ := os.Open("text.txt") //test.txt的内容是“world”

	n3, err := buf3.ReadFrom(fileRead)
	fmt.Printf("Read the file into buf3 done with [%d] bytes readed, and err=%s\n", n3, err)
	fmt.Println("buf3 data now is\n ", buf3.String())

	fmt.Println("buf3's current capacity is :", buf2.Cap())
	buf2.Grow(100)
	fmt.Println("buf3's current capacity after Grow 100 is :", buf2.Cap())

}
