package main

import "fmt"

//  the visitor design pattern is a way of separating an algorithm from an object structure on which it operates.
// A practical result of this separation is the ability to add new operations to existent object structures without modifying the structures.
// It is one way to follow the open/closed principle.

// the main object struct should implement this interface
type SystemAcceptor interface {
	Accept(SystemVisitor) error
}

// the behavior implement in visitor
type SystemVisitor interface {
	VisitFile(*file) error
	VisitDirectory(*directory) error
}

type SystemItem interface {
	Entry() string
	Size() int
}

// file object structure
type file struct {
	entry string
	size  int
}

func (f *file) Entry() string {
	return f.entry
}

func (f *file) Size() int {
	return f.size
}

func (f *file) Accept(v SystemVisitor) error {
	return v.VisitFile(f)
}

// directory object structure
type directory struct {
	entry string
	// either file or diectory
	items []SystemItem
}

func (d *directory) Entry() string {
	return d.entry
}

func (d *directory) Size() int {
	totalSize := 0
	for _, f := range d.items {
		totalSize += f.Size()
	}
	return totalSize
}

func (d *directory) Accept(v SystemVisitor) error {
	return v.VisitDirectory(d)
}

// Implement the visitor interface (implement the behavior)
type ListVisitor struct{}

func (v ListVisitor) VisitFile(f *file) error {
	fmt.Printf("\tfile:%s, size:%d\n", f.Entry(), f.Size())
	return nil
}

func (v ListVisitor) VisitDirectory(d *directory) error {
	fmt.Printf("directory:[%s], size:[%d]\n", d.Entry(), d.Size())

	for _, item := range d.items {
		switch item.(type) {
		case *file:
			v.VisitFile(item.(*file))
		case *directory:
			v.VisitDirectory(item.(*directory))
		default:
			fmt.Println("Not supported item type")
			continue
		}
	}

	return nil
}

func main() {
	var rootDirectory SystemAcceptor = &directory{
		"/",
		[]SystemItem{
			&directory{"/tmp", []SystemItem{
				&file{"temp", 20},
				&file{"log.txt", 1024},
			}},
			&directory{"/etc", []SystemItem{
				&directory{"/redis", []SystemItem{
					&file{"/redis.conf", 688},
				}},
				&directory{"/system", []SystemItem{
					&file{"init.d", 2048},
				}},
			}},
		},
	}

	var visitor SystemVisitor

	// configure the visitor as list visitor
	visitor = ListVisitor{}

	rootDirectory.Accept(visitor)
}
