package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Contains CS 677 student information.
type student struct {
	// Username used to sign into USF services
	usfUsername string

	// What I should call you -- help me with pronunciation if needed! :-)
	preferredName string

	// she/her, he/his, they/their, ze/zir, etc.
	pronouns string

	// Detailed information about you
	bio string
}

func main() {
	// Create "you" -- a student :-)
	you := student{}

	/* --- The following information is REQUIRED. --- */

	you.usfUsername = ""

	/* --- The following information is optional. --- */

	you.preferredName = ""
	you.pronouns = ""

	you.bio = `
	Tell me something about yourself (if you want), and anything I can do
	to help you succeed this semester. This can include things that help
	me accommodate your learning better (e.g., 'I commute from Sacramento
	every day in a horse-drawn carriage, so sometimes I might be a bit
	late') or if you'd like to stop by and chat about something feel free
	to let me know here.

	If you aren't currently registered for the class, please provide
	details.

	If you aren't going to write anything here, set 'you.bio' to be a
	blank string or delete this declaration entirely.
	`

	err := printStudent(&you)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Prints a CS 677 student's information (if filled out)
// Returns an error if the student's username is not set, which is the minimum
// required information.
func printStudent(s *student) error {
	if s.usfUsername == "" {
		return errors.New("No username specified!")
	}

	fmt.Printf("Username:       %s\n", s.usfUsername)

	if s.preferredName != "" {
		fmt.Printf("Preferred Name: %s\n", s.preferredName)
	}

	if s.pronouns != "" {
		fmt.Printf("Pronouns:       %s\n", s.pronouns)
	}

	if s.bio != "" {
		/* Ignore 'about' info if the first few words haven't been changed: */
		if strings.HasPrefix(strings.TrimSpace(s.bio), "Tell me something") == false {
			fmt.Printf("-------------------------------------\n")
			fmt.Printf("About me: %s\n", s.bio)
		}
	}

	return nil // This means no error occurred
}
