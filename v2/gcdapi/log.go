// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Log functionality.
// API Version: 1.3

package gcdapi

import (
	"context"
	"github.com/wirepair/gcd/v2/gcdmessage"
)

// Log entry.
type LogLogEntry struct {
	Source           string                 `json:"source"`                     // Log entry source.
	Level            string                 `json:"level"`                      // Log entry severity.
	Text             string                 `json:"text"`                       // Logged text.
	Timestamp        float64                `json:"timestamp"`                  // Timestamp when this entry was added.
	Url              string                 `json:"url,omitempty"`              // URL of the resource if known.
	LineNumber       int                    `json:"lineNumber,omitempty"`       // Line number in the resource.
	StackTrace       *RuntimeStackTrace     `json:"stackTrace,omitempty"`       // JavaScript stack trace.
	NetworkRequestId string                 `json:"networkRequestId,omitempty"` // Identifier of the network request associated with this entry.
	WorkerId         string                 `json:"workerId,omitempty"`         // Identifier of the worker associated with this entry.
	Args             []*RuntimeRemoteObject `json:"args,omitempty"`             // Call arguments.
}

// Violation configuration setting.
type LogViolationSetting struct {
	Name      string  `json:"name"`      // Violation type.
	Threshold float64 `json:"threshold"` // Time threshold to trigger upon.
}

// Issued when new message was logged.
type LogEntryAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		Entry *LogLogEntry `json:"entry"` // The entry.
	} `json:"Params,omitempty"`
}

type Log struct {
	target gcdmessage.ChromeTargeter
}

func NewLog(target gcdmessage.ChromeTargeter) *Log {
	c := &Log{target: target}
	return c
}

// Clears the log.
func (c *Log) Clear(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Log.clear"})
}

// Disables log domain, prevents further log entries from being reported to the client.
func (c *Log) Disable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Log.disable"})
}

// Enables log domain, sends the entries collected so far to the client by means of the `entryAdded` notification.
func (c *Log) Enable(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Log.enable"})
}

type LogStartViolationsReportParams struct {
	// Configuration for violations.
	Config []*LogViolationSetting `json:"config"`
}

// StartViolationsReportWithParams - start violation reporting.
func (c *Log) StartViolationsReportWithParams(ctx context.Context, v *LogStartViolationsReportParams) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Log.startViolationsReport", Params: v})
}

// StartViolationsReport - start violation reporting.
// config - Configuration for violations.
func (c *Log) StartViolationsReport(ctx context.Context, config []*LogViolationSetting) (*gcdmessage.ChromeResponse, error) {
	var v LogStartViolationsReportParams
	v.Config = config
	return c.StartViolationsReportWithParams(ctx, &v)
}

// Stop violation reporting.
func (c *Log) StopViolationsReport(ctx context.Context) (*gcdmessage.ChromeResponse, error) {
	return c.target.SendDefaultRequest(ctx, &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Log.stopViolationsReport"})
}
