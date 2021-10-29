package meural

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

var base_url string = "https://api.meural.com/v0"

type Meural struct {
	Token   string `json:"token"`
	Session string `json:"session"`
}

func Authenticate(username, password string) (Meural, error) {
	m := Meural{}
	payload := strings.NewReader(fmt.Sprintf("username=%s&password=%s",
		url.QueryEscape(username),
		url.QueryEscape(password),
	))

	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/authenticate", base_url), payload)

	if err != nil {
		fmt.Println(err)
		return m, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return m, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return m, err
	}

	err = json.Unmarshal(body, &m)
	return m, nil
}

func AddToGallery(galleryID string, URL string, fileName string, r Meural) error {
	url := fmt.Sprintf("https://api.meural.com/v1/galleries/%s/items", galleryID)
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	// file, err := os.Open(filePath)
	// defer file.Close()
	part, err := writer.CreateFormFile("image", fileName)
	_, err = io.Copy(part, bytes.NewReader(data))

	if err != nil {
		fmt.Println(err)
		return err
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Add("authority", "api.meural.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("authorization", fmt.Sprintf("Token %s", r.Token))
	req.Header.Add("x-meural-source-platform", "web")
	req.Header.Add("accept", "*/*")
	req.Header.Add("origin", "https://my.meural.netgear.com")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// fmt.Println(string(body))
	return nil
}
