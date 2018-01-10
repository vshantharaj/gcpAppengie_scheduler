package appeng

import (
	"bytes"
	"fmt"
	"gcpschedular/computeservice"
	"log"
	"net/http"
	"net/mail"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
	"google.golang.org/appengine"
	glog "google.golang.org/appengine/log"
)

func startvmHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	params := mux.Vars(r)
	client := getAuthClient(r)
	comService := computeservice.New(*client)
	comService.SetProject(ctx)
	glog.Infof(ctx, "params", params["servername"], params["zone"])
	status, _ := comService.StartVM(params["servername"], params["zone"], ctx)
	fmt.Fprint(w, "starting VM", status)
}
func stopvmHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	params := mux.Vars(r)

	client := getAuthClient(r)
	comService := computeservice.New(*client)
	comService.SetProject(ctx)
	glog.Infof(ctx, "params", params["servername"], params["zone"])
	status, _ := comService.StopVM(params["servername"], params["zone"], ctx)
	fmt.Fprint(w, "stopping VM", status)

}
func incomingMail(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	defer r.Body.Close()

	asda, err := mail.ReadMessage(r.Body)
	if err != nil {
		glog.Errorf(ctx, "error %v", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(asda.Body)
	glog.Infof(ctx, "value %v", buf.String())

}

func getAuthClient(r *http.Request) *http.Client {
	//projectmap := map[int]cloudresourcemanager.Project{}
	ctx := appengine.NewContext(r)
	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		log.Fatal("Error getting Client", err)
	}

	return client
}
