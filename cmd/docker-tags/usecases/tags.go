package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	APIURL    = "https://hub.docker.com"
	PATH_FMT  = "%s/v2/repositories/%s/%s/tags/"
	NAMESPACE = "library"
)

func Tags(ctx context.Context, image string, p map[string]any) ([]string, error) {
	if len(image) == 0 {
		return nil, ErrNoImageName
	}
	url, err := url.Parse(fmt.Sprintf(PATH_FMT, APIURL, NAMESPACE, image))
	if err != nil {
		return nil, ErrInvalidURL(url.String())
	}
	q := url.Query()
	q.Set("page_size", fmt.Sprint(pageSize(p)))
	url.RawQuery = q.Encode()

	fmt.Println(url)

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, ErrInvalidRequest(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, ErrInvalidStatus(resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var res Response
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	tags := make([]string, 0, len(res.Results))
	for _, r := range res.Results {
		tags = append(tags, r.Name)
	}
	return tags, nil
}

func pageSize(p map[string]any) int {
	if size, ok := p["size"]; ok {
		return size.(int)
	}
	return 10
}
