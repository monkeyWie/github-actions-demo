package service

import "testing"

func TestUser_GetName(t *testing.T) {
	type fields struct {
		name string
		age  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test",
			fields: fields{
				name: "admin",
				age:  18,
			},
			want: "admin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				name: tt.fields.name,
				age:  tt.fields.age,
			}
			if got := user.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
