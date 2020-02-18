// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"encoding/json"
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// GetEntryHandlerFunc turns a function with the right signature into a get entry handler
type GetEntryHandlerFunc func(GetEntryParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEntryHandlerFunc) Handle(params GetEntryParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetEntryHandler interface for that can handle valid get entry params
type GetEntryHandler interface {
	Handle(GetEntryParams, interface{}) middleware.Responder
}

// NewGetEntry creates a new http.Handler for the get entry operation
func NewGetEntry(ctx *middleware.Context, handler GetEntryHandler) *GetEntry {
	return &GetEntry{Context: ctx, Handler: handler}
}

/*GetEntry swagger:route GET /entries/{entry_id} getEntry

GetEntry get entry API

*/
type GetEntry struct {
	Context *middleware.Context
	Handler GetEntryHandler
}

func (o *GetEntry) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEntryParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetEntryDefaultBody get entry default body
// swagger:model GetEntryDefaultBody
type GetEntryDefaultBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// message
	// Required: true
	Message string `json:"message"`
}

// Validate validates this get entry default body
func (o *GetEntryDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEntryDefaultBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.RequiredString("getEntry default"+"."+"message", "body", string(o.Message)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEntryDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEntryDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetEntryDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetEntryOKBody get entry o k body
// swagger:model GetEntryOKBody
type GetEntryOKBody struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// creators
	Creators []*CreatorsItems0 `json:"creators"`

	// description
	Description string `json:"description,omitempty"`

	// files
	Files []*FilesItems0 `json:"files"`

	// id
	ID string `json:"id,omitempty"`

	// kind
	// Enum: [quote image]
	Kind string `json:"kind,omitempty"`

	// published at
	PublishedAt string `json:"published_at,omitempty"`

	// source
	Source *GetEntryOKBodySource `json:"source,omitempty"`

	// titles
	Titles []string `json:"titles"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this get entry o k body
func (o *GetEntryOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreators(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEntryOKBody) validateCreators(formats strfmt.Registry) error {

	if swag.IsZero(o.Creators) { // not required
		return nil
	}

	for i := 0; i < len(o.Creators); i++ {
		if swag.IsZero(o.Creators[i]) { // not required
			continue
		}

		if o.Creators[i] != nil {
			if err := o.Creators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEntryOK" + "." + "creators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetEntryOKBody) validateFiles(formats strfmt.Registry) error {

	if swag.IsZero(o.Files) { // not required
		return nil
	}

	for i := 0; i < len(o.Files); i++ {
		if swag.IsZero(o.Files[i]) { // not required
			continue
		}

		if o.Files[i] != nil {
			if err := o.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEntryOK" + "." + "files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var getEntryOKBodyTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["quote","image"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getEntryOKBodyTypeKindPropEnum = append(getEntryOKBodyTypeKindPropEnum, v)
	}
}

const (

	// GetEntryOKBodyKindQuote captures enum value "quote"
	GetEntryOKBodyKindQuote string = "quote"

	// GetEntryOKBodyKindImage captures enum value "image"
	GetEntryOKBodyKindImage string = "image"
)

// prop value enum
func (o *GetEntryOKBody) validateKindEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getEntryOKBodyTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetEntryOKBody) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(o.Kind) { // not required
		return nil
	}

	// value enum
	if err := o.validateKindEnum("getEntryOK"+"."+"kind", "body", o.Kind); err != nil {
		return err
	}

	return nil
}

func (o *GetEntryOKBody) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(o.Source) { // not required
		return nil
	}

	if o.Source != nil {
		if err := o.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getEntryOK" + "." + "source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEntryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEntryOKBody) UnmarshalBinary(b []byte) error {
	var res GetEntryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetEntryOKBodySource get entry o k body source
// swagger:model GetEntryOKBodySource
type GetEntryOKBodySource struct {

	// entity
	Entity *GetEntryOKBodySourceEntity `json:"entity,omitempty"`

	// titles
	Titles []string `json:"titles"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this get entry o k body source
func (o *GetEntryOKBodySource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEntryOKBodySource) validateEntity(formats strfmt.Registry) error {

	if swag.IsZero(o.Entity) { // not required
		return nil
	}

	if o.Entity != nil {
		if err := o.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getEntryOK" + "." + "source" + "." + "entity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEntryOKBodySource) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEntryOKBodySource) UnmarshalBinary(b []byte) error {
	var res GetEntryOKBodySource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetEntryOKBodySourceEntity get entry o k body source entity
// swagger:model GetEntryOKBodySourceEntity
type GetEntryOKBodySourceEntity struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// files
	Files []*GetEntryOKBodySourceEntityFilesItems0 `json:"files"`

	// id
	ID string `json:"id,omitempty"`

	// kind
	// Enum: [person organization]
	Kind string `json:"kind,omitempty"`

	// names
	Names []string `json:"names"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// urls
	Urls []string `json:"urls"`
}

// Validate validates this get entry o k body source entity
func (o *GetEntryOKBodySourceEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEntryOKBodySourceEntity) validateFiles(formats strfmt.Registry) error {

	if swag.IsZero(o.Files) { // not required
		return nil
	}

	for i := 0; i < len(o.Files); i++ {
		if swag.IsZero(o.Files[i]) { // not required
			continue
		}

		if o.Files[i] != nil {
			if err := o.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getEntryOK" + "." + "source" + "." + "entity" + "." + "files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var getEntryOKBodySourceEntityTypeKindPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["person","organization"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getEntryOKBodySourceEntityTypeKindPropEnum = append(getEntryOKBodySourceEntityTypeKindPropEnum, v)
	}
}

const (

	// GetEntryOKBodySourceEntityKindPerson captures enum value "person"
	GetEntryOKBodySourceEntityKindPerson string = "person"

	// GetEntryOKBodySourceEntityKindOrganization captures enum value "organization"
	GetEntryOKBodySourceEntityKindOrganization string = "organization"
)

// prop value enum
func (o *GetEntryOKBodySourceEntity) validateKindEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getEntryOKBodySourceEntityTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *GetEntryOKBodySourceEntity) validateKind(formats strfmt.Registry) error {

	if swag.IsZero(o.Kind) { // not required
		return nil
	}

	// value enum
	if err := o.validateKindEnum("getEntryOK"+"."+"source"+"."+"entity"+"."+"kind", "body", o.Kind); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEntryOKBodySourceEntity) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEntryOKBodySourceEntity) UnmarshalBinary(b []byte) error {
	var res GetEntryOKBodySourceEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetEntryOKBodySourceEntityFilesItems0 get entry o k body source entity files items0
// swagger:model GetEntryOKBodySourceEntityFilesItems0
type GetEntryOKBodySourceEntityFilesItems0 struct {

	// content type
	ContentType string `json:"content_type,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// filename
	Filename string `json:"filename,omitempty"`

	// height
	Height int64 `json:"height,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// source url
	SourceURL string `json:"source_url,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// width
	Width int64 `json:"width,omitempty"`
}

// Validate validates this get entry o k body source entity files items0
func (o *GetEntryOKBodySourceEntityFilesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetEntryOKBodySourceEntityFilesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEntryOKBodySourceEntityFilesItems0) UnmarshalBinary(b []byte) error {
	var res GetEntryOKBodySourceEntityFilesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}