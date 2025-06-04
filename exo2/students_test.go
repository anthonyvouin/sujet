package students

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNewStudent(t *testing.T) {
	type args struct {
		name  string
		age   int
		grade float64
	}
	tests := []struct {
		name    string
		args    args
		want    *Student
		wantErr bool
	}{
		{
			name: "Test",
			args: args{name: "Test", age: 44, grade: 18.5},
			want: &Student{Name: "Test", Age: 44, Grade: 18.5},
		},
		{
			name:    "Empty",
			args:    args{},
			wantErr: true,
		},
		{
			name: "John Doe",
			args: args{name: "DOE John", age: 42, grade: 20},
			want: &Student{Name: "DOE John", Age: 42, Grade: 20},
		},
		{
			name:    "Invalid Age (too low)",
			args:    args{name: "AGE John", age: -10, grade: 20},
			wantErr: true,
		},
		{
			name:    "Invalid Age (too high)",
			args:    args{name: "AGE John", age: 100, grade: 20},
			wantErr: true,
		},
		{
			name: "Min Age (1)",
			args: args{name: "MIN John", age: 1, grade: 13},
			want: &Student{Name: "MIN John", Age: 1, Grade: 13},
		},
		{
			name: "Max Age (99)",
			args: args{name: "MAX John", age: 99, grade: 15},
			want: &Student{Name: "MAX John", Age: 99, Grade: 15},
		},
		{
			name:    "Invalid grade (too high)",
			args:    args{name: "GRADE John", age: 10, grade: 35},
			wantErr: true,
		},
		{
			name:    "Invalid grade (too low)",
			args:    args{name: "GRADE John", age: 10, grade: -5},
			wantErr: true,
		},
		{
			name: "Min grade (0)",
			args: args{name: "MIN John", age: 12, grade: 0},
			want: &Student{Name: "MIN John", Age: 12, Grade: 0},
		},
		{
			name: "Min grade (20)",
			args: args{name: "MAX John", age: 14, grade: 20},
			want: &Student{Name: "MAX John", Age: 14, Grade: 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStudent(tt.args.name, tt.args.age, tt.args.grade)
			if tt.wantErr && err == nil {
				t.Errorf("NewStudent(\"%s\", %d, %.2f) : an error was expected", tt.args.name, tt.args.age, tt.args.grade)
				return
			} else if !tt.wantErr && err != nil {
				t.Errorf("NewStudent(\"%s\", %d, %.2f) : no error was expected", tt.args.name, tt.args.age, tt.args.grade)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStudent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudentList_AddStudents(t *testing.T) {
	type fields struct {
		students []Student
	}
	type args struct {
		students []Student
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Student
	}{
		{
			name: "Four plus one",
			fields: fields{
				students: []Student{
					{Name: "Test1", Age: 44, Grade: 18.5},
					{Name: "Test2", Age: 33, Grade: 20},
					{Name: "Test3", Age: 22, Grade: 14},
					{Name: "Test4", Age: 11, Grade: 16},
				},
			},
			args: args{
				students: []Student{
					{Name: "New", Age: 44, Grade: 18.5},
				},
			},
			want: []Student{
				{Name: "Test1", Age: 44, Grade: 18.5},
				{Name: "Test2", Age: 33, Grade: 20},
				{Name: "Test3", Age: 22, Grade: 14},
				{Name: "Test4", Age: 11, Grade: 16},
				{Name: "New", Age: 44, Grade: 18.5},
			},
		},
		{
			name: "Three plus three",
			fields: fields{
				students: []Student{
					{Name: "Test1", Age: 44, Grade: 18.5},
					{Name: "Test2", Age: 33, Grade: 20},
					{Name: "Test3", Age: 22, Grade: 14},
				},
			},
			args: args{
				students: []Student{
					{Name: "New1", Age: 25, Grade: 13},
					{Name: "New2", Age: 22, Grade: 14},
					{Name: "New3", Age: 29, Grade: 15},
				},
			},
			want: []Student{
				{Name: "Test1", Age: 44, Grade: 18.5},
				{Name: "Test2", Age: 33, Grade: 20},
				{Name: "Test3", Age: 22, Grade: 14},
				{Name: "New1", Age: 25, Grade: 13},
				{Name: "New2", Age: 22, Grade: 14},
				{Name: "New3", Age: 29, Grade: 15},
			},
		},
		{
			name: "Empty plus ten",
			fields: fields{
				students: []Student{},
			},
			args: args{
				students: []Student{
					{Name: "New1", Age: 25, Grade: 13},
					{Name: "New2", Age: 22, Grade: 14},
					{Name: "New3", Age: 29, Grade: 2},
					{Name: "New4", Age: 14, Grade: 15},
					{Name: "New5", Age: 10, Grade: 15},
					{Name: "New6", Age: 19, Grade: 4},
					{Name: "New7", Age: 38, Grade: 15},
					{Name: "New8", Age: 29, Grade: 3},
					{Name: "New9", Age: 40, Grade: 15},
					{Name: "New10", Age: 50, Grade: 9},
				},
			},
			want: []Student{
				{Name: "New1", Age: 25, Grade: 13},
				{Name: "New2", Age: 22, Grade: 14},
				{Name: "New3", Age: 29, Grade: 2},
				{Name: "New4", Age: 14, Grade: 15},
				{Name: "New5", Age: 10, Grade: 15},
				{Name: "New6", Age: 19, Grade: 4},
				{Name: "New7", Age: 38, Grade: 15},
				{Name: "New8", Age: 29, Grade: 3},
				{Name: "New9", Age: 40, Grade: 15},
				{Name: "New10", Age: 50, Grade: 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StudentList{
				students: tt.fields.students,
			}
			s.AddStudents(tt.args.students...)
			if !reflect.DeepEqual(s.students, tt.want) {
				t.Errorf("AddStudents() got = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestStudentList_RemoveStudents(t *testing.T) {
	type fields struct {
		students []Student
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Student
	}{
		{
			name: "Four",
			fields: fields{
				students: []Student{
					{Name: "Test1", Age: 44, Grade: 18.5},
					{Name: "Test2", Age: 33, Grade: 20},
					{Name: "Test3", Age: 22, Grade: 14},
					{Name: "Test4", Age: 11, Grade: 16},
				},
			},
			args: args{
				name: "Test3",
			},
			want: []Student{
				{Name: "Test1", Age: 44, Grade: 18.5},
				{Name: "Test2", Age: 33, Grade: 20},
				{Name: "Test4", Age: 11, Grade: 16},
			},
		},
		{
			name: "Ten",
			fields: fields{
				students: []Student{
					{Name: "New1", Age: 25, Grade: 13},
					{Name: "New2", Age: 22, Grade: 14},
					{Name: "New3", Age: 29, Grade: 2},
					{Name: "New4", Age: 14, Grade: 15},
					{Name: "New5", Age: 10, Grade: 15},
					{Name: "New6", Age: 19, Grade: 4},
					{Name: "New7", Age: 38, Grade: 15},
					{Name: "New8", Age: 29, Grade: 3},
					{Name: "New9", Age: 40, Grade: 15},
					{Name: "New10", Age: 50, Grade: 9},
				},
			},
			args: args{
				name: "New7",
			},
			want: []Student{
				{Name: "New1", Age: 25, Grade: 13},
				{Name: "New2", Age: 22, Grade: 14},
				{Name: "New3", Age: 29, Grade: 2},
				{Name: "New4", Age: 14, Grade: 15},
				{Name: "New5", Age: 10, Grade: 15},
				{Name: "New6", Age: 19, Grade: 4},
				{Name: "New8", Age: 29, Grade: 3},
				{Name: "New9", Age: 40, Grade: 15},
				{Name: "New10", Age: 50, Grade: 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StudentList{
				students: tt.fields.students,
			}
			s.RemoveStudent(tt.args.name)
			if !reflect.DeepEqual(s.students, tt.want) {
				t.Errorf("AddStudents() got = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestStudentList_Sort(t *testing.T) {
	type fields struct {
		students []Student
	}
	tests := []struct {
		name   string
		fields fields
		want   StudentList
	}{
		{
			name: "Test",
			fields: fields{
				students: []Student{
					{Name: "Test1", Age: 44, Grade: 18.5},
					{Name: "Test2", Age: 33, Grade: 20},
					{Name: "Test3.1", Age: 22, Grade: 14},
					{Name: "Test3.2", Age: 22, Grade: 14},
					{Name: "Test4", Age: 11, Grade: 16},
				},
			},
			want: StudentList{
				students: []Student{
					{Name: "Test2", Age: 33, Grade: 20},
					{Name: "Test1", Age: 44, Grade: 18.5},
					{Name: "Test4", Age: 11, Grade: 16},
					{Name: "Test3.1", Age: 22, Grade: 14},
					{Name: "Test3.2", Age: 22, Grade: 14},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StudentList{
				students: tt.fields.students,
			}
			if got := s.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudentList_Print(t *testing.T) {
	type fields struct {
		students []Student
	}
	tests := []struct {
		name    string
		fields  fields
		wantOut string
	}{
		{
			name: "Four students",
			fields: fields{
				students: []Student{
					{Name: "Test1", Age: 44, Grade: 18.5},
					{Name: "Test2", Age: 33, Grade: 20},
					{Name: "Test3", Age: 22, Grade: 14},
					{Name: "Test4", Age: 11, Grade: 16},
				},
			},
			wantOut: "Test1 (44): 18.5\n" +
				"Test2 (33): 20.0\n" +
				"Test3 (22): 14.0\n" +
				"Test4 (11): 16.0\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StudentList{
				students: tt.fields.students,
			}
			out := &bytes.Buffer{}
			s.Print(out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("Print() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
