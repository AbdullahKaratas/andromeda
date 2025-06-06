// SPDX-FileCopyrightText: Copyright 2025 SAP SE or an SAP affiliate company
//
// SPDX-License-Identifier: Apache-2.0

package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sapcc/go-api-declarations/cadf"
	"github.com/sapcc/go-bits/audittools"

	"github.com/sapcc/andromeda/internal/auth"
	"github.com/sapcc/andromeda/internal/config"
	"github.com/sapcc/andromeda/internal/policy"
)

type auditController struct {
	Auditor audittools.Auditor
}

func NewAuditController() *auditController {
	ctx := context.Background()
	auditor, err := audittools.NewAuditor(ctx, audittools.AuditorOpts{
		Observer: audittools.Observer{
			TypeURI: "service/gtm",
			Name:    "Andromeda",
			ID:      audittools.GenerateUUID(),
		},
		QueueName:     config.Global.Audit.QueueName,
		ConnectionURL: config.Global.Audit.TransportURL,
	})
	if err != nil {
		panic(err)
	}

	return &auditController{auditor}
}

// auditResponseWriter is a wrapper of regular ResponseWriter
type auditResponseWriter struct {
	http.ResponseWriter
	controller *auditController
	request    *http.Request
}

// AuditResource is an audittools.EventRenderer.
type AuditResource struct {
	project     string
	resource    string
	routeParams middleware.RouteParams
}

// Render implements the audittools.EventRenderer interface.
func (a AuditResource) Render() cadf.Resource {
	id := ""
	var attachments []cadf.Attachment
	for _, routeParam := range a.routeParams {
		attachments = append(attachments, cadf.Attachment{
			Name:    routeParam.Name,
			Content: routeParam.Value,
		})
		// Last route param is our target id
		id = routeParam.Value
	}
	res := cadf.Resource{
		TypeURI:     fmt.Sprintf("gtm/%s", a.resource),
		ID:          id,
		ProjectID:   a.project,
		Attachments: attachments,
	}

	return res
}

func (arw *auditResponseWriter) WriteHeader(code int) {
	arw.ResponseWriter.WriteHeader(code)

	mr := middleware.MatchedRouteFrom(arw.request)
	resource := strings.Split(policy.RuleFromHTTPRequest(arw.request), ":")[1]
	user, err := auth.UserForRequest(arw.request)
	projectID, _ := auth.ProjectScopeForRequest(arw.request)
	if err != nil {
		log.Error(err.Error())
		return
	}

	arw.controller.Auditor.Record(audittools.EventParameters{
		Time:       time.Now(),
		Request:    arw.request,
		User:       user,
		ReasonCode: code,
		Action:     cadf.GetAction(arw.request.Method),
		Target: AuditResource{
			projectID,
			resource,
			mr.Params,
		},
	})
}

func (ac *auditController) NewAuditResponseWriter(w http.ResponseWriter, r *http.Request) *auditResponseWriter {
	return &auditResponseWriter{w, ac, r}
}

// AuditHandler provides the audit handling.
func (ac *auditController) AuditHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			next.ServeHTTP(w, r)
			return
		}

		qrw := ac.NewAuditResponseWriter(w, r)
		next.ServeHTTP(qrw, r)
	})
}
