package cron

import (
	"log"

	"github.com/robfig/cron"
)

func Start() {
	c := cron.New()

	addNews()
	deleteNews()

	if err := c.AddFunc("@hourly", addNews); err != nil {
		log.Println(err)
	}
	if err := c.AddFunc("@midnight", deleteNews); err != nil {
		log.Println(err)
	}

	c.Start()
}
