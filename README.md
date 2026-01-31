# Introduce 
We're going to build an RSS feed aggregator in Go! We'll call it "Gator"
It's a CLI tool that allows users to:
- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

# Learning Goals
Learn how to integrate a Go application with a PostgreSQL database
Practice using your SQL skills to query and migrate a database (using sqlc and goose, two lightweight tools for typesafe SQL in Go)
Learn how to write a long-running service that continuously fetches new posts from RSS feeds and stores them in the database

# Commands:
- gator login - sets the current user in the config
- gator register - adds a new user to the database
- gator users - lists all the users in the database
- gator reset - Report back to the user about whether or not it was successful with an appropriate exit code.
- gator users - list all users
- gator agg - fetch the RSS feed of a website and store its content in a structured format in our database
- gator addfeed - add new feeds to user
- gator feeds - show all feeds on database
- gator follow - follow a feed to a user
- gator following - show all feeds of a user's following
- gator unfollow - unfollow a feed from user
- gator browse - show all posts of a user

# install
## Tech Stack
- Go 1.25+
- PostgreSQL
- github.com/lib/pq
- github.com/google/uuid
- goose 
- sqlc

# Config
We'll use a single JSON file to keep track of two things:
- Who is currently logged in
- The connection credentials for the PostgreSQL database

- Manually create a config file in your home directory, ~/.gatorconfig.json, with the following content:
`json
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
`

# install postgresql
```
sudo apt update
sudo apt install postgresql postgresql-contrib
```

# Goose Migrations
`go install github.com/pressly/goose/v3/cmd/goose@latest`

- Up one migration
`goose postgres <connection_string> up` 

- Down one migration
`goose postgres <connection_string> down`

# Install sqlc
`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

