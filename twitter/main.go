package twitter

import (
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

func (t *Twitter) createClient() {
	oauthConfig := oauth1.NewConfig(t.Config.ConsumerKey, t.Config.ConsumerSecret)
	oauthToken := oauth1.NewToken(t.Config.AccessToken, t.Config.AccessSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, oauthToken)

	t.Client = twitter.NewClient(httpClient)
}

// DeleteOldTweets destroys Tweets older than Config.DaysToKeeps
func (t *Twitter) DeleteOldTweets() (deleted, skipped int, err error) {
	now := time.Now()
	deleted = 0
	skipped = 0

	tweets, _, err := t.Client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		Count:           t.Config.TweetCount,
		ScreenName:      t.Config.ScreenName,
		IncludeRetweets: twitter.Bool(t.Config.IncludeRetweets),
	})
	if err != nil {
		return deleted, skipped, err
	}

	for _, tweet := range tweets {
		createdAt, err := tweet.CreatedAtTime()
		if err != nil {
			return deleted, skipped, err
		}
		daysAgo := now.Sub(createdAt).Hours() / 24
		if int(daysAgo) >= t.Config.DaysToKeep {
			_, _, err := t.Client.Statuses.Destroy(tweet.ID, nil)
			if err != nil {
				return deleted, skipped, err
			}
			deleted++
		} else {
			skipped++
		}
	}

	return deleted, skipped, nil
}

// New returns a new Twitter struct
func New() (*Twitter, error) {
	var t Twitter

	c, err := config.New()
	if err != nil {
		return nil, err
	}

	t.setConfig(c)
	t.createClient()

	return &t, nil
}
