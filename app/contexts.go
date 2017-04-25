// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "api": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/intervention-engine/ie/design
// --out=$(GOPATH)/src/github.com/intervention-engine/ie
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
	"time"
)

// CreateCareTeamContext provides the care_team create action context.
type CreateCareTeamContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateCareTeamPayload
}

// NewCreateCareTeamContext parses the incoming request URL and body, performs validations and creates the
// context used by the care_team controller create action.
func NewCreateCareTeamContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateCareTeamContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateCareTeamContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createCareTeamPayload is the care_team create action payload.
type createCareTeamPayload struct {
	// Timestamp for care team creation
	CreatedAt *time.Time `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Unique care team ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Care team leader
	Leader *string `form:"leader,omitempty" json:"leader,omitempty" xml:"leader,omitempty"`
	// Care team name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createCareTeamPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Leader == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "leader"))
	}
	return
}

// Publicize creates CreateCareTeamPayload from createCareTeamPayload
func (payload *createCareTeamPayload) Publicize() *CreateCareTeamPayload {
	var pub CreateCareTeamPayload
	if payload.CreatedAt != nil {
		pub.CreatedAt = payload.CreatedAt
	}
	if payload.ID != nil {
		pub.ID = payload.ID
	}
	if payload.Leader != nil {
		pub.Leader = *payload.Leader
	}
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// CreateCareTeamPayload is the care_team create action payload.
type CreateCareTeamPayload struct {
	// Timestamp for care team creation
	CreatedAt *time.Time `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Unique care team ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Care team leader
	Leader string `form:"leader" json:"leader" xml:"leader"`
	// Care team name
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateCareTeamPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Leader == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "leader"))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateCareTeamContext) OK(r *Careteam) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.careteam+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *CreateCareTeamContext) OKLink(r *CareteamLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.careteam+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateCareTeamContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateCareTeamContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CreateCareTeamContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateCareTeamContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// DeleteCareTeamContext provides the care_team delete action context.
type DeleteCareTeamContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewDeleteCareTeamContext parses the incoming request URL and body, performs validations and creates the
// context used by the care_team controller delete action.
func NewDeleteCareTeamContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteCareTeamContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteCareTeamContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeleteCareTeamContext) OK(r *Careteam) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.careteam+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *DeleteCareTeamContext) OKLink(r *CareteamLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.careteam+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteCareTeamContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteCareTeamContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteCareTeamContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteCareTeamContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListCareTeamContext provides the care_team list action context.
type ListCareTeamContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListCareTeamContext parses the incoming request URL and body, performs validations and creates the
// context used by the care_team controller list action.
func NewListCareTeamContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListCareTeamContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListCareTeamContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListCareTeamContext) OK(r CareteamCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.careteam+json; type=collection")
	if r == nil {
		r = CareteamCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListCareTeamContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListCareTeamContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListCareTeamContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowCareTeamContext provides the care_team show action context.
type ShowCareTeamContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowCareTeamContext parses the incoming request URL and body, performs validations and creates the
// context used by the care_team controller show action.
func NewShowCareTeamContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowCareTeamContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowCareTeamContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowCareTeamContext) OK(r *Careteam) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.careteam+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowCareTeamContext) OKLink(r *CareteamLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.careteam+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowCareTeamContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowCareTeamContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowCareTeamContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// UpdateCareTeamContext provides the care_team update action context.
type UpdateCareTeamContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      string
	Payload *UpdateCareTeamPayload
}

// NewUpdateCareTeamContext parses the incoming request URL and body, performs validations and creates the
// context used by the care_team controller update action.
func NewUpdateCareTeamContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateCareTeamContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateCareTeamContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// updateCareTeamPayload is the care_team update action payload.
type updateCareTeamPayload struct {
	// Timestamp for care team creation
	CreatedAt *time.Time `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// Unique care team ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Care team leader
	Leader *string `form:"leader,omitempty" json:"leader,omitempty" xml:"leader,omitempty"`
	// Care team name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateCareTeamPayload) Validate() (err error) {
	if payload.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "id"))
	}
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Leader == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "leader"))
	}
	if payload.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "createdAt"))
	}
	return
}

// Publicize creates UpdateCareTeamPayload from updateCareTeamPayload
func (payload *updateCareTeamPayload) Publicize() *UpdateCareTeamPayload {
	var pub UpdateCareTeamPayload
	if payload.CreatedAt != nil {
		pub.CreatedAt = *payload.CreatedAt
	}
	if payload.ID != nil {
		pub.ID = *payload.ID
	}
	if payload.Leader != nil {
		pub.Leader = *payload.Leader
	}
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// UpdateCareTeamPayload is the care_team update action payload.
type UpdateCareTeamPayload struct {
	// Timestamp for care team creation
	CreatedAt time.Time `form:"createdAt" json:"createdAt" xml:"createdAt"`
	// Unique care team ID
	ID string `form:"id" json:"id" xml:"id"`
	// Care team leader
	Leader string `form:"leader" json:"leader" xml:"leader"`
	// Care team name
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateCareTeamPayload) Validate() (err error) {
	if payload.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "id"))
	}
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	if payload.Leader == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "leader"))
	}

	return
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateCareTeamContext) OK(r *Careteam) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.careteam+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *UpdateCareTeamContext) OKLink(r *CareteamLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.careteam+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateCareTeamContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateCareTeamContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateCareTeamContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *UpdateCareTeamContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListPatientContext provides the patient list action context.
type ListPatientContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Page    *int
	PerPage *int
	SortBy  *string
}

// NewListPatientContext parses the incoming request URL and body, performs validations and creates the
// context used by the patient controller list action.
func NewListPatientContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListPatientContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListPatientContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPage := req.Params["page"]
	if len(paramPage) > 0 {
		rawPage := paramPage[0]
		if page, err2 := strconv.Atoi(rawPage); err2 == nil {
			tmp2 := page
			tmp1 := &tmp2
			rctx.Page = tmp1
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("page", rawPage, "integer"))
		}
		if rctx.Page != nil {
			if *rctx.Page < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`page`, *rctx.Page, 1, true))
			}
		}
	}
	paramPerPage := req.Params["per_page"]
	if len(paramPerPage) > 0 {
		rawPerPage := paramPerPage[0]
		if perPage, err2 := strconv.Atoi(rawPerPage); err2 == nil {
			tmp4 := perPage
			tmp3 := &tmp4
			rctx.PerPage = tmp3
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("per_page", rawPerPage, "integer"))
		}
		if rctx.PerPage != nil {
			if *rctx.PerPage < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError(`per_page`, *rctx.PerPage, 1, true))
			}
		}
	}
	paramSortBy := req.Params["sort_by"]
	if len(paramSortBy) > 0 {
		rawSortBy := paramSortBy[0]
		rctx.SortBy = &rawSortBy
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListPatientContext) OK(r PatientCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.patient+json; type=collection")
	if r == nil {
		r = PatientCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListPatientContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListPatientContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListPatientContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowPatientContext provides the patient show action context.
type ShowPatientContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID string
}

// NewShowPatientContext parses the incoming request URL and body, performs validations and creates the
// context used by the patient controller show action.
func NewShowPatientContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowPatientContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowPatientContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowPatientContext) OK(r *Patient) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.patient+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowPatientContext) OKLink(r *PatientLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.patient+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowPatientContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowPatientContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowPatientContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}
