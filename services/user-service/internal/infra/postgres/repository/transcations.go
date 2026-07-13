package repository

import (
	"context"
	"errors"
	"financial_assistant/services/user-service/internal/entities"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrTransactionNotFound = errors.New("transaction not found")

type TransactionsRepo struct {
	pool *pgxpool.Pool
}

func NewTransactionsRepo(pool *pgxpool.Pool) *TransactionsRepo {
	return &TransactionsRepo{
		pool: pool,
	}
}

func (r TransactionsRepo) Upsert(ctx context.Context, transaction *entities.Transaction) error {
	query := `
		INSERT INTO transactions (id, user_id, item, price, class, note, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
		ON CONFLICT (id) DO UPDATE SET
			item = EXCLUDED.item,
			price = EXCLUDED.price,
			class = EXCLUDED.class,
			note = EXCLUDED.note
		WHERE transactions.user_id = EXCLUDED.user_id
	`

	res, err := r.pool.Exec(ctx, query, transaction.ID, transaction.UserID, transaction.Item, transaction.Price, transaction.Class, transaction.Note, entities.TransactionStatusPending)
	if err != nil {
		return fmt.Errorf("error upserting transaction: %w", err)
	}

	if res.RowsAffected() == 0 {
		return ErrTransactionNotFound
	}

	return nil
}

func (r TransactionsRepo) GetByID(ctx context.Context, id string, userID string) (*entities.Transaction, error) {
	query := `SELECT id, user_id, item, price, class, note FROM transactions WHERE id = $1 AND user_id = $2`

	var transaction entities.Transaction
	err := r.pool.QueryRow(ctx, query, id, userID).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Item,
		&transaction.Price,
		&transaction.Class,
		&transaction.Note,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTransactionNotFound
		}
		return nil, fmt.Errorf("error getting transaction: %w", err)
	}

	return &transaction, nil
}

func (r TransactionsRepo) List(ctx context.Context, userID string) ([]entities.Transaction, error) {
	query := `SELECT id, user_id, item, price, class, note FROM transactions WHERE user_id = $1`

	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error listing transactions: %w", err)
	}
	defer rows.Close()

	var transactions []entities.Transaction
	for rows.Next() {
		var transaction entities.Transaction
		if err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.Item,
			&transaction.Price,
			&transaction.Class,
			&transaction.Note,
		); err != nil {
			return nil, fmt.Errorf("error scanning transaction: %w", err)
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating transactions: %w", err)
	}

	return transactions, nil
}

func (r TransactionsRepo) Delete(ctx context.Context, id string, userID string) error {
	query := `DELETE FROM transactions WHERE id = $1 AND user_id = $2`

	res, err := r.pool.Exec(ctx, query, id, userID)
	if err != nil {
		return fmt.Errorf("error deleting transaction: %w", err)
	}

	if res.RowsAffected() == 0 {
		return ErrTransactionNotFound
	}

	return nil
}

func (r *TransactionsRepo) GetPendingTransactions(ctx context.Context) ([]entities.Transaction, error) {
	query := `
		UPDATE transactions
		SET status = $1
		WHERE status = $2
		RETURNING id, user_id, item, price, class, note
	`

	rows, err := r.pool.Query(ctx, query, string(entities.TransactionStatusProcessing), string(entities.TransactionStatusPending))
	if err != nil {
		return nil, fmt.Errorf("error getting pending transactions: %w", err)
	}
	defer rows.Close()

	var result []entities.Transaction
	for rows.Next() {
		var transaction entities.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Item, &transaction.Price, &transaction.Class, &transaction.Note); err != nil {
			return nil, fmt.Errorf("error scanning pending transactions: %w", err)
		}
		result = append(result, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating pending transactions: %w", err)
	}

	return result, nil
}

func (r *TransactionsRepo) MarkTransactionsProcessed(ctx context.Context, ids []uuid.UUID, success bool) error {
	if len(ids) == 0 {
		return nil
	}

	status := entities.TransactionStatusPending
	if success {
		status = entities.TransactionStatusComplete
	}

	idStrings := make([]string, len(ids))
	for i, id := range ids {
		idStrings[i] = id.String()
	}

	query := `UPDATE transactions SET status = $1 WHERE id = ANY($2::uuid[])`
	if _, err := r.pool.Exec(ctx, query, string(status), idStrings); err != nil {
		return fmt.Errorf("error updating transactions status: %w", err)
	}

	return nil
}
