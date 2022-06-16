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
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIOperationClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"rg1",
		"apimService1",
		"echo-api",
		"operationId",
		"*",
		armapimanagement.OperationUpdateContract{
			Properties: &armapimanagement.OperationUpdateContractProperties{
				TemplateParameters: []*armapimanagement.ParameterContract{},
				Request: &armapimanagement.RequestContract{
					QueryParameters: []*armapimanagement.ParameterContract{
						{
							Name:         to.Ptr("param1"),
							Type:         to.Ptr("string"),
							Description:  to.Ptr("A sample parameter that is required and has a default value of \"sample\"."),
							DefaultValue: to.Ptr("sample"),
							Required:     to.Ptr(true),
							Values: []*string{
								to.Ptr("sample")},
						}},
				},
				Responses: []*armapimanagement.ResponseContract{
					{
						Description:     to.Ptr("Returned in all cases."),
						Headers:         []*armapimanagement.ParameterContract{},
						Representations: []*armapimanagement.RepresentationContract{},
						StatusCode:      to.Ptr[int32](200),
					},
					{
						Description:     to.Ptr("Server Error."),
						Headers:         []*armapimanagement.ParameterContract{},
						Representations: []*armapimanagement.RepresentationContract{},
						StatusCode:      to.Ptr[int32](500),
					}},
				Method:      to.Ptr("GET"),
				DisplayName: to.Ptr("Retrieve resource"),
				URLTemplate: to.Ptr("/resource"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
