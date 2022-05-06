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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateApiOperation.json
func ExampleAPIOperationClient_Update() {
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
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<service-name>",
		"<api-id>",
		"<operation-id>",
		"<if-match>",
		armapimanagement.OperationUpdateContract{
			Properties: &armapimanagement.OperationUpdateContractProperties{
				TemplateParameters: []*armapimanagement.ParameterContract{},
				Request: &armapimanagement.RequestContract{
					QueryParameters: []*armapimanagement.ParameterContract{
						{
							Name:         to.Ptr("<name>"),
							Type:         to.Ptr("<type>"),
							Description:  to.Ptr("<description>"),
							DefaultValue: to.Ptr("<default-value>"),
							Required:     to.Ptr(true),
							Values: []*string{
								to.Ptr("sample")},
						}},
				},
				Responses: []*armapimanagement.ResponseContract{
					{
						Description:     to.Ptr("<description>"),
						Headers:         []*armapimanagement.ParameterContract{},
						Representations: []*armapimanagement.RepresentationContract{},
						StatusCode:      to.Ptr[int32](200),
					},
					{
						Description:     to.Ptr("<description>"),
						Headers:         []*armapimanagement.ParameterContract{},
						Representations: []*armapimanagement.RepresentationContract{},
						StatusCode:      to.Ptr[int32](500),
					}},
				Method:      to.Ptr("<method>"),
				DisplayName: to.Ptr("<display-name>"),
				URLTemplate: to.Ptr("<urltemplate>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
