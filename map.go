package main

func displayName(firstName, lastName string) string {
	return firstName + " " + lastName
}

func userResponseFromDBModel(u User) UserResponse {
	var emails []EmailResponse
	for _, e := range u.Emails {
		emails = append(emails, emailReponseFromDBModel(e))
	}

	return UserResponse{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		DisplayName: displayName(u.FirstName, u.LastName),
		Emails:      emails,
	}
}

func emailReponseFromDBModel(e Email) EmailResponse {
	return EmailResponse{
		Address: e.Address,
		Primary: e.Primary,
	}

}

func userDBModelFromCreateRequest(r CreateUserRequest) User {
	return User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Emails: []Email{
			{
				Address: r.Email,
			},
		},
	}
}
