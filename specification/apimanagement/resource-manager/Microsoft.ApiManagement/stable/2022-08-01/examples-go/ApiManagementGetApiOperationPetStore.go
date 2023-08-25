package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiOperationPetStore.json
func ExampleAPIOperationClient_Get_apiManagementGetApiOperationPetStore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIOperationClient().Get(ctx, "rg1", "apimService1", "swagger-petstore", "loginUser", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationContract = armapimanagement.OperationContract{
	// 	Name: to.Ptr("loginUser"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis/operations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/swagger-petstore/operations/loginUser"),
	// 	Properties: &armapimanagement.OperationContractProperties{
	// 		Description: to.Ptr(""),
	// 		TemplateParameters: []*armapimanagement.ParameterContract{
	// 			{
	// 				Name: to.Ptr("username"),
	// 				Type: to.Ptr("string"),
	// 				Description: to.Ptr("The user name for login"),
	// 				Required: to.Ptr(true),
	// 				Values: []*string{
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("password"),
	// 				Type: to.Ptr("string"),
	// 				Description: to.Ptr("The password for login in clear text"),
	// 				Required: to.Ptr(true),
	// 				Values: []*string{
	// 				},
	// 		}},
	// 		Request: &armapimanagement.RequestContract{
	// 			Headers: []*armapimanagement.ParameterContract{
	// 			},
	// 			QueryParameters: []*armapimanagement.ParameterContract{
	// 			},
	// 			Representations: []*armapimanagement.RepresentationContract{
	// 			},
	// 		},
	// 		Responses: []*armapimanagement.ResponseContract{
	// 			{
	// 				Description: to.Ptr("successful operation"),
	// 				Headers: []*armapimanagement.ParameterContract{
	// 					{
	// 						Name: to.Ptr("X-Rate-Limit"),
	// 						Type: to.Ptr("integer"),
	// 						Description: to.Ptr("calls per hour allowed by the user"),
	// 						Values: []*string{
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("X-Expires-After"),
	// 						Type: to.Ptr("string"),
	// 						Description: to.Ptr("date in UTC when token expires"),
	// 						Values: []*string{
	// 						},
	// 				}},
	// 				Representations: []*armapimanagement.RepresentationContract{
	// 					{
	// 						ContentType: to.Ptr("application/xml"),
	// 						SchemaID: to.Ptr("5ba91a35f373b513a0bf31c6"),
	// 						TypeName: to.Ptr("UserLoginGet200ApplicationXmlResponse"),
	// 					},
	// 					{
	// 						ContentType: to.Ptr("application/json"),
	// 						SchemaID: to.Ptr("5ba91a35f373b513a0bf31c6"),
	// 						TypeName: to.Ptr("UserLoginGet200ApplicationJsonResponse"),
	// 				}},
	// 				StatusCode: to.Ptr[int32](200),
	// 			},
	// 			{
	// 				Description: to.Ptr("Invalid username/password supplied"),
	// 				Headers: []*armapimanagement.ParameterContract{
	// 				},
	// 				Representations: []*armapimanagement.RepresentationContract{
	// 					{
	// 						ContentType: to.Ptr("application/xml"),
	// 					},
	// 					{
	// 						ContentType: to.Ptr("application/json"),
	// 				}},
	// 				StatusCode: to.Ptr[int32](400),
	// 		}},
	// 		Method: to.Ptr("GET"),
	// 		DisplayName: to.Ptr("Logs user into the system"),
	// 		URLTemplate: to.Ptr("/user/login?username={username}&password={password}"),
	// 	},
	// }
}
