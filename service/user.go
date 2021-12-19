package service

import (
	"context"
	"time"

	pb "github.com/Shahboz4131/to-do-service/genproto"
	l "github.com/Shahboz4131/to-do-service/pkg/logger"
	"github.com/Shahboz4131/to-do-service/storage"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewTaskService ...
func NewTaskService(storage storage.IStorage, log l.Logger) *TaskService {
	return &TaskService{
		storage: storage,
		logger:  log,
	}
}

func (s *TaskService) Create(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	id, err := uuid.NewV4()
	if err != nil {
		s.logger.Error("failed while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}
	layout := "2006-01-02"
	_, err = time.Parse(layout, req.Deadline)

	if err != nil {
		s.logger.Error("failed while parsing deadlime", l.Error(err))
		return nil, status.Error(codes.Internal, "failed parse deadline")
	}

	req.Id = id.String()

	task, err := s.storage.Task().Create(*req)
	if err != nil {
		s.logger.Error("failed to create task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create task")
	}

	return &task, nil
}

func (s *TaskService) Get(ctx context.Context, req *pb.ByIdReq) (*pb.Task, error) {
	task, err := s.storage.Task().Get(req.GetId())
	if err != nil {
		s.logger.Error("failed to get task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get task")
	}

	return &task, nil
}

func (s *TaskService) List(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {
	tasks, count, err := s.storage.Task().List(req.Page, req.Limit)
	if err != nil {
		s.logger.Error("failed to list tasks", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to list tasks")
	}

	return &pb.ListResp{
		Tasks: tasks,
		Count: count,
	}, nil
}

func (s *TaskService) Update(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	task, err := s.storage.Task().Update(*req)
	if err != nil {
		s.logger.Error("failed to update task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update task")
	}

	return &task, nil
}

func (s *TaskService) Delete(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyRes, error) {
	err := s.storage.Task().Delete(req.Id)
	if err != nil {
		s.logger.Error("failed to delete task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete task")
	}

	return &pb.EmptyRes{}, nil
}

func (s *TaskService) Overdue(ctx context.Context, req *pb.OverdueReq) (*pb.OverdueResp, error) {
	layout := "2006-01-02"

	deadline, err := time.Parse(layout, req.Timed)
	if err != nil {
		s.logger.Error("failed while parsing deadlime", l.Error(err))
		return nil, status.Error(codes.Internal, "failed parse deadline")
	}

	tasks, count, err := s.storage.Task().Overdue(deadline, req.Limit, req.Page)
	if err != nil {
		s.logger.Error("failed to list tasks", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to list tasks")
	}

	return &pb.OverdueResp{
		Overres: tasks,
		Count:   count,
	}, nil
}
