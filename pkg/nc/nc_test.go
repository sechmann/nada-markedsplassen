package nc_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/navikt/nada-backend/pkg/nc"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetTeamGoogleProjects(t *testing.T) {
	testCases := []struct {
		name        string
		naisCluster string
		body        *nc.Response
		err         error
		expectErr   bool
		expect      map[string]nc.TeamOutput
	}{
		{
			name:        "should return team project",
			naisCluster: "env-1",
			body: &nc.Response{
				Data: nc.Data{
					Teams: nc.Teams{
						Nodes: []nc.Team{
							{
								Slug:             "team-1",
								GoogleGroupEmail: "team-1@nav.no",
								Environments: []nc.Environment{
									{
										Name:         "env-1",
										GcpProjectID: "gcp-project-1",
									},
								},
							},
							{
								Slug:             "team-2",
								GoogleGroupEmail: "team-1@nav.no",
								Environments: []nc.Environment{
									{
										Name:         "env-2",
										GcpProjectID: "gcp-project-2",
									},
								},
							},
						},
						PageInfo: nc.PageInfo{
							HasNextPage: false,
						},
					},
				},
			},
			expect: map[string]nc.TeamOutput{
				"team-1": {
					GroupEmail:   "team-1@nav.no",
					GCPProjectID: "gcp-project-1",
				},
			},
		},
		{
			name:        "should return empty if no team project",
			naisCluster: "env-1",
			body: &nc.Response{
				Data: nc.Data{
					Teams: nc.Teams{
						Nodes: []nc.Team{
							{
								Slug:             "team-1",
								GoogleGroupEmail: "team-1@nav.no",
								Environments: []nc.Environment{
									{
										Name:         "env-2",
										GcpProjectID: "gcp-project-2",
									},
								},
							},
						},
					},
				},
			},
			expect: map[string]nc.TeamOutput{},
		},
		{
			name:      "should return error",
			err:       assert.AnError,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/query", r.URL.Path)
				assert.Equal(t, "Bearer super-secret", r.Header.Get("Authorization"))
				assert.Equal(t, http.MethodPost, r.Method)

				if tc.err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(tc.body)
				assert.NoError(t, err)
			}))

			client := nc.New(testServer.URL, "super-secret", tc.naisCluster, http.DefaultClient)
			got, err := client.GetTeamGoogleProjects(context.Background())
			if tc.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expect, got)
			}
		})
	}
}
