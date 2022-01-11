package main

// ioutil : File Input Output Utility
import (
	"fmt"
	"io/ioutil"
)

// Use struct      - to stored page data in memory
type Page struct {
	Title string
	Body  []byte // as blob text
}

// Use save method - to stored page data in persistent storage
func (p *Page) save() error {
	fileName := p.Title + ".txt"
	return ioutil.WriteFile(fileName, p.Body, 0600) // args : filename, content, permissions bits
}

// ioutil.WriteFile - write slice of bytes into file - return an error

// Load a specific page from a file
func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	} else {
		return &Page{
			Title: title,
			Body:  body,
		}, nil
	}
}

func main() {
	p1 := &Page{
		Title: "TestPage",
		Body:  []byte("This is a sample page"),
	}

	p1.save()

	p2, err := loadPage("TestPage")
	if err == nil { // Succesfully loaded the page
		fmt.Println(p2.Title)
		fmt.Println(string(p2.Body))
	}
}

// Permission Bits ?
// 1 digit number directory (0-1)
// d : directory
// 3 digit of permissions   (0-7)
// o : owner
// g : group
// a : all others

// Each permission represent of rwx (read, write, and executable)
// r = 4 100
// w = 2 010
// x = 1 001

// For Ex :
// Owner : rwx = 4 + 2 + 1 = 7
// Group : r-x = 4 + 1 = 5
// Other : r-x = 4 + 1 = 5

// Resources : https://wiki.archlinux.org/index.php/File_permissions_and_attributes#Changing_permissions
