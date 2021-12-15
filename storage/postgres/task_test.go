package postgres

import (
	"reflect"
	"testing"

	pb "github.com/Shahboz4131/to-do-service/genproto"
)

// func TestTaskRepo_Create(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		input   pb.Task
// 		want    pb.Task
// 		wantErr bool
// 	}{
// 		{
// 			name: "successful",
// 			input: pb.Task{
// 				Assignee: "Rustam",
// 				Title:    "Turgunov",
// 				Summary:  "summmary",
// 				Deadline: "2020-10-10",
// 				Status:   "active",
// 			},
// 			want: pb.Task{
// 				Assignee: "Rustam",
// 				Title:    "Turgunov",
// 				Summary:  "summmary",
// 				Deadline: "2020-10-10",
// 				Status:   "active",
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			got, err := pgRepo.Create(tc.input)
// 			if err != nil {
// 				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
// 			}
// 			got.Id = 0
// 			// got.Deadline = got.Deadline[:len(got.Deadline)-10]
// 			if !reflect.DeepEqual(tc.want, got) {
// 				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
// 			}
// 		})
// 	}
// }

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
			// tc.want.Id = got.Id
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

// func TestTaskRepo_List(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		inputPage   int64
// 		inputLimit   int64
// 		want1    []*pb.Task
// 		want2	int64
// 		wantErr bool
// 	}{
// 		{
// 			name: "successful",
// 			inputPage:  2,
// 			inputLimit:  2,
// 			want1: []pb.Task{
// 				Assignee: "Rustam",
// 				Title:    "Turgunov",
// 				Summary:  "summmary",
// 				Deadline: "2020-10-10",
// 				Status:   "active",
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			got, err := pgRepo.Get(input)
// 			if err != nil {
// 				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.wantErr, err)
// 			}
// 			// got.Id = 0
// 			// got.Deadline = got.Deadline[:len(got.Deadline)-10]
// 			if !reflect.DeepEqual(tc.want, got) {
// 				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
// 			}
// 		})
// 	}
// }
// func TestUserRepo_Get(t *testing.T) {

// }

// func TestUserRepo_List(t *testing.T) {

// }
