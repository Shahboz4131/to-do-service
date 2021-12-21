package service

import (
	"context"
	"log"
	"reflect"
	"testing"

	pb "github.com/Shahboz4131/to-do-service/genproto"
)

func TestTaskService_Create(t *testing.T) {
	tests := []struct {
		name  string
		input pb.Task
		want  pb.Task
	}{
		{
			name: "successful",
			input: pb.Task{
				Assignee: "assignee",
				Title:    "title",
				Summary:  "summary",
				Deadline: "2000-10-10",
				Status:   "status",
			},
			want: pb.Task{
				Assignee:  "assignee",
				Title:     "title",
				Summary:   "summary",
				Deadline:  "2000-10-10T00:00:00Z",
				Status:    "status",
				CreatedAt: "",
				UpdatedAt: "",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Create(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to create task", err)
			}
			tc.want.Id = got.Id
			tc.want.CreatedAt = got.CreatedAt
			tc.want.UpdatedAt = got.UpdatedAt
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskService_Get(t *testing.T) {
	for_get := pb.Task{
		Assignee: "assignee",
		Title:    "title",
		Summary:  "summary",
		Deadline: "2020-10-10",
		Status:   "status",
	}

	get_got, err := client.Create(context.Background(), &for_get)
	if err != nil {
		log.Fatalf("Did not create task %v", err)
	}

	tests := []struct {
		name  string
		input pb.ByIdReq
		want  pb.Task
	}{
		{
			name: "successfull",
			input: pb.ByIdReq{
				Id: get_got.Id,
			},
			want: *get_got,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Get(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to get task", err)
			}
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskService_Update(t *testing.T) {
	for_get := pb.Task{
		Assignee: "assignee",
		Title:    "title",
		Summary:  "summary",
		Deadline: "2020-10-10",
		Status:   "status",
	}

	for_got, err := client.Create(context.Background(), &for_get)
	if err != nil {
		log.Fatalf("Did not create task %v", err)
	}

	tests := []struct {
		name  string
		input pb.Task
	}{
		{
			name:  "successfull",
			input: *for_got,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := client.Update(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to update task", err)
			}
		})
	}
}

func TestTaskService_Delete(t *testing.T) {
	for_get := pb.Task{
		Assignee: "assignee",
		Title:    "title",
		Summary:  "summary",
		Deadline: "2020-10-10",
		Status:   "status",
	}

	for_got, err := client.Create(context.Background(), &for_get)
	if err != nil {
		log.Fatalf("Did not create task %v", err)
	}

	tests := []struct {
		name  string
		input pb.ByIdReq
	}{
		{
			name: "successfull",
			input: pb.ByIdReq{
				Id: for_got.Id,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := client.Delete(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to delete task", err)
			}
		})
	}
}

func TestTaskService_List(t *testing.T) {
	tests := []struct {
		name  string
		input pb.ListReq
		want  pb.ListResp
	}{
		{
			name: "successfull",
			input: pb.ListReq{
				Page:  2,
				Limit: 1,
			},
			want: pb.ListResp{
				Tasks: []*pb.Task{
					{
						Id:        "4e16d07e-9ad9-4c38-867d-44d46f6fef7e",
						Assignee:  "assignee",
						Title:     "title",
						Summary:   "summary",
						Deadline:  "2020-10-10T00:00:00Z",
						Status:    "status",
						CreatedAt: "2021-12-21T05:11:51.265358Z",
						UpdatedAt: "2021-12-21T05:11:51.265358Z",
					},
				},
				Count: 14,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.List(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to delete task", err)
			}

			if !reflect.DeepEqual(tc.want, pb.ListResp{
				Tasks: got.Tasks,
				Count: got.Count,
			}) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskService_Overdue(t *testing.T) {
	tests := []struct {
		name  string
		input pb.OverdueReq
		want  pb.OverdueResp
	}{
		{
			name: "successfull",
			input: pb.OverdueReq{
				Timed: "2019-10-10",
				Page:  1,
				Limit: 1,
			},
			want: pb.OverdueResp{
				Overres: []*pb.Task{
					{
						Id:        "ca8ab028-300b-47f2-981d-eb013c6246f3",
						Assignee:  "assignee",
						Title:     "title",
						Summary:   "summary",
						Deadline:  "2000-10-10T00:00:00Z",
						Status:    "status",
						CreatedAt: "2021-12-21T06:43:38.865676Z",
						UpdatedAt: "2021-12-21T06:43:38.865676Z",
					},
				},
				Count: 1,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := client.Overdue(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to delete task", err)
			}

			if !reflect.DeepEqual(tc.want, pb.OverdueResp{
				Overres: got.Overres,
				Count:   got.Count,
			}) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}
