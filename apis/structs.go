package apis

type (
	Spotify struct {
		ID     string
		Secret string
		Token  string
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
