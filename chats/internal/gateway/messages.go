package gateway

import (
	"context"
	"bytes"
	"io"
	"net/http"
	"encoding/json"

	"github.com/bd878/lesnotes_bot/chats/internal/domain"
	"github.com/bd878/lesnotes_bot/internal/logger"
	galleryMessages "github.com/bd878/gallery/server/messages/pkg/model"
)

type MessagesGateway struct {
	client *http.Client
	url   string
}

func NewMessagesGateway(client *http.Client, url string) MessagesGateway {
	return MessagesGateway{client: client, url: url}
}

func (g MessagesGateway) Save(ctx context.Context, message *domain.Message) (int32, error) {
	data, err := json.Marshal(message.Message)
	if err != nil {
		logger.Log.Debug(err)
		return 0, err
	}

	buff := bytes.NewReader(data)

	req, err := http.NewRequestWithContext(ctx, "post", g.url + "/messages/v2/send", buff)
	if err != nil {
		return 0, err
	}

	req.Header["Content-Type"] = []string{"application/json"}

	resp, err := g.client.Do(req)
	if err != nil {
		logger.Log.Debug(err)
		return 0, err
	}

	data, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		logger.Log.Debug(err)
		return 0, err
	}

	logger.Log.Debugf("%s\n", data)

	var result galleryMessages.NewMessageServerResponse
	if err := json.Unmarshal(data, &result); err != nil {
		logger.Log.Debug(err)
		return 0, err
	}

	return result.Message.ID, nil
}
