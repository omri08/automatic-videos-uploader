package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"youtube/models"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

const missingClientSecretsMessage = `
Please configure OAuth 2.0
`

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("youtube-go-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open("file")
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}
func checkParams(path, title, description, category, keywords, privacy *string) {
	if *path == "" {
		log.Fatalf("You must provide a path of a video file to upload")
	}
	if *title == "" {
		*title = "Title"
	}
	if *description == "" {
		*description = "description"
	}
	if *category == "" {
		*category = "22"
	}

	if *privacy == "" {
		*privacy = "unlisted"
	}

}
func uploadSingle(service *youtube.Service, vid models.Video) {
	privacy := vid.Privacy.String()
	checkParams(&vid.Path, &vid.Title, &vid.Description, &vid.Category, &vid.Keywords, &privacy)
	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       vid.Title,
			Description: vid.Description,
			CategoryId:  vid.Category,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: privacy},
	}

	// The API returns a 400 Bad Request response if tags is an empty string.
	if strings.Trim(vid.Keywords, "") != "" {
		upload.Snippet.Tags = strings.Split(vid.Keywords, ",")
	}

	part := []string{"snippet,status"}
	call := service.Videos.Insert(part, upload)

	file, err := os.Open(vid.Path)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error opening %v: %v", vid.Path, err)
	}

	fmt.Println("Uploading.... ")
	response, err := call.Media(file).Do()
	handleError(err, "")
	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)

}

// UploadToYoutube uses the youtube api to upload a video to youtube
func UploadToYoutube(list []models.Video) {
	ctx := context.Background()

	b, err := ioutil.ReadFile(`C:\Users\Omri\go\src\youtube\config\client_secret.json`)
	handleError(err, "Unable to read client secret file:")

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/youtube-go-quickstart.json
	config, err := google.ConfigFromJSON(b, youtube.YoutubeUploadScope, youtube.YoutubeScope)
	handleError(err, "Unable to read client secret file:")

	client := getClient(ctx, config)
	service, err := youtube.New(client)
	handleError(err, "Error creating YouTube client")
	for _, vid := range list {
		uploadSingle(service, vid)
	}
}
