package cron

import "github.com/robfig/cron"

func Start() {
	c := cron.New()

	// addNews()
	// deleteNews()

	c.AddFunc("@hourly", addNews)
	c.AddFunc("@midnight", deleteNews)

	c.Start()
}
