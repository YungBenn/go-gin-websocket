package test

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"workshop-web-golang/internal/controller"
// )



// func TestGetStudents(t *testing.T)  {
// 	// initDB("test111")

//     req, err := http.NewRequest("GET", "/students", nil)
//     if err != nil {
//         t.Fatal(err)
//     }

//     rr := httptest.NewRecorder()
//     handler := http.HandlerFunc()

//     handler.ServeHTTP(rr, req)

//     // Check the status code is what we expect.
//     if status := rr.Code; status != http.StatusOK {
//         t.Errorf("handler returned wrong status code: got %v want %v",
//             status, http.StatusOK)
//     }

//     var response []Jobs
//     if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
//         t.Errorf("got invalid response, expected list of jobs, got: %v", rr.Body.String())
//     }

//     if len(response) < 1 {
//         t.Errorf("expected at least 1 job, got %v", len(response))
//     }

//     for _, job := range response {
//         if job.SourcePath == "" {
//             t.Errorf("expected job id %d to  have a source path, was empty", job.JobID)
//         }
//     }

// }

