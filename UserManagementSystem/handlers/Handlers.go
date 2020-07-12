package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	. "../dataloaders"
	. "../models"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// page object made of Page Model
	page := Page{
		ID: 7,
		Name: "UserPage",
		Description: "Users List",
		URI: "/users",
	}

	// data loader
	users := LoadUsers()
	interests := LoadInterests()
	interestsMappings := LoadInterestMappings()

	// data
	var newUsers []User
	for _, user := range users {
		for _, interestMapping := range interestsMappings {
			if user.ID == interestMapping.UserID{
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interest = append(user.Interest, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}
	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	w.Write([]byte(data))

}
