package container

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInitcont(t *testing.T) {
	type testStruct struct {
		name     string
		filename string
	}

	tests := []testStruct{
		{
			name:     "env normal",
			filename: ".env",
		},
		{
			name:     "env test",
			filename: ".env.test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Initcont(tt.filename)
		})
	}
}

func TestAppsInit(t *testing.T) {
	type appsInitTest struct {
		name string
		want Apps
	}
	tests := []appsInitTest{
		{
			name: ".env.test read",
			want: Apps{
				Name:           "go-simple-user-crud",
				Host:           "0.0.0.0",
				Version:        "v1",
				SwaggerAddress: "0.0.0.0:8080",
				HttpPort:       8080,
				SecretJwt:      "this-is-secret-jwt",
				CtxTimeout:     500,
			},
		},
	}

	Initcont(".env.test")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AppsInit()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppsInit() = %v, want %v", got, tt.want)
			}

			fmt.Printf("Struct %#v \n\n", got)
		})
	}
}

func TestLoggerInit(t *testing.T) {
	tests := []struct {
		name string
		want Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoggerInit(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoggerInit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitContainer(t *testing.T) {
	type args struct {
		containters []string
	}
	tests := []struct {
		name string
		args args
		want *Container
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitContainer(tt.args.containters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}
