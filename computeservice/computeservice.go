package computeservice

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/api/compute/v1"
	"google.golang.org/appengine"
	glog "google.golang.org/appengine/log"
)

//Computeservice type
type Computeservice struct {
	*compute.Service
	client  *http.Client
	project string
}

//SetProject extract the environment id of the current applicaiton
func (cs *Computeservice) SetProject(ctx context.Context) string {

	cs.project = appengine.AppID(ctx)
	return cs.project

}

//New factory method to createa  compute service
func New(client http.Client) *Computeservice {
	cs := Computeservice{}
	cs.client = &client
	cs.Service, _ = compute.New(&client)
	return &cs
}

//GetVM help extract the complete instance object based on instance name and zone
func (cs *Computeservice) GetVM(instancename, zone string, ctx context.Context) *compute.Instance {
	inst, err := cs.Instances.Get(cs.project, zone, instancename).Do()

	if err != nil {
		// log.Fatal("error getting the vm", err)
		glog.Errorf(ctx, "error getting the vm", err)
	}

	return inst
}

//StartVM helps start a instance based on instance name and zone
func (cs *Computeservice) StartVM(instancename, zone string, ctx context.Context) (bool, error) {
	issucces := false

	inst, err := cs.Instances.Get(cs.project, zone, instancename).Do()
	glog.Infof(ctx, "project name %v", cs.project)
	if err != nil {
		glog.Errorf(ctx, "error getting the vm %v", err)
	}
	if inst.Status == "RUNNING" {
		issucces = true
	} else if inst.Status == "TERMINATED" {
		op, err := cs.Instances.Start(cs.project, zone, instancename).Do()
		if err != nil {
			log.Fatal("error Starting the VM: ", err)
		}
		cs.waitzoneOpertaton(op, zone)
		issucces = true
	} else {
		log.Panicln("Current Status is ", inst.Status)
	}

	return issucces, err
}

// StopVM method used to stop VM based on instance name and zone
func (cs *Computeservice) StopVM(instancename, zone string, ctx context.Context) (bool, error) {
	issucces := false

	inst, err := cs.Instances.Get(cs.project, zone, instancename).Do()

	if err != nil {
		log.Fatal("error getting the vm", err)
	}

	if inst.Status == "Terminated" {
		issucces = true
	} else if inst.Status != "Terminated" {
		op, err := cs.Instances.Stop(cs.project, zone, instancename).Do()
		if err != nil {
			log.Fatal("error stoppign the VM: ", err)
		}
		cs.waitzoneOpertaton(op, zone)
		issucces = true
	} else {
		log.Panicln("Current Status is ", inst.Status)
	}
	return issucces, err
}

// Wait operations for completing invoked task

func (cs *Computeservice) waitzoneOpertaton(op *compute.Operation, zone string) *compute.Operation {
	return cs.waitprojzoneOpertaton(op, cs.project, zone)
}

func (cs *Computeservice) waitprojzoneOpertaton(op *compute.Operation, project, zone string) *compute.Operation {
	ch := make(chan *compute.Operation, 1)

	go cs.waitforoperation(op, project, zone, ch)
	select {
	case finalOp := <-ch:
		return finalOp
	}

}
func (cs *Computeservice) waitforoperation(op *compute.Operation, project, zone string, ch chan *compute.Operation) {

	for {
		time.Sleep(10 * time.Second)
		fmt.Printf("operation name: %v Status: %v \n ", op.Name, op.Status)
		opres, err := cs.ZoneOperations.Get(project, zone, op.Name).Do()
		if err != nil {
			log.Fatal(err)
		}
		if opres != nil {
			fmt.Println("status ", opres.Status)
			if opres.Status == "DONE" {
				ch <- opres
			}
		}
	}

}
