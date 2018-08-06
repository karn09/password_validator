package loader

import (
	"reflect"
	"testing"
)

func Test_commonPasswords_SortCommon(t *testing.T) {
	type fields struct {
		list []string
	}
	tests := []struct {
		name   string
		fields fields
		want   *commonPasswords
	}{
		{
			name:   "sort a short list",
			fields: fields{list: []string{"b", "c", "d", "a"}},
			want:   &commonPasswords{list: []string{"a", "b", "c", "d"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &commonPasswords{
				list: tt.fields.list,
			}
			if got := s.sortCommon(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonPasswords.sortCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadCommon(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *commonPasswords
		wantErr bool
	}{
		{
			name: "load example success",
			args: args{path: "./loader_test_ex.txt"},
			want: &commonPasswords{list: []string{
				"zxcd",
				"reww",
				"54what",
				"how",
			}},
			wantErr: false,
		},
		{
			name:    "load example fails",
			args:    args{path: "./i_dont_exist.txt"},
			want:    &commonPasswords{list: []string{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadCommon(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadCommon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commonPasswords_isCommon(t *testing.T) {
	type fields struct {
		list []string
	}
	type args struct {
		pass string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "true if string is found",
			fields: fields{list: []string{
				"zxcd",
				"reww",
				"54what",
				"how",
			}},
			args: args{pass: "how"},
			want: true,
		},
		{
			name: "false if string is not found",
			fields: fields{list: []string{
				"zxcd",
				"reww",
				"54what",
				"how",
			}},
			args: args{pass: "123"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &commonPasswords{
				list: tt.fields.list,
			}
			if got := s.isCommon(tt.args.pass); got != tt.want {
				t.Errorf("commonPasswords.isCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commonPasswords_isUnderMinimumLength(t *testing.T) {
	type fields struct {
		list []string
	}
	type args struct {
		pass []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "true when passed string as byte array below min length (8)",
			fields: fields{list: []string{}},
			args:   args{pass: []byte("short")},
			want:   true,
		},
		{
			name:   "false when passed string as byte array and meets min length (8)",
			fields: fields{list: []string{}},
			args:   args{pass: []byte("justright")},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &commonPasswords{
				list: tt.fields.list,
			}
			if got := s.isUnderMinimumLength(tt.args.pass); got != tt.want {
				t.Errorf("commonPasswords.isUnderMinimumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commonPasswords_isOverMaximumLength(t *testing.T) {
	type fields struct {
		list []string
	}
	type args struct {
		pass []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "false when passed string as byte array and is not greater than max length (64)",
			fields: fields{list: []string{}},
			args:   args{pass: []byte("justright")},
			want:   false,
		},
		{
			name:   "true when greater than max length",
			fields: fields{list: []string{}},
			args:   args{pass: []byte("lllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllll")},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &commonPasswords{
				list: tt.fields.list,
			}
			if got := s.isOverMaximumLength(tt.args.pass); got != tt.want {
				t.Errorf("commonPasswords.isOverMaximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commonPasswords_containsInvalidASCII(t *testing.T) {
	type fields struct {
		list []string
	}
	type args struct {
		pass []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  []rune
	}{
		{
			name:   "false when containing all valid characters",
			fields: fields{list: []string{}},
			args:   args{pass: []byte("short")},
			want:   false,
			want1:  []rune("short"),
		},
		{
			name:   "true when containing any invalid characters",
			fields: fields{list: []string{}},
			args:   args{pass: []byte("£short")},
			want:   true,
			want1:  []rune("*short"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &commonPasswords{
				list: tt.fields.list,
			}
			got, got1 := s.containsInvalidASCII(tt.args.pass)
			if got != tt.want {
				t.Errorf("commonPasswords.containsInvalidASCII() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("commonPasswords.containsInvalidASCII() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_commonPasswords_IsValid(t *testing.T) {
	type fields struct {
		list []string
	}
	type args struct {
		pass []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "not valid when containing invalid ascii chars.",
			fields:  fields{list: []string{}},
			args:    args{pass: []byte("£short")},
			want:    "*short",
			wantErr: true,
		},
		{
			name:    "not valid when password over 64 chars.",
			fields:  fields{list: []string{}},
			args:    args{pass: []byte("lllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllll")},
			want:    "lllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllll",
			wantErr: true,
		},
		{
			name:    "not valid when password under 8 chars.",
			fields:  fields{list: []string{}},
			args:    args{pass: []byte("short")},
			want:    "short",
			wantErr: true,
		},
		{
			name:    "valid password.",
			fields:  fields{list: []string{}},
			args:    args{pass: []byte("zxlkas12s")},
			want:    "zxlkas12s",
			wantErr: false,
		},
		{
			name:    "common password.",
			fields:  fields{list: []string{"password"}},
			args:    args{pass: []byte("password")},
			want:    "password",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &commonPasswords{
				list: tt.fields.list,
			}
			got, err := s.IsValid(tt.args.pass)
			if (err != nil) != tt.wantErr {
				t.Errorf("commonPasswords.IsValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("commonPasswords.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
