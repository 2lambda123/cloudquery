// Code generated by codegen; DO NOT EDIT.

package client

var GcpServices = map[string]bool{
    // Non discoverable services
	"livestream.googleapis.com": true,
	"cloudiot.googleapis.com": true,
	// Discoverable services
	{{- range .}}
      "{{.}}": true,
	{{- end}}
}
