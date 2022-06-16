package armlocks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlocks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Authorization/stable/2020-05-01/examples/ManagementLocks_CreateOrUpdateAtResourceLevel.json
func ExampleManagementLocksClient_CreateOrUpdateAtResourceLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armlocks.NewManagementLocksClient("subscriptionId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdateAtResourceLevel(ctx,
		"resourcegroupname",
		"Microsoft.Storage",
		"parentResourcePath",
		"storageAccounts",
		"teststorageaccount",
		"testlock",
		armlocks.ManagementLockObject{
			Properties: &armlocks.ManagementLockProperties{
				Level: to.Ptr(armlocks.LockLevelReadOnly),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
