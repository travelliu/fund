// Copyright © 2020 Bin Liu <bin.liu@enmotech.com>

package databases

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestTimeInt64_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		t       TimeInt64
		want    []byte
		wantErr bool
	}{
		{
			name: "test",
			t:    TimeInt64(time.Unix(1594196275, 0)),
			want: []byte("\"1594196275000\""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got, err := tt.t.MarshalJSON()
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			// }
			// fmt.Println(string(got))
		})
	}
}

func TestTimeInt64_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		t       TimeInt64
		args    args
		wantErr bool
	}{
		{
			name:    "test",
			args:    args{b: []byte("1594196275000")},
			wantErr: false,
			t:       TimeInt64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModel_MarshalJSON(t *testing.T) {
	var (
	// deleteTime = TimeInt64(time.Unix(1594196275, 0))
	)
	type args struct {
		m Model
	}
	tests := []struct {
		name    string
		t       TimeInt64
		args    args
		wantErr bool
		want    string
	}{
		// {
		// 	name: "Normal",
		// 	args: args{m: Model{
		// 		ID:        11111111111,
		// 		CreatedAt: TimeInt64(time.Unix(1594196275, 0)),
		// 		UpdatedAt: TimeInt64(time.Unix(1594196275, 0)),
		// 		DeletedAt: &deleteTime,
		// 	}},
		// 	wantErr: false,
		// 	want:    "{\"id\":\"11111111111\",\"CreatedAt\":\"1594196275000\",\"UpdatedAt\":\"1594196275000\",\"DeletedAt\":\"1594196275000\"}",
		// },
		{
			name: "DeletedAtIsNull",
			args: args{m: Model{
				ID:        11111111111,
				CreatedAt: TimeInt64(time.Unix(1594196275, 0)),
				UpdatedAt: TimeInt64(time.Unix(1594196275, 0)),
			}},
			wantErr: false,
			want:    "{\"id\":\"11111111111\",\"createdAt\":\"1594196275000\",\"updatedAt\":\"1594196275000\"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(b), tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(b), tt.want)
			}
		})
	}
}
