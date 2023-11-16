package memo

import (
	"database/sql"

	"github.com/koki120/go-rest-api/domain/entity"
	"github.com/koki120/go-rest-api/interface/i_memo"
)

type MemoRepository struct {
	db *sql.DB
}

func NewMemoRepository(db *sql.DB) i_memo.Repository {
	return &MemoRepository{
		db: db,
	}
}

func (r *MemoRepository) FindByID(userID string, memoID string) (entity.Memo, error) {
	query := `
		SELECT memo_id, title, body, created_at, updated_at
		FROM memos
		WHERE user_id = ? AND id = ?
	`
	row := r.db.QueryRow(query, userID, memoID)

	return ParseMemoEntity(row)
}

func (r *MemoRepository) Create(userID string, memo entity.Memo) error {
	query := `
		INSERT INTO memos (user_id, title, body)
		VALUES (?, ?, ?)
	`
	_, err := r.db.Exec(query, userID, memo.Title, memo.Body)
	if err != nil {
		return err
	}
	return nil
}

func (r *MemoRepository) Update(memo entity.Memo) error {
	query := `
		UPDATE memos
		SET title = ?, body = ?
		WHERE AND memo_id = ?
	`
	_, err := r.db.Exec(query, memo.Title, memo.Body, memo.MemoID)
	if err != nil {
		return err
	}
	return nil
}

func (r *MemoRepository) Delete(memoID string) error {
	query := `
		DELETE FROM memos
		WHERE memo_id = ?
	`
	_, err := r.db.Exec(query, memoID)
	return err
}

func (r *MemoRepository) Search() ([]entity.Memo, int, error) {
	panic("unimplemented")
}

func ParseMemoEntity(row *sql.Row) (entity.Memo, error) {
	var memo entity.Memo
	err := row.Scan(&memo.MemoID, &memo.Title, &memo.Body, &memo.CreatedAt, &memo.UpdatedAt)
	if err != nil {
		return entity.Memo{}, err
	}
	return memo, nil
}
