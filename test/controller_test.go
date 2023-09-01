package test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"workshop-web-golang/internal/controller"
	"workshop-web-golang/internal/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupDBConfig() *db.ConfigDB {
	config := &db.ConfigDB{
		Host:     "localhost",
		User:     "postgres",
		Password: "example",
		DBName:   "workshop_len",
		Port:     "5432",
		SSLMode:  "disable",
	}

	return config
}

func TestStudentsController_CreateStudents(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		r *gin.Engine
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		payload        []byte
		expectedStatus int
	}{
		{
			name: "ValidPayload",
			fields: fields{
				DB: db.ConnectDB(setupDBConfig()),
			},
			args: args{
				r: func() *gin.Engine {
					return gin.Default()
				}(),
			},
			payload: []byte(`{
				"nama": "John Doe",
				"jurusan": "Computer Science",
				"status": "Aktif"
			}`),
			expectedStatus: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &controller.StudentsController{
				DB: tt.fields.DB,
			}

			r := tt.args.r

			r.POST("/students", sc.CreateStudents)

			bodyReader := bytes.NewReader(tt.payload)
			req, err := http.NewRequest(http.MethodPost, "/students", bodyReader)
			if err != nil {
				t.Errorf("Error creating request: %s", err)
			}

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)
			log.Println(w)
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestStudentsController_GetStudents(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		r *gin.Engine
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErr        bool
		expectedStatus int
	}{
		{
			name: "Successful case",
			fields: fields{
				DB: db.ConnectDB(setupDBConfig()),
			},
			args: args{
				r: func() *gin.Engine {
					return gin.Default()
				}(),
			},
			wantErr:        false,
			expectedStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &controller.StudentsController{
				DB: tt.fields.DB,
			}

			r := tt.args.r

			r.GET("/students", sc.GetStudents)

			req, _ := http.NewRequest(http.MethodGet, "/students", nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestStudentsController_GetStudentByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		r *gin.Engine
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		expectedStatus int
		studentID      string
	}{
		{
			name: "ValidStudentID",
			fields: fields{
				DB: db.ConnectDB(setupDBConfig()),
			},
			args: args{
				r: func() *gin.Engine {
					return gin.Default()
				}(),
			},
			expectedStatus: http.StatusOK,
			studentID:      "2",
		},
		{
			name: "InvalidStudentID",
			fields: fields{
				DB: db.ConnectDB(setupDBConfig()),
			},
			args: args{
				r: func() *gin.Engine {
					return gin.Default()
				}(),
			},
			expectedStatus: http.StatusOK,
			studentID:      "999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &controller.StudentsController{
				DB: tt.fields.DB,
			}

			r := tt.args.r

			r.GET("/students/:id", sc.GetStudentByID)

			req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/students/%s", tt.studentID), nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if tt.name == "InvalidStudentID" {
				if w.Code == tt.expectedStatus {
					t.Errorf("Expected status code %d, but got %d", tt.expectedStatus, w.Code)
				}

				expectedID := tt.studentID
				if actualID := w.Body.String(); strings.Contains(actualID, expectedID) {
					t.Errorf("Expected response to contain ID %s, but got %s", expectedID, actualID)
				}
			} else {
				if w.Code != tt.expectedStatus {
					t.Errorf("Expected status code %d, but got %d", tt.expectedStatus, w.Code)
				}

				// Memeriksa parameter ID yang diproses dengan benar
				expectedID := tt.studentID
				if actualID := w.Body.String(); !strings.Contains(actualID, expectedID) {
					t.Errorf("Expected response to contain ID %s, but got %s", expectedID, actualID)
				}
			}

		})
	}
}

func TestStudentsController_DeleteStudentByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		r *gin.Engine
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		expectedStatus int
		studentID      string
	}{
		{
			name: "ValidStudentID",
			fields: fields{
				DB: db.ConnectDB(setupDBConfig()),
			},
			args: args{
				r: func() *gin.Engine {
					return gin.Default()
				}(),
			},
			expectedStatus: http.StatusOK,
			studentID:      "17",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := &controller.StudentsController{
				DB: tt.fields.DB,
			}

			r := tt.args.r

			r.DELETE("/students/:id", sc.DeleteStudentByID)

			req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/students/%s", tt.studentID), nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)
			log.Println(w)
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
