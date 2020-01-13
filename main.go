package main

import (
	"github.com/jadametz/goTweetDelete/twitter"
	"os"
	log "github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("ENVIRONMENT") == "production" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetOutput(os.Stdout)
	log.Info("Starting goTweetDelete run")

	t, err := twitter.New()
	if err != nil {
		log.WithError(err).Fatal("Unable to create Twitter struct")
	}

	deleted, skipped, err := t.DeleteOldTweets()
	if err != nil {
		log.WithError(err).Error("Issue deleting Tweets")
	}
	log.WithFields(log.Fields{
		"deleted": deleted,
		"skipped": skipped,
	}).Info("Tweet evaluation complete")

	log.Info("Completed goTweetDelete run")
}
