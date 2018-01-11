package appeng

import (
	"fmt"
	"gcpAppengie_scheduler/computeservice"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
	"google.golang.org/appengine"
	glog "google.golang.org/appengine/log"
)

func startvmHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	status := startVM(r, params["zone"], params["servername"])
	// client := getAuthClient(r)
	// comService := computeservice.New(*client)
	// comService.SetProject(ctx)
	// glog.Infof(ctx, "params", params["servername"], params["zone"])
	// status, _ := comService.StartVM(params["servername"], params["zone"], ctx)
	fmt.Fprint(w, "starting VM", status)
}
func stopvmHandler(w http.ResponseWriter, r *http.Request) {
	// ctx := appengine.NewContext(r)
	params := mux.Vars(r)
	status := stopVM(r, params["zone"], params["servername"])
	// client := getAuthClient(r)
	// comService := computeservice.New(*client)
	// comService.SetProject(ctx)
	// glog.Infof(ctx, "params", params["servername"], params["zone"])
	// status, _ := comService.StopVM(params["servername"], params["zone"], ctx)
	fmt.Fprint(w, "stopping VM", status)

}
func incomingMail(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	defer r.Body.Close()
	mymail, err := mail.ReadMessage(r.Body)
	if err != nil {
		glog.Errorf(ctx, "error %v", err)

	}
	from := mymail.Header.Get("From")
	if strings.Contains(from, "vikram.bs@gmail.com") {
		glog.Errorf(ctx, "the mail address '%v' is Not authorized", from)
		//http.Error(w, "no access", 500)
		log.Fatal("error")
	}
	subject := mymail.Header.Get("Subject")
	command, zone, instancename := getCommand(subject)

	switch command {
	case "start":
		status := startVM(r, zone, instancename)
		glog.Infof(ctx, "StartingVM %v", status)
	case "stop":
		status := stopVM(r, zone, instancename)
		glog.Infof(ctx, "StoppingVM %v", status)

	}

	//glog.Infof(ctx, "val %v,%v,%v", command, zone, instancename)
	// select command{
	// case "start":{

	// }
	// }

	// buf := new(bytes.Buffer)
	// buf.ReadFrom(mymail.Body)
	// glog.Infof(ctx, "value %v", buf.String())

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

func getCommand(subject string) (string, string, string) {
	var command, zone, instancename string
	comval := strings.Split(subject, "#")
	if len(comval) > 1 {
		command = comval[0]
		tmp := strings.Split(comval[1], ":")
		if len(tmp) > 1 {
			zone = tmp[0]
			instancename = tmp[1]
		}
	}

	return command, zone, instancename

}

func startVM(r *http.Request, zone, instancename string) bool {
	ctx := appengine.NewContext(r)
	client := getAuthClient(r)
	comService := computeservice.New(*client)
	comService.SetProject(ctx)
	//glog.Infof(ctx, "params", params["servername"], params["zone"])
	status, _ := comService.StartVM(instancename, zone, ctx)
	return status

}

func stopVM(r *http.Request, zone, instancename string) bool {
	ctx := appengine.NewContext(r)
	client := getAuthClient(r)
	comService := computeservice.New(*client)
	comService.SetProject(ctx)
	//glog.Infof(ctx, "params", params["servername"], params["zone"])
	status, _ := comService.StopVM(instancename, zone, ctx)
	return status

}
