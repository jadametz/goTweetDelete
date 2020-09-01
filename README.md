# goTweetDelete

![](https://github.com/jadametz/goTweetDelete/workflows/Docker%20Image%20CI/badge.svg)

goTweetDelete deletes old Tweets after X (30 by default) days.

## Getting Started

goTweetDelete authorizes as _you_ and therefore requires you to have access to [the Twitter API](https://developer.twitter.com/en/docs/basics/getting-started).

### Configuration

The following environment variables exist for configuration:

|Variable|Description|Required|Default|
|--------|-----------|--------|-------|
|`ACCESSSECRET`|Twitter API credential|Y||
|`ACCESSTOKEN`|Twitter API credential|Y||
|`CONSUMERKEY`|Twitter API credential|Y||
|`CONSUMERSECRET`|Twitter API credential|Y||
|`DAYSTOKEEP`|The number of days to keep Tweets before they're deleted|N|`30`|
|`IGNOREID`|The Tweet ID of a Tweet you'd like to be ignored (e.g. a pinned Tweet)|N||
|`TWEETCOUNT`|This is an API detail - the number of Tweets that'll be retrieved, `3200` is the max|N|`3200`|
|`INCLUDERETWEETS`|Whether RT's should be included in the search/deletion|N|`true`|
|`SCREENNAME`|Your Twitter handle|Y||

## Running the app

### Go

> Note: Golang 1.13 was used for development. No other version has been tested at this time.

```bash
# clone this repository
$ git clone git@github.com:jadametz/goTweetDelete.git
$ cd goTweetDelete

# build the app
$ go build

# export all of the necessary configuration
# e.g.
# export ACCESSSECRET=foo
# export ACCESSTOKEN=bar

# run the app!
$ ./goTweetDelete
```

### Docker

```bash
$ docker run --name gotweetdelete \
  --rm \
  -e ACCESSSECRET=... \
  -e ACCESSTOKEN=... \
  -e CONSUMERKEY=... \
  -e CONSUMERSECRET=... \
  -e SCREENNAME=... \
  jadametz/gotweetdelete
```

#### docker-compose

A `docker-compose.yml` is provided for ease of use and assumes that `.env` exists with the necessary variables.

```bash
$ docker-compose up
```
