package profile

import (
	"fmt"
	"net/url"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/user/role"
)

// Identify returns the ID associated with the specified user profile.
func Identify(c pc.PrismaCloudClient, email string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s: %s", singular, email)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, "name")

	var ans []role.NameId
	if _, err := c.Communicate("GET", path, nil, nil, &ans); err != nil {
		return "", err
	}

	for _, o := range ans {
		if o.Name == email {
			return o.Name, nil
		}
	}

	return "", pc.ObjectNotFoundError
}

// List returns the user profiles.
func List(c pc.PrismaCloudClient) ([]UserProfile, error) {
	c.Log(pc.LogAction, "(get) list of %s", plural)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, "v2")
	path = append(path, Suffix...)

	var ans []UserProfile
	if _, err := c.Communicate("GET", path, nil, nil, &ans); err != nil {
		return nil, err
	}

	return ans, nil
}

// Get returns all information about an user profile using its email.
func Get(c pc.PrismaCloudClient, email string) (UserProfile, error) {
	c.Log(pc.LogAction, "(get) %s email:%s", singular, email)

	ans := UserProfile{}

	path := make([]string, 0, len(Suffix)+2)
	path = append(path, "v2")
	path = append(path, Suffix...)
	path = append(path, url.QueryEscape(url.QueryEscape(email)))

	if _, err := c.Communicate("GET", path, nil, nil, &ans); err != nil {
		return ans, err
	}

	return ans, nil
}

// Create makes a new user role on the Prisma Cloud platform.
func Create(c pc.PrismaCloudClient, obj UserProfile) error {
	return createUpdate(false, c, obj)
}

// Update modifies information related to an existing user role.
func Update(c pc.PrismaCloudClient, obj UserProfile) error {
	return createUpdate(true, c, obj)
}

// Delete removes an existing user role using its ID.
func Delete(c pc.PrismaCloudClient, email string) error {
	c.Log(pc.LogAction, "(delete) %s email:%s", singular, email)

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, url.QueryEscape(url.QueryEscape(email)))
	_, err := c.Communicate("DELETE", path, nil, nil, nil)
	return err
}

func createUpdate(exists bool, c pc.PrismaCloudClient, obj UserProfile) error {
	var (
		logMsg strings.Builder
		method string
	)

	logMsg.Grow(30)
	logMsg.WriteString("(")
	if exists {
		logMsg.WriteString("update")
		method = "PUT"
	} else {
		logMsg.WriteString("create")
		method = "POST"
	}
	logMsg.WriteString(") ")

	logMsg.WriteString(" ")
	logMsg.WriteString(singular)
	if exists {
		fmt.Fprintf(&logMsg, ": %s", obj.Email)
	}

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, "v2")
	path = append(path, Suffix...)
	if exists {
		path = append(path, url.QueryEscape(url.QueryEscape(obj.Email)))
	}

	_, err := c.Communicate(method, path, nil, obj, nil)
	return err
}
