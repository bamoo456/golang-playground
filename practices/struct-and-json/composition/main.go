package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() string {
	return p.firstName + p.lastName
}

func (p person) greeting() string {
	return fmt.Sprintf("Hi %s", p.fullName())
}

type attendee struct {
	// embed person into attendee struct
	person
	courtesyTitles string
}

// rewrite person greeting
func (atd attendee) greeting() string {
	return fmt.Sprintf("Hi %s %s", atd.courtesyTitles, atd.lastName)
}

func main() {
	george := person{"Geroge", "Chen"}
	attendeeGeorge := attendee{george, "Mrs."}
	// also demo struct pointer
	david := &person{"David", "Lee"}
	fmt.Println("====demo struct pointer====")
	fmt.Printf("type david: %T \n", david)
	fmt.Println("david.fullName(): ", david.fullName())

	fmt.Println("====demo embeded struct and composition====")
	// use firstName and lastName directly from its person struct
	fmt.Println(attendeeGeorge.firstName, attendeeGeorge.lastName)

	fmt.Println("Use person greeting:", george.greeting())
	fmt.Println("Use attendee greeting:", attendeeGeorge.greeting())
	fmt.Println("Use attendee's person greeting:", attendeeGeorge.person.greeting())
}
