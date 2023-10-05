package build

import (
	extensionv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/controller/openapi/builder"
	"k8s.io/kube-openapi/pkg/spec3"
)

type ConverterV3 interface {
	Convert(crd *extensionv1.CustomResourceDefinition) ([]*spec3.OpenAPI, error)
}

type OpenApiV3 struct {
}

func (s OpenApiV3) Convert(crd *extensionv1.CustomResourceDefinition) ([]*spec3.OpenAPI, error) {
	var crdSpecs []*spec3.OpenAPI
	for _, v := range crd.Spec.Versions {
		if !v.Served {
			continue
		}
		// Defaults are not pruned here, but before being served.
		sw, err := builder.BuildOpenAPIV3(crd, v.Name, builder.Options{V2: true, SkipFilterSchemaForKubectlOpenAPIV2Validation: true, StripValueValidation: true, StripNullable: true, AllowNonStructural: false})
		if err != nil {
			return nil, err
		}
		crdSpecs = append(crdSpecs, sw)
	}
	//builder.MergeSpecs(c.staticSpec, crdSpecs...)
	return crdSpecs, nil
}

func NewOpenApiV3Converter() ConverterV3 {
	return OpenApiV3{}
}
