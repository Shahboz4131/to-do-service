package postgres

import (
	"reflect"
	"testing"

	pb "github.com/Shahboz4131/to-do-service/genproto"
)

func TestTaskRepo_Create(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.Task
		want    pb.Task
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.Task{
				Assignee: "Ructam",
				Title:    "Turgunov",
				Summary:  "summmary",
				Deadline: "2020-10-10",
				Status:   "active",
			},
			want: pb.Task{
				Assignee: "Ructam",
				Title:    "Turgunov",
				Summary:  "summmary",
				Deadline: "2020-10-10T00:00:00Z",
				Status:   "active",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Create(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
			got.Id = 0
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Get(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		want    pb.Task
		wantErr bool
	}{
		{
			name:  "successful",
			input: 8,
			want: pb.Task{
				Id:       8,
				Assignee: "Rustam",
				Title:    "Turgunov",
				Summary:  "summmary",
				Deadline: "2020-10-10T00:00:00Z",
				Status:   "active",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Get(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Update(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.Task
		want    pb.Task
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.Task{
				Id:       3,
				Assignee: "Rustan",
				Title:    "turgunov",
				Summary:  "summmary",
				Deadline: "2020-10-10",
				Status:   "active",
			},
			want: pb.Task{
				Id:       3,
				Assignee: "Rustan",
				Title:    "turgunov",
				Summary:  "summmary",
				Deadline: "2020-10-10T00:00:00Z",
				Status:   "active",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := pgRepo.Update(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Delete(t *testing.T) {
	tests := []struct {
		name    string
		input   int64
		wantErr bool
	}{
		{
			name:    "successful",
			input:   22,
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := pgRepo.Delete(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
		})
	}
}

func TestRepo_List(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.ListReq
		want    pb.ListResp
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.ListReq{
				Page:  2,
				Limit: 2,
			},
			want: pb.ListResp{
				Tasks: []*pb.Task{
					{
						Id:       8,
						Assignee: "Rustam",
						Title:    "Turgunov",
						Summary:  "summmary",
						Deadline: "2020-10-10T00:00:00Z",
						Status:   "active",
					},
					{
						Id:       11,
						Assignee: "Rustam",
						Title:    "Turgunov",
						Summary:  "summmary",
						Deadline: "2020-10-10T00:00:00Z",
						Status:   "active",
					},
				},
				Count: 47,
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, count, err := pgRepo.List(tc.input.Page, tc.input.Limit)
			tc.want.Count = count
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
			if !reflect.DeepEqual(tc.want, pb.ListResp{
				Tasks: got,
				Count: count,
			}) {
				t.Fatalf("%s: expected: %v, got: %v ", tc.name, tc.want, got)
			}
		})
	}
}

func TestRepo_Overdue(t *testing.T) {
	tests := []struct {
		name    string
		input   pb.OverdueReq
		want    pb.OverdueResp
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.OverdueReq{
				Timed: "2021-10-09",
				Page:  2,
				Limit: 2,
			},
			want: pb.OverdueResp{
				Overres: []*pb.Task{
					{
						Id:       8,
						Assignee: "Rustam",
						Title:    "Turgunov",
						Summary:  "summmary",
						Deadline: "2020-10-10T00:00:00Z",
						Status:   "active",
					},
					{
						Id:       11,
						Assignee: "Rustam",
						Title:    "Turgunov",
						Summary:  "summmary",
						Deadline: "2020-10-10T00:00:00Z",
						Status:   "active",
					},
				},
				Count: 47,
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, count, err := pgRepo.Overdue(tc.input.Timed, tc.input.Page, tc.input.Limit)
			tc.want.Count = count
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}
			if !reflect.DeepEqual(tc.want, pb.OverdueResp{
				Overres: got,
				Count:   count,
			}) {
				t.Fatalf("%s: expected: %v, got: %v ", tc.name, tc.want, got)
			}
		})
	}
}
