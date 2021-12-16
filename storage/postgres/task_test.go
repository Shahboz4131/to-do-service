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
				Assignee: "Rustam",
				Title:    "Turgunov",
				Summary:  "summmary",
				Deadline: "2020-10-10",
				Status:   "active",
			},
			want: pb.Task{
				Assignee: "Rustam",
				Title:    "Turgunov",
				Summary:  "summmary",
				Deadline: "2020-10-10T00:00:00Z",
				Status:   "active",
			},
			wantErr: false,
		},
		{
			name: "unsuccessful",
			input: pb.Task{
				Assignee: "Rustammmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm",
				Title:    "Turgunov",
				Summary:  "summmary",
				Deadline: "2020-10-10",
				Status:   "active",
			},
			want: pb.Task{
				Assignee: "Rustammmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm",
				Title:    "Turgunov",
				Summary:  "summmary",
				Deadline: "2020-10-10T00:00:00Z",
				Status:   "active",
			},
			wantErr: true,
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
			input: 13,
			want: pb.Task{
				Id:       13,
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
				Id:       4,
				Assignee: "Rustam",
				Title:    "Turgunov",
				Summary:  "summmary",
				Deadline: "2020-10-10",
				Status:   "active",
			},
			want: pb.Task{
				Id:       4,
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
		name  string
		input int64
		// want    pb.EmptyRes
		wantErr bool
	}{
		{
			name:  "successful",
			input: 1,
			// want:    pb.EmptyRes{},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := pgRepo.Delete(tc.input)
			if err != nil {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
			}

			// if !reflect.DeepEqual(tc.want, got) {
			// 	t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			// }
		})
	}
}
