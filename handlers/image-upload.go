package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadImages(files *[]*multipart.FileHeader) (*[]string, error) {
	images := make([]string, 0)
	for _, file := range *files {
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()
		url, err := uploadImage(src, file.Filename)
		if err != nil {
			return nil, err
		}
		images = append(images, url)
	}
	return &images, nil
}

func uploadImage(file io.Reader, filename string) (string, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	fw, err := w.CreateFormFile("file", filename)
	if err != nil {
		return "", err
	}
	if _, err := io.Copy(fw, file); err != nil {
		return "", err
	}

	fw, err = w.CreateFormField("fileName")
	if err != nil {
		return "", err
	}
	if _, err = fw.Write([]byte(filename)); err != nil {
		return "", err
	}
	w.Close()

	req, err := http.NewRequest("POST", "https://upload.imagekit.io/api/v1/files/upload", &buf)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(os.Getenv("IMAGEKIT_API_KEY")+":"))))
	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("image API responded with status code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var jsonRes ImageKitResponse
	if err := json.Unmarshal(body, &jsonRes); err != nil {
		return "", err
	}
	return jsonRes.URL, nil
}
