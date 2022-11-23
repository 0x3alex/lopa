package apis

type (
	//Spotify Structs for the Spotify API requests
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
	// Urban Structs for the Urban API requests
	Urban struct {
		Word       string
		Definition string
		Link       string
		ThumbsUp   int
		ThumbsDown int
		Author     string
		Date       string
		Example    string
	}
)
