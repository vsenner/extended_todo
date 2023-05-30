package task_service

import (
	//"database/sql"
	db "extended_todo/routing"
	"fmt"

	"github.com/pkg/errors"
)

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

func GetAll(cardID int) ([]Task, error) {
	rows, err := db.DB.Query(fmt.Sprintf("SELECT * FROM task WHERE card_id=%d", cardID))
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

func GetOne(taskID int) Task {

	t := Task{}

	if err := db.DB.QueryRow("SELECT * FROM task WHERE id=$1", taskID).Scan(&t.ID, &t.CardID, &t.Title, &t.Description, &t.Start, &t.Percent, &t.Deadline, &t.Completed); err != nil {
	}

	fmt.Print("GOOD")

	return t

}

func Add(cardID int, title string, description string, start string, deadline string, percent int) Task {
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

	t := Task{}

	if err := db.DB.QueryRow(query).Scan(&t.ID, &t.CardID, &t.Title, &t.Description, &t.Start, &t.Percent, &t.Deadline, &t.Completed); err != nil {
	}

	fmt.Print("GOOD")

	return t
}

func Remove(taskID int) (bool, error) {
	result, err := db.DB.Exec("DELETE FROM task WHERE id=$1", taskID)
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

func ChangeTitle(taskID int, title string) Task {

	t := Task{}

	if err := db.DB.QueryRow("UPDATE task SET title='$1' WHERE id=$2 RETURNING *", title, taskID).Scan(&t.ID, &t.CardID, &t.Title, &t.Description, &t.Start, &t.Percent, &t.Deadline, &t.Completed); err != nil {
	}

	fmt.Print("GOOD")

	return t
}

func ChangeDescription(taskID int, description string) Task {

	t := Task{}

	if err := db.DB.QueryRow("UPDATE task SET description='$1' WHERE id=$2 RETURNING *", description, taskID).Scan(&t.ID, &t.CardID, &t.Title, &t.Description, &t.Start, &t.Percent, &t.Deadline, &t.Completed); err != nil {
	}

	fmt.Print("GOOD")

	return t
}

func ChangeStart(taskID int, start string) Task {

	t := Task{}

	if err := db.DB.QueryRow("UPDATE task SET start='$1' WHERE id=$2 RETURNING *", start, taskID).Scan(&t.ID, &t.CardID, &t.Title, &t.Description, &t.Start, &t.Percent, &t.Deadline, &t.Completed); err != nil {
	}

	fmt.Print("GOOD")

	return t
}

func ChangeDeadline(taskID int, deadline string) Task {

	t := Task{}

	if err := db.DB.QueryRow("UPDATE task SET deadline='$1' WHERE id=$2 RETURNING *", deadline, taskID).Scan(&t.ID, &t.CardID, &t.Title, &t.Description, &t.Start, &t.Percent, &t.Deadline, &t.Completed); err != nil {
	}

	fmt.Print("GOOD")

	return t
}

func ChangeComplete(taskID int, status bool) Task {

	t := Task{}

	if err := db.DB.QueryRow("UPDATE task SET completed=$1 WHERE id=$2 RETURNING *", status, taskID).Scan(&t.ID, &t.CardID, &t.Title, &t.Description, &t.Start, &t.Percent, &t.Deadline, &t.Completed); err != nil {
	}

	fmt.Print("GOOD")

	return t
}

func ChangeCard(taskID int, cardID int) Task {

	t := Task{}

	if err := db.DB.QueryRow("UPDATE task SET card_id=$1 WHERE id=$2 RETURNING *", cardID, taskID).Scan(&t.ID, &t.CardID, &t.Title, &t.Description, &t.Start, &t.Percent, &t.Deadline, &t.Completed); err != nil {
	}

	fmt.Print("GOOD")

	return t
}

func ChangePercent(taskID int, percent int) Task {

	t := Task{}

	if err := db.DB.QueryRow("UPDATE task SET percent=$1 WHERE id=$2 RETURNING *", percent, taskID).Scan(&t.ID, &t.CardID, &t.Title, &t.Description, &t.Start, &t.Percent, &t.Deadline, &t.Completed); err != nil {
	}

	fmt.Print("GOOD")

	return t
}
