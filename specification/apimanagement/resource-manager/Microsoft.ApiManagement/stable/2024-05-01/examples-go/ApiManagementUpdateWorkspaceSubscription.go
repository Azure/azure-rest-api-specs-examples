package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdateWorkspaceSubscription.json
func ExampleWorkspaceSubscriptionClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceSubscriptionClient().Update(ctx, "rg1", "apimService1", "wks1", "testsub", "*", armapimanagement.SubscriptionUpdateParameters{
		Properties: &armapimanagement.SubscriptionUpdateParameterProperties{
			DisplayName: to.Ptr("testsub"),
		},
	}, &armapimanagement.WorkspaceSubscriptionClientUpdateOptions{Notify: nil,
		AppType: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SubscriptionContract = armapimanagement.SubscriptionContract{
	// 	Name: to.Ptr("testsub"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/subscriptions"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/subscriptions/testsub"),
	// 	Properties: &armapimanagement.SubscriptionContractProperties{
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T17:59:06.223Z"); return t}()),
	// 		DisplayName: to.Ptr("testsub"),
	// 		OwnerID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/5931a75ae4bbd512a88c680b"),
	// 		Scope: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/products/5600b59475ff190048060002"),
	// 		State: to.Ptr(armapimanagement.SubscriptionStateSubmitted),
	// 	},
	// }
}
