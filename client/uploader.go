package client

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func SetProfileImage(Client *SlackClient, path string) {
	fmt.Printf("Setting profile image to file locatated at %v", path)
	values := map[string]io.Reader{
		"image": mustOpen(path),
	}
	err := Upload(Client, values)
	if err != nil {
		panic(err)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

// Upload function
// https://stackoverflow.com/questions/20205796/post-data-using-the-content-type-multipart-form-data
func Upload(Client *SlackClient, values map[string]io.Reader) (err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", Client.BaseURL, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+Client.Token)
	// Submit the request
	res, err := Client.HTTPClient.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	fmt.Printf("Response: %v", res.StatusCode)

	return
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}
