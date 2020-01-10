package main

import (
	"github.com/jadametz/goTweetDelete/twitter"
)

func main() {
	t := twitter.New()
	t.DeleteOldTweets()
}
