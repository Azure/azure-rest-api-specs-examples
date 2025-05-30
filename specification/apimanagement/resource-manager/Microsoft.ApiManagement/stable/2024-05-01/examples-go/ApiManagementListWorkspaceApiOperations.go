package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListWorkspaceApiOperations.json
func ExampleWorkspaceAPIOperationClient_NewListByAPIPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceAPIOperationClient().NewListByAPIPager("rg1", "apimService1", "wks1", "57d2ef278aa04f0888cba3f3", &armapimanagement.WorkspaceAPIOperationClientListByAPIOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
		Tags: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.OperationCollection = armapimanagement.OperationCollection{
		// 	Count: to.Ptr[int64](5),
		// 	Value: []*armapimanagement.OperationContract{
		// 		{
		// 			Name: to.Ptr("57d2ef278aa04f0ad01d6cdc"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/apis/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/57d2ef278aa04f0888cba3f3/operations/57d2ef278aa04f0ad01d6cdc"),
		// 			Properties: &armapimanagement.OperationContractProperties{
		// 				Method: to.Ptr("POST"),
		// 				DisplayName: to.Ptr("CancelOrder"),
		// 				URLTemplate: to.Ptr("/?soapAction=http://tempuri.org/IFazioService/CancelOrder"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("57d2ef278aa04f0ad01d6cda"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/apis/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/57d2ef278aa04f0888cba3f3/operations/57d2ef278aa04f0ad01d6cda"),
		// 			Properties: &armapimanagement.OperationContractProperties{
		// 				Method: to.Ptr("POST"),
		// 				DisplayName: to.Ptr("GetMostRecentOrder"),
		// 				URLTemplate: to.Ptr("/?soapAction=http://tempuri.org/IFazioService/GetMostRecentOrder"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("57d2ef278aa04f0ad01d6cd9"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/apis/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/57d2ef278aa04f0888cba3f3/operations/57d2ef278aa04f0ad01d6cd9"),
		// 			Properties: &armapimanagement.OperationContractProperties{
		// 				Method: to.Ptr("POST"),
		// 				DisplayName: to.Ptr("GetOpenOrders"),
		// 				URLTemplate: to.Ptr("/?soapAction=http://tempuri.org/IFazioService/GetOpenOrders"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("57d2ef278aa04f0ad01d6cdb"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/apis/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/57d2ef278aa04f0888cba3f3/operations/57d2ef278aa04f0ad01d6cdb"),
		// 			Properties: &armapimanagement.OperationContractProperties{
		// 				Method: to.Ptr("POST"),
		// 				DisplayName: to.Ptr("GetOrder"),
		// 				URLTemplate: to.Ptr("/?soapAction=http://tempuri.org/IFazioService/GetOrder"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("57d2ef278aa04f0ad01d6cd8"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/apis/operations"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/57d2ef278aa04f0888cba3f3/operations/57d2ef278aa04f0ad01d6cd8"),
		// 			Properties: &armapimanagement.OperationContractProperties{
		// 				Method: to.Ptr("POST"),
		// 				DisplayName: to.Ptr("submitOrder"),
		// 				URLTemplate: to.Ptr("/?soapAction=http://tempuri.org/IFazioService/submitOrder"),
		// 			},
		// 	}},
		// }
	}
}
