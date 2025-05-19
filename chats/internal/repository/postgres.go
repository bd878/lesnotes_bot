package repository

import (
	"fmt"
	"context"

	"github.com/go-telegram/bot/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/bd878/lesnotes_bot/internal/i18n"
	"github.com/bd878/lesnotes_bot/chats/internal/domain"
)

type ChatsRepository struct {
	tableName string
	pool *pgxpool.Pool
}

func NewChatsRepository(tableName string, pool *pgxpool.Pool) *ChatsRepository {
	return &ChatsRepository{tableName: tableName, pool: pool}
}

func (r ChatsRepository) FindChat(ctx context.Context, chatID int64) (*domain.Chat, error) {
	const query = "SELECT lang, type, title, username, first_name, last_name, is_forum FROM %s WHERE id = $1 LIMIT 1"

	chat := &domain.Chat{
		Chat: &models.Chat{
			ID: chatID,
		},
	}

	var (
		chatType, lang string
	)

	err := r.pool.QueryRow(ctx, r.table(query), chatID).Scan(&lang, &chatType, &chat.Chat.Title,
		&chat.Chat.Username, &chat.Chat.FirstName, &chat.Chat.LastName, &chat.Chat.IsForum)
	if err != nil {
		return nil, err
	}

	chat.Chat.Type = models.ChatType(chatType)
	chat.Lang = i18n.LangFromString(lang)

	return chat, nil
}

func (r ChatsRepository) Save(ctx context.Context, chat *domain.Chat) error {
	const query = "INSERT INTO %s (lang, id, type, title, username, first_name, last_name, is_forum) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	_, err := r.pool.Exec(ctx, r.table(query), chat.Lang.String(), chat.Chat.ID, chat.Chat.Type, chat.Chat.Title,
		chat.Chat.Username, chat.Chat.FirstName, chat.Chat.LastName, chat.Chat.IsForum)

	return err
}

func (r ChatsRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}