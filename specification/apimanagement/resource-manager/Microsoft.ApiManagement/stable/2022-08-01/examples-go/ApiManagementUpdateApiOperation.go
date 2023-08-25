package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateApiOperation.json
func ExampleAPIOperationClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIOperationClient().Update(ctx, "rg1", "apimService1", "echo-api", "operationId", "*", armapimanagement.OperationUpdateContract{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationContract = armapimanagement.OperationContract{
	// 	Name: to.Ptr("57d2ef278aa04f0ad01d6cdc"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis/operations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/57d2ef278aa04f0888cba3f3/operations/57d2ef278aa04f0ad01d6cdc"),
	// 	Properties: &armapimanagement.OperationContractProperties{
	// 		TemplateParameters: []*armapimanagement.ParameterContract{
	// 		},
	// 		Request: &armapimanagement.RequestContract{
	// 			Description: to.Ptr("IFazioService_CancelOrder_InputMessage"),
	// 			Headers: []*armapimanagement.ParameterContract{
	// 			},
	// 			QueryParameters: []*armapimanagement.ParameterContract{
	// 			},
	// 			Representations: []*armapimanagement.RepresentationContract{
	// 				{
	// 					ContentType: to.Ptr("text/xml"),
	// 					SchemaID: to.Ptr("6980a395-f08b-4a59-8295-1440cbd909b8"),
	// 					TypeName: to.Ptr("CancelOrder"),
	// 			}},
	// 		},
	// 		Responses: []*armapimanagement.ResponseContract{
	// 			{
	// 				Description: to.Ptr("IFazioService_CancelOrder_OutputMessage"),
	// 				Headers: []*armapimanagement.ParameterContract{
	// 				},
	// 				Representations: []*armapimanagement.RepresentationContract{
	// 					{
	// 						ContentType: to.Ptr("text/xml"),
	// 						SchemaID: to.Ptr("6980a395-f08b-4a59-8295-1440cbd909b8"),
	// 						TypeName: to.Ptr("CancelOrderResponse"),
	// 				}},
	// 				StatusCode: to.Ptr[int32](200),
	// 		}},
	// 		Method: to.Ptr("POST"),
	// 		DisplayName: to.Ptr("CancelOrder"),
	// 		URLTemplate: to.Ptr("/?soapAction=http://tempuri.org/IFazioService/CancelOrder"),
	// 	},
	// }
}
