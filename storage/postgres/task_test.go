package postgres

import (
	"reflect"
	"testing"
	"time"

	pb "github.com/Shahboz4131/to-do-service/genproto"
)

func TestTaskRepo_Create(t *testing.T) {
	tests := []struct {
		id      string
		name    string
		input   pb.Task
		want    pb.Task
		wantErr bool
	}{
		{
			name: "successful",
			input: pb.Task{
				Id:       "83cefcca-1c83-4337-93a3-35baf7a88cfd",
				Assignee: "Ructam",
				Title:    "Turgunov",
				Summary:  "summmary",
				Deadline: "2020-10-10",
				Status:   "active",
			},
			want: pb.Task{
				Id:        "83cefcca-1c83-4337-93a3-35baf7a88cfd",
				Assignee:  "Ructam",
				Title:     "Turgunov",
				Summary:   "summmary",
				Deadline:  "2020-10-10T00:00:00Z",
				Status:    "active",
				CreatedAt: "",
				UpdatedAt: "",
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

			got.CreatedAt = ""
			got.UpdatedAt = ""
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Get(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    pb.Task
		wantErr bool
	}{
		{
			name:  "successful",
			input: "83cefcca-1c83-4337-93a3-35baf7a88cfd",
			want: pb.Task{
				Id:        "83cefcca-1c83-4337-93a3-35baf7a88cfd",
				Assignee:  "Ructam",
				Title:     "Turgunov",
				Summary:   "summmary",
				Deadline:  "2020-10-10T00:00:00Z",
				Status:    "active",
				CreatedAt: "2021-12-19T05:16:03.898394Z",
				UpdatedAt: "2021-12-19T05:16:03.898394Z",
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
				Id:        "83cefcca-1c83-4337-93a3-35baf7a88cfd",
				Assignee:  "Ructam",
				Title:     "Turgunov",
				Summary:   "summmary",
				Deadline:  "2020-10-10T00:00:00Z",
				Status:    "actived",
				CreatedAt: "2021-12-19T05:16:03.898394Z",
				UpdatedAt: "",
			},
			want: pb.Task{
				Id:        "83cefcca-1c83-4337-93a3-35baf7a88cfd",
				Assignee:  "Ructam",
				Title:     "Turgunov",
				Summary:   "summmary",
				Deadline:  "2020-10-10T00:00:00Z",
				Status:    "actived",
				CreatedAt: "2021-12-19T05:16:03.898394Z",
				UpdatedAt: "",
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
			got.UpdatedAt = ""
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestTaskRepo_Delete(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "successful",
			input:   "83cefcca-1c83-4337-93a3-35baf7a88cfd",
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
				Page:  1,
				Limit: 1,
			},
			want: pb.ListResp{
				Tasks: []*pb.Task{
					{
						Id:        "83cefcca-1c83-4337-93a3-35baf7a88cfd",
						Assignee:  "Ructam",
						Title:     "Turgunov",
						Summary:   "summmary",
						Deadline:  "2020-10-10T00:00:00Z",
						Status:    "actived",
						CreatedAt: "2021-12-19T05:16:03.898394Z",
						UpdatedAt: "2021-12-19T05:48:14.898872Z",
					},
				},
				Count: 1,
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
				Timed: "2022-10-10",
				Page:  1,
				Limit: 1,
			},
			want: pb.OverdueResp{
				Overres: []*pb.Task{
					{
						Id:        "83cefcca-1c83-4337-93a3-35baf7a88cfd",
						Assignee:  "Ructam",
						Title:     "Turgunov",
						Summary:   "summmary",
						Deadline:  "2020-10-10T00:00:00Z",
						Status:    "actived",
						CreatedAt: "2021-12-19T05:16:03.898394Z",
						UpdatedAt: "2021-12-19T05:48:14.898872Z",
					},
				},
				Count: 1,
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {

		layout := "2006-01-02"

		tim, err := time.Parse(layout, tc.input.Timed)
		if err != nil {
			t.Fatalf("time parsing error %v", err)
		}

		t.Run(tc.name, func(t *testing.T) {
			got, count, err := pgRepo.Overdue(tim, tc.input.Page, tc.input.Limit)
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
