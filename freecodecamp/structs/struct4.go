package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (authI authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s:%s",
		authI.username,
		authI.password,
	)
}

func main() {
	f := authenticationInfo{
		username: "Me",
		password: "Me123",
	}
	fmt.Println(f.getBasicAuth())

}
