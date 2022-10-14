// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package http

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}