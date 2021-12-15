package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Shahboz4131/to-do-service/genproto"
	l "github.com/Shahboz4131/to-do-service/pkg/logger"
	"github.com/Shahboz4131/to-do-service/storage"
)

// UserService is an object that implements user interface.
// type UserService struct {
// 	storage storage.IStorage
// 	logger  l.Logger
// }

type TaskService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewUserService ...
// func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
// 	return &UserService{
// 		storage: storage.NewStoragePg(db),
// 		logger:  log,
// 	}
// }

// NewTaskService ...
func NewTaskService(storage storage.IStorage, log l.Logger) *TaskService {
	return &TaskService{
		storage: storage,
		logger:  log,
	}
}

// func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
// 	user, err := s.storage.User().Create(*req)
// 	if err != nil {
// 		s.logger.Error("failed to create user", l.Error(err))
// 		return nil, status.Error(codes.Internal, "failed to create user")
// 	}

// 	return &user, nil
// }

func (s *TaskService) Create(ctx context.Context, req *pb.Task) (*pb.Task, error) {
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

// func (s *UserService) List(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {
// 	users, count, err := s.storage.User().List(req.Page, req.Limit)
// 	if err != nil {
// 		s.logger.Error("failed to list users", l.Error(err))
// 		return nil, status.Error(codes.Internal, "failed to list users")
// 	}

// 	return &pb.ListResp{
// 		Users: users,
// 		Count: count,
// 	}, nil
// }

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

// func (s *UserService) Update(ctx context.Context, req *pb.User) (*pb.User, error) {
// 	user, err := s.storage.User().Update(*req)
// 	if err != nil {
// 		s.logger.Error("failed to update user", l.Error(err))
// 		return nil, status.Error(codes.Internal, "failed to update user")
// 	}

// 	return &user, nil
// }

func (s *TaskService) Update(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	task, err := s.storage.Task().Update(*req)
	if err != nil {
		s.logger.Error("failed to update task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update task")
	}

	return &task, nil
}

// func (s *UserService) Delete(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
// 	err := s.storage.User().Delete(req.Id)
// 	if err != nil {
// 		s.logger.Error("failed to delete user", l.Error(err))
// 		return nil, status.Error(codes.Internal, "failed to delete user")
// 	}

// 	return &pb.EmptyResp{}, nil
// }

func (s *TaskService) Delete(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyRes, error) {
	err := s.storage.Task().Delete(req.Id)
	if err != nil {
		s.logger.Error("failed to delete task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete task")
	}

	return &pb.EmptyRes{}, nil
}

func (s *TaskService) Overdue(ctx context.Context, req *pb.OverdueReq) (*pb.OverdueResp, error) {
	tasks, count, err := s.storage.Task().Overdue(req.Timed, req.Limit, req.Page)
	if err != nil {
		s.logger.Error("failed to list tasks", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to list tasks")
	}

	return &pb.OverdueResp{
		Overres: tasks,
		Count:   count,
	}, nil
}
