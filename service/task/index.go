package service

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type TaskService struct {
	db *sql.DB
}

type Task struct {
	ID          int    `json:"id"`
	CardID      int    `json:"card_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Start       string `json:"start"`
	Percent     int    `json:"percent"`
	Deadline    string `json:"deadline"`
	Completed   bool   `json:"completed"`
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{
		db: db,
	}
}

func (ts *TaskService) GetAll(cardID int) ([]Task, error) {
	rows, err := ts.db.Query(fmt.Sprintf("SELECT * FROM task WHERE card_id=%d", cardID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.CardID, &task.Title, &task.Description, &task.Start, &task.Percent, &task.Deadline, &task.Completed)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "error iterating over rows")
	}

	return tasks, nil
}

func (ts *TaskService) GetOne(taskID int) (*Task, error) {
	row := ts.db.QueryRow(fmt.Sprintf("SELECT * FROM task WHERE id=%d", taskID))

	var task Task
	err := row.Scan(&task.ID, &task.CardID, &task.Title, &task.Description, &task.Start, &task.Percent, &task.Deadline, &task.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("task not found")
		}
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &task, nil
}

func (ts *TaskService) Add(cardID int, title string, description string, start string, deadline string, percent int) (*Task, error) {
	var query string
	if description != "" && start != "" && deadline != "" {
		query = fmt.Sprintf(`INSERT INTO task (card_id, title, description, start, deadline, percent, completed)
			VALUES (%d, '%s', '%s', '%s', '%s', %d, false) RETURNING *`, cardID, title, description, start, deadline, percent)
	} else if description != "" && start != "" {
		query = fmt.Sprintf(`INSERT INTO task (card_id, title, description, start, percent, completed)
			VALUES (%d, '%s', '%s', '%s', %d, false) RETURNING *`, cardID, title, description, start, percent)
	} else if description != "" && deadline != "" {
		query = fmt.Sprintf(`INSERT INTO task (card_id, title, description, deadline, percent, completed)
			VALUES (%d, '%s', '%s', '%s', %d, false) RETURNING *`, cardID, title, description, deadline, percent)
	} else if start != "" && deadline != "" {
		query = fmt.Sprintf(`INSERT INTO task (card_id, title, start, deadline, percent, completed)
			VALUES (%d, '%s', '%s', '%s', %d, false) RETURNING *`, cardID, title, start, deadline, percent)
	} else if description != "" {
		query = fmt.Sprintf(`INSERT INTO task (card_id, title, description, percent, completed)
			VALUES (%d, '%s', '%s', %d, false) RETURNING *`, cardID, title, description, percent)
	} else if start != "" {
		query = fmt.Sprintf(`INSERT INTO task (card_id, title, start, percent, completed)
			VALUES (%d, '%s', '%s', %d, false) RETURNING *`, cardID, title, start, percent)
	} else if deadline != "" {
		query = fmt.Sprintf(`INSERT INTO task (card_id, title, deadline, percent, completed)
			VALUES (%d, '%s', '%s', %d, false) RETURNING *`, cardID, title, deadline, percent)
	} else {
		query = fmt.Sprintf(`INSERT INTO task (card_id, title, percent, completed)
			VALUES (%d, '%s', %d, false) RETURNING *`, cardID, title, percent)
	}

	row := ts.db.QueryRow(query)

	var task Task
	err := row.Scan(&task.ID, &task.CardID, &task.Title, &task.Description, &task.Start, &task.Percent, &task.Deadline, &task.Completed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &task, nil
}

func (ts *TaskService) Remove(taskID int) (bool, error) {
	result, err := ts.db.Exec(fmt.Sprintf("DELETE FROM task WHERE id=%d", taskID))
	if err != nil {
		return false, errors.Wrap(err, "failed to execute query")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, errors.Wrap(err, "failed to get rows affected")
	}

	if rowsAffected == 0 {
		return false, errors.New("task not found")
	}

	return true, nil
}

func (ts *TaskService) ChangeTitle(taskID int, title string) (*Task, error) {
	row := ts.db.QueryRow(fmt.Sprintf("UPDATE task SET title='%s' WHERE id=%d RETURNING *", title, taskID))

	var task Task
	err := row.Scan(&task.ID, &task.CardID, &task.Title, &task.Description, &task.Start, &task.Percent, &task.Deadline, &task.Completed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &task, nil
}

func (ts *TaskService) ChangeDescription(taskID int, description string) (*Task, error) {
	row := ts.db.QueryRow(fmt.Sprintf("UPDATE task SET description='%s' WHERE id=%d RETURNING *", description, taskID))

	var task Task
	err := row.Scan(&task.ID, &task.CardID, &task.Title, &task.Description, &task.Start, &task.Percent, &task.Deadline, &task.Completed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &task, nil
}

func (ts *TaskService) ChangeStart(taskID int, start string) (*Task, error) {
	row := ts.db.QueryRow(fmt.Sprintf("UPDATE task SET start='%s' WHERE id=%d RETURNING *", start, taskID))

	var task Task
	err := row.Scan(&task.ID, &task.CardID, &task.Title, &task.Description, &task.Start, &task.Percent, &task.Deadline, &task.Completed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &task, nil
}

func (ts *TaskService) ChangeDeadline(taskID int, deadline string) (*Task, error) {
	row := ts.db.QueryRow(fmt.Sprintf("UPDATE task SET deadline='%s' WHERE id=%d RETURNING *", deadline, taskID))

	var task Task
	err := row.Scan(&task.ID, &task.CardID, &task.Title, &task.Description, &task.Start, &task.Percent, &task.Deadline, &task.Completed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &task, nil
}

func (ts *TaskService) ChangeComplete(taskID int, status bool) (*Task, error) {
	row := ts.db.QueryRow(fmt.Sprintf("UPDATE task SET completed=%v WHERE id=%d RETURNING *", status, taskID))

	var task Task
	err := row.Scan(&task.ID, &task.CardID, &task.Title, &task.Description, &task.Start, &task.Percent, &task.Deadline, &task.Completed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &task, nil
}

func (ts *TaskService) ChangeCard(taskID int, cardID int) (*Task, error) {
	row := ts.db.QueryRow(fmt.Sprintf("UPDATE task SET card_id=%d WHERE id=%d RETURNING *", cardID, taskID))

	var task Task
	err := row.Scan(&task.ID, &task.CardID, &task.Title, &task.Description, &task.Start, &task.Percent, &task.Deadline, &task.Completed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &task, nil
}

func (ts *TaskService) ChangePercent(taskID int, percent int) (*Task, error) {
	row := ts.db.QueryRow(fmt.Sprintf("UPDATE task SET percent=%d WHERE id=%d RETURNING *", percent, taskID))

	var task Task
	err := row.Scan(&task.ID, &task.CardID, &task.Title, &task.Description, &task.Start, &task.Percent, &task.Deadline, &task.Completed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan row")
	}

	return &task, nil
}
