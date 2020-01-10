package twitter

import (
	"fmt"
	"log"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/jadametz/goTweetDelete/config"
)

// Twitter is the all of the data necessary to delete old Tweets
type Twitter struct {
	Client *twitter.Client
	Config *config.Config
}

func (t *Twitter) setConfig(c *config.Config) {
	t.Config = c
}

func (t *Twitter) setClient() {
	oauthConfig := oauth1.NewConfig(t.Config.ConsumerKey, t.Config.ConsumerSecret)
	oauthToken := oauth1.NewToken(t.Config.AccessToken, t.Config.AccessSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, oauthToken)

	t.Client = twitter.NewClient(httpClient)
}

// DeleteOldTweets destroys Tweets older than Config.DaysToKeeps
func (t *Twitter) DeleteOldTweets() error {
	now := time.Now()
	deleted := 0
	skipped := 0

	tweets, _, err := t.Client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		Count:           t.Config.TweetCount,
		ScreenName:      t.Config.ScreenName,
		IncludeRetweets: twitter.Bool(t.Config.IncludeRetweets),
	})
	if err != nil {
		return err
	}

	for _, tweet := range tweets {
		createdAt, err := tweet.CreatedAtTime()
		if err != nil {
			log.Fatal(err.Error())
		}
		daysAgo := now.Sub(createdAt).Hours() / 24
		if int(daysAgo) >= t.Config.DaysToKeep {
			fmt.Printf("Deleting: %s - %s", tweet.CreatedAt, tweet.Text)
			_, _, err := t.Client.Statuses.Destroy(tweet.ID, nil)
			if err != nil {
				return err
			}
			deleted++
		} else {
			skipped++
		}
	}

	fmt.Println("Tweets:")
	fmt.Printf("	Deleted: %d\n", deleted)
	fmt.Printf("	Skipped: %d\n", skipped)
	return nil
}

// New returns a new Twitter struct
func New() *Twitter {
	var t Twitter

	c := config.New()
	t.setConfig(c)
	t.setClient()

	return &t
}
