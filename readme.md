# Google Cloud Platform Scheduled Start/Stop of Compute Instance	

## Objective

>*The Key/hidden objective is to better understand Google cloud platform services and explore this utilizing golang.
But to build "something..." i'll need to define **"something"** hence.. the below objective...*

* To be able to Schedule start/ stop of a Compute instance 
* To be able start/stop the Compute instance without access to Google cloud console

## Overview

One of the benefits of Cloud Platform is that it provides a great cost benefit by paying only for what we use. Basically if I have a test/Dev environment which i am not using.. then I can just shutdown the environment and don’t have to pay for it...(don’t have to pay for compute.. would still pay for storage). But what if I forget to shutdown by end of day... ? ( i have burnt my trial periods (and some amount of my credit card) in the past) so I decided to try implement my own scheduler using google's platform.

I manage to implement this by utilizing AppEngine standard a fully managed Google Cloud Platform Service to schedule a start and stop a Compute instance. I also extend this to tap into Mail API feature of AppEngine to start/stop on demand too...

## Tech Spec

* Programming Language : Golang
* Cloud Platform Services
  * AppEngine Standard
    * Service
    * Mail Api
    * Cron Job

### Design

* Appengine service : This hosts the service which provides the ability  has 3 routes (same URL)
  * Two of the routes are used API service to start/stop*
  * The third route is to support Mail API
* Mail API : Appengine has this cool support for Mail API, basically you can provide a specific route which can receive emails as a web request
* Cron Job: One of the AppEngine features is to provide ability to schedule cron jobs - by doing so Cron Jobs can make a HTTP Get Call to specific App Engine Service based on time.
* Choice of Programming Language : Golang- no technical reason for this... this is one of my new goto language which I like to use...

### Security

#### How does the AppEngine service call GCP API to stop/start compute instance?

When you create AppEngine service a Default Service account is allocated to that project, which gets "Editor" role. Appengine Service taps into this service to make these calls. To avoid Appengine from being exploited reduce the access privileges for the service account to the lowest desired...

#### How do I secure the cron URL?

AppEngine URLs are public facing and if we don’t secure the URL anyone can tap into them to start/stop a compute instance... which is not what we want...

AppEngine has a feature to enforce a link to be accessed admin only... which is what i use to secure the link

``` YAML
handlers:
- url: /.*
  script: _go_app
  login: admin
```

#### Can anyone send email to start/stop?

There is no default feature which I can find in GCP which lets me control this. But I implement this as part of the authorization aspect of my code.

### *Disclaimers*

>* Ignore the rest API design i know that it violates several restful design concepts - which is not what this demo/sample is about...

### *Reference*

* [AppEngine](https://cloud.google.com/appengine/docs/)
* [Cron Job](https://cloud.google.com/appengine/docs/standard/go/config/cron)
* [Mail API](https://cloud.google.com/appengine/docs/standard/go/mail/)
* [Golang -*my new love*](https://golang.org/)


