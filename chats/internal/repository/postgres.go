package repository

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ChatsRepository struct {
	tableName string
	pool *pgxpool.Pool
}

func NewChatsRepository(tableName string, pool *pgxpool.Pool) ChatsRepository {
	return ChatsRepository{tableName: tableName, pool: pool}
}

func (r ChatsRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}