package postgres

import (
	"testing"

	"github.com/Shahboz4131/to-do-service/config"
	pb "github.com/Shahboz4131/to-do-service/genproto"
	"github.com/Shahboz4131/to-do-service/pkg/db"
	"github.com/Shahboz4131/to-do-service/storage/repo"

	"github.com/stretchr/testify/suite"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	CleanupFunc func()
	Repository  repo.TaskStorageI
}

func (suite *TaskRepositoryTestSuite) SetupSuite() {
	pgPool, cleanup := db.ConnectDBForSuite(config.Load())

	suite.Repository = NewTaskRepo(pgPool)
	suite.CleanupFunc = cleanup
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *TaskRepositoryTestSuite) TestTaskCRUD() {
	id := "0d512776-60ed-4980-b8a3-6904a2234fd4"
	assignee := "aa"

	task := pb.Task{
		Id:       id,
		Assignee: "aa",
		Title:    "tt",
		Summary:  "ss",
		Deadline: "2020-10-10",
		Status:   "st",
	}

	_ = suite.Repository.Delete(id)

	task, err := suite.Repository.Create(task)
	suite.Nil(err)

	getTask, err := suite.Repository.Get(task.Id)
	suite.Nil(err)
	suite.NotNil(getTask, "task must not be nil")
	suite.Equal(assignee, getTask.Assignee, "assignee must match")

	task.Title = "ttt"
	updatedTask, err := suite.Repository.Update(task)
	suite.Nil(err)

	getTask, err = suite.Repository.Get(id)
	suite.Nil(err)
	suite.NotNil(getTask)
	suite.Equal(task.Title, updatedTask.Title)

	listTasks, _, err := suite.Repository.List(1, 10)
	suite.Nil(err)
	suite.NotEmpty(listTasks)
	suite.Equal(task.Title, listTasks[0].Title)

	err = suite.Repository.Delete(id)
	suite.Nil(err)
}

func (suite *TaskRepositoryTestSuite) TearDownSuite() {
	suite.CleanupFunc()
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
