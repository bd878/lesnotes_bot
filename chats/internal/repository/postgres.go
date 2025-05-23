package repository

import (
	"fmt"
	"context"

	"github.com/go-telegram/bot/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/bd878/lesnotes_bot/internal/ddd"
	"github.com/bd878/lesnotes_bot/internal/i18n"
	"github.com/bd878/lesnotes_bot/chats/internal/domain"
)

type ChatRepository struct {
	tableName string
	pool *pgxpool.Pool
}

var _ domain.ChatRepository = (*ChatRepository)(nil)

func NewChatRepository(tableName string, pool *pgxpool.Pool) *ChatRepository {
	return &ChatRepository{tableName: tableName, pool: pool}
}

func (r ChatRepository) Load(ctx context.Context, chatID int64) (*domain.Chat, error) {
	const query = "SELECT lang, type, token, login, password, title, username, first_name, last_name, is_forum FROM %s WHERE id = $1 LIMIT 1"

	chat := &domain.Chat{
		Chat: &models.Chat{
			ID: chatID,
		},
	}

	var (
		chatType, lang string
	)

	err := r.pool.QueryRow(ctx, r.table(query), chatID).Scan(&lang, &chatType, &chat.Token, &chat.Login, &chat.Password, &chat.Chat.Title,
		&chat.Chat.Username, &chat.Chat.FirstName, &chat.Chat.LastName, &chat.Chat.IsForum)
	if err != nil {
		return nil, err
	}

	chat.Aggregate = ddd.NewAggregate(domain.ChatAggregate)

	chat.Chat.Type = models.ChatType(chatType)
	chat.Lang = i18n.LangFromString(lang)

	return chat, nil
}

func (r ChatRepository) Save(ctx context.Context, chat *domain.Chat) error {
	const query = "INSERT INTO %s (lang, id, type, token, login, password, title, username, first_name, last_name, is_forum) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"

	_, err := r.pool.Exec(ctx, r.table(query), chat.Lang.String(), chat.Chat.ID, chat.Chat.Type, chat.Token, chat.Login, chat.Password, chat.Chat.Title,
		chat.Chat.Username, chat.Chat.FirstName, chat.Chat.LastName, chat.Chat.IsForum)

	return err
}

func (r ChatRepository) UpdateToken(ctx context.Context, chatID int64, token string) error {
	const query = "UPDATE %s SET token = $2 WHERE id = $1"

	_, err := r.pool.Exec(ctx, r.table(query), chatID, token)

	return err
}

func (r ChatRepository) Remove(ctx context.Context, chatID int64) error {
	const query = "DELETE FROM %s WHERE id = $1"

	_, err := r.pool.Exec(ctx, r.table(query), chatID)

	return err
}

func (r ChatRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}