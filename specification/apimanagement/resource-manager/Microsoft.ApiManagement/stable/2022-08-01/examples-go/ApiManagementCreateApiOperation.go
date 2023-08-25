package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiOperation.json
func ExampleAPIOperationClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIOperationClient().CreateOrUpdate(ctx, "rg1", "apimService1", "PetStoreTemplate2", "newoperations", armapimanagement.OperationContract{
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
	}, &armapimanagement.APIOperationClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationContract = armapimanagement.OperationContract{
	// 	Name: to.Ptr("newoperations"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis/operations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/PetStoreTemplate2/operations/newoperations"),
	// 	Properties: &armapimanagement.OperationContractProperties{
	// 		Description: to.Ptr("This can only be done by the logged in user."),
	// 		TemplateParameters: []*armapimanagement.ParameterContract{
	// 		},
	// 		Request: &armapimanagement.RequestContract{
	// 			Description: to.Ptr("Created user object"),
	// 			Headers: []*armapimanagement.ParameterContract{
	// 			},
	// 			QueryParameters: []*armapimanagement.ParameterContract{
	// 			},
	// 			Representations: []*armapimanagement.RepresentationContract{
	// 				{
	// 					ContentType: to.Ptr("application/json"),
	// 					SchemaID: to.Ptr("592f6c1d0af5840ca8897f0c"),
	// 					TypeName: to.Ptr("User"),
	// 			}},
	// 		},
	// 		Responses: []*armapimanagement.ResponseContract{
	// 			{
	// 				Description: to.Ptr("successful operation"),
	// 				Headers: []*armapimanagement.ParameterContract{
	// 				},
	// 				Representations: []*armapimanagement.RepresentationContract{
	// 					{
	// 						ContentType: to.Ptr("application/xml"),
	// 					},
	// 					{
	// 						ContentType: to.Ptr("application/json"),
	// 				}},
	// 				StatusCode: to.Ptr[int32](200),
	// 		}},
	// 		Method: to.Ptr("POST"),
	// 		DisplayName: to.Ptr("createUser2"),
	// 		URLTemplate: to.Ptr("/user1"),
	// 	},
	// }
}
