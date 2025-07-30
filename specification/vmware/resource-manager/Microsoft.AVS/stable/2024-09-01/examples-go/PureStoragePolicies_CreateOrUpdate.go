package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2024-09-01/PureStoragePolicies_CreateOrUpdate.json
func ExamplePureStoragePoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPureStoragePoliciesClient().BeginCreateOrUpdate(ctx, "group1", "cloud1", "storagePolicy1", armavs.PureStoragePolicy{
		Properties: &armavs.PureStoragePolicyProperties{
			StoragePolicyDefinition: to.Ptr("storagePolicyDefinition1"),
			StoragePoolID:           to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/PureStorage.Block/storagePools/storagePool1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.PureStoragePoliciesClientCreateOrUpdateResponse{
	// 	PureStoragePolicy: &armavs.PureStoragePolicy{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/pureStoragePolicies/storagePolicy1"),
	// 		Name: to.Ptr("storagePolicy1"),
	// 		Properties: &armavs.PureStoragePolicyProperties{
	// 			StoragePolicyDefinition: to.Ptr("storagePolicyDefinition1"),
	// 			StoragePoolID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/PureStorage.Block/storagePools/storagePool1"),
	// 			ProvisioningState: to.Ptr(armavs.PureStoragePolicyProvisioningStateSucceeded),
	// 		},
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/pureStoragePolicies"),
	// 	},
	// }
}
