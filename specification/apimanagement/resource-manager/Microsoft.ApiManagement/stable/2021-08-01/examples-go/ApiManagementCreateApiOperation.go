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
	}
	ctx := context.Background()
	client, err := armapimanagement.NewAPIOperationClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"rg1",
		"apimService1",
		"PetStoreTemplate2",
		"newoperations",
		armapimanagement.OperationContract{
			Properties: &armapimanagement.OperationContractProperties{
				Description:        to.Ptr("This can only be done by the logged in user."),
				TemplateParameters: []*armapimanagement.ParameterContract{},
				Request: &armapimanagement.RequestContract{
					Description:     to.Ptr("Created user object"),
					Headers:         []*armapimanagement.ParameterContract{},
					QueryParameters: []*armapimanagement.ParameterContract{},
					Representations: []*armapimanagement.RepresentationContract{
						{
							ContentType: to.Ptr("application/json"),
							SchemaID:    to.Ptr("592f6c1d0af5840ca8897f0c"),
							TypeName:    to.Ptr("User"),
						}},
				},
				Responses: []*armapimanagement.ResponseContract{
					{
						Description: to.Ptr("successful operation"),
						Headers:     []*armapimanagement.ParameterContract{},
						Representations: []*armapimanagement.RepresentationContract{
							{
								ContentType: to.Ptr("application/xml"),
							},
							{
								ContentType: to.Ptr("application/json"),
							}},
						StatusCode: to.Ptr[int32](200),
					}},
				Method:      to.Ptr("POST"),
				DisplayName: to.Ptr("createUser2"),
				URLTemplate: to.Ptr("/user1"),
			},
		},
		&armapimanagement.APIOperationClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
