package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetSubscription.json
func ExampleSubscriptionClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubscriptionClient().Get(ctx, "rg1", "apimService1", "5931a769d8d14f0ad8ce13b8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SubscriptionContract = armapimanagement.SubscriptionContract{
	// 	Name: to.Ptr("5931a769d8d14f0ad8ce13b8"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/subscriptions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/subscriptions/5931a769d8d14f0ad8ce13b8"),
	// 	Properties: &armapimanagement.SubscriptionContractProperties{
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-02T17:59:06.223Z"); return t}()),
	// 		DisplayName: to.Ptr("Unlimited"),
	// 		OwnerID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/5931a75ae4bbd512a88c680b"),
	// 		Scope: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/5600b59475ff190048060002"),
	// 		State: to.Ptr(armapimanagement.SubscriptionStateSubmitted),
	// 	},
	// }
}
