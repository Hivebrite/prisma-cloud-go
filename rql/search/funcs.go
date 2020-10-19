package search

import (
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/rql/history"
)

// Identify returns the ID for the given account group.
func Identify(c pc.PrismaCloudClient, query string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s: %s", singular, query)

	listing, err := history.List(c, history.Recent, 0)
	if err != nil {
		return "", err
	}

	for _, o := range listing {
		if o.Model.Query == query {
			return o.Model.Id, nil
		}
	}

	return "", pc.ObjectNotFoundError
}

func Create(c pc.PrismaCloudClient, query history.Query) error {
	var (
		logMsg strings.Builder
	)

	logMsg.Grow(30)
	logMsg.WriteString("(create)")
	logMsg.WriteString(singular)

	c.Log(pc.LogAction, logMsg.String())

	path := make([]string, 0, len(Suffix)+1)
	path = append(path, Suffix...)
	path = append(path, query.SearchType)

	_, err := c.Communicate("POST", path, nil, query, nil)
	return err
}
