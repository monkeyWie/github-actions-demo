package service

import "testing"

func TestUser_GetName(t *testing.T) {
	type fields struct {
		Name string
		Age  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test",
			fields: fields{
				Name: "admin",
				Age:  18,
			},
			want: "admin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if got := user.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetAge(t *testing.T) {
	type fields struct {
		Name string
		Age  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test",
			fields: fields{
				Name: "admin",
				Age:  18,
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}
			if got := user.GetAge(); got != tt.want {
				t.Errorf("GetAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
