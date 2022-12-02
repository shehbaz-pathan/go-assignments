package main

import "fmt"

type ProfilePicture struct {
	ImageName string
	ImagePath string
}
type Profile struct {
	Name string
	Username string
	Designation string
	ContactNumber string
	profilePicture ProfilePicture
}

func changeName(name string, p *Profile) {
	p.Name = name
}
func changeUsername(username string, p *Profile) {
        p.Username = username
}
func changeDesignation(designation string, p *Profile) {
        p.Designation = designation
}
func changeContactNumber(number string, p *Profile) {
        p.ContactNumber = number
}
func (p *Profile) UpdateProfile(changeType func(string,*Profile), change string) Profile {
	changeType(change,p)
	updated_p:=*p
	return updated_p
}
func main() {
	p1:=Profile{
		Name: "Shehbaz",
		Username: "shehbaz.khan@infracloud.io",
		Designation: "SRE",
		ContactNumber: "9922717346",
		profilePicture: ProfilePicture{
			ImageName: "shehbaz.jpg",
			ImagePath: "/some/path/",
		},
	}
	p2:= Profile{
                Name: "Pratheesh",
                Username: "Pratheesh@infracloud.io",
                Designation: "SRE",
                ContactNumber: "1234567890",
                profilePicture: ProfilePicture{
                        ImageName: "pratheesh.jpg",
                        ImagePath: "/some/path/",
                },
        }
	fmt.Println("Profile 1 Before update:",p1)
	p1=p1.UpdateProfile(changeDesignation,"Site Reliability Engineer")
	fmt.Println("Profile 1 After update:",p1)
	fmt.Println("Profile 2 Before update:",p2)
        p2=p2.UpdateProfile(changeUsername,"pratheesh.p@infracloud.io")
        fmt.Println("Profile 2 After update:",p2)

}



