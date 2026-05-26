package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementGetWorkspaceProduct.json
func ExampleWorkspaceProductClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceProductClient().Get(ctx, "rg1", "apimService1", "wks1", "unlimited", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.WorkspaceProductClientGetResponse{
	// 	ProductContract: armapimanagement.ProductContract{
	// 		Name: to.Ptr("unlimited"),
	// 		Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/products"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/products/unlimited"),
	// 		Properties: &armapimanagement.ProductContractProperties{
	// 			Description: to.Ptr("Subscribers have completely unlimited access to the API. Administrator approval is required."),
	// 			ApprovalRequired: to.Ptr(true),
	// 			DisplayName: to.Ptr("Unlimited"),
	// 			State: to.Ptr(armapimanagement.ProductStatePublished),
	// 			SubscriptionRequired: to.Ptr(true),
	// 			SubscriptionsLimit: to.Ptr[int32](1),
	// 		},
	// 	},
	// }
}
