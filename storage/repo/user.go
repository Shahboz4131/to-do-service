package repo

import (
	pb "github.com/Shahboz4131/to-do-service/genproto"
)

// TaskStorageI ...
type TaskStorageI interface {
	Create(pb.Task) (pb.Task, error)
	Get(id int64) (pb.Task, error)
	List(page, limit int64) ([]*pb.Task, int64, error)
	Update(pb.Task) (pb.Task, error)
	Delete(id int64) error
	Overdue(timed string, page, limit int64) ([]*pb.Task, int64, error)
}
