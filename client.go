package nest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/blaskovicz/go-nest/models"
)

const (
	DefaultAccessTokenURI = "https://api.home.nest.com/oauth2/access_token"
	DefaultRootURI        = "https://developer-api.nest.com"
	envVarPrefix          = "NEST"
)

type Client struct {
	secret         string
	id             string
	rootURI        string
	accessTokenURI string
	accessToken    string
}

func New() *Client {
	rootURI := os.Getenv(envVarPrefix + "_ROOT_URI")
	if rootURI == "" {
		rootURI = DefaultRootURI
	}
	accessTokenURI := os.Getenv(envVarPrefix + "_ACCESS_TOKEN_URI")
	if accessTokenURI == "" {
		accessTokenURI = DefaultAccessTokenURI
	}
	accessToken := os.Getenv(envVarPrefix + "_ACCESS_TOKEN")
	secret := os.Getenv(envVarPrefix + "_CLIENT_SECRET")
	id := os.Getenv(envVarPrefix + "_CLIENT_ID")

	c := &Client{}
	return c.
		SetAccessToken(accessToken).
		SetRootURI(rootURI).
		SetAccessTokenURI(accessTokenURI).
		SetSecret(secret).
		SetID(id)
}

func (c *Client) SetRootURI(rootURI string) *Client {
	c.rootURI = rootURI
	return c
}
func (c *Client) SetAccessToken(accessToken string) *Client {
	c.accessToken = accessToken
	return c
}
func (c *Client) SetAccessTokenURI(accessTokenURI string) *Client {
	c.accessTokenURI = accessTokenURI
	return c
}
func (c *Client) SetSecret(secret string) *Client {
	c.secret = secret
	return c
}
func (c *Client) SetID(id string) *Client {
	c.id = id
	return c
}
func (c *Client) AccessToken() string    { return c.accessToken }
func (c *Client) RootURI() string        { return c.rootURI }
func (c *Client) AccessTokenURI() string { return c.accessTokenURI }
func (c *Client) ID() string             { return c.id }
func (c *Client) Secret() string         { return c.secret }

func (c *Client) uri(path string, pathArgs ...interface{}) string {
	return fmt.Sprintf("%s%s", c.rootURI, fmt.Sprintf(path, pathArgs...))
}

func (c *Client) CreateAccessToken(code string) (*models.CreateAccessTokenResponse, error) {
	formData := url.Values{
		"client_id":     []string{c.id},
		"client_secret": []string{c.secret},
		"code":          []string{code},
		"grant_type":    []string{"authorization_code"},
	}.Encode()
	req, err := http.NewRequest("POST", c.accessTokenURI, strings.NewReader(formData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var at models.CreateAccessTokenResponse
	if err = c.do(req, &at); err != nil {
		return nil, err
	}
	return &at, nil
}

// https://developers.nest.com/documentation/cloud/api-camera
func (c *Client) ListCameras() ([]*models.Camera, error) {
	req, err := http.NewRequest("GET", c.uri("/devices/cameras"), nil)
	if err != nil {
		return nil, err
	}
	cams := map[string]*models.Camera{}
	if err = c.do(req, &cams); err != nil {
		return nil, err
	}
	cams2 := []*models.Camera{}
	for k := range cams {
		cams2 = append(cams2, cams[k])
	}
	return cams2, nil
}

func (c *Client) UpdateCameraIsStreaming(cameraID string, isStreaming bool) (*models.Camera, error) {
	body, err := json.Marshal(map[string]interface{}{"is_streaming": isStreaming})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", c.uri("/devices/cameras/%s", cameraID), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	var cam models.Camera
	if err = c.do(req, &cam); err != nil {
		return nil, err
	}
	return &cam, nil
}

// do a request, return the undread response if no errors and 200 OK
func (c *Client) doWithResponse(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	if c.accessToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := (&http.Client{
		CheckRedirect: redirectPolicyFunc(req.Header.Get("Authorization")),
	}).Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %s", err)
	}
	// error
	if resp.StatusCode != http.StatusOK {
		rawBody, err := ioutil.ReadAll(resp.Body)
		//fmt.Printf("TODO: %#v", string(rawBody))
		if err != nil {
			return nil, fmt.Errorf("failed to read %s response payload: %s", resp.Status, err)
		}
		if rawBody != nil && len(rawBody) != 0 {
			var e models.ErrorResponse
			if err = json.Unmarshal(rawBody, &e); err != nil {
				return nil, fmt.Errorf("failed to decode %s error payload (%s): %s", resp.Status, string(rawBody), err)
			}
			return nil, fmt.Errorf("request failed with %s: %v", resp.Status, e)
		}
		return nil, fmt.Errorf("request failed with %s: %s", resp.Status, string(rawBody))
	}
	return resp, nil
}

func (c *Client) do(req *http.Request, decodeTarget interface{}) error {
	if decodeTarget != nil {
		if decodeKind := reflect.TypeOf(decodeTarget).Kind(); decodeKind != reflect.Ptr {
			return fmt.Errorf("invalid decode target type %s (need %s)", decodeKind.String(), reflect.Ptr.String())
		}
	}
	resp, err := c.doWithResponse(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if decodeTarget != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %s", err)
		}
		err = json.Unmarshal(body, decodeTarget)
		if err != nil {
			return fmt.Errorf("failed to decode payload: %s\n%s", err, string(body))
		}
	}
	return nil
}

// grabbed from my proxy:
// https://github.com/blaskovicz/garmin-nest-api-proxy/blob/master/cmd/web/main.go#L38
func redirectPolicyFunc(authHeader string) func(*http.Request, []*http.Request) error {
	return func(r *http.Request, via []*http.Request) error {
		//logrus.Infof("==> %s (policy)", r.URL.String())
		// nest api proxies to firebase services somewhere else
		r.Header.Set("Authorization", authHeader)
		if via != nil && len(via) == 10 {
			return fmt.Errorf("too many redirects")
		}
		return nil
	}
}
