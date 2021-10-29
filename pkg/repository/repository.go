package repository

import (
	"bufio"
	"fmt"
	"os"

	"gopkg.in/masci/flickr.v2"
)

func Authenticate() (*flickr.FlickrClient, error) {
	client := flickr.NewFlickrClient(os.Getenv("FLICKR_KEY"), os.Getenv("FLICKR_SECRET"))
	// first, get a request token
	requestTok, _ := flickr.GetRequestToken(client)

	// build the authorizatin URL
	url, _ := flickr.GetAuthorizeUrl(client, requestTok)
	fmt.Println(fmt.Sprintf("Visit the following URL to get your authorization code: \n%s", url))

	// ask user to hit the authorization url with
	// their browser, authorize this application and coming
	// back with the confirmation token
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter authorization code: ")
	code, _ := reader.ReadString('\n')

	// finally, get the access token, setup the client and start making requests
	accessTok, _ := flickr.GetAccessToken(client, requestTok, code)
	client.OAuthToken = accessTok.OAuthToken
	client.OAuthTokenSecret = accessTok.OAuthTokenSecret

	return client, nil
}
