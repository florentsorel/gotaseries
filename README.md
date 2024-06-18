# Gotaseries

Gotaseries is a Go client for [Betaseries API](https://developers.betaseries.com/docs/making-requests).

## Installation

```bash
go get github.com/florentsorel/gotaseries
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/florentsorel/gotaseries"
)

func main() {
	betaseries := gotaseries.NewClient("YOUR_API_KEY")
	// You can set locale for each request globally instead of passing it to each params
	betaseries.Locale = gotaseries.LocaleEN
	// If request need authentication, you must set your token
	betaseries.Token = "YOUR_TOKEN"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	show, err := betaseries.Shows.Display(ctx, gotaseries.ShowsDisplayParams{
		ID:     gotaseries.Int(1161),
		Locale: gotaseries.Locale(gotaseries.LocaleFR),
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", show)

	date, _ := time.Parse("2006-01-02", "2024-04-03")
	shows, err := betaseries.Shows.List(context.Background(), gotaseries.ShowsListParams{
		Order: gotaseries.Order(gotaseries.OrderPopularity),
		Since: gotaseries.Time(date),
	})
	if err != nil {
		log.Fatalln(err)
	}

	for _, s := range shows {
		fmt.Println(s.Title)
	}
}
```

## Endpoints
<details>
  <summary>Badges</summary>

  - [ ] Returns badge details (GET /badges/badge)
</details>

<details>
  <summary>Collections</summary>

  - [ ] Display collection's data (GET /collections/collection) 
  - [ ] Delete a collection from the identified user [Premium feature] (DELETE /collections/collection) - Token
  - [ ] Create/Update a collection for the identified user [Premium feature] (POST /collections/collection) - Token
  - [ ] Display the list of all collections of the member (GET /collections/list) - Token or ID parameter
</details>

<details>
  <summary>Comments</summary>

  - [ ] Get comments (GET /comments/comments)
  - [ ] Retrieve a given comment (GET /comments/comment)
  - [ ] Create or edit a comment for the specified item (POST /comments/comment) - Token
  - [ ] Delete a comment from the identified user (DELETE /comments/comment) - Token
  - [ ] Retrieve the replies of a given comment (GET /comments/replies)
  - [ ] Create a comment for an event (POST /comments/comment_event) - Token
  - [ ] Subscribe the member to email notifications for the given item (POST /comments/subscription) - Token
  - [ ] Unsubscribe the member from email notifications for the given item (DELETE /comments/subscription) - Token
  - [ ] Add a vote for the user for the given comment (POST /comments/thumb) - Token
  - [ ] Remove the user's vote for the given comment (DELETE /comments/thumb) - Token
  - [ ] Retrieve the status of comments on the given item (closed or open) (GET /comments/status)
</details>
<details>
  <summary>Friends</summary>

  - [ ] Rate an episode (POST /episodes/note) - Token
  - [ ] Remove a rating (DELETE /episodes/note) - Token
  - [ ] Retrieve the list of episodes to watch. (GET /episodes/list) - Token
  - [ ] Mark an episode as downloaded (POST /episodes/downloaded) - Token
  - [ ] Remove the downloaded mark (DELETE /episodes/downloaded) - Token
  - [ ] Mark an episode as watched (POST /episodes/watched) - Token
  - [ ] Unmark an episode as watched (DELETE /episodes/watched) - Token
  - [ ] Display information of an episode (GET /episodes/display)
  - [ ] Retrieve episode information (GET /episodes/scraper) - Token
  - [ ] Retrieve episode information (GET /episodes/search) - Token
  - [ ] Retrieve the latest aired episode (GET /episodes/latest) - Token
  - [ ] Retrieve the next episode (GET /episodes/next) - Token
  - [ ] Mark an episode as not to watch (POST /episodes/hidden) - Token
  - [ ] Remove an episode from the hidden list (DELETE /episodes/hidden) - Token
  - [ ] Retrieve the list of watched and unrated episodes (GET /episodes/unrated) - Token
</details>
<details>
  <summary>Friends</summary>

  - [ ] Retrieves friends List (GET /friends/list) - Token or ID parameter
  - [ ] Retrieves sent requests (GET /friends/sent) - Token
  - [ ] Adds a friend (POST /friends/friend) - Token
  - [ ] Removes a friend (DELETE /friends/friend) - Token
  - [ ] Blocks a user (POST /friends/block) - Token
  - [ ] Unblocks a user (DELETE /friends/block) - Token
</details>
<details>
<summary>Members</summary>

  - [ ] Deletes filter (DELETE /profile-filters/filter) - Token
  - [ ] Retrieves member options (GET /members/options) - Token
  - [ ] Standard member authentication (POST /members/auth)
  - [ ] OAuth Authentication (POST /members/oauth)
  - [ ] OAuth2 Access Token (POST /members/access_token)
  - [ ] Returns member information (GET /members/infos)
  - [ ] Returns available usernames (GET /members/username)
  - [ ] Modifies user option (POST /members/option) - Token
  - [ ] Checks token activity (GET /members/is_active) - Token
  - [ ] Destroys active token (DELETE /members/destroy) - Token
  - [ ] Displays member badges (GET /members/badges)
  - [ ] Displays latest notifications (GET /members/notifications) - Token
  - [ ] Deletes a notification (DELETE /members/notification) - Token
  - [ ] Creates new member account (POST /members/signup)
  - [ ] Member search (GET /members/search)
  - [ ] Searches members among friends (GET /members/sync) - Token
  - [ ] Password reset email (POST /members/lost)
  - [ ] Uploads and replaces user avatar (POST /members/avatar) - Token
  - [ ] Deletes user avatar (DELETE /members/avatar) - Token
  - [ ] Uploads user banner (POST /members/banner) - Token
  - [ ] Remove the banner (DELETE /members/banner) - Token
  - [ ] Change the locale (POST /members/locale) - Token
  - [ ] Retrieve the email address (GET /members/email) - Token
  - [ ] Change the email address (POST /members/email) - Token
  - [ ] Change the password (POST /members/password) - Token
  - [ ] Returns yearly member statistics (GET /members/year)
  - [ ] Initiates account deletion process (POST /members/delete) - Token
</details>
<details>
  <summary>Messages</summary>

  - [ ] Retrieve the member's inbox (GET /messages/inbox) - Token
  - [ ] Retrieve a discussion (GET /messages/discussion) - Token
  - [ ] Mark a message as read (POST /messages/read) - Token
  - [ ] Delete a message (DELETE /messages/message) - Token
  - [ ] Send a message (POST /messages/message) - Token
</details>
<details>
  <summary>Movies</summary>

  - [ ] Show movie details (GET /movies/movie)
  - [ ] Add or update a movie (POST /movies/movie) - Token
  - [ ] Remove a movie (DELETE /movies/movie) - Token
  - [ ] Display the list of all movies (GET /movies/list)
  - [ ] Display all movies of a member (GET /movies/member) - Token or ID parameter
  - [ ] Display a random movie (GET /movies/random)
  - [ ] Search for a movie (GET /movies/search)
  - [ ] Retrieve movie information (GET /movies/scraper) - Token
  - [ ] Display all available genres (GET /movies/genres)
  - [ ] Rate a movie (POST /movies/note) - Token
  - [ ] Remove a movie rating (DELETE /movies/note) - Token
  - [ ] Retrieve similar movies (GET /movies/similars)
  - [ ] Retrieve the cast of the movie. (GET /movies/characters)
  - [ ] Retrieve favorite movies (GET /movies/favorites) - Token or ID parameter
  - [ ] Add a favorite movie (POST /movies/favorite) - Token
  - [ ] Remove a favorite movie (DELETE /movies/favorite) - Token
  - [ ] Display upcoming movies (GET /movies/upcoming)
  - [ ] Display movies to discover (GET /movies/discover)
  - [ ] Display blog articles about the movie (GET /movies/articles)
</details>
<details>
  <summary>News</summary>

  - [ ] Display the latest news (GET /news/last)
</details>
<details>
  <summary>Authentication</summary>

  - [ ] Retrieve an access token with the code provided by OAuth 2 authentication (POST /oauth2/access_token)
  - [ ] Retrieve a code to present to the user for identification on a remote device (e.g., Television) (POST /oauth2/device)
</details>
<details>
  <summary>Persons</summary>

  - [ ] Display details of the actor (GET /persons/person)
  - [ ] Display news articles (GET /persons/articles)
</details>
<details>
  <summary>Pictures</summary>

  - [ ] Return a picture of the member (GET /pictures/picture)
  - [ ] Return a picture of the episode (GET /pictures/episodes)
  - [ ] Return a picture of the show (GET /pictures/shows)
  - [ ] Return an image of the badge (32x32) (GET /pictures/badges)
  - [ ] Return an image of the character (GET /pictures/characters)
  - [ ] Return an image of the person (GET /pictures/persons)
  - [ ] Return an image of the movie (GET /pictures/movies)
  - [ ] Return an image of the show's season (GET /pictures/seasons)
  - [ ] Return an image of the SVOD or VOD platform (GET /pictures/platforms)
</details>
<details>
  <summary>Planning</summary>

  - [ ] Display all episodes broadcasted (GET /planning/general)
  - [ ] Display the schedule (GET /planning/member) - Token or ID parameter
  - [ ] Display only the first episode of the upcoming series (GET /planning/incoming)
</details>
<details>
  <summary>Platforms</summary>

  - [ ] Display the SVOD and VOD platforms available in the country (GET /platforms/list)
  - [ ] Display the different services a user can have (GET /platforms/services) - Token or ID parameter
  - [ ] Add the service to the user's subscriptions (POST /platforms/service) - Token
  - [ ] Remove the service from the user's subscriptions (DELETE /platforms/service) - Token
</details>
<details>
  <summary>Polls</summary>

  - [ ] Display the latest active poll (GET /polls/last)
  - [ ] Display the details of a poll (GET /polls/poll)
  - [ ] Display the latest active poll (GET /polls/target)
  - [ ] Display all polls (GET /polls/list)
  - [ ] Send a response to a poll (POST /polls/answer)
</details>
<details>
  <summary>Reports</summary>
  
  - [ ] Create a report for the element (POST /reports/report) - Token
  - [ ] Request an update for the element (POST /reports/update) - Token
</details>
<details>
  <summary>Search</summary>

  - [ ] Return search results for all types of elements. (GET /search/all)
  - [ ] Return series search results with advanced filters. (GET /search/shows)
  - [ ] Return movie search results with advanced filters. (GET /search/movies)
</details>
<details>
  <summary>Seasons</summary>

  - [ ] Mark all episodes of a season as watched (POST /seasons/watched) - Token
  - [ ] Remove all episodes of a season from watched (DELETE /seasons/watched) - Token
  - [ ] Mark all episodes of a season as hidden (POST /seasons/hide) - Token
  - [ ] Remove all episodes of a season from hidden (DELETE /seasons/hide) - Token
  - [ ] Rate a season (POST /seasons/note) - Token
  - [ ] Remove a rating from a season (DELETE /seasons/note) - Token
</details>
<details>
  <summary>Shows</summary>

  - [x] Rate a series (POST /shows/note) - Token
  - [x] Delete a series rating (DELETE /shows/note) - Token
  - [x] Search for a series, with member information if a token is provided (GET shows/search)
  - [x] Display information about a series (GET /shows/display)
  - [x] Display the list of all series (GET /shows/list)
  - [x] Display a random series (GET /shows/random)
  - [x] Display episodes of a series (GET /shows/episodes)
  - [x] Add a series to the member's account (POST /shows/show) - Token
  - [x] Remove a series from the member's account (DELETE /shows/show) - Token
  - [x] Archive a series in the member's account (POST /shows/archive) - Token
  - [x] Remove a series from the archives of the member's account (DELETE /shows/archive) - Token
  - [x] Create a series recommendation from a member to a friend (POST /shows/recommendation) - Token
  - [x] Delete a sent or received series recommendation (DELETE /shows/recommendation) - Token
  - [x] Change the status of a received series recommendation (PUT /shows/recommendation) - Token
  - [x] Retrieve recommendations received by the identified user (GET /shows/recommendation) - Token
  - [x] Retrieve series marked as similar (GET /shows/similars)
  - [x] Retrieve videos associated with the series (GET /shows/videos)
  - [x] Retrieve characters of the series (GET /shows/characters)
  - [x] Retrieve images of the series (GET /shows/pictures)
  - [x] Retrieve the favorite series of the member (GET /shows/favorites) - Token or ID parameter
  - [x] Add a favorite series to the profile of the identified member (POST /shows/favorite) - Token
  - [x] Remove a favorite series from the profile of the identified member (DELETE /shows/favorite) - Token
  - [x] Update tags for the given series of the identified member (POST /shows/tags) - Token
  - [x] Display the list of all series of the member with tags (GET /shows/member) - Token or ID parameter
  - [x] Display the list of series to discover (GET /shows/discover)
  - [x] Display the list of series to discover on major SVoD platforms (GET /shows/discover_platforms)
  - [x] Display the list of available series genres (GET /shows/genres)
  - [x] Display the seasons of the series (GET /shows/seasons)
  - [ ] Display blog articles that talk about the series (GET /shows/articles)
  - [ ] Retrieve the list of finished and unrated series (GET /shows/unrated) - Token
</details>
<details>
  <summary>Subtitles</summary>

  - [ ] Display the latest subtitles retrieved by BetaSeries (GET /subtitles/last)
  - [ ] Display subtitles for a given show (GET /subtitles/show)
  - [ ] Display subtitles for a given episode (GET /subtitles/episode)
  - [ ] Display subtitles for a season or all seasons (GET /subtitles/season)
  - [ ] Reports subtitles as incorrect to be removed from the list. (POST /subtitles/report) - Token
</details>
<details>
  <summary>Tags</summary>

  - [ ] Display all tags created by the connected member (GET /tags/list) - Token
  - [ ] Add a tag (or several) for the show (or movie) for the connected member (POST /tags/show) - Token
  - [ ] Remove a tag for the show (or movie) for the connected member (DELETE /tags/show) - Token
</details>
<details>
  <summary>Timeline</summary>

  - [ ] Display the latest events on the site (GET /timeline/home)
  - [ ] Display the latest events of the friends of the identified member (GET /timeline/feed) - Token
  - [ ] Display the latest events of the friends of the identified member (GET /timeline/friends) - Token
  - [ ] Display the latest events of the specified member (GET /timeline/member)
  - [ ] Display a particular event (GET /timeline/event)
  - [ ] Display the latest events of the connected member about the specified show (GET /timeline/show) - Token
</details>
