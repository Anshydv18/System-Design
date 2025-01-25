package main

import "fmt"

//single responsibility , one struct have multiple responsibility

/*type User struct {
	FirstName string
	LastName  string
	Email     string
}

func (u User) GetFullName() string {
	return u.FirstName + u.LastName
}

func (u User) GetNameWithThePrefix() string {
	return "the " + u.FirstName
}

func (u User) sendEmailMessage() {
	//send email message
}

func (u User) saveToDB() {
	//saving to the database
}

func main() {
	user := User{
		FirstName: "Ansh",
		LastName:  "Yadav",
	}

	fullname := user.GetFullName()
	thePrefix := user.GetNameWithThePrefix()
	fmt.Println(fullname, thePrefix)
}

/* the class or struct have multiple responsbility to  change it
the User have responsibility to savetodb , GetfullName,  the Namewiththeprefix()
Now, let's discuss about the solution about the sperate the class of struct

*/

type User struct {
	FirstName string
	LastName  string
	Email     string
}

type UserRespository struct{}

func (r UserRespository) FullName(user User) string {
	return user.FirstName + user.LastName
}

func (r UserRespository) GetFullNameWithThePrefix(user User) string {
	return "the " + user.FirstName
}

type sendEmail struct{}

func (r sendEmail) sendEmailMessage(user User) string {
	return "sended the mail to" + user.Email
}

func main() {
	userData := User{
		FirstName: "Ansh",
		LastName:  "Yadav",
		Email:     "anshyadav21002@gmail.com",
	}

	userRepo := UserRespository{}
	fullName := userRepo.FullName(userData)
	NamePrefix := userRepo.GetFullNameWithThePrefix(userData)
	fmt.Println(fullName, NamePrefix)

	sendEmail := sendEmail{}
	res := sendEmail.sendEmailMessage(userData)
	fmt.Println(res)
}
