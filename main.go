package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joooostb/flickr-meural-sync/pkg/meural"
	"github.com/joooostb/flickr-meural-sync/pkg/repository"
	"github.com/joooostb/flickr/photosets"
)

func main() {
	m, err := meural.Authenticate(os.Getenv("MEURAL_EMAIL"), os.Getenv("MEURAL_PASSWORD"))

	if err != nil {
		log.Fatalln("Failed to authenticate with Netgear Meural: %w", err)
	}

	c, _ := repository.Authenticate()

	r, _ := photosets.GetList(c, true, "", 1)
	album, _ := selectAlbum(r)
	p, _ := photosets.GetPhotos(c, true, r.Photosets.Items[album].Id, "", 1)

	for i, v := range p.Photoset.Photos {
		fmt.Println(fmt.Sprintf("Uploading image %d/%d: %s (%s)", i+1, len(p.Photoset.Photos), v.Title, v.URLO))
		meural.AddToGallery(os.Getenv("MEURAL_GALLERY_ID"), v.URLO, fmt.Sprintf("%s.jpg", v.Title), m)
	}
	fmt.Println("Sync completed 😃")
}

func selectAlbum(v *photosets.PhotosetsListResponse) (int, error) {
	fmt.Println("Select a Flickr Album you'd like to sync:")
	for i, v := range v.Photosets.Items {
		fmt.Println(fmt.Sprintf("[%d] %s", i, v.Title))
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Album number: ")
	s, _ := reader.ReadString('\n')
	album, err := strconv.Atoi(strings.Split(s, "\n")[0])

	if err != nil {
		fmt.Println(fmt.Errorf("Not a album valid choice, exitting: %w", err))
		selectAlbum(v)
	}

	return album, err
}
