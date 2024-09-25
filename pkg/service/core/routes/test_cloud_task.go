package routes

import (
	"context"
	"fmt"
	"io"
	"net/http"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	taskspb "cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
)

func OrderCreateWorkstationTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the request body
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Get the authentication cookie
		cookie, err := r.Cookie("nada_session")
		if err != nil {
			http.Error(w, "Missing session cookie", http.StatusUnauthorized)
			return
		}

		// Create a Cloud Task to handle the request asynchronously
		task, err := createHTTPTaskWithAuthCookie(r.Context(), "nada-dev-db2e", "europe-west1", "test-async-arch-markedsplassen", "data.ansatt.dev.nav.no/api/workstations", cookie.Value, reqBody)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create Cloud Task:%v\n%v\n%v", err, len(cookie.Value), string(reqBody)), http.StatusInternalServerError)
			return
		}

		// Respond that the task was created successfully
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(fmt.Sprintf("Task created: %s\n%s\n%s", task.GetName(), len(cookie.Value), string(reqBody))))
	}
}

// createHTTPTaskWithAuthCookie constructs a task with a authorization cookie
// and HTTP target then adds it to a Queue.
func createHTTPTaskWithAuthCookie(ctx context.Context, projectID, locationID, queueID, url, authCookie string, body []byte) (*taskspb.Task, error) {
	client, err := cloudtasks.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewClient: %w", err)
	}
	defer client.Close()

	// Build the Task queue path.
	queuePath := fmt.Sprintf("projects/%s/locations/%s/queues/%s", projectID, locationID, queueID)

	// Build the Task payload.
	// https://godoc.org/google.golang.org/genproto/googleapis/cloud/tasks/v2#CreateTaskRequest
	req := &taskspb.CreateTaskRequest{
		Parent: queuePath,
		Task: &taskspb.Task{
			// https://godoc.org/google.golang.org/genproto/googleapis/cloud/tasks/v2#HttpRequest
			MessageType: &taskspb.Task_HttpRequest{
				HttpRequest: &taskspb.HttpRequest{
					HttpMethod: taskspb.HttpMethod_POST,
					Url:        url,
					Body:       body, // Pass the original request body
					Headers: map[string]string{
						"Cookie": fmt.Sprintf("nada-session=%s", authCookie), // Pass the authentication cookie
					},
				},
			},
		},
	}

	createdTask, err := client.CreateTask(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cloudtasks.CreateTask: %w\nqueue:%v", err, queuePath)
	}

	return createdTask, nil
}
