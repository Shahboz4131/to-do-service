package repo

import (
	"time"

	pb "github.com/Shahboz4131/to-do-service/genproto"
)

// TaskStorageI ...
type TaskStorageI interface {
	Create(pb.Task) (pb.Task, error)
	Get(id string) (pb.Task, error)
	List(page, limit int64) ([]*pb.Task, int64, error)
	Update(pb.Task) (pb.Task, error)
	Delete(id string) error
	Overdue(timed time.Time, page, limit int64) ([]*pb.Task, int64, error)
}
