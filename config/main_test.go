package config

import "testing"

func TestConfig_ShouldIgnoreId(t *testing.T) {
	type fields struct {
		IgnoreIDs []int64
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "One ID provided and found",
			fields: fields{IgnoreIDs: []int64{123}},
			args:   args{id: int64(123)},
			want:   true,
		},
		{
			name:   "One ID provided but not found",
			fields: fields{IgnoreIDs: []int64{123}},
			args:   args{id: int64(321)},
			want:   false,
		},
		{
			name:   "Multiple IDs provided but not found",
			fields: fields{IgnoreIDs: []int64{123, 456}},
			args:   args{id: int64(321)},
			want:   false,
		},
		{
			name:   "Multiple IDs provided and one found",
			fields: fields{IgnoreIDs: []int64{123, 456}},
			args:   args{id: int64(456)},
			want:   true,
		},
		{
			name:   "No IDs provided",
			fields: fields{IgnoreIDs: nil},
			args:   args{id: int64(456)},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				IgnoreIDs: tt.fields.IgnoreIDs,
			}
			if got := c.ShouldIgnoreId(tt.args.id); got != tt.want {
				t.Errorf("Config.ShouldIgnoreId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_ShouldIgnoreOnSubstrings(t *testing.T) {
	type fields struct {
		IgnoreSubstrings []string
	}
	type args struct {
		body string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "No substrings provided",
			args: args{body: "Test body"},
			want: false,
		},
		{
			name:   "One substring provided (matching)",
			fields: fields{IgnoreSubstrings: []string{"bod"}},
			args:   args{body: "Test body"},
			want:   true,
		},
		{
			name:   "One substring provided (not matching)",
			fields: fields{IgnoreSubstrings: []string{"bodi"}},
			args:   args{body: "Test body"},
			want:   false,
		},
		{
			name:   "Multiple substring provided (matching)",
			fields: fields{IgnoreSubstrings: []string{"bodi", "foo", "bar", "bod"}},
			args:   args{body: "Test body"},
			want:   true,
		},
		{
			name:   "Multiple substring provided (not matching)",
			fields: fields{IgnoreSubstrings: []string{"bodi", "foo", "bar", "2bod"}},
			args:   args{body: "Test body"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				IgnoreSubstrings: tt.fields.IgnoreSubstrings,
			}
			if got := c.ShouldIgnoreOnSubstrings(tt.args.body); got != tt.want {
				t.Errorf("Config.ShouldIgnoreOnSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
