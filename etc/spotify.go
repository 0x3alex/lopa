package etc

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

func NewSpotify(id, secret string) *Spotify {
	return &Spotify{
		ID:     id,
		Secret: secret,
	}
}

func (spotify *Spotify) encode() string {
	raw := fmt.Sprintf("%v:%v", spotify.ID, spotify.Secret)
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(raw))
}

func (spotify *Spotify) Auth() bool {

	req := gorequest.New()
	req.Post("https://accounts.spotify.com/api/token")
	req.Set("Authorization", spotify.encode())
	req.Send("grant_type=client_credentials")
	_, body, _ := req.End()

	m := make(map[string]interface{})
	if e := json.Unmarshal([]byte(body), &m); e != nil {
		return false
	}
	val, ok := m["access_token"]
	spotify.Token = val.(string)
	return ok
}

func (spotify *Spotify) artistAlbums(id string) string {
	req := gorequest.New()
	url := fmt.Sprintf(
		"https://api.spotify.com/v1/artists/%s/albums",
		id)
	req.Get(url)
	req.Set("Authorization", "Bearer "+spotify.Token)
	req.Set("Accept", "application/json")
	req.Set("Content-Type", "application/json")
	_, body, _ := req.End()

	var result map[string]any
	err := json.Unmarshal([]byte(body), &result)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	if result["items"] == nil {
		return ""
	}
	var resultAlbum, last string
	items := result["items"].([]interface{})
	for _, value := range items {

		name := value.(map[string]interface{})["name"]
		//the api seems to return some albums twice
		if last == name {
			continue
		}
		last = name.(string)
		release := value.(map[string]interface{})["release_date"]
		resultAlbum += fmt.Sprintf("%s - %s\n",
			name, release)
	}

	return resultAlbum
}

func (spotify *Spotify) artistTopTracks(id string) string {
	return ""
}

func (spotify *Spotify) SearchArtist(artist string) *Artist {
	req := gorequest.New()
	url := fmt.Sprintf(
		"https://api.spotify.com/v1/search?q=%s&type=artist&limit=%d",
		artist, 1)
	req.Get(url)
	req.Set("Authorization", "Bearer "+spotify.Token)
	req.Set("Accept", "application/json")
	req.Set("Content-Type", "application/json")
	_, body, _ := req.End()

	var result map[string]any
	err := json.Unmarshal([]byte(body), &result)
	if err != nil {
		return nil
	}
	if result["artists"] == nil {
		return nil
	}
	artists := result["artists"].(map[string]any)
	var resultArtist Artist
	var genres, id string
	for key, v1 := range artists {
		if key == "items" {
			for _, v2 := range v1.([]interface{}) {
				info := v2.(map[string]interface{})
				id = info["id"].(string)
				//get genres
				for _, v3 := range info["genres"].([]interface{}) {
					genres += "-" + v3.(string) + "\n"
				}
				//get followers
				followers := info["followers"].(map[string]interface{})["total"]
				resultArtist.Follower = int(followers.(float64))
				resultArtist.Genres = genres
				//name
				resultArtist.Name = info["name"].(string)
				//popularity
				resultArtist.Popularity = int(info["popularity"].(float64))
				//get the first image url
				for _, v3 := range info["images"].([]interface{}) {
					resultArtist.Image = v3.(map[string]interface{})["url"].(string)
					break
				}
			}
		}

	}
	resultArtist.Albums = spotify.artistAlbums(id)
	return &resultArtist
}

func (spotify *Spotify) albumTracks(id string) string {
	req := gorequest.New()
	url := fmt.Sprintf(
		"https://api.spotify.com/v1/albums/%s/tracks", id)
	req.Get(url)
	req.Set("Authorization", "Bearer "+spotify.Token)
	req.Set("Accept", "application/json")
	req.Set("Content-Type", "application/json")
	_, body, _ := req.End()

	var result map[string]any
	err := json.Unmarshal([]byte(body), &result)
	if err != nil {
		return ""
	}
	if result["items"] == nil {
		return ""
	}
	tracks := result["items"].([]interface{})

	var resultTracks string
	for _, val := range tracks {
		name := val.(map[string]interface{})["name"].(string)
		trackNumber := int(val.(map[string]interface{})["track_number"].(float64))
		duration := val.(map[string]interface{})["duration_ms"].(float64)
		duration = (duration / 1000) / 60
		resultTracks +=
			fmt.Sprintf("%d) %s - %.2f minutes\n", trackNumber, name, duration)

	}
	return resultTracks
}

func (spotify *Spotify) SearchAlbum(album string, count int) []Album {
	req := gorequest.New()
	url := fmt.Sprintf(
		"https://api.spotify.com/v1/search?q=%s&type=album&limit=%d",
		album, count)
	req.Get(url)
	req.Set("Authorization", "Bearer "+spotify.Token)
	req.Set("Accept", "application/json")
	req.Set("Content-Type", "application/json")
	_, body, _ := req.End()

	var result map[string]any
	err := json.Unmarshal([]byte(body), &result)
	if err != nil {
		return nil
	}
	if result["albums"] == nil {
		return nil
	}
	albums := result["albums"].(map[string]any)
	var albumResults []Album
	var id string
	for key, v1 := range albums {
		//loop over items of albums
		if key == "items" {
			//loop over albums go ger the info
			for _, v2 := range v1.([]interface{}) {
				info := v2.(map[string]interface{})

				var img string
				//get the first image url
				for _, v3 := range info["images"].([]interface{}) {
					img = v3.(map[string]interface{})["url"].(string)
					break
				}
				id = info["id"].(string)
				album := Album{
					AlbumName:   info["name"].(string),
					Release:     info["release_date"].(string),
					TotalTracks: int(info["total_tracks"].(float64)),
					Image:       img,
				}
				//get artists
				for _, v3 := range info["artists"].([]interface{}) {
					//loop, because there can be more than one artist
					for key2, v4 := range v3.(map[string]interface{}) {
						if key2 == "name" {
							album.AristName += v4.(string) + ", "
						}

					}

				}
				album.Tracks = spotify.albumTracks(id)
				albumResults = append(albumResults, album)
			}
		}
	}

	return albumResults
}
