package main

import "fmt"

type ProfileMaker interface {
	UpdateProfilePicture(ProfilePicture)
	CheckDuplicateProfile(Profile) bool
}
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
func (p *Profile) UpdateProfilePicture(pf ProfilePicture) {
	p.profilePicture.ImageName = pf.ImageName
	p.profilePicture.ImagePath = pf.ImagePath
}
func (p Profile) CheckDuplicateProfile(p1 Profile) bool {
	if p == p1 {
		return true
	} else {
		return false
	}
}
func main() {
	pf1:= ProfilePicture{
		ImageName: "shehbaz.jpg",
		ImagePath: "/some/path/",
	}
	pf2:= ProfilePicture{
		ImageName: "shehbaz.jpg",
		ImagePath: "/some/path/",
	}
	p1:=Profile{
		Name: "Shehbaz",
		Username: "shehbaz.khan@infracloud.io",
		Designation: "SRE",
		ContactNumber: "9922717346",
		profilePicture: ProfilePicture{},
	}
	p2:= Profile{
                Name: "Shehbaz",
                Username: "shehbaz.khan@infracloud.io",
                Designation: "SRE",
                ContactNumber: "9922717346",
                profilePicture: ProfilePicture{},
        }
	var pm1 ProfileMaker = &p1
	var pm2 ProfileMaker = &p2
	fmt.Println("Slack profile 1 before updating profile picture",p1)
	pm1.UpdateProfilePicture(pf1)
	fmt.Println("Slack profile 1 after updating profile picture",p1)
	fmt.Println("Slack profile 2 before updating profile picture",p2)
        pm2.UpdateProfilePicture(pf2)
        fmt.Println("Slack profile 2 after updating profile picture",p2)
	fmt.Println("=====================================================================")
	if pm1.CheckDuplicateProfile(p2) {
		fmt.Println("profile 1 and 2 is same")
	} else {
		fmt.Println("profile 1 and 2 is not same")
	}
}



