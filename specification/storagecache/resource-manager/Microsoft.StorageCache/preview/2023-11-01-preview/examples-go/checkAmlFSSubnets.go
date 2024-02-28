package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/checkAmlFSSubnets.json
func ExampleManagementClient_CheckAmlFSSubnets() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewManagementClient().CheckAmlFSSubnets(ctx, &armstoragecache.ManagementClientCheckAmlFSSubnetsOptions{AmlFilesystemSubnetInfo: &armstoragecache.AmlFilesystemSubnetInfo{
		FilesystemSubnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/fsSub"),
		SKU: &armstoragecache.SKUName{
			Name: to.Ptr("AMLFS-Durable-Premium-125"),
		},
		StorageCapacityTiB: to.Ptr[float32](16),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
