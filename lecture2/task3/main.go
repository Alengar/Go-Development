package main

import "goDeveloping/lecture2/task3/people"

func main() {
	friend := people.NewFriend("Nurym")
	colleague := people.NewColleague("Adil")
	senior := people.NewSenior("Gaziz")

	peopleList := []people.Behavior{friend, colleague, senior}

	for _, person := range peopleList {
		person.Act()
	}
}
