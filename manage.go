package main

import (
	"fmt"
	"strings"
)

func createPupil(text string) string {
	s := strings.Split(text, " ")
	if len(s) < 3 {
		return "Not enough arguments provided. (createpupil <name> <class>)"
	}

	name := strings.ReplaceAll(s[1], "_", " ")
	class := s[2]

	if DB.Get("class:"+class) == "active" {
		DB.Set("pupil:"+name, class)
		return fmt.Sprintf("Created pupil with name '%s' and class '%s'.", name, class)
	}

	return "Class doesn't exist or isn't active."
}

func createClass(text string) string {
	s := strings.Split(text, " ")
	if len(s) < 2 {
		return "Not enough arguments. (createclass <class>)"
	}

	DB.Set("class:"+s[1], "active")

	return fmt.Sprintf("Class set with identifier '%s'.", s[1])
}

func deletePupil(text string) string {
	s := strings.Split(text, " ")
	if len(s) < 2 {
		return "Not enough arguments. (deletepupil <pupil>)"
	}

	DB.Set("pupil:"+s[1], "unactive")

	return fmt.Sprintf("'%s' is now unactive.", s[1])
}

func deleteClass(text string) string {
	s := strings.Split(text, " ")
	if len(s) < 2 {
		return "Not enough arguments. (deleteclass <class>)"
	}

	DB.Set("class:"+s[1], "unactive")

	return fmt.Sprintf("'%s' is now unactive.", s[1])
}

func setAttendance(text string) string {
	s := strings.Split(text, " ")
	if len(s) < 3 {
		return "Not enough arguments. (setattendance <pupil> <attendance>"
	}

	pupil := strings.ReplaceAll(s[1], "_", " ")

	DB.Set("attendance:"+pupil, s[2])
	return fmt.Sprintf("Pupil '%s' now has the attendance of '%s'.", pupil, s[2])
}

func getAttendance(text string) string {
	s := strings.Split(text, " ")
	if len(s) < 2 {
		return "Not enough arguments. (getattendance <pupil>)"
	}

	name := strings.ReplaceAll(s[1], "_", " ")

	attendance := DB.Get("attendance:" + name)
	class := DB.Get("pupil:" + name)
	return fmt.Sprintf("Pupil '%s' (in class '%s') has the attendance of '%s'.", name, class, attendance)
}
