//go:generate goagen bootstrap -d github.com/intervention-engine/ie/design

package main

import (
	"log"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/intervention-engine/ie/app"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	// Create service
	service := goa.New("api")
	sess, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatalln("dialing mongo failed for session at mongodb://localhost:27017")
	}
	defer sess.Close()

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	service.Use(WithMongoService(sess))

	// Mount "patient" controller
	pc := NewPatientController(service)
	app.MountPatientController(service, pc)
	ct := NewCareTeamController(service)
	app.MountCareTeamController(service, ct)
	sc := NewSwaggerController(service)
	app.MountSwaggerController(service, sc)

	// Start service
	if err := service.ListenAndServe(":3001"); err != nil {
		service.LogError("startup", "err", err)
	}
}