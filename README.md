# EventBot

A bot to manage events via twitter, hosted [@organiserbot](https://twitter.com/organiserbot).

Very much a work in progress.

[![Build Status](https://travis-ci.org/ryankurte/eventbot.svg?branch=master)](https://travis-ci.org/ryankurte/eventbot)

## Configuration
Command line args (see `./eventbot --help`) or via environmental vars (prefix command line args with `EBOT_` ie. `EBOT_TWITTER_USER=boop`)

### Example environmental config

```bash
# EventBot environmental configuration

# Twitter username
export EBOT_TWITTERUSER=""

# Twitter application tokens
export EBOT_TWITTERAPIKEY=""
export EBOT_TWITTERAPISECRET=""

# Access tokens for a given user
export EBOT_TWITTERACCESSTOKEN=""
export EBOT_TWITTERTOKENSECRET=""

# Watson credentials and workspace
export EBOT_WATSONUSER=""
export EBOT_WATSONPASS=""
export EBOT_WATSONWS=""

```