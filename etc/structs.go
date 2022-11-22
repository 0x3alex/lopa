package etc

type (
	Spotify struct {
		ID     string
		Secret string
		Token  string
	}
	Bot struct {
		Name          string `json:name`
		Token         string `json:token`
		Prefix        string `json:prefix`
		Status        string `json:status`
		SpotifyID     string `json:spotifyid`
		SpotifySecret string `json:spotifysecret`
		Spotify       *Spotify
	}
	Command struct {
		Name string
		Desc string
	}
	Artist struct {
		Name       string
		Follower   int
		Genres     string
		Image      string
		Albums     string
		TopTracks  string
		Popularity int
	}

	Album struct {
		Release     string
		AlbumName   string
		AristName   string
		Image       string
		TotalTracks int
		Tracks      string
	}
)
