package postgres

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	pb "github.com/Shahboz4131/to-do-service/genproto"
)

type taskRepo struct {
	db *sqlx.DB
}

// NewTaskRepo ...
func NewTaskRepo(db *sqlx.DB) *taskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(task pb.Task) (pb.Task, error) {
	var id string
	// var uid uuid.UUID
	err := r.db.QueryRow(`
        INSERT INTO tasks(id, assignee, title, summary, deadline, status, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8) returning id`, task.Id, task.Assignee, task.Title, task.Summary, task.Deadline, task.Status, time.Now().UTC(), time.Now().UTC()).Scan(&id)
	if err != nil {
		return pb.Task{}, err
	}

	task, err = r.Get(id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Get(id string) (pb.Task, error) {
	var task pb.Task

	err := r.db.QueryRow(`
        SELECT id::varchar, assignee, title, summary, deadline, status, created_at, updated_at FROM tasks
        WHERE id=$1 and deleted_at is  null`, id).Scan(&task.Id, &task.Assignee, &task.Title, &task.Summary, &task.Deadline, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) List(page, limit int64) ([]*pb.Task, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(
		`SELECT id, assignee, title, summary, deadline, status, created_at, updated_at FROM tasks where deleted_at is  null OFFSET $1  LIMIT $2`,
		offset, limit)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() // nolint:errcheck

	var (
		tasks []*pb.Task
		count int64
	)

	for rows.Next() {
		var task pb.Task
		err = rows.Scan(&task.Id, &task.Assignee, &task.Title, &task.Summary, &task.Deadline, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}

		tasks = append(tasks, &task)
	}

	err = r.db.QueryRow(`SELECT count(id) FROM tasks where deleted_at is  null`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	return tasks, count, nil
}

func (r *taskRepo) Update(task pb.Task) (pb.Task, error) {
	result, err := r.db.Exec(`UPDATE tasks SET assignee=$2, title=$3, summary=$4, deadline=$5, status=$6, updated_at=$7 WHERE id=$1 and deleted_at is  null`,
		task.Id, task.Assignee, task.Title, task.Summary, task.Deadline, task.Status, time.Now().UTC())
	if err != nil {
		return pb.Task{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Task{}, sql.ErrNoRows
	}

	task, err = r.Get(task.Id)
	if err != nil {
		return pb.Task{}, err
	}

	return task, nil
}

func (r *taskRepo) Delete(id string) error {
	result, err := r.db.Exec(`UPDATE tasks SET deleted_at = $2 where id = $1 and deleted_at is null `, id, time.Now().UTC())
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *taskRepo) Overdue(t time.Time, limit, page int64) ([]*pb.Task, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Queryx(
		`SELECT id, assignee, title, summary, deadline, status, created_at, updated_at FROM tasks where deadline < $1 and deleted_at is null  LIMIT $2 OFFSET $3`,
		t, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() // nolint:errcheck

	var (
		tasks []*pb.Task
		count int64
	)

	for rows.Next() {
		var task pb.Task
		err = rows.Scan(&task.Id, &task.Assignee, &task.Title, &task.Summary, &task.Deadline, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}

		tasks = append(tasks, &task)
	}

	err = r.db.QueryRow(`SELECT count(id) FROM tasks where deadline < $1 and deleted_at is null`, t).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return tasks, count, nil
}
