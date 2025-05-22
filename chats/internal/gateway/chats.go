package gateway

import (
	"context"
	"bytes"
	"io"
	"net/http"
	"encoding/json"

	"github.com/bd878/lesnotes_bot/internal/logger"
	galleryUsers "github.com/bd878/gallery/server/users/pkg/model"
)

type ChatsGateway struct {
	client *http.Client
	url   string
}

func NewChatsGateway(client *http.Client, url string) ChatsGateway {
	return ChatsGateway{client: client, url: url}
}

func (g ChatsGateway) Signup(ctx context.Context, login, password string) (token string, err error) {
	user := &galleryUsers.User{
		Name: login,
		Password: password,
	}

	data, err := json.Marshal(user)
	if err != nil {
		logger.Log.Debug(err)
		return "", err
	}

	buff := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, "POST", g.url + "/users/v2/signup", buff)
	if err != nil {
		logger.Log.Debug(err)
		return "", err
	}

	req.Header["Content-Type"] = []string{"application/json"}

	resp, err := g.client.Do(req)
	if err != nil {
		logger.Log.Debug(err)
		return "", err
	}

	data, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		logger.Log.Debug(err)
		return "", err
	}

	var result galleryUsers.SignupJsonUserServerResponse
	if err := json.Unmarshal(data, &result); err != nil {
		logger.Log.Debug(err)
		return "", err
	}

	return result.Token, nil
}

func (g ChatsGateway) Login(ctx context.Context, login, password string) (token string, err error) {
	return "", nil
}