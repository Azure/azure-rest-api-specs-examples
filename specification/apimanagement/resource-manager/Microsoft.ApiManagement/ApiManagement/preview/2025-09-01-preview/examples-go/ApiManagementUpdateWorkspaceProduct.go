package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementUpdateWorkspaceProduct.json
func ExampleWorkspaceProductClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceProductClient().Update(ctx, "rg1", "apimService1", "wks1", "testproduct", "*", armapimanagement.ProductUpdateParameters{
		Properties: &armapimanagement.ProductUpdateProperties{
			DisplayName: to.Ptr("Test Template ProductName 4"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.WorkspaceProductClientUpdateResponse{
	// 	ProductContract: armapimanagement.ProductContract{
	// 		Name: to.Ptr("testproduct"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/products"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/products/testproduct"),
	// 		Properties: &armapimanagement.ProductContractProperties{
	// 			Description: to.Ptr("Subscribers have completely unlimited access to the API. Administrator approval is required."),
	// 			ApprovalRequired: to.Ptr(true),
	// 			DisplayName: to.Ptr("Test Template ProductName 4"),
	// 			State: to.Ptr(armapimanagement.ProductStatePublished),
	// 			SubscriptionRequired: to.Ptr(true),
	// 			SubscriptionsLimit: to.Ptr[int32](1),
	// 		},
	// 	},
	// }
}
