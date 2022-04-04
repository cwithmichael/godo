package mysql

import (
	"database/sql"
	"errors"
	"github.com/cwithmichael/godo/pkg/models"
)

// TodoModel type which wraps a sql.DB connection pool.
type TodoModel struct {
	DB *sql.DB
}

// Insert a new todo into the database.
func (m *TodoModel) Insert(title, content string, userID int) (int, error) {
	stmt := `INSERT INTO todo (title, content, created, completed, user_id)
			VALUES(?, ?, UTC_TIMESTAMP(), ?, ?)`

	result, err := m.DB.Exec(stmt, title, content, false, userID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Delete will delete a specific todo based on its id.
func (m *TodoModel) Delete(id int) error {
	stmt := `DELETE FROM todo WHERE id = ?`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil

}

// Update will update a specific todo based on its id.
func (m *TodoModel) Update(id int, title, content string, completed bool) error {
	stmt := `UPDATE todo SET title = ?, content = ?, completed = ? WHERE id = ?`

	_, err := m.DB.Exec(stmt, title, content, completed, id)
	if err != nil {
		return err
	}
	return nil
}

// Get will return a specific todo based on its id.
func (m *TodoModel) Get(id int) (*models.Todo, error) {
	stmt := `SELECT id, title, content, created, completion_date, completed, user_id FROM todo
    WHERE id = ?`
	t := &models.Todo{}

	err := m.DB.QueryRow(stmt, id).Scan(&t.ID, &t.Title, &t.Content, &t.Created, &t.CompletionDate, &t.Completed, &t.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err

	}

	return t, nil

}

// Latest will return the 10 most recently created todos.
func (m *TodoModel) Latest(userID int) ([]*models.Todo, error) {
	// Write the SQL statement we want to execute.
	stmt := `SELECT id, title, content, created, completion_date, completed FROM todo
    WHERE user_id=? ORDER BY created DESC LIMIT 10 `

	rows, err := m.DB.Query(stmt, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []*models.Todo{}

	for rows.Next() {
		t := &models.Todo{}
		err = rows.Scan(&t.ID, &t.Title, &t.Content, &t.Created, &t.CompletionDate, &t.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
