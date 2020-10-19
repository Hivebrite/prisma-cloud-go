package search

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/rql/history"
)

// Identify returns the ID for the given account group.
func Identify(c pc.PrismaCloudClient, query string) (string, error) {
	c.Log(pc.LogAction, "(get) id for %s: %s", "RQL search not saved", query)

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
