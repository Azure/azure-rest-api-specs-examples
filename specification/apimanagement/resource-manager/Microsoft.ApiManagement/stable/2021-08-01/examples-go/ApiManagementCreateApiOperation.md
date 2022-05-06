Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapimanagement%2Farmapimanagement%2Fv0.5.0/sdk/resourcemanager/apimanagement/armapimanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiOperation.json
func ExampleAPIOperationClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIOperationClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<operation-id>",
		armapimanagement.OperationContract{
			Properties: &armapimanagement.OperationContractProperties{
				Description:        to.Ptr("<description>"),
				TemplateParameters: []*armapimanagement.ParameterContract{},
				Request: &armapimanagement.RequestContract{
					Description:     to.Ptr("<description>"),
					Headers:         []*armapimanagement.ParameterContract{},
					QueryParameters: []*armapimanagement.ParameterContract{},
					Representations: []*armapimanagement.RepresentationContract{
						{
							ContentType: to.Ptr("<content-type>"),
							SchemaID:    to.Ptr("<schema-id>"),
							TypeName:    to.Ptr("<type-name>"),
						}},
				},
				Responses: []*armapimanagement.ResponseContract{
					{
						Description: to.Ptr("<description>"),
						Headers:     []*armapimanagement.ParameterContract{},
						Representations: []*armapimanagement.RepresentationContract{
							{
								ContentType: to.Ptr("<content-type>"),
							},
							{
								ContentType: to.Ptr("<content-type>"),
							}},
						StatusCode: to.Ptr[int32](200),
					}},
				Method:      to.Ptr("<method>"),
				DisplayName: to.Ptr("<display-name>"),
				URLTemplate: to.Ptr("<urltemplate>"),
			},
		},
		&armapimanagement.APIOperationClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
