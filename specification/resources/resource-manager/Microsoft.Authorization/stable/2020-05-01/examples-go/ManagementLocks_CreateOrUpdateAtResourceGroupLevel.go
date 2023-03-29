package armlocks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlocks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Authorization/stable/2020-05-01/examples/ManagementLocks_CreateOrUpdateAtResourceGroupLevel.json
func ExampleManagementLocksClient_CreateOrUpdateAtResourceGroupLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlocks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagementLocksClient().CreateOrUpdateAtResourceGroupLevel(ctx, "resourcegroupname", "testlock", armlocks.ManagementLockObject{
		Properties: &armlocks.ManagementLockProperties{
			Level: to.Ptr(armlocks.LockLevelReadOnly),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagementLockObject = armlocks.ManagementLockObject{
	// 	Name: to.Ptr("testlock"),
	// 	Type: to.Ptr("Microsoft.Authorization/locks"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourcegroupname/providers/Microsoft.Authorization/locks/testlock"),
	// 	Properties: &armlocks.ManagementLockProperties{
	// 		Level: to.Ptr(armlocks.LockLevelReadOnly),
	// 	},
	// }
}
