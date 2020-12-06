package api

import (
	"github.com/go-gitea/lgtm/cache"
	"github.com/go-gitea/lgtm/model"
	"github.com/go-gitea/lgtm/router/middleware/session"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetTeams gets the list of user teams.
func GetTeams(c *gin.Context) {
	user := session.User(c)
	teams, err := cache.GetTeams(c, user)
	if err != nil {
		logrus.Errorf("Error getting teams for user %s. %s", user.Login, err)
		c.String(500, "Error getting team list")
		return
	}
	teams = append(teams, &model.Team{
		Login:  user.Login,
		Avatar: user.Avatar,
	})
	c.JSON(200, teams)
}
