# Google Cloud Platform Scheduled Start/Stop of Compute Instance

## Goal Objective

>*The Key/hidden objective is to better understand Google cloud platform services and explore this utilizing golang.
But to build "something..." i'll need to define **"something"** hence.. the below objective...*

* To be able to Schedule start/ stop of a Compute instance 
* To be able start/stop the Compute instance without access to Google cloud console

## Overview

One of the benifits of Cloud Platform is that it provides a great cost benifit by paying only for what we use. Baiscally if I have a test/Dev environment which i am not using.. then I can jsut shutdown the environment and dont have to pay forit...(dont have to pay for compute.. would still pay for storage). But what if I forget to shutdown by end of day... ? ( i have burnt my trail periods (and some amout of my creadit card) in the past) so I decided to try impliment my own scheduler using google's platform.

This implimentaiton utlizes AppEngine standard a fully managed Google Cloud Platform Service to scheulde a job to start and stop a Compute instance. I also extend this to tap into Mail API feature of AppEngine to start/stop on demand too...

### Design

* Appengine service : This hosts the service which provides the ability  has 3 routes (same URL)
  * Two of the routes are used API service to start/stop*

## Tech Spec -

* Programming Language : Golang
* Cloud Platform Services
  * AppEngine Standard
    * Service
    * Mail Api
    * Cron Job

### Security

#### Authentiaiton/Authorization

##### Secure the cron URL :

### Debugging


>* Ignore the rest API design i know that it vailotes several restful design concepts - which is not what this demo/sample is about...

### *Refrence*
* [Cron Job](https://cloud.google.com/appengine/docs/standard/go/config/cron)
* [Mail API](https://cloud.google.com/appengine/docs/standard/go/mail/)
* [Golang -*my new love*](https://golang.org/)

