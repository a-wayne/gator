# Gator
### A simple RSS blog aggregator written in Go for the CLI

## Requirements
```
go 1.24.1
postgresql 16
```

## Installation
Gator can be installed using **go install**
```
go install github.com/a-wayne/gator
```

## Configuration
Gator expects a configuration file located at **~/.gatorconfig.json** .
You must specify your PostgreSQL connection string with the following format:

**{"db_url":"<CONNECTION_STRING>"}**

Example **~/.gatorconfig.json**
```
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"}
```

## Usage
There are several commands you can use to interact with the program.

### Register
```
gator register <USERNAME>
```
This command will register a new user in the system. When you register a new user, that user is automatically logged in.

### Login
```
gator login <USERNAME>
```
This command is used to login as a user in the system. 

### AddFeed
```
gator addfeed  <FEED_NAME> <FEED_URL>
```
This command is used to add a new RSS feed to be tracked. This new feed will be associated with the current user.

### Agg
```
gator agg <INTERVAL>
```
This command is used to fetch posts from the configured feeds and save them. This command is intended to be left running in the background in order to collect posts.

### Browse
```
gator browse (<NUM_POSTS>)
```
This command is used to browse the posts collected for the current user. There is an optional **NUM_POSTS** argument, which will determine how many posts are returned. If **NUM_POSTS** is left out, the default (2) will be used.

### Follow
```
gator follow <FEED_URL>
```
This command allows the current user to follow a feed added by another user.

### Following
```
gator following
```
This command is used to see what feeds the logged in user is following.
