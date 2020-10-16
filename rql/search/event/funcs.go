package config

import (
	"strings"

	pc "github.com/hivebrite/prisma-cloud-go"
	"github.com/hivebrite/prisma-cloud-go/rql/history"
)

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

	_, err := c.Communicate("POST", path, nil, query, nil)
	return err
}
