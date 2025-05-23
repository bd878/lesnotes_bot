package gateway

import (
	"context"
	"errors"
	"bytes"
	"io"
	"net/http"
	"encoding/json"

	"github.com/bd878/lesnotes_bot/internal/logger"
	"github.com/bd878/lesnotes_bot/chats/internal/domain"
	gallerymodel "github.com/bd878/gallery/server/pkg/model"
	galleryUsers "github.com/bd878/gallery/server/users/pkg/model"
)

type ChatGateway struct {
	client *http.Client
	url   string
}

var _ domain.ChatGateway = (*ChatGateway)(nil)

func NewChatGateway(client *http.Client, url string) ChatGateway {
	return ChatGateway{client: client, url: url}
}

func (g ChatGateway) Signup(ctx context.Context, login, password string) (token string, err error) {
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

	if result.Status == "error" {
		logger.Log.Debug(result)
		return "", errors.New(result.Description)
	}

	logger.Log.Debug(result)
	return result.Token, nil
}

func (g ChatGateway) Auth(ctx context.Context, token string) error {
	data, err := json.Marshal(&gallerymodel.JSONServerRequest{
		Token: token,
	})
	if err != nil {
		logger.Log.Debug(err)
		return err
	}

	buff := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, "POST", g.url + "/users/v2/auth", buff)
	if err != nil {
		logger.Log.Debug(err)
		return err
	}

	req.Header["Content-Type"] = []string{"application/json"}

	resp, err := g.client.Do(req)
	if err != nil {
		logger.Log.Debug(err)
		return err
	}

	data, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		logger.Log.Debug(err)
		return err
	}

	var result galleryUsers.ServerAuthorizeResponse
	if err := json.Unmarshal(data, &result); err != nil {
		logger.Log.Debug(err)
		return err
	}

	if result.Status == "error" {
		logger.Log.Debug(result)
		return errors.New(result.Description)
	}

	if result.Expired {
		logger.Log.Debug(result)
		return domain.ErrExpired
	}

	return nil
}

func (g ChatGateway) Login(ctx context.Context, login, password string) (token string, err error) {
	data, err := json.Marshal(&galleryUsers.User{
		Name: login,
		Password: password,
	})
	if err != nil {
		logger.Log.Debug(err)
		return "", err
	}

	buff := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, "POST", g.url + "/users/v2/login", buff)
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

	var result galleryUsers.LoginJsonUserServerResponse
	if err := json.Unmarshal(data, &result); err != nil {
		logger.Log.Debug(err)
		return "", err
	}

	if result.Status == "error" {
		logger.Log.Debug(result)
		return "", errors.New(result.Description)
	}

	return result.Token, nil
}

func (g ChatGateway) Delete(ctx context.Context, login, password, token string) error {
	data, err := json.Marshal(&galleryUsers.DeleteUserJsonRequest{
		Name: login,
		Password: password,
		Token: token,
	})
	if err != nil {
		logger.Log.Debug(err)
		return err
	}

	buff := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, "POST", g.url + "/users/v2/delete", buff)
	if err != nil {
		logger.Log.Debug(err)
		return err
	}

	req.Header["Content-Type"] = []string{"application/json"}

	resp, err := g.client.Do(req)
	if err != nil {
		logger.Log.Debug(err)
		return err
	}

	data, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		logger.Log.Debug(err)
		return err
	}

	var result galleryUsers.DeleteUserJsonServerResponse
	if err := json.Unmarshal(data, &result); err != nil {
		logger.Log.Debug(err)
		return err
	}

	if result.Status == "error" {
		logger.Log.Debug(result)
		return errors.New(result.Description)
	}

	return nil
}