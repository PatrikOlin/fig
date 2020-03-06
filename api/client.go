package api		   

import (
	"net/http"
	"net/url"		   
	"encoding/json"	   
	"io"			   
	"bytes"			   
	"strconv"
	"fig/models"
	"time"
	"fmt"
)			

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	httpClient *http.Client
}
			
			
func NewBasicClient(urlString string) *Client {
	baseUrl, _ := url.Parse(urlString)
	fmt.Println(baseUrl)
	return &Client{
		BaseURL: baseUrl,
		UserAgent: "fig",

		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},

	}
}

func (c *Client) GetArticles(numOfArticles int) ([]models.Article, error) {

	req, err := c.newRequest("GET", "/articles?amount=" + strconv.Itoa(numOfArticles), nil)
	fmt.Println(numOfArticles)
	if err != nil {
		return nil, err
	}

	var articles []models.Article
	_, err = c.do(req, &articles)


	return articles, err
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.BaseURL.String() + path, buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
