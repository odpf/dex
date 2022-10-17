package project

import (
	"net/http"

	"github.com/gorilla/mux"
	shieldv1beta1 "go.buf.build/odpf/gwv/odpf/proton/odpf/shield/v1beta1"
)

// Routes installs project management APIs to router.
func Routes(r *mux.Router, shieldClient shieldv1beta1.ShieldServiceClient) {
	r.HandleFunc("/projects", handleListProjects(shieldClient)).Methods(http.MethodGet)
	r.HandleFunc("/projects/{id}", handleGetProject(shieldClient)).Methods(http.MethodGet)
}
