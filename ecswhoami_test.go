package ecswhoami

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestLookup(t *testing.T) {
	_, err := Lookup()
	if err != ErrEnvNotSet {
		t.Errorf("expected %v, got %v", ErrEnvNotSet, err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
			"DockerId": "43481a6ce4842eec8fe72fc28500c6b52edcc0917f105b83379f88cac1ff3946",
			"Name": "nginx-curl",
			"DockerName": "ecs-nginx-5-nginx-curl-ccccb9f49db0dfe0d901",
			"Image": "nrdlngr/nginx-curl",
			"ImageID": "sha256:2e00ae64383cfc865ba0a2ba37f61b50a120d2d9378559dcd458dc0de47bc165",
			"Labels": {
				"com.amazonaws.ecs.cluster": "default",
				"com.amazonaws.ecs.container-name": "nginx-curl",
				"com.amazonaws.ecs.task-arn": "arn:aws:ecs:us-east-2:012345678910:task/9781c248-0edd-4cdb-9a93-f63cb662a5d3",
				"com.amazonaws.ecs.task-definition-family": "nginx",
				"com.amazonaws.ecs.task-definition-version": "5"
			},
			"DesiredStatus": "RUNNING",
			"KnownStatus": "RUNNING",
			"Limits": {
				"CPU": 512,
				"Memory": 512
			},
			"CreatedAt": "2018-02-01T20:55:10.554941919Z",
			"StartedAt": "2018-02-01T20:55:11.064236631Z",
			"Type": "NORMAL",
			"Networks": [
				{
					"NetworkMode": "awsvpc",
					"IPv4Addresses": [
						"10.0.2.106"
					]
				}
			]
		}`)
	}))
	defer ts.Close()

	os.Setenv(envVar, ts.URL)

	m, err := Lookup()
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	expected := Metadata{
		DockerID:   "43481a6ce4842eec8fe72fc28500c6b52edcc0917f105b83379f88cac1ff3946",
		Name:       "nginx-curl",
		DockerName: "ecs-nginx-5-nginx-curl-ccccb9f49db0dfe0d901",
		Image:      "nrdlngr/nginx-curl",
		ImageID:    "sha256:2e00ae64383cfc865ba0a2ba37f61b50a120d2d9378559dcd458dc0de47bc165",
		Labels: map[string]string{
			"com.amazonaws.ecs.cluster":                 "default",
			"com.amazonaws.ecs.container-name":          "nginx-curl",
			"com.amazonaws.ecs.task-arn":                "arn:aws:ecs:us-east-2:012345678910:task/9781c248-0edd-4cdb-9a93-f63cb662a5d3",
			"com.amazonaws.ecs.task-definition-family":  "nginx",
			"com.amazonaws.ecs.task-definition-version": "5",
		},
		DesiredStatus: "RUNNING",
		KnownStatus:   "RUNNING",
		Limits: Limits{
			CPU:    512,
			Memory: 512,
		},
	}

	if m.DockerID != m.DockerID {
		t.Errorf("expected %v, got %v", expected.DockerID, m.DockerID)
	}

	if m.Name != m.Name {
		t.Errorf("expected %v, got %v", expected.Name, m.Name)
	}

	if m.DockerName != m.DockerName {
		t.Errorf("expected %v, got %v", expected.DockerName, m.DockerName)
	}

	if m.Image != m.Image {
		t.Errorf("expected %v, got %v", expected.Image, m.Image)
	}

	if m.ImageID != m.ImageID {
		t.Errorf("expected %v, got %v", expected.ImageID, m.ImageID)
	}

	if m.DesiredStatus != m.DesiredStatus {
		t.Errorf("expected %v, got %v", expected.DesiredStatus, m.DesiredStatus)
	}

	if m.KnownStatus != m.KnownStatus {
		t.Errorf("expected %v, got %v", expected.KnownStatus, m.KnownStatus)
	}

	if m.Limits != m.Limits {
		t.Errorf("expected %v, got %v", expected.Limits, m.Limits)
	}

}
