// Code generated by gin-swagger; DO NOT EDIT.

package infrastructure_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/mikkeloscar/gin-swagger/api"
	"github.com/mikkeloscar/gin-swagger/tracing"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/mikkeloscar/gin-swagger/example/models"
)

// CreateInfrastructureAccountEndpoint executes the core logic of the related
// route endpoint.
func CreateInfrastructureAccountEndpoint(handler func(ctx *gin.Context, params *CreateInfrastructureAccountParams) *api.Response) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		span := opentracing.SpanFromContext(tracing.Context(ctx))

		// attach tags to opentracing span
		if span != nil {
			ext.HTTPMethod.Set(span, ctx.Request.Method)
			ext.HTTPUrl.Set(span, ctx.Request.URL.String())
		}

		// generate params from request
		params := NewCreateInfrastructureAccountParams()
		err := params.readRequest(ctx)
		if err != nil {
			errObj := err.(*errors.CompositeError)
			problem := api.Problem{
				Title:  "Unprocessable Entity.",
				Status: int(errObj.Code()),
				Detail: errObj.Error(),
			}

			// attach tags to opentracing span
			if span != nil {
				ext.HTTPStatusCode.Set(span, uint16(problem.Status))
			}

			ctx.Writer.Header().Set("Content-Type", "application/problem+json")
			ctx.JSON(problem.Status, problem)
			return
		}

		resp := handler(ctx, params)

		// attach tags to opentracing span
		if span != nil {
			ext.HTTPStatusCode.Set(span, uint16(resp.Code))
		}

		switch resp.Code {
		case http.StatusNoContent:
			ctx.AbortWithStatus(resp.Code)
		default:
			ctx.JSON(resp.Code, resp.Body)
		}
	}
}

// NewCreateInfrastructureAccountParams creates a new CreateInfrastructureAccountParams object
// with the default values initialized.
func NewCreateInfrastructureAccountParams() *CreateInfrastructureAccountParams {
	var ()
	return &CreateInfrastructureAccountParams{}
}

// CreateInfrastructureAccountParams contains all the bound params for the create infrastructure account operation
// typically these are obtained from a http.Request
//
// swagger:parameters createInfrastructureAccount
type CreateInfrastructureAccountParams struct {

	/*Account that will be created.
	  Required: true
	  In: body
	*/
	InfrastructureAccount *models.InfrastructureAccount
}

// readRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *CreateInfrastructureAccountParams) readRequest(ctx *gin.Context) error {
	var res []error
	formats := strfmt.NewFormats()

	if runtime.HasBody(ctx.Request) {
		var body models.InfrastructureAccount
		if err := ctx.BindJSON(&body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("infrastructureAccount", "body", ""))
			} else {
				res = append(res, errors.NewParseError("infrastructureAccount", "body", "", err))
			}

		} else {
			if err := body.Validate(formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.InfrastructureAccount = &body
			}
		}

	} else {
		res = append(res, errors.Required("infrastructureAccount", "body", ""))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// vim: ft=go
